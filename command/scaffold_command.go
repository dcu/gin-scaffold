package command

import (
	"fmt"
)

type ScaffoldCommand struct {
}

func (command *ScaffoldCommand) Help() {
	fmt.Printf(`Usage:
	gin-scaffold scaffold <controller name> <field name>:<field type> ...

Description:
	The gin-scaffold scaffold command creates a new controller and model with the given fields.

Example:
	gin-scaffold controller Post Title:string Body:string
`)
}

func (command *ScaffoldCommand) Execute(args []string) {
	modelCommand := &ModelCommand{}
	modelCommand.Execute(args)

	controllerCommand := &ControllerCommand{}
	controllerCommand.Execute([]string{modelCommand.ModelNamePlural})
}
