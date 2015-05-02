package command

import (
	//"fmt"
	"github.com/dcu/gin-scaffold/template"
	"strings"
)

type ModelCommand struct {
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
	modelName := args[0]
	fields := processFields(args[1:])

	renderer := template.NewRenderer("model.go.tmpl", modelName, fields)
	renderer.Render()
}
