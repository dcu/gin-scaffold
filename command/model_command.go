package command

import (
	//"fmt"
	"bitbucket.org/pkg/inflect"
	"github.com/dcu/gin-scaffold/template"
	"path/filepath"
	"strings"
)

type ModelCommand struct {
	PackageName        string
	ModelName          string
	ModelNamePlural    string
	InstanceName       string
	InstanceNamePlural string
	TemplateName       string
	Fields             map[string]string
}

// Converts "<fieldname>:<type>" to {"<fieldname>": "<type>"}
func processFields(args []string) map[string]string {
	fields := map[string]string{}
	for _, arg := range args {
		fieldNameAndType := strings.SplitN(arg, ":", 2)
		fields[fieldNameAndType[0]] = fieldNameAndType[1]
	}

	return fields
}

func (command *ModelCommand) Execute(args []string) {
	command.ModelName = args[0]
	command.ModelNamePlural = inflect.Pluralize(command.ModelName)

	command.Fields = processFields(args[1:])
	command.InstanceName = inflect.CamelizeDownFirst(command.ModelName)
	command.InstanceNamePlural = inflect.Pluralize(command.InstanceName)
	command.PackageName = template.PackageName()

	outputPath := filepath.Join("models", inflect.Underscore(command.ModelName)+".go")

	builder := template.NewBuilder("model.go.tmpl") //, modelName, fields)
	builder.Write(outputPath, command)
}
