// Modified for Classic Beans.

package ui

import (
	"regexp"
	"strings"
	"testing"

	"github.com/50lo/classic-beans/pkg/bean"
	"github.com/50lo/classic-beans/pkg/config"
	"github.com/charmbracelet/lipgloss"
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

func TestRenderTreeNestedHeaderAlignment(t *testing.T) {
	output := stripANSI(RenderTree(taggedTreeFixture(), config.Default(), 12, true, 180))
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

	rootRow := findRenderedRow(t, lines[2:], "beans-a1")
	taskRow := findRenderedRow(t, lines[2:], "beans-a2")
	bugRow := findRenderedRow(t, lines[2:], "beans-a3")

	if got := displayIndex(rootRow, "Import issues from Markdown"); got != titleCol {
		t.Fatalf("root title misaligned: title index=%d want %d\nrow=%q", got, titleCol, rootRow)
	}
	if got := displayIndex(taskRow, " design "); got != tagsCol {
		t.Fatalf("task tag misaligned: tag index=%d want %d\nrow=%q", got, tagsCol, taskRow)
	}
	if got := displayIndex(bugRow, " review "); got != tagsCol {
		t.Fatalf("bug tag misaligned: tag index=%d want %d\nrow=%q", got, tagsCol, bugRow)
	}
}

func TestRenderTreeTaggedRowsStayWithinPaneWidth(t *testing.T) {
	const paneWidth = 145
	output := stripANSI(RenderTree(taggedTreeFixture(), config.Default(), 12, true, paneWidth))
	lines := strings.Split(strings.TrimRight(output, "\n"), "\n")
	if len(lines) < 5 {
		t.Fatalf("expected header and three rendered rows, got %d lines", len(lines))
	}

	for _, line := range lines {
		if width := lipgloss.Width(line); width > paneWidth-1 {
			t.Fatalf("line exceeds safe pane width: got %d want <= %d\nline=%q", width, paneWidth-1, line)
		}
	}
}

func taggedTreeFixture() []*TreeNode {
	parent := &bean.Bean{ID: "beans-a1", Title: "Import issues from Markdown", Type: "epic", Status: "in-progress"}
	task := &bean.Bean{
		ID: "beans-a2", Type: "task", Status: "todo", Parent: parent.ID, Tags: []string{"design"},
		Title: "Preserve old issue IDs, parent links, metadata, and all discussion when importing Markdown files from an archived project",
	}
	bug := &bean.Bean{
		ID: "beans-a3", Title: "Skip unreadable files", Type: "bug", Status: "todo", Parent: parent.ID, Tags: []string{"review"},
	}
	return BuildTree([]*bean.Bean{task, bug}, []*bean.Bean{parent, task, bug}, func([]*bean.Bean) {}, nil)
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
