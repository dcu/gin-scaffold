package command

import (
	"fmt"
	"os"
	"path/filepath"

	"bitbucket.org/pkg/inflect"
	"github.com/dcu/gin-scaffold/template"
)

type ControllerCommand struct {
	PackageName        string
	ControllerName     string
	ModelName          string
	ModelNamePlural    string
	InstanceName       string
	InstanceNamePlural string
	RoutePath          string
	TemplateName       string
	Fields             map[string]string
}

func (command *ControllerCommand) Help() {
	fmt.Printf(`Usage:
	gin-scaffold controller <controller name>

Description:
	The gin-scaffold controller command creates a new controller.

Example:
	gin-scaffold controller Posts
`)
}

func (command *ControllerCommand) Execute(args []string) {
	if len(args) == 0 {
		command.Help()
		os.Exit(2)
	}
	command.ControllerName = args[0]
	command.RoutePath = inflect.Underscore(command.ControllerName)
	command.ModelName = inflect.Singularize(command.ControllerName)
	command.ModelNamePlural = inflect.Pluralize(command.ModelName)

	command.InstanceName = inflect.CamelizeDownFirst(command.ModelName)
	command.InstanceNamePlural = inflect.Pluralize(command.InstanceName)
	command.PackageName = template.PackageName()

	outputPath := filepath.Join("controllers", inflect.Underscore(command.ControllerName)+".go")
	builder := template.NewBuilder("controller.go.tmpl")
	builder.WriteToPath(outputPath, command)

	outputPath = filepath.Join("controllers", inflect.Underscore(command.ControllerName)+"_helpers.go")
	builder = template.NewBuilder("controller_helpers.go.tmpl")
	builder.WriteToPath(outputPath, command)

	outputPath = filepath.Join("controllers", "suite_test.go")
	builder = template.NewBuilder("suite_test.go.tmpl")
	builder.WriteToPath(outputPath, command)

	// templates folder
	path := filepath.Join("templates", inflect.Underscore(command.ControllerName))
	must(os.MkdirAll(path, 00755))

	/*
	// TODO: index
	outputPath = filepath.Join(path, "index.tmpl")
	builder = template.NewBuilder("template_index.tmpl")
	builder.WriteToPath(outputPath, command)

	// TODO: show
	outputPath = filepath.Join(path, "show.tmpl")
	builder = template.NewBuilder("template_show.tmpl")
	builder.WriteToPath(outputPath, command)

	// TODO: edit
	outputPath = filepath.Join(path, "edit.tmpl")
	builder = template.NewBuilder("template_edit.tmpl")
	builder.WriteToPath(outputPath, command)
	*/

	command.insertIntoRoutes()
}

func (command *ControllerCommand) insertIntoRoutes() {
	builder := template.NewBuilder("controller_router.go.tmpl")
	builder.InsertAfterToPath("controllers/router.go", "func Setup(", command)
}
