package beancore

import (
	"encoding/hex"
	"errors"
	"fmt"
	"hash/fnv"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/hmans/beans/internal/search"
	"github.com/hmans/beans/pkg/bean"
	"github.com/hmans/beans/pkg/config"
)

const BeansDir = ".beans"
const ArchiveDir = "archive"

var ErrNotFound = errors.New("bean not found")

type ETagMismatchError struct {
	Provided string
	Current  string
}

func (e *ETagMismatchError) Error() string {
	return fmt.Sprintf("etag mismatch: provided %s, current is %s", e.Provided, e.Current)
}

type ETagRequiredError struct{}

func (e *ETagRequiredError) Error() string {
	return "if-match etag is required (set require_if_match: false in config to disable)"
}

type Core struct {
	root   string
	config *config.Config

	mu    sync.RWMutex
	beans map[string]*bean.Bean
	dirty map[string]bool

	searchIndex *search.Index

	warnWriter io.Writer
}

func New(root string, cfg *config.Config) *Core {
	return &Core{
		root:       root,
		config:     cfg,
		beans:      make(map[string]*bean.Bean),
		dirty:      make(map[string]bool),
		warnWriter: os.Stderr,
	}
}

func (c *Core) SetWarnWriter(w io.Writer) {
	c.warnWriter = w
}

func (c *Core) logWarn(format string, args ...any) {
	if c.warnWriter != nil {
		fmt.Fprintf(c.warnWriter, "warning: "+format+"\n", args...)
	}
}

func (c *Core) Root() string {
	return c.root
}

func (c *Core) Config() *config.Config {
	return c.config
}

func (c *Core) Load() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.loadFromDisk()
}

func (c *Core) loadFromDisk() error {
	c.migrateLegacyDirs()

	c.beans = make(map[string]*bean.Bean)
	c.dirty = make(map[string]bool)

	err := filepath.WalkDir(c.root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() && path != c.root {
			if strings.HasPrefix(d.Name(), ".") {
				return filepath.SkipDir
			}
			return nil
		}

		if !strings.HasSuffix(d.Name(), ".md") {
			return nil
		}

		b, loadErr := c.loadBean(path)
		if loadErr != nil {
			return fmt.Errorf("loading %s: %w", path, loadErr)
		}

		c.beans[b.ID] = b
		return nil
	})
	if err != nil {
		return err
	}

	if c.searchIndex != nil {
		c.searchIndex.Close()
		c.searchIndex = nil

		if err := c.ensureSearchIndexLocked(); err != nil {
			c.logWarn("failed to reinitialize search index after reload: %v", err)
		}
	}

	return nil
}

func (c *Core) loadBean(path string) (*bean.Bean, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	b, err := bean.Parse(f)
	if err != nil {
		return nil, err
	}

	relPath, err := filepath.Rel(c.root, path)
	if err != nil {
		return nil, err
	}
	b.Path = relPath

	filename := filepath.Base(path)
	b.ID, b.Slug = bean.ParseFilename(filename)

	if b.Type == "" {
		b.Type = "task"
	}
	if b.Priority == "" {
		b.Priority = "normal"
	}
	if b.Tags == nil {
		b.Tags = []string{}
	}
	if b.Blocking == nil {
		b.Blocking = []string{}
	}
	if b.CreatedAt == nil {
		if b.UpdatedAt != nil {
			b.CreatedAt = b.UpdatedAt
		} else {
			info, statErr := os.Stat(path)
			if statErr == nil {
				modTime := info.ModTime().UTC().Truncate(time.Second)
				b.CreatedAt = &modTime
			}
		}
	}
	if b.UpdatedAt == nil {
		b.UpdatedAt = b.CreatedAt
	}

	return b, nil
}

