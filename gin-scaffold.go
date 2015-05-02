package main

import (
	"fmt"
	"github.com/dcu/gin-scaffold/command"
	"os"
)

var (
	Commands = map[string]command.Base{
		"init":     &command.InitCommand{},
		"model":    &command.ModelCommand{},
		"scaffold": &command.ScaffoldCommand{},
	}
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <command> <args>\n", os.Args[0])
		return
	}

	commandName := os.Args[1]
	commandArgs := []string{}

	if len(os.Args) > 2 {
		commandArgs = os.Args[2:]
	}

	command := Commands[commandName]
	command.Execute(commandArgs)
}
