package main

import (
	"fmt"
	"github.com/dcu/gin-scaffold/command"
	"os"
)

var (
	Commands = map[string]command.Base{
		"init":       &command.InitCommand{},
		"model":      &command.ModelCommand{},
		"controller": &command.ControllerCommand{},
		"scaffold":   &command.ScaffoldCommand{},
	}
)

func printUsageAndExit() {
	fmt.Printf("Usage: %s <command> <args>\n", os.Args[0])
	os.Exit(0)
}

func main() {
	if len(os.Args) < 2 {
		printUsageAndExit()
	}

	commandName := os.Args[1]
	commandArgs := []string{}

	if len(os.Args) > 2 {
		commandArgs = os.Args[2:]
	}

	command := Commands[commandName]
	if command == nil {
		printUsageAndExit()
	}

	command.Execute(commandArgs)
}
