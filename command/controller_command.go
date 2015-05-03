package command

import (
	//"fmt"
	"bitbucket.org/pkg/inflect"
	"bufio"
	"github.com/dcu/gin-scaffold/template"
	"os"
	"path/filepath"
	"strings"
)

type ControllerCommand struct {
	PackageName        string
	ControllerName     string
	ModelName          string
	InstanceName       string
	InstanceNamePlural string
	RoutePath          string
	TemplateName       string
	Fields             map[string]string
}

func (command *ControllerCommand) Execute(args []string) {
	command.ControllerName = args[0]
	command.RoutePath = inflect.Underscore(command.ControllerName)
	command.ModelName = inflect.Singularize(command.ControllerName)

	command.InstanceName = inflect.CamelizeDownFirst(command.ControllerName)
	command.InstanceNamePlural = inflect.Pluralize(command.InstanceName)
	command.PackageName = template.PackageName()

	outputPath := filepath.Join("controllers", inflect.Underscore(command.ControllerName)+".go")

	builder := template.NewBuilder("controller.go.tmpl")
	builder.WriteToPath(outputPath, command)

	command.insertIntoRoutes()
}

func (command *ControllerCommand) insertIntoRoutes() {
	newFilePath := "controllers/router.go.new"
	targetFilePath := "controllers/router.go"

	file, err := os.Open(targetFilePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	outputFile, err := os.Create(newFilePath)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	writer := bufio.NewWriter(outputFile)

	for scanner.Scan() {
		line := scanner.Text()

		writer.WriteString(line + "\n")
		if strings.HasPrefix(line, "func Setup(") {
			builder := template.NewBuilder("controller_router.go.tmpl")
			builder.Write(writer, command)
		}
	}

	writer.Flush()
	outputFile.Close()

	os.Rename(newFilePath, targetFilePath)
}
