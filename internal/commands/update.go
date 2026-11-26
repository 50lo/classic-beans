// Modified for Classic Beans.

package commands

import (
	"errors"
	"fmt"
	"strings"

	"github.com/50lo/classic-beans/internal/output"
	"github.com/50lo/classic-beans/internal/ui"
	"github.com/50lo/classic-beans/pkg/bean"
	"github.com/50lo/classic-beans/pkg/beancore"
	"github.com/50lo/classic-beans/pkg/config"
	"github.com/spf13/cobra"
)

var (
	updateStatus          string
	updateType            string
	updatePriority        string
	updateTitle           string
	updateBody            string
	updateBodyFile        string
	updateBodyReplaceOld  string
	updateBodyReplaceNew  string
	updateBodyAppend      string
	updateParent          string
	updateRemoveParent    bool
	updateBlocking        []string
	updateRemoveBlocking  []string
	updateBlockedBy       []string
	updateRemoveBlockedBy []string
	updateTag             []string
	updateRemoveTag       []string
	updateIfMatch         string
	updateJSON            bool
)

var updateCmd = &cobra.Command{
	Use:     "update <id>",
	Aliases: []string{"u"},
	Short:   "Update a bean's properties",
	Long:    `Updates one or more properties of an existing bean.`,
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		b, err := core.Get(args[0])
		if err != nil {
			return cmdError(updateJSON, output.ErrNotFound, "failed to find bean: %v", err)
		}

		wasArchived := false
		if b == nil {
			unarchived, unarchiveErr := core.LoadAndUnarchive(args[0])
			if unarchiveErr != nil {
				return cmdError(updateJSON, output.ErrNotFound, "bean not found: %s", args[0])
			}
			b = unarchived
			wasArchived = true
		}

		var changes []string

		var ifMatch *string
		if updateIfMatch != "" {
			ifMatch = &updateIfMatch
		}

		changed, fieldChanges, err := applyFieldUpdates(cmd, b)
		if err != nil {
			return cmdError(updateJSON, output.ErrValidation, "%s", err)
		}
		changes = append(changes, fieldChanges...)

		if changed {
			if err := core.Update(b, ifMatch); err != nil {
				return mutationError(updateJSON, err)
			}
		}

		if len(changes) == 0 {
			return cmdError(updateJSON, output.ErrValidation,
				"no changes specified (use --status, --type, --priority, --title, --body, --parent, --blocking, --blocked-by, --tag, or their --remove-* variants)")
		}

		if updateJSON {
			msg := "Bean updated"
			if wasArchived {
				msg = "Bean unarchived and updated"
			}
			return output.Success(b, msg)
		}

		if wasArchived {
			fmt.Println(ui.Success.Render("Unarchived and updated ") + ui.ID.Render(b.ID) + " " + ui.Muted.Render(b.Path))
		} else {
			fmt.Println(ui.Success.Render("Updated ") + ui.ID.Render(b.ID) + " " + ui.Muted.Render(b.Path))
		}
		return nil
	},
}

func applyFieldUpdates(cmd *cobra.Command, b *bean.Bean) (bool, []string, error) {
	var changes []string
	changed := false

	if cmd.Flags().Changed("status") {
		if !cfg.IsValidStatus(updateStatus) {
			return false, nil, fmt.Errorf("invalid status: %s (must be %s)", updateStatus, cfg.StatusList())
		}
		b.Status = updateStatus
		changes = append(changes, "status")
		changed = true
	}

	if cmd.Flags().Changed("type") {
		if !cfg.IsValidType(updateType) {
			return false, nil, fmt.Errorf("invalid type: %s (must be %s)", updateType, cfg.TypeList())
		}
		b.Type = updateType
		changes = append(changes, "type")
		changed = true
	}

	if cmd.Flags().Changed("priority") {
		if !cfg.IsValidPriority(updatePriority) {
			return false, nil, fmt.Errorf("invalid priority: %s (must be %s)", updatePriority, cfg.PriorityList())
		}
		b.Priority = updatePriority
		changes = append(changes, "priority")
		changed = true
	}

	if cmd.Flags().Changed("title") {
		b.Title = updateTitle
		changes = append(changes, "title")
		changed = true
	}

	if cmd.Flags().Changed("body") || cmd.Flags().Changed("body-file") {
		body, err := resolveContent(updateBody, updateBodyFile)
		if err != nil {
			return false, nil, err
		}
		b.Body = body
		changes = append(changes, "body")
		changed = true
	} else if cmd.Flags().Changed("body-replace-old") || cmd.Flags().Changed("body-append") {
		if cmd.Flags().Changed("body-replace-old") {
			newBody, err := applyBodyReplace(b.Body, bean.UnescapeBody(updateBodyReplaceOld), bean.UnescapeBody(updateBodyReplaceNew))
			if err != nil {
				return false, nil, err
			}
			b.Body = newBody
		}

		if cmd.Flags().Changed("body-append") {
			appendText, err := resolveAppendContent(updateBodyAppend)
			if err != nil {
				return false, nil, err
			}
			b.Body = applyBodyAppend(b.Body, appendText)
		}
		changes = append(changes, "body")
		changed = true
	}

	if len(updateTag) > 0 {
		if err := applyTags(b, updateTag); err != nil {
			return false, nil, err
		}
		changes = append(changes, "tags")
		changed = true
	}
	if len(updateRemoveTag) > 0 {
		for _, t := range updateRemoveTag {
			b.RemoveTag(t)
		}
		changes = append(changes, "tags")
		changed = true
	}

	if cmd.Flags().Changed("parent") {
		b.Parent = updateParent
		changes = append(changes, "parent")
		changed = true
	} else if updateRemoveParent {
		b.Parent = ""
		changes = append(changes, "parent")
		changed = true
	}

	if len(updateBlocking) > 0 {
		b.Blocking = append(b.Blocking, updateBlocking...)
		changes = append(changes, "blocking")
		changed = true
	}
	if len(updateRemoveBlocking) > 0 {
		for _, id := range updateRemoveBlocking {
			b.RemoveBlocking(id)
		}
		changes = append(changes, "blocking")
		changed = true
	}

	if len(updateBlockedBy) > 0 {
		b.BlockedBy = append(b.BlockedBy, updateBlockedBy...)
		changes = append(changes, "blocked-by")
		changed = true
	}
	if len(updateRemoveBlockedBy) > 0 {
		for _, id := range updateRemoveBlockedBy {
			b.RemoveBlockedBy(id)
		}
		changes = append(changes, "blocked-by")
		changed = true
	}

	return changed, changes, nil
}