func (c *Core) ensureSearchIndexLocked() error {
	if c.searchIndex != nil {
		return nil
	}

	idx, err := search.NewIndex()
	if err != nil {
		return fmt.Errorf("initializing search index: %w", err)
	}

	c.searchIndex = idx

	allBeans := make([]*bean.Bean, 0, len(c.beans))
	for _, b := range c.beans {
		allBeans = append(allBeans, b)
	}
	if err := c.searchIndex.IndexBeans(allBeans); err != nil {
		return fmt.Errorf("populating search index: %w", err)
	}

	return nil
}

func (c *Core) Search(query string) ([]*bean.Bean, error) {
	c.mu.Lock()
	if err := c.ensureSearchIndexLocked(); err != nil {
		c.mu.Unlock()
		return nil, err
	}
	idx := c.searchIndex
	c.mu.Unlock()

	ids, err := idx.Search(query, search.DefaultSearchLimit)
	if err != nil {
		return nil, err
	}

	c.mu.RLock()
	defer c.mu.RUnlock()

	result := make([]*bean.Bean, 0, len(ids))
	for _, id := range ids {
		if b, ok := c.beans[id]; ok {
			result = append(result, b)
		}
	}
	return result, nil
}

func (c *Core) All() []*bean.Bean {
	c.mu.RLock()
	defer c.mu.RUnlock()

	result := make([]*bean.Bean, 0, len(c.beans))
	for _, b := range c.beans {
		result = append(result, b)
	}
	return result
}

func (c *Core) Get(id string) (*bean.Bean, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if b, ok := c.beans[id]; ok {
		return b, nil
	}

	if c.config != nil && c.config.Beans.Prefix != "" && !strings.HasPrefix(id, c.config.Beans.Prefix) {
		if b, ok := c.beans[c.config.Beans.Prefix+id]; ok {
			return b, nil
		}
	}

	return nil, ErrNotFound
}

func (c *Core) NormalizeID(id string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if _, ok := c.beans[id]; ok {
		return id, true
	}

	if c.config != nil && c.config.Beans.Prefix != "" && !strings.HasPrefix(id, c.config.Beans.Prefix) {
		fullID := c.config.Beans.Prefix + id
		if _, ok := c.beans[fullID]; ok {
			return fullID, true
		}
	}

	return id, false
}

type UpdateOption func(*updateOptions)

type updateOptions struct {
	persist bool
}

func defaultUpdateOptions() updateOptions {
	return updateOptions{persist: true}
}

func WithPersist(persist bool) UpdateOption {
	return func(o *updateOptions) {
		o.persist = persist
	}
}

func (c *Core) IsDirty(id string) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.dirty[id]
}

func (c *Core) DirtyIDs() []string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	ids := make([]string, 0, len(c.dirty))
	for id := range c.dirty {
		ids = append(ids, id)
	}
	return ids
}

func (c *Core) HasDirty() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.dirty) > 0
}

