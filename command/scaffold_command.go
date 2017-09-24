package command

import (
	"fmt"
	"os"
	"strings"
	"bitbucket.org/pkg/inflect"
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

func remove(strings []string, search string) []string {
	result := []string{}
	for _, str := range strings {
		if str != search {
			result = append(result, str)
		}
	}
	return result
}

func (command *ScaffoldCommand) Execute(args []string) {
	// --option
	optionStrings := []string{}
	removeStrings := []string{}

	if len(args) == 0 {
		command.Help()
		os.Exit(2)
	} else if len(args) > 2 {
		// check args
		for _, value := range args {
			res := strings.HasPrefix(value, "--")

			if res {
				optionStrings = append(optionStrings, value)
				removeStrings = append(removeStrings, value)
			}
		}
	}

	// remove from argument
	for _, value := range removeStrings {
		args = remove(args, value)
	}

	args[0] = inflect.Singularize(args[0])
	modelCommand := &ModelCommand{}
	modelCommand.Execute(args)

	controllerCommand := &ControllerCommand{}
	// controllerCommand.Execute([]string{modelCommand.ModelNamePlural})
	controllerCommand.ExecuteWithOption([]string{modelCommand.ModelNamePlural}, optionStrings)
}
