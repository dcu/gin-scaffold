package main

import (
	"fmt"
	"github.com/dcu/gin-scaffold/command"
	"os"
)

func printUsageAndExit() {
	fmt.Printf("Gin Scaffold\n")
	for commandName, command := range command.Commands {
		fmt.Printf("\nCommand `%s`:\n\n", commandName)
		command.Help()
	}
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

	command := command.FindCommand(commandName)
	if command == nil {
		printUsageAndExit()
	}

	command.Execute(commandArgs)
}
