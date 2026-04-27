package ui

import (
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"testing"

	"github.com/charmbracelet/lipgloss"
	"github.com/hmans/beans/pkg/bean"
	"github.com/hmans/beans/pkg/beancore"
	"github.com/hmans/beans/pkg/config"
)

func TestBuildTree(t *testing.T) {
	// Create test beans with parent relationships:
	// milestone1
	//   └── epic1
	//       └── task1
	// task2 (orphan)

	milestone1 := &bean.Bean{ID: "m1", Title: "Milestone 1", Type: "milestone"}
	epic1 := &bean.Bean{ID: "e1", Title: "Epic 1", Type: "epic", Parent: "m1"}
	task1 := &bean.Bean{ID: "t1", Title: "Task 1", Type: "task", Parent: "e1"}
	task2 := &bean.Bean{ID: "t2", Title: "Task 2", Type: "task"} // orphan

	allBeans := []*bean.Bean{milestone1, epic1, task1, task2}

	// Identity sort function (no sorting)
	noSort := func(b []*bean.Bean) {}

	t.Run("all beans matched", func(t *testing.T) {
		tree := BuildTree(allBeans, allBeans, noSort, nil)

		// Should have 2 root nodes: milestone1 and task2
		if len(tree) != 2 {
			t.Errorf("expected 2 root nodes, got %d", len(tree))
		}

		// Find milestone node
		var milestoneNode *TreeNode
		for _, n := range tree {
			if n.Bean.ID == "m1" {
				milestoneNode = n
				break
			}
		}
		if milestoneNode == nil {
			t.Fatal("milestone node not found")
		}
		if !milestoneNode.Matched {
			t.Error("milestone should be marked as matched")
		}

		// Milestone should have epic as child
		if len(milestoneNode.Children) != 1 {
			t.Errorf("milestone should have 1 child, got %d", len(milestoneNode.Children))
		}
		epicNode := milestoneNode.Children[0]
		if epicNode.Bean.ID != "e1" {
			t.Errorf("expected epic child, got %s", epicNode.Bean.ID)
		}

		// Epic should have task as child
		if len(epicNode.Children) != 1 {
			t.Errorf("epic should have 1 child, got %d", len(epicNode.Children))
		}
		taskNode := epicNode.Children[0]
		if taskNode.Bean.ID != "t1" {
			t.Errorf("expected task child, got %s", taskNode.Bean.ID)
		}
	})

	t.Run("filter leaf only - ancestors included", func(t *testing.T) {
		// Only task1 matched, but ancestors should be included
		matchedBeans := []*bean.Bean{task1}
		tree := BuildTree(matchedBeans, allBeans, noSort, nil)

		// Should have 1 root: milestone (as ancestor)
		if len(tree) != 1 {
			t.Errorf("expected 1 root node, got %d", len(tree))
		}

		milestoneNode := tree[0]
		if milestoneNode.Bean.ID != "m1" {
			t.Errorf("expected milestone as root, got %s", milestoneNode.Bean.ID)
		}
		if milestoneNode.Matched {
			t.Error("milestone should NOT be marked as matched (it's an ancestor)")
		}

		// Should have epic as child (also ancestor)
		if len(milestoneNode.Children) != 1 {
			t.Fatalf("milestone should have 1 child, got %d", len(milestoneNode.Children))
		}
		epicNode := milestoneNode.Children[0]
		if epicNode.Matched {
			t.Error("epic should NOT be marked as matched (it's an ancestor)")
		}

		// Task should be matched
		if len(epicNode.Children) != 1 {
			t.Fatalf("epic should have 1 child, got %d", len(epicNode.Children))
		}
		taskNode := epicNode.Children[0]
		if !taskNode.Matched {
			t.Error("task should be marked as matched")
		}
	})

	t.Run("filter middle - ancestors included", func(t *testing.T) {
		// Only epic1 matched
		matchedBeans := []*bean.Bean{epic1}
		tree := BuildTree(matchedBeans, allBeans, noSort, nil)

		// Should have 1 root: milestone (ancestor)
		if len(tree) != 1 {
			t.Errorf("expected 1 root node, got %d", len(tree))
		}

		milestoneNode := tree[0]
		if milestoneNode.Matched {
			t.Error("milestone should NOT be marked as matched")
		}

		epicNode := milestoneNode.Children[0]
		if !epicNode.Matched {
			t.Error("epic should be marked as matched")
		}

		// Epic should have no children (task1 was not matched)
		if len(epicNode.Children) != 0 {
			t.Errorf("epic should have 0 children (task not matched), got %d", len(epicNode.Children))
		}
	})

	t.Run("orphan bean", func(t *testing.T) {
		matchedBeans := []*bean.Bean{task2}
		tree := BuildTree(matchedBeans, allBeans, noSort, nil)

		if len(tree) != 1 {
			t.Errorf("expected 1 root node, got %d", len(tree))
		}
		if tree[0].Bean.ID != "t2" {
			t.Errorf("expected task2 as root, got %s", tree[0].Bean.ID)
		}
		if !tree[0].Matched {
			t.Error("task2 should be marked as matched")
		}
	})

	t.Run("broken parent link", func(t *testing.T) {
		// Bean with parent that doesn't exist
		brokenBean := &bean.Bean{ID: "broken", Title: "Broken", Parent: "nonexistent"}
		matchedBeans := []*bean.Bean{brokenBean}
		allBeansWithBroken := append(allBeans, brokenBean)

		tree := BuildTree(matchedBeans, allBeansWithBroken, noSort, nil)

		// Should be treated as root (parent not found)
		if len(tree) != 1 {
			t.Errorf("expected 1 root node, got %d", len(tree))
		}
		if tree[0].Bean.ID != "broken" {
			t.Errorf("expected broken bean as root, got %s", tree[0].Bean.ID)
		}
	})
}

