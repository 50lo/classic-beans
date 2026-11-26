// Modified for Classic Beans.

package commands

import (
	_ "embed"
	"os"
	"text/template"

	"github.com/50lo/classic-beans/pkg/config"
	"github.com/spf13/cobra"
)

//go:embed prompt.tmpl
var agentPromptTemplate string

type promptData struct {
	Types      []config.TypeConfig
	Statuses   []config.StatusConfig
	Priorities []config.PriorityConfig
}

var primeCmd = &cobra.Command{
	Use:   "prime",
	Short: "Output instructions for AI coding agents",
	Long:  `Outputs instructions for coding agents on how to manage project issues with cbeans.`,
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		if beansPath == "" && configPath == "" {
			cwd, err := os.Getwd()
			if err != nil {
				return nil
			}
			configFile, err := config.FindConfig(cwd)
			if err != nil || configFile == "" {
				return nil
			}
		}

		tmpl, err := template.New("prompt").Parse(agentPromptTemplate)
		if err != nil {
			return err
		}

		data := promptData{
			Types:      config.DefaultTypes,
			Statuses:   config.DefaultStatuses,
			Priorities: config.DefaultPriorities,
		}

		return tmpl.Execute(cmd.OutOrStdout(), data)
	},
}

func RegisterPrimeCmd(root *cobra.Command) {
	root.AddCommand(primeCmd)
}