func (c *Core) Create(b *bean.Bean, opts ...UpdateOption) error {
	o := defaultUpdateOptions()
	for _, opt := range opts {
		opt(&o)
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if b.ID == "" {
		prefix := ""
		length := 4
		if c.config != nil {
			prefix = c.config.Beans.Prefix
			if c.config.Beans.IDLength > 0 {
				length = c.config.Beans.IDLength
			}
		}
		id, err := bean.NewID(prefix, length)
		if err != nil {
			return fmt.Errorf("generating bean ID: %w", err)
		}
		b.ID = id
	}

	now := time.Now().UTC().Truncate(time.Second)
	b.CreatedAt = &now
	b.UpdatedAt = &now

	if o.persist {
		if err := c.saveToDisk(b); err != nil {
			return err
		}
	} else {
		c.dirty[b.ID] = true
	}

	c.beans[b.ID] = b

	if c.searchIndex != nil {
		if err := c.searchIndex.IndexBean(b); err != nil {
			c.logWarn("failed to index bean %s: %v", b.ID, err)
		}
	}

	return nil
}

func (c *Core) Update(b *bean.Bean, ifMatch *string, opts ...UpdateOption) error {
	o := defaultUpdateOptions()
	for _, opt := range opts {
		opt(&o)
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	storedBean, ok := c.beans[b.ID]
	if !ok {
		return ErrNotFound
	}

	requireIfMatch := c.config != nil && c.config.Beans.RequireIfMatch

	if requireIfMatch && (ifMatch == nil || *ifMatch == "") {
		return &ETagRequiredError{}
	}

	if ifMatch != nil && *ifMatch != "" {
		var currentETag string
		if storedBean.Path != "" && !c.dirty[b.ID] {
			diskPath := filepath.Join(c.root, storedBean.Path)
			content, err := os.ReadFile(diskPath)
			if err != nil {
				currentETag = storedBean.ETag()
			} else {
				h := fnv.New64a()
				h.Write(content)
				currentETag = hex.EncodeToString(h.Sum(nil))
			}
		} else {
			currentETag = storedBean.ETag()
		}

		if currentETag != *ifMatch {
			return &ETagMismatchError{
				Provided: *ifMatch,
				Current:  currentETag,
			}
		}
	}

	now := time.Now().UTC().Truncate(time.Second)
	b.UpdatedAt = &now

	if o.persist {
		if err := c.saveToDisk(b); err != nil {
			return err
		}
		delete(c.dirty, b.ID)
	} else {
		c.dirty[b.ID] = true
	}

	c.beans[b.ID] = b

	if c.searchIndex != nil {
		if err := c.searchIndex.IndexBean(b); err != nil {
			c.logWarn("failed to update bean %s in search index: %v", b.ID, err)
		}
	}

	return nil
}

func (c *Core) saveToDisk(b *bean.Bean) error {
	var path string
	if b.Path != "" {
		path = filepath.Join(c.root, b.Path)
	} else {
		filename := bean.BuildFilename(b.ID, b.Slug)
		path = filepath.Join(c.root, filename)
		b.Path = filename
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("creating directory: %w", err)
	}

	content, err := b.Render()
	if err != nil {
		return err
	}

	if err := os.WriteFile(path, content, 0644); err != nil {
		return fmt.Errorf("writing file: %w", err)
	}

	return nil
}

func (c *Core) Delete(id string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	targetID := id
	targetBean, ok := c.beans[id]

	if !ok && c.config != nil && c.config.Beans.Prefix != "" && !strings.HasPrefix(id, c.config.Beans.Prefix) {
		fullID := c.config.Beans.Prefix + id
		if b, found := c.beans[fullID]; found {
			targetID = fullID
			targetBean = b
			ok = true
		}
	}

	if !ok {
		return ErrNotFound
	}

	path := filepath.Join(c.root, targetBean.Path)
	if err := os.Remove(path); err != nil {
		return err
	}

	delete(c.beans, targetID)

	if c.searchIndex != nil {
		if err := c.searchIndex.DeleteBean(targetID); err != nil {
			c.logWarn("failed to remove bean %s from search index: %v", targetID, err)
		}
	}

	return nil
}

func (c *Core) Archive(id string) error {
	c.mu.Lock()

	targetBean, targetID, err := c.findBeanLocked(id)
	if err != nil {
		c.mu.Unlock()
		return err
	}

	if c.isArchivedPath(targetBean.Path) {
		c.mu.Unlock()
		return nil
	}

	archivePath := filepath.Join(c.root, ArchiveDir)
	if err := os.MkdirAll(archivePath, 0755); err != nil {
		c.mu.Unlock()
		return fmt.Errorf("creating archive directory: %w", err)
	}

	oldPath := filepath.Join(c.root, targetBean.Path)
	newRelPath := filepath.Join(ArchiveDir, filepath.Base(targetBean.Path))
	newPath := filepath.Join(c.root, newRelPath)

	if err := os.Rename(oldPath, newPath); err != nil {
		c.mu.Unlock()
		return fmt.Errorf("moving bean to archive: %w", err)
	}

	targetBean.Path = newRelPath
	c.beans[targetID] = targetBean
	c.mu.Unlock()

	return nil
}

func (c *Core) Unarchive(id string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	targetBean, targetID, err := c.findBeanLocked(id)
	if err != nil {
		return err
	}

	if !c.isArchivedPath(targetBean.Path) {
		return nil
	}

	oldPath := filepath.Join(c.root, targetBean.Path)
	newRelPath := filepath.Base(targetBean.Path)
	newPath := filepath.Join(c.root, newRelPath)

	if err := os.Rename(oldPath, newPath); err != nil {
		return fmt.Errorf("moving bean from archive: %w", err)
	}

	targetBean.Path = newRelPath
	c.beans[targetID] = targetBean

	return nil
}

func (c *Core) IsArchived(id string) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	b, _, err := c.findBeanLocked(id)
	if err != nil {
		return false
	}

	return c.isArchivedPath(b.Path)
}

func (c *Core) isArchivedPath(path string) bool {
	return strings.HasPrefix(path, ArchiveDir+string(filepath.Separator)) ||
		strings.HasPrefix(path, ArchiveDir+"/")
}

func (c *Core) normalizeID(id string) string {
	if c.config != nil && c.config.Beans.Prefix != "" && !strings.HasPrefix(id, c.config.Beans.Prefix) {
		return c.config.Beans.Prefix + id
	}
	return id
}

func (c *Core) findBeanLocked(id string) (*bean.Bean, string, error) {
	if b, ok := c.beans[id]; ok {
		return b, id, nil
	}

	fullID := c.normalizeID(id)
	if fullID != id {
		if b, ok := c.beans[fullID]; ok {
			return b, fullID, nil
		}
	}

	return nil, "", ErrNotFound
}

func (c *Core) GetFromArchive(id string) (*bean.Bean, error) {
	fullID := c.normalizeID(id)

	archiveDir := filepath.Join(c.root, ArchiveDir)
	if _, err := os.Stat(archiveDir); os.IsNotExist(err) {
		return nil, nil
	}

	entries, err := os.ReadDir(archiveDir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".md") {
			continue
		}

		fileID, _ := bean.ParseFilename(entry.Name())
		if fileID == fullID {
			path := filepath.Join(archiveDir, entry.Name())
			return c.loadBean(path)
		}
	}

	return nil, nil
}

