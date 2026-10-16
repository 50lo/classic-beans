package commands

import (
	"fmt"
	"slices"
	"strings"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/hmans/beans/internal/output"
	"github.com/hmans/beans/internal/ui"
	"github.com/hmans/beans/pkg/bean"
	"github.com/spf13/cobra"
)

var (
	showJSON     bool
	showRaw      bool
	showBodyOnly bool
	showETagOnly bool
)

var showCmd = &cobra.Command{
	Use:   "show <id> [id...]",
	Short: "Show a bean's contents",
	Long:  `Displays the full contents of one or more beans, including front matter and body.`,
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var beans []*bean.Bean
		for _, id := range args {
			b, err := core.Get(id)
			if err != nil {
				if showJSON {
					return output.Error(output.ErrNotFound, err.Error())
				}
				return fmt.Errorf("failed to find bean: %w", err)
			}
			beans = append(beans, b)
		}

		if showJSON {
			if len(beans) == 1 {
				return output.SuccessSingle(beans[0])
			}
			return output.SuccessMultiple(beans)
		}

		if showRaw {
			for i, b := range beans {
				if i > 0 {
					fmt.Print("\n---\n\n")
				}
				content, err := b.Render()
				if err != nil {
					return fmt.Errorf("failed to render bean: %w", err)
				}
				fmt.Print(string(content))
			}
			return nil
		}

		if showBodyOnly {
			for i, b := range beans {
				if i > 0 {
					fmt.Print("\n---\n\n")
				}
				fmt.Print(b.Body)
			}
			return nil
		}

		if showETagOnly {
			for i, b := range beans {
				if i > 0 {
					fmt.Println()
				}
				fmt.Print(b.ETag())
			}
			return nil
		}

		for i, b := range beans {
			if i > 0 {
				fmt.Println()
				fmt.Println(ui.Muted.Render(strings.Repeat("═", 60)))
				fmt.Println()
			}
			showStyledBean(b)
		}

		return nil
	},
}

func showStyledBean(b *bean.Bean) {
	statusCfg := cfg.GetStatus(b.Status)
	statusColor := "gray"
	if statusCfg != nil {
		statusColor = statusCfg.Color
	}
	isArchive := cfg.IsArchiveStatus(b.Status)

	var header strings.Builder
	header.WriteString(ui.ID.Render(b.ID))
	header.WriteString(" ")
	header.WriteString(ui.RenderStatusWithColor(b.Status, statusColor, isArchive))
	if implicitStatus, implicitStatusFrom := core.ImplicitStatus(b.ID); implicitStatus != "" {
		header.WriteString(" ")
		header.WriteString(ui.Muted.Render("↑" + implicitStatus + " (from " + implicitStatusFrom + ")"))
	}

	if b.Type != "" {
		typeCfg := cfg.GetType(b.Type)
		typeColor := "gray"
		if typeCfg != nil {
			typeColor = typeCfg.Color
		}
		header.WriteString(" ")
		header.WriteString(ui.RenderTypeWithColor(b.Type, typeColor))
	}

	if b.Priority != "" {
		priorityCfg := cfg.GetPriority(b.Priority)
		priorityColor := "gray"
		if priorityCfg != nil {
			priorityColor = priorityCfg.Color
		}
		header.WriteString(" ")
		header.WriteString(ui.RenderPriorityWithColor(b.Priority, priorityColor))
	}
	if len(b.Tags) > 0 {
		header.WriteString("  ")
		header.WriteString(ui.Muted.Render(strings.Join(b.Tags, ", ")))
	}
	if b.CreatedAt != nil {
		header.WriteString("  ")
		header.WriteString(ui.Muted.Render("created " + b.CreatedAt.Format("2006-01-02 15:04 UTC")))
	}
	if b.UpdatedAt != nil {
		header.WriteString("  ")
		header.WriteString(ui.Muted.Render("updated " + b.UpdatedAt.Format("2006-01-02 15:04 UTC")))
	}
	if relationships := formatRelationships(b); relationships != "" {
		header.WriteString("\n")
		header.WriteString(relationships)
	}

	header.WriteString("\n")
	header.WriteString(ui.Muted.Render(strings.Repeat("─", 50)))
	header.WriteString("\n")
	header.WriteString(ui.Title.Render(b.Title))

	if b.Parent != "" {
		header.WriteString("\n")
		header.WriteString(ui.Muted.Render(strings.Repeat("─", 50)))
		header.WriteString("\n")
		header.WriteString(ui.Muted.Render("parent:") + " " + ui.ID.Render(b.Parent))
	}

	header.WriteString("\n")
	header.WriteString(ui.Muted.Render(strings.Repeat("─", 50)))

	headerBox := lipgloss.NewStyle().
		MarginBottom(1).
		Render(header.String())

	fmt.Println(headerBox)

	if b.Body != "" {
		renderer, err := glamour.NewTermRenderer(
			glamour.WithAutoStyle(),
			glamour.WithWordWrap(80),
		)
		if err != nil {
			fmt.Printf("failed to create renderer: %v\n", err)
			return
		}

		rendered, err := renderer.Render(b.Body)
		if err != nil {
			fmt.Printf("failed to render markdown: %v\n", err)
			return
		}

		fmt.Print(rendered)
	}
}

func formatRelationships(b *bean.Bean) string {
	var parts []string

	blockedBy := slices.Clone(b.BlockedBy)
	blocks := slices.Clone(b.Blocking)
	for _, link := range core.FindIncomingLinks(b.ID) {
		switch link.LinkType {
		case "blocking":
			blockedBy = append(blockedBy, link.FromBean.ID)
		case "blocked_by":
			blocks = append(blocks, link.FromBean.ID)
		}
	}

	appendList := func(label string, ids []string) {
		if len(ids) == 0 {
			return
		}
		slices.Sort(ids)
		ids = slices.Compact(ids)
		for i, id := range ids {
			ids[i] = ui.ID.Render(id)
		}
		parts = append(parts, ui.Muted.Render(label)+" "+strings.Join(ids, ", "))
	}
	appendList("Blocks:", blocks)
	appendList("Blocked by:", blockedBy)
	return strings.Join(parts, "\n")
}

func RegisterShowCmd(root *cobra.Command) {
	showCmd.Flags().BoolVar(&showJSON, "json", false, "Output as JSON")
	showCmd.Flags().BoolVar(&showRaw, "raw", false, "Output raw markdown without styling")
	showCmd.Flags().BoolVar(&showBodyOnly, "body-only", false, "Output only the body content")
	showCmd.Flags().BoolVar(&showETagOnly, "etag-only", false, "Output only the etag")
	showCmd.MarkFlagsMutuallyExclusive("json", "raw", "body-only", "etag-only")
	root.AddCommand(showCmd)
}
