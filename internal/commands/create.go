package commands

import (
	"fmt"
	"strings"

	"github.com/hmans/beans/internal/output"
	"github.com/hmans/beans/internal/ui"
	"github.com/hmans/beans/pkg/bean"
	"github.com/hmans/beans/pkg/config"
	"github.com/spf13/cobra"
)

var (
	createStatus    string
	createType      string
	createPriority  string
	createBody      string
	createBodyFile  string
	createTag       []string
	createParent    string
	createBlocking  []string
	createBlockedBy []string
	createPrefix    string
	createJSON      bool
)

var createCmd = &cobra.Command{
	Use:     "create [title]",
	Aliases: []string{"c", "new"},
	Short:   "Create a new bean",
	Long:    `Creates a new bean (issue) with a generated ID and optional title.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		title := strings.Join(args, " ")
		if title == "" {
			title = "Untitled"
		}

		if createStatus != "" && !cfg.IsValidStatus(createStatus) {
			return cmdError(createJSON, output.ErrInvalidStatus, "invalid status: %s (must be %s)", createStatus, cfg.StatusList())
		}
		if createType != "" && !cfg.IsValidType(createType) {
			return cmdError(createJSON, output.ErrValidation, "invalid type: %s (must be %s)", createType, cfg.TypeList())
		}
		if createPriority != "" && !cfg.IsValidPriority(createPriority) {
			return cmdError(createJSON, output.ErrValidation, "invalid priority: %s (must be %s)", createPriority, cfg.PriorityList())
		}

		body, err := resolveContent(createBody, createBodyFile)
		if err != nil {
			return cmdError(createJSON, output.ErrFileError, "%s", err)
		}

		status := cfg.GetDefaultStatus()
		if createStatus != "" {
			status = createStatus
		}

		beanType := cfg.GetDefaultType()
		if createType != "" {
			beanType = createType
		}

		b := &bean.Bean{
			Title:     title,
			Slug:      bean.Slugify(title),
			Status:    status,
			Type:      beanType,
			Priority:  createPriority,
			Body:      body,
			Tags:      createTag,
			Parent:    createParent,
			Blocking:  createBlocking,
			BlockedBy: createBlockedBy,
		}

		if createPrefix != "" {
			savedPrefix := cfg.Beans.Prefix
			cfg.Beans.Prefix = createPrefix
			if err := core.Create(b); err != nil {
				cfg.Beans.Prefix = savedPrefix
				return cmdError(createJSON, output.ErrFileError, "failed to create bean: %v", err)
			}
			cfg.Beans.Prefix = savedPrefix
		} else {
			if err := core.Create(b); err != nil {
				return cmdError(createJSON, output.ErrFileError, "failed to create bean: %v", err)
			}
		}

		if createJSON {
			return output.Success(b, "Bean created")
		}

		fmt.Println(ui.Success.Render("Created ") + ui.ID.Render(b.ID) + " " + ui.Muted.Render(b.Path))
		return nil
	},
}

func RegisterCreateCmd(root *cobra.Command) {
	statusNames := make([]string, len(config.DefaultStatuses))
	for i, s := range config.DefaultStatuses {
		statusNames[i] = s.Name
	}
	typeNames := make([]string, len(config.DefaultTypes))
	for i, t := range config.DefaultTypes {
		typeNames[i] = t.Name
	}
	priorityNames := make([]string, len(config.DefaultPriorities))
	for i, p := range config.DefaultPriorities {
		priorityNames[i] = p.Name
	}

	createCmd.Flags().StringVarP(&createStatus, "status", "s", "", "Initial status ("+strings.Join(statusNames, ", ")+")")
	createCmd.Flags().StringVarP(&createType, "type", "t", "", "Bean type ("+strings.Join(typeNames, ", ")+")")
	createCmd.Flags().StringVarP(&createPriority, "priority", "p", "", "Priority level ("+strings.Join(priorityNames, ", ")+")")
	createCmd.Flags().StringVarP(&createBody, "body", "d", "", "Body content (use '-' to read from stdin)")
	createCmd.Flags().StringVar(&createBodyFile, "body-file", "", "Read body from file")
	createCmd.Flags().StringArrayVar(&createTag, "tag", nil, "Add tag (can be repeated)")
	createCmd.Flags().StringVar(&createParent, "parent", "", "Parent bean ID")
	createCmd.Flags().StringArrayVar(&createBlocking, "blocking", nil, "ID of bean this blocks (can be repeated)")
	createCmd.Flags().StringArrayVar(&createBlockedBy, "blocked-by", nil, "ID of bean that blocks this one (can be repeated)")
	createCmd.Flags().StringVar(&createPrefix, "prefix", "", "Custom ID prefix (overrides config prefix)")
	createCmd.Flags().BoolVar(&createJSON, "json", false, "Output as JSON")
	createCmd.MarkFlagsMutuallyExclusive("body", "body-file")
	root.AddCommand(createCmd)
}
