package main

import (
	"fmt"
	"github.com/dcu/gin-scaffold/command"
	"os"
	"strings"
)

func printUsageAndExit() {
	fmt.Printf("Gin Scaffold\n")
	for commandName, command := range command.Commands {
		fmt.Printf("\nCommand `%s`:\n\n", commandName)
		command.Help()
	}
	os.Exit(0)
}

func checkGOPATH() {
	if os.Getenv("GOPATH") == "" {
		fmt.Printf("$GOPATH is not defined. Exiting.\n")
		os.Exit(2)
	}

	wd, _ := os.Getwd()
	if !strings.HasPrefix(wd, os.Getenv("GOPATH")) {
		fmt.Printf("%s is not in the $GOPATH. Exiting.\n", wd)
		os.Exit(3)
	}
}

func main() {
	checkGOPATH()

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
