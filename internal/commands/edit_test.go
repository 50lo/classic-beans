package commands

import (
	"testing"
)

func TestResolveEditor(t *testing.T) {
	t.Run("EDITOR takes precedence", func(t *testing.T) {
		t.Setenv("EDITOR", "vim")
		t.Setenv("VISUAL", "emacs")
		got := resolveEditor()
		if got != "vim" {
			t.Errorf("resolveEditor() = %q, want %q", got, "vim")
		}
	})

	t.Run("VISUAL used when EDITOR is empty", func(t *testing.T) {
		t.Setenv("EDITOR", "")
		t.Setenv("VISUAL", "code")
		got := resolveEditor()
		if got != "code" {
			t.Errorf("resolveEditor() = %q, want %q", got, "code")
		}
	})

	t.Run("nano used when both are empty", func(t *testing.T) {
		t.Setenv("EDITOR", "")
		t.Setenv("VISUAL", "")
		got := resolveEditor()
		if got != "nano" {
			t.Errorf("resolveEditor() = %q, want %q", got, "nano")
		}
	})

	t.Run("EDITOR with path", func(t *testing.T) {
		t.Setenv("EDITOR", "/usr/local/bin/nvim")
		t.Setenv("VISUAL", "")
		got := resolveEditor()
		if got != "/usr/local/bin/nvim" {
			t.Errorf("resolveEditor() = %q, want %q", got, "/usr/local/bin/nvim")
		}
	})
}

func TestBuildEditorCmd(t *testing.T) {
	t.Run("simple editor", func(t *testing.T) {
		cmd, err := buildEditorCmd("vim", "/path/to/file.md")
		if err != nil {
			t.Fatalf("buildEditorCmd() error = %v", err)
		}
		if cmd.Path == "" {
			t.Error("expected command path to be set")
		}
		got := append([]string{cmd.Args[0]}, cmd.Args[1:]...)
		want := []string{"vim", "/path/to/file.md"}
		if len(got) != len(want) {
			t.Errorf("buildEditorCmd args = %v, want %v", got, want)
		}
		for i := range want {
			if got[i] != want[i] {
				t.Errorf("buildEditorCmd args[%d] = %q, want %q", i, got[i], want[i])
			}
		}
	})

	t.Run("editor with flags", func(t *testing.T) {
		cmd, err := buildEditorCmd("code --wait", "/path/to/file.md")
		if err != nil {
			t.Fatalf("buildEditorCmd() error = %v", err)
		}
		got := cmd.Args
		want := []string{"code", "--wait", "/path/to/file.md"}
		if len(got) != len(want) {
			t.Errorf("buildEditorCmd args = %v, want %v", got, want)
		}
		for i := range want {
			if got[i] != want[i] {
				t.Errorf("buildEditorCmd args[%d] = %q, want %q", i, got[i], want[i])
			}
		}
	})

	t.Run("editor with multiple flags", func(t *testing.T) {
		cmd, err := buildEditorCmd("nvim -f -c startinsert", "/path/to/file.md")
		if err != nil {
			t.Fatalf("buildEditorCmd() error = %v", err)
		}
		got := cmd.Args
		want := []string{"nvim", "-f", "-c", "startinsert", "/path/to/file.md"}
		if len(got) != len(want) {
			t.Errorf("buildEditorCmd args = %v, want %v", got, want)
		}
		for i := range want {
			if got[i] != want[i] {
				t.Errorf("buildEditorCmd args[%d] = %q, want %q", i, got[i], want[i])
			}
		}
	})

	t.Run("editor with quoted argument", func(t *testing.T) {
		cmd, err := buildEditorCmd(`nvim -c "set ft=markdown"`, "/path/to/file.md")
		if err != nil {
			t.Fatalf("buildEditorCmd() error = %v", err)
		}
		got := cmd.Args
		want := []string{"nvim", "-c", "set ft=markdown", "/path/to/file.md"}
		if len(got) != len(want) {
			t.Errorf("buildEditorCmd args = %v, want %v", got, want)
		}
		for i := range want {
			if got[i] != want[i] {
				t.Errorf("buildEditorCmd args[%d] = %q, want %q", i, got[i], want[i])
			}
		}
	})

	t.Run("editor path can be quoted", func(t *testing.T) {
		cmd, err := buildEditorCmd(`"/Applications/Visual Studio Code.app/Contents/Resources/app/bin/code" --wait`, "/path/to/file.md")
		if err != nil {
			t.Fatalf("buildEditorCmd() error = %v", err)
		}
		got := cmd.Args
		want := []string{"/Applications/Visual Studio Code.app/Contents/Resources/app/bin/code", "--wait", "/path/to/file.md"}
		if len(got) != len(want) {
			t.Errorf("buildEditorCmd args = %v, want %v", got, want)
		}
		for i := range want {
			if got[i] != want[i] {
				t.Errorf("buildEditorCmd args[%d] = %q, want %q", i, got[i], want[i])
			}
		}
	})

	t.Run("empty string falls back to nano", func(t *testing.T) {
		cmd, err := buildEditorCmd("", "/path/to/file.md")
		if err != nil {
			t.Fatalf("buildEditorCmd() error = %v", err)
		}
		got := cmd.Args
		want := []string{"nano", "/path/to/file.md"}
		if len(got) != len(want) {
			t.Errorf("buildEditorCmd args = %v, want %v", got, want)
		}
		for i := range want {
			if got[i] != want[i] {
				t.Errorf("buildEditorCmd args[%d] = %q, want %q", i, got[i], want[i])
			}
		}
	})

	t.Run("unterminated quote returns error", func(t *testing.T) {
		if _, err := buildEditorCmd(`nvim -c "set ft=markdown`, "/path/to/file.md"); err == nil {
			t.Fatal("expected error for unterminated quote, got nil")
		}
	})
}

func TestEditCmdRequiresExactlyOneArg(t *testing.T) {
	editCmd := editCmd
	if err := editCmd.Args(editCmd, nil); err == nil {
		t.Error("expected error for zero arguments, got nil")
	}
	if err := editCmd.Args(editCmd, []string{"a", "b"}); err == nil {
		t.Error("expected error for two arguments, got nil")
	}
	if err := editCmd.Args(editCmd, []string{"abc123"}); err != nil {
		t.Errorf("expected no error for one argument, got %v", err)
	}
}
