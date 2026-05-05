package commands

import (
	"fmt"
	"os"
	"os/exec"
	"unicode"

	"github.com/spf13/cobra"
)

func resolveEditor() string {
	if editor := os.Getenv("EDITOR"); editor != "" {
		return editor
	}
	if visual := os.Getenv("VISUAL"); visual != "" {
		return visual
	}
	return "nano"
}

func splitEditorCommand(editor string) ([]string, error) {
	var parts []string
	var current []rune
	var quote rune
	var escaped bool

	flush := func() {
		if len(current) == 0 {
			return
		}
		parts = append(parts, string(current))
		current = current[:0]
	}

	for _, r := range editor {
		if escaped {
			current = append(current, r)
			escaped = false
			continue
		}

		switch {
		case quote == '\'':
			if r == '\'' {
				quote = 0
				continue
			}
			current = append(current, r)
		case quote == '"':
			switch r {
			case '"':
				quote = 0
			case '\\':
				escaped = true
			default:
				current = append(current, r)
			}
		case unicode.IsSpace(r):
			flush()
		case r == '\'' || r == '"':
			quote = r
		case r == '\\':
			escaped = true
		default:
			current = append(current, r)
		}
	}

	if escaped {
		return nil, fmt.Errorf("unterminated escape sequence")
	}
	if quote != 0 {
		return nil, fmt.Errorf("unterminated quoted string")
	}

	flush()
	return parts, nil
}

func buildEditorCmd(editor string, filePath string) (*exec.Cmd, error) {
	parts, err := splitEditorCommand(editor)
	if err != nil {
		return nil, err
	}
	if len(parts) == 0 {
		parts = []string{"nano"}
	}
	return exec.Command(parts[0], append(parts[1:], filePath)...), nil
}

var editCmd = &cobra.Command{
	Use:   "edit <id>",
	Short: "Open a bean in your editor",
	Long: `Opens the markdown file for the given bean in your text editor.

The editor is resolved in this order:
  1. $EDITOR environment variable
  2. $VISUAL environment variable
  3. nano (fallback)`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		b, err := core.Get(args[0])
		if err != nil {
			return fmt.Errorf("failed to find bean: %w", err)
		}

		editor := resolveEditor()
		filePath := core.FullPath(b)

		editExec, err := buildEditorCmd(editor, filePath)
		if err != nil {
			return fmt.Errorf("invalid editor command: %w", err)
		}
		editExec.Stdin = os.Stdin
		editExec.Stdout = os.Stdout
		editExec.Stderr = os.Stderr

		if err := editExec.Run(); err != nil {
			return fmt.Errorf("editor exited with error: %w", err)
		}

		return nil
	},
}

func RegisterEditCmd(root *cobra.Command) {
	root.AddCommand(editCmd)
}
