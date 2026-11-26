// Modified for Classic Beans.

package main

import "github.com/50lo/classic-beans/internal/commands"

func main() {
	root := commands.NewRootCmd()
	commands.RegisterCoreCommands(root)
	commands.Execute(root)
}
