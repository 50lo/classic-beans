// Modified for Classic Beans.

package commands

import (
	"fmt"

	"github.com/50lo/classic-beans/internal/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintln(cmd.OutOrStdout(), version.String())
	},
}

func RegisterVersionCmd(root *cobra.Command) {
	root.AddCommand(versionCmd)
}