func TestTreeNodeToJSON(t *testing.T) {
	b := &bean.Bean{
		ID:       "test-id",
		Slug:     "test-slug",
		Path:     "test.md",
		Title:    "Test Title",
		Status:   "todo",
		Type:     "task",
		Priority: "high",
		Tags:     []string{"tag1", "tag2"},
		Body:     "Test body content",
	}

	node := &TreeNode{
		Bean:    b,
		Matched: true,
		Children: []*TreeNode{
			{
				Bean:    &bean.Bean{ID: "child-id", Title: "Child"},
				Matched: false,
			},
		},
	}

	t.Run("without full body", func(t *testing.T) {
		json := node.ToJSON(false)
		if json.ID != "test-id" {
			t.Errorf("expected id 'test-id', got %s", json.ID)
		}
		if json.Body != "" {
			t.Error("body should be empty when includeFull is false")
		}
		if !json.Matched {
			t.Error("matched should be true")
		}
		if len(json.Children) != 1 {
			t.Errorf("expected 1 child, got %d", len(json.Children))
		}
	})

	t.Run("with full body", func(t *testing.T) {
		json := node.ToJSON(true)
		if json.Body != "Test body content" {
			t.Errorf("expected body content, got %s", json.Body)
		}
	})
}

func TestRenderTreeHeaderAlignment(t *testing.T) {
	t.Run("title aligns without tags", func(t *testing.T) {
		nodes := []*TreeNode{
			{
				Bean: &bean.Bean{
					ID:     "beans-a1",
					Title:  "Example title",
					Status: "todo",
					Type:   "task",
				},
				Matched: true,
			},
		}

		output := stripANSI(RenderTree(nodes, config.Default(), 12, false, 100))
		lines := strings.Split(strings.TrimRight(output, "\n"), "\n")
		if len(lines) < 3 {
			t.Fatalf("expected at least 3 lines, got %d", len(lines))
		}

		header := lines[0]
		row := lines[2]

		if got, want := displayIndex(header, "TITLE"), displayIndex(row, "Example title"); got != want {
			t.Fatalf("TITLE column misaligned: header index=%d row index=%d\nheader=%q\nrow=%q", got, want, header, row)
		}
	})

	t.Run("title and tags align with tags column", func(t *testing.T) {
		nodes := []*TreeNode{
			{
				Bean: &bean.Bean{
					ID:     "beans-a1",
					Title:  "Example title",
					Status: "todo",
					Type:   "task",
					Tags:   []string{"alpha"},
				},
				Matched: true,
			},
		}

		output := stripANSI(RenderTree(nodes, config.Default(), 12, true, 160))
		lines := strings.Split(strings.TrimRight(output, "\n"), "\n")
		if len(lines) < 3 {
			t.Fatalf("expected at least 3 lines, got %d", len(lines))
		}

		header := lines[0]
		row := lines[2]

		if got, want := displayIndex(header, "TITLE"), displayIndex(row, "Example title"); got != want {
			t.Fatalf("TITLE column misaligned: header index=%d row index=%d\nheader=%q\nrow=%q", got, want, header, row)
		}
		if got, want := displayIndex(header, "TAGS"), displayIndex(row, " alpha "); got != want {
			t.Fatalf("TAGS column misaligned: header index=%d row index=%d\nheader=%q\nrow=%q", got, want, header, row)
		}
	})
}

