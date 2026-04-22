package commands

import (
	"github.com/spf13/cobra"
)

func RegisterCoreCommands(root *cobra.Command) {
	RegisterArchiveCmd(root)
	RegisterCheckCmd(root)
	RegisterCreateCmd(root)
	RegisterDeleteCmd(root)
	RegisterInitCmd(root)
	RegisterListCmd(root)
	RegisterPrimeCmd(root)
	RegisterRoadmapCmd(root)
	RegisterShowCmd(root)
	RegisterUpdateCmd(root)
	RegisterVersionCmd(root)
}