func (c *Core) LoadAndUnarchive(id string) (*bean.Bean, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	b, targetID, err := c.findBeanLocked(id)
	if err != nil {
		return nil, ErrNotFound
	}

	if !c.isArchivedPath(b.Path) {
		return b, nil
	}

	oldPath := filepath.Join(c.root, b.Path)
	newRelPath := filepath.Base(b.Path)
	newPath := filepath.Join(c.root, newRelPath)

	if err := os.Rename(oldPath, newPath); err != nil {
		return nil, fmt.Errorf("moving bean from archive: %w", err)
	}

	b.Path = newRelPath
	c.beans[targetID] = b

	return b, nil
}

func (c *Core) Init() error {
	if err := os.MkdirAll(c.root, 0755); err != nil {
		return err
	}
	return writeGitignore(c.root)
}

func (c *Core) FullPath(b *bean.Bean) string {
	return filepath.Join(c.root, b.Path)
}

func (c *Core) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.searchIndex != nil {
		if err := c.searchIndex.Close(); err != nil {
			return err
		}
		c.searchIndex = nil
	}

	return nil
}

func Init(dir string) error {
	beansPath := filepath.Join(dir, BeansDir)
	if err := os.MkdirAll(beansPath, 0755); err != nil {
		return err
	}
	return writeGitignore(beansPath)
}

func writeGitignore(beansDir string) error {
	content := "# Generated by beans init\n.conversations/\n"
	return os.WriteFile(filepath.Join(beansDir, ".gitignore"), []byte(content), 0644)
}

