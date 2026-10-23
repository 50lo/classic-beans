package commands

import (
	"encoding/json"
	"io"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/hmans/beans/pkg/bean"
	"github.com/hmans/beans/pkg/beancore"
	"github.com/hmans/beans/pkg/config"
)

func setupShowTest(t *testing.T, beans ...*bean.Bean) {
	t.Helper()
	oldCore, oldCfg := core, cfg
	oldJSON, oldRaw, oldBody, oldETag := showJSON, showRaw, showBodyOnly, showETagOnly
	cfg = config.Default()
	core = beancore.New(t.TempDir(), cfg)
	showJSON, showRaw, showBodyOnly, showETagOnly = false, false, false, false
	t.Cleanup(func() {
		if err := core.Close(); err != nil {
			t.Errorf("close core: %v", err)
		}
		core, cfg = oldCore, oldCfg
		showJSON, showRaw, showBodyOnly, showETagOnly = oldJSON, oldRaw, oldBody, oldETag
	})
	for _, b := range beans {
		if err := core.Create(b); err != nil {
			t.Fatalf("create bean: %v", err)
		}
	}
}

func runShowTest(t *testing.T, ids ...string) string {
	t.Helper()
	f, err := os.CreateTemp(t.TempDir(), "stdout")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	oldStdout := os.Stdout
	os.Stdout = f
	defer func() { os.Stdout = oldStdout }()
	if err := showCmd.RunE(showCmd, ids); err != nil {
		t.Fatalf("show: %v", err)
	}
	if _, err := f.Seek(0, io.SeekStart); err != nil {
		t.Fatal(err)
	}
	data, err := io.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}
	return string(data)
}

func TestShowRelationships(t *testing.T) {
	tests := []struct {
		name   string
		bean   *bean.Bean
		others []*bean.Bean
		want   string
	}{
		{
			name: "no relationships",
			bean: &bean.Bean{ID: "test-main"},
		},
		{
			name: "parent only",
			bean: &bean.Bean{ID: "test-main", Parent: "test-parent"},
		},
		{
			name: "blocked by only including missing target",
			bean: &bean.Bean{ID: "test-main", BlockedBy: []string{"test-missing"}},
			want: "Blocked by: test-missing",
		},
		{
			name: "blocks only",
			bean: &bean.Bean{ID: "test-main", Blocking: []string{"test-target"}},
			want: "Blocks: test-target",
		},
		{
			name: "incoming relationships only",
			bean: &bean.Bean{ID: "test-main"},
			others: []*bean.Bean{
				{ID: "test-blocker", Blocking: []string{"test-main"}},
				{ID: "test-target", BlockedBy: []string{"test-main"}},
				{ID: "test-child", Parent: "test-main"},
			},
			want: "Blocks: test-target\nBlocked by: test-blocker",
		},
		{
			name: "combined relationships sorted and deduplicated",
			bean: &bean.Bean{
				ID: "test-main", Parent: "test-parent",
				BlockedBy: []string{"test-blocker-z", "test-blocker-a"},
				Blocking:  []string{"test-target-z", "test-target-a"},
			},
			others: []*bean.Bean{
				{ID: "test-blocker-a", Status: "completed", Blocking: []string{"test-main"}},
				{ID: "test-blocker-b", Blocking: []string{"test-main"}},
				{ID: "test-target-a", BlockedBy: []string{"test-main"}},
				{ID: "test-target-b", BlockedBy: []string{"test-main"}},
			},
			want: "Blocks: test-target-a, test-target-b, test-target-z\nBlocked by: test-blocker-a, test-blocker-b, test-blocker-z",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.bean.Title = "Test bean"
			tt.bean.Status = "todo"
			setupShowTest(t, append([]*bean.Bean{tt.bean}, tt.others...)...)
			before, err := tt.bean.Render()
			if err != nil {
				t.Fatal(err)
			}
			if got := formatRelationships(tt.bean); got != tt.want {
				t.Errorf("relationships = %q, want %q", got, tt.want)
			}
			got := runShowTest(t, tt.bean.ID)
			lines := strings.Split(got, "\n")
			for i := range lines {
				lines[i] = strings.TrimRight(lines[i], " ")
			}
			separator := strings.Repeat("─", 50)
			wantHeader := lines[0] + "\n"
			if tt.want != "" {
				wantHeader += tt.want + "\n"
			}
			wantHeader += separator + "\n"
			if tt.bean.Parent != "" {
				wantHeader += "parent: " + tt.bean.Parent + "\n" + separator + "\n"
			}
			wantHeader += "Test bean\n" + separator + "\n"
			if !strings.HasPrefix(strings.Join(lines, "\n"), wantHeader) {
				t.Errorf("show output has incorrect header layout, want prefix %q:\n%s", wantHeader, got)
			}
			for _, label := range []string{"Blocked by:", "Blocks:"} {
				if !strings.Contains(tt.want, label) && strings.Contains(got, label) {
					t.Errorf("unexpected %s in output:\n%s", label, got)
				}
			}
			after, err := tt.bean.Render()
			if err != nil {
				t.Fatal(err)
			}
			if string(before) != string(after) {
				t.Error("show mutated bean data")
			}
		})
	}
}

func TestShowRelationshipsOutputModes(t *testing.T) {
	b := &bean.Bean{
		ID: "test-main", Title: "Test bean", Status: "todo", Body: "Bean body\n",
		BlockedBy: []string{"test-blocker"}, Blocking: []string{"test-target"},
	}
	setupShowTest(t, b)
	tests := []struct {
		name string
		flag *bool
		want func() string
	}{
		{"raw", &showRaw, func() string {
			data, err := b.Render()
			if err != nil {
				t.Fatal(err)
			}
			return string(data)
		}},
		{"body only", &showBodyOnly, func() string { return b.Body }},
		{"etag only", &showETagOnly, b.ETag},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			*tt.flag = true
			defer func() { *tt.flag = false }()
			if got, want := runShowTest(t, b.ID), tt.want(); got != want {
				t.Errorf("show = %q, want %q", got, want)
			}
		})
	}
	t.Run("json", func(t *testing.T) {
		showJSON = true
		defer func() { showJSON = false }()
		var got bean.Bean
		if err := json.Unmarshal([]byte(runShowTest(t, b.ID)), &got); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(got.BlockedBy, b.BlockedBy) || !reflect.DeepEqual(got.Blocking, b.Blocking) {
			t.Errorf("JSON relationships = %v / %v, want %v / %v", got.BlockedBy, got.Blocking, b.BlockedBy, b.Blocking)
		}
	})
}