func TestRenderTreeHeaderAlignmentWithTmpFixtureData(t *testing.T) {
	fixtureRoot := filepath.Join("..", "..", "tmp", ".beans")
	testCfg := config.Default()
	core := beancore.New(fixtureRoot, testCfg)
	if err := core.Load(); err != nil {
		t.Fatalf("loading tmp fixture beans: %v", err)
	}

	parentID := "fp-6qkk"
	matched, err := core.FilteredBeans(&beancore.Filter{ParentID: &parentID})
	if err != nil {
		t.Fatalf("filtering tmp fixture beans: %v", err)
	}

	allBeans := core.All()
	implicitStatuses := make(map[string]string, len(allBeans))
	for _, b := range allBeans {
		if status, _ := core.ImplicitStatus(b.ID); status != "" {
			implicitStatuses[b.ID] = status
		}
	}

	sortFn := func(beans []*bean.Bean) {
		bean.SortByStatusPriorityAndType(beans, testCfg.StatusNames(), testCfg.PriorityNames(), testCfg.TypeNames())
		sort.SliceStable(beans, func(i, j int) bool {
			return beans[i].ID < beans[j].ID
		})
	}

	tree := BuildTree(matched, allBeans, sortFn, implicitStatuses)

	maxIDWidth := 2
	for _, b := range allBeans {
		if len(b.ID) > maxIDWidth {
			maxIDWidth = len(b.ID)
		}
	}
	maxIDWidth += 2

	hasTags := false
	for _, b := range matched {
		if len(b.Tags) > 0 {
			hasTags = true
			break
		}
	}
	if !hasTags {
		t.Fatal("expected tmp fixture subtree to include tagged matched beans")
	}

	output := stripANSI(RenderTree(tree, testCfg, maxIDWidth, hasTags, 180))
	lines := strings.Split(strings.TrimRight(output, "\n"), "\n")
	if len(lines) < 4 {
		t.Fatalf("expected multiple rendered lines, got %d", len(lines))
	}

	header := lines[0]
	titleCol := displayIndex(header, "TITLE")
	tagsCol := displayIndex(header, "TAGS")
	if titleCol < 0 || tagsCol < 0 {
		t.Fatalf("expected TITLE and TAGS headers, got %q", header)
	}

	rootRow := findRenderedRow(t, lines[2:], "fp-6qkk")
	phaseERow := findRenderedRow(t, lines[2:], "fp-4qw6")
	bugRow := findRenderedRow(t, lines[2:], "fp-j9jk")

	if got := displayIndex(rootRow, "v4 source-host runtime and webhook refactor"); got != titleCol {
		t.Fatalf("root title misaligned: title index=%d want %d\nrow=%q", got, titleCol, rootRow)
	}
	if got := displayIndex(phaseERow, " design-rev"); got != tagsCol {
		t.Fatalf("phase E tag misaligned: tag index=%d want %d\nrow=%q", got, tagsCol, phaseERow)
	}
	if got := displayIndex(bugRow, " code-review"); got != tagsCol {
		t.Fatalf("bug tag misaligned: tag index=%d want %d\nrow=%q", got, tagsCol, bugRow)
	}
}

func TestRenderTreeTaggedRowsStayWithinPaneWidth(t *testing.T) {
	fixtureRoot := filepath.Join("..", "..", "tmp", ".beans")
	testCfg := config.Default()
	core := beancore.New(fixtureRoot, testCfg)
	if err := core.Load(); err != nil {
		t.Fatalf("loading tmp fixture beans: %v", err)
	}

	parentID := "fp-3i6d"
	matched, err := core.FilteredBeans(&beancore.Filter{ParentID: &parentID})
	if err != nil {
		t.Fatalf("filtering tmp fixture beans: %v", err)
	}

	allBeans := core.All()
	implicitStatuses := make(map[string]string, len(allBeans))
	for _, b := range allBeans {
		if status, _ := core.ImplicitStatus(b.ID); status != "" {
			implicitStatuses[b.ID] = status
		}
	}

	sortFn := func(beans []*bean.Bean) {
		bean.SortByStatusPriorityAndType(beans, testCfg.StatusNames(), testCfg.PriorityNames(), testCfg.TypeNames())
	}

	tree := BuildTree(matched, allBeans, sortFn, implicitStatuses)

	maxIDWidth := 2
	for _, b := range allBeans {
		if len(b.ID) > maxIDWidth {
			maxIDWidth = len(b.ID)
		}
	}
	maxIDWidth += 2

	const paneWidth = 145
	output := stripANSI(RenderTree(tree, testCfg, maxIDWidth, true, paneWidth))
	lines := strings.Split(strings.TrimRight(output, "\n"), "\n")
	if len(lines) < 3 {
		t.Fatalf("expected rendered output, got %d lines", len(lines))
	}

	for _, line := range lines {
		if width := lipgloss.Width(line); width > paneWidth-1 {
			t.Fatalf("line exceeds safe pane width: got %d want <= %d\nline=%q", width, paneWidth-1, line)
		}
	}
}

var ansiPattern = regexp.MustCompile(`\x1b\[[0-9;]*m`)

func stripANSI(s string) string {
	return ansiPattern.ReplaceAllString(s, "")
}

func findRenderedRow(t *testing.T, rows []string, id string) string {
	t.Helper()
	for _, row := range rows {
		if strings.Contains(row, id) {
			return row
		}
	}
	t.Fatalf("row for %s not found", id)
	return ""
}

func displayIndex(line, needle string) int {
	idx := strings.Index(line, needle)
	if idx < 0 {
		return -1
	}
	return lipgloss.Width(line[:idx])
}