func isConflictError(err error) bool {
	var mismatchErr *beancore.ETagMismatchError
	var requiredErr *beancore.ETagRequiredError
	return errors.As(err, &mismatchErr) || errors.As(err, &requiredErr)
}

func mutationError(jsonOutput bool, err error) error {
	if isConflictError(err) {
		return cmdError(jsonOutput, output.ErrConflict, "%s", err)
	}
	return cmdError(jsonOutput, output.ErrValidation, "%s", err)
}

func RegisterUpdateCmd(root *cobra.Command) {
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

	updateCmd.Flags().StringVarP(&updateStatus, "status", "s", "", "New status ("+strings.Join(statusNames, ", ")+")")
	updateCmd.Flags().StringVarP(&updateType, "type", "t", "", "New type ("+strings.Join(typeNames, ", ")+")")
	updateCmd.Flags().StringVarP(&updatePriority, "priority", "p", "", "New priority ("+strings.Join(priorityNames, ", ")+", or empty to clear)")
	updateCmd.Flags().StringVar(&updateTitle, "title", "", "New title")
	updateCmd.Flags().StringVarP(&updateBody, "body", "d", "", "New body (use '-' to read from stdin)")
	updateCmd.Flags().StringVar(&updateBodyFile, "body-file", "", "Read body from file")
	updateCmd.Flags().StringVar(&updateBodyReplaceOld, "body-replace-old", "", "Text to find and replace (requires --body-replace-new)")
	updateCmd.Flags().StringVar(&updateBodyReplaceNew, "body-replace-new", "", "Replacement text (requires --body-replace-old)")
	updateCmd.Flags().StringVar(&updateBodyAppend, "body-append", "", "Text to append to body (use '-' for stdin)")
	updateCmd.Flags().StringVar(&updateParent, "parent", "", "Set parent bean ID")
	updateCmd.Flags().BoolVar(&updateRemoveParent, "remove-parent", false, "Remove parent")
	updateCmd.Flags().StringArrayVar(&updateBlocking, "blocking", nil, "ID of bean this blocks (can be repeated)")
	updateCmd.Flags().StringArrayVar(&updateRemoveBlocking, "remove-blocking", nil, "ID of bean to unblock (can be repeated)")
	updateCmd.Flags().StringArrayVar(&updateBlockedBy, "blocked-by", nil, "ID of bean that blocks this one (can be repeated)")
	updateCmd.Flags().StringArrayVar(&updateRemoveBlockedBy, "remove-blocked-by", nil, "ID of blocker bean to remove (can be repeated)")
	updateCmd.Flags().StringArrayVar(&updateTag, "tag", nil, "Add tag (can be repeated)")
	updateCmd.Flags().StringArrayVar(&updateRemoveTag, "remove-tag", nil, "Remove tag (can be repeated)")
	updateCmd.Flags().StringVar(&updateIfMatch, "if-match", "", "Only update if etag matches (optimistic locking)")
	updateCmd.MarkFlagsMutuallyExclusive("parent", "remove-parent")
	updateCmd.Flags().BoolVar(&updateJSON, "json", false, "Output as JSON")
	updateCmd.MarkFlagsMutuallyExclusive("body", "body-file", "body-replace-old")
	updateCmd.MarkFlagsMutuallyExclusive("body", "body-file", "body-append")
	updateCmd.MarkFlagsRequiredTogether("body-replace-old", "body-replace-new")
	root.AddCommand(updateCmd)
}