func (c *Core) migrateLegacyDirs() {
	for _, name := range []string{"worktrees", "conversations"} {
		oldPath := filepath.Join(c.root, name)
		newPath := filepath.Join(c.root, "."+name)
		if _, err := os.Stat(oldPath); err == nil {
			if _, err := os.Stat(newPath); os.IsNotExist(err) {
				_ = os.Rename(oldPath, newPath)
			}
		}
	}
}

type Filter struct {
	Status                  []string
	ExcludeStatus           []string
	Type                    []string
	ExcludeType             []string
	Priority                []string
	ExcludePriority         []string
	Tags                    []string
	ExcludeTags             []string
	Search                  *string
	HasParent               *bool
	NoParent                *bool
	ParentID                *string
	HasBlocking             *bool
	NoBlocking              *bool
	IsBlocked               *bool
	ExcludeImplicitTerminal *bool
	ExcludeArchived         *bool
}

func BoolPtr(b bool) *bool { return &b }

func (c *Core) FilteredBeans(f *Filter) ([]*bean.Bean, error) {
	var beans []*bean.Bean

	if f != nil && f.Search != nil && *f.Search != "" {
		var err error
		beans, err = c.Search(*f.Search)
		if err != nil {
			return nil, err
		}
	} else {
		beans = c.All()
	}

	if f == nil {
		return beans, nil
	}

	var result []*bean.Bean
	for _, b := range beans {
		if !c.matchesFilter(b, f) {
			continue
		}
		result = append(result, b)
	}

	return result, nil
}

func (c *Core) matchesFilter(b *bean.Bean, f *Filter) bool {
	if len(f.Status) > 0 && !contains(f.Status, b.Status) {
		return false
	}
	if len(f.ExcludeStatus) > 0 && contains(f.ExcludeStatus, b.Status) {
		return false
	}
	if len(f.Type) > 0 && !contains(f.Type, b.Type) {
		return false
	}
	if len(f.ExcludeType) > 0 && contains(f.ExcludeType, b.Type) {
		return false
	}
	if len(f.Priority) > 0 && !contains(f.Priority, b.Priority) {
		return false
	}
	if len(f.ExcludePriority) > 0 && contains(f.ExcludePriority, b.Priority) {
		return false
	}
	if len(f.Tags) > 0 && !hasAnyTag(b, f.Tags) {
		return false
	}
	if len(f.ExcludeTags) > 0 && hasAnyTag(b, f.ExcludeTags) {
		return false
	}
	if f.HasParent != nil && *f.HasParent && b.Parent == "" {
		return false
	}
	if f.NoParent != nil && *f.NoParent && b.Parent != "" {
		return false
	}
	if f.ParentID != nil && *f.ParentID != "" {
		fullID, _ := c.NormalizeID(*f.ParentID)
		if b.Parent != fullID {
			return false
		}
	}
	if f.HasBlocking != nil && *f.HasBlocking && len(b.Blocking) == 0 {
		return false
	}
	if f.NoBlocking != nil && *f.NoBlocking && len(b.Blocking) > 0 {
		return false
	}
	if f.IsBlocked != nil {
		blocked := c.IsBlocked(b.ID)
		if *f.IsBlocked != blocked {
			return false
		}
	}
	if f.ExcludeImplicitTerminal != nil && *f.ExcludeImplicitTerminal {
		if status, _ := c.ImplicitStatus(b.ID); status != "" {
			return false
		}
	}
	if f.ExcludeArchived != nil && *f.ExcludeArchived {
		if c.isArchivedPath(b.Path) {
			return false
		}
	}

	return true
}

func contains(slice []string, s string) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}
	return false
}

func hasAnyTag(b *bean.Bean, tags []string) bool {
	for _, t := range tags {
		for _, bt := range b.Tags {
			if t == bt {
				return true
			}
		}
	}
	return false
}
