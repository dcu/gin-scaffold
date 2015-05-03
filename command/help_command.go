package command

import (
	"fmt"
)

type HelpCommand struct {
}

func (command *HelpCommand) Help() {
	fmt.Printf(`Usage:
	gin-scaffold help <command name>

Description:
	The gin-scaffold help command prints help about the given command.

Example:
	gin-scaffold help init
`)
}

func (command *HelpCommand) Execute(args []string) {
	if len(args) == 0 {
		command.Help()
		return
	}

	targetCommand := FindCommand(args[0])

	if targetCommand == nil {
		command.Help()
	} else {
		targetCommand.Help()
	}
}
