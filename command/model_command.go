package command

import (
	"bitbucket.org/pkg/inflect"
	"fmt"
	"github.com/dcu/gin-scaffold/template"
	"path/filepath"
	"strings"
)

// ModelCommand generates files related to model.
type ModelCommand struct {
	PackageName        string
	ModelName          string
	ModelNamePlural    string
	InstanceName       string
	InstanceNamePlural string
	TemplateName       string
	Fields             map[string]string
}

// Help prints a help message for this command.
func (command *ModelCommand) Help() {
	fmt.Printf(`Usage:
	gin-scaffold model <model name> <field name>:<field type> ...

Description:
	The gin-scaffold model command creates a new model with the given fields.

Example:
	gin-scaffold model Post Title:string Body:string 
`)
}

func findFieldType(name string) string {
	switch name {
	case "text":
		{
		name = "string"
		}
	case "float":
		{
		name = "float64"
		}
	case "boolean":
		{
		name = "bool"
		}
	case "integer":
		{
		name = "int"
		}
	case
		"time",
		"datetime":
		{
		name = "int64"
		}
	}

	return name
}

// Converts "<fieldname>:<type>" to {"<fieldname>": "<type>"}
func processFields(args []string) map[string]string {
	fields := map[string]string{}
	for _, arg := range args {
		fieldNameAndType := strings.SplitN(arg, ":", 2)
		fields[inflect.Titleize(fieldNameAndType[0])] = findFieldType(fieldNameAndType[1])
	}

	return fields
}

// Execute runs this command.
func (command *ModelCommand) Execute(args []string) {
	command.ModelName = inflect.Titleize(args[0])
	command.ModelNamePlural = inflect.Pluralize(command.ModelName)

	command.Fields = processFields(args[1:])
	command.InstanceName = inflect.CamelizeDownFirst(command.ModelName)
	command.InstanceNamePlural = inflect.Pluralize(command.InstanceName)
	command.PackageName = template.PackageName()

	outputPath := filepath.Join("models", inflect.Underscore(command.ModelName)+".go")

	builder := template.NewBuilder("model.go.tmpl")
	builder.WriteToPath(outputPath, command)

	outputPath = filepath.Join("models", inflect.Underscore(command.ModelName)+"_dbsession.go")
	builder = template.NewBuilder("model_dbsession.go.tmpl")
	builder.WriteToPath(outputPath, command)
}
