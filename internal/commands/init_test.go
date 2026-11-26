package commands

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/50lo/classic-beans/pkg/config"
	"gopkg.in/yaml.v3"
)

func TestInitWritesCLIConfig(t *testing.T) {
	dir := t.TempDir()
	t.Chdir(dir)
	t.Setenv("BEANS_PATH", "")

	root := NewRootCmd()
	RegisterInitCmd(root)
	root.SetArgs([]string{"init"})
	if err := root.Execute(); err != nil {
		t.Fatal(err)
	}

	data, err := os.ReadFile(config.ConfigFileName)
	if err != nil {
		t.Fatal(err)
	}
	var sections map[string]yaml.Node
	if err := yaml.Unmarshal(data, &sections); err != nil {
		t.Fatal(err)
	}
	if len(sections) != 2 || sections["project"].Kind != yaml.MappingNode || sections["beans"].Kind != yaml.MappingNode {
		t.Fatalf("expected only project and beans settings, got:\n%s", data)
	}

	cfg, err := config.Load(config.ConfigFileName)
	if err != nil {
		t.Fatal(err)
	}
	name := filepath.Base(dir)
	if cfg.Project.Name != name || cfg.Beans != config.DefaultWithPrefix(name+"-").Beans {
		t.Fatalf("unexpected project or beans defaults:\n%s", data)
	}
}
