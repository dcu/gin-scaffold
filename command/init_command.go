package command

import (
	//"fmt"
	"github.com/dcu/gin-scaffold/template"
	"os"
	"path/filepath"
	"strings"
)

var (
	dirsToCreate = []string{
		"models",
		"controllers",
		"helpers",
		"config",
		"path/filepath",
	}
)

type InitCommand struct {
	ProjectDir         string
	DatabaseNamePrefix string
}

func (command *InitCommand) Execute(args []string) {
	projectDir, err := filepath.Abs(args[0])
	if err != nil {
		panic(err)
	}

	command.ProjectDir = projectDir
	command.DatabaseNamePrefix = filepath.Base(projectDir)
	command.createLayout()

	command.installFiles("helpers")
	command.installFiles("config")
}

func (command *InitCommand) installFiles(dirName string) {
	helperFiles, err := filepath.Glob(template.TemplatePath(filepath.Join(dirName, "*.tmpl")))
	if err != nil {
		panic(err)
	}

	for _, templateFile := range helperFiles {
		builder := template.NewBuilder(templateFile)

		fileName := filepath.Base(templateFile)
		fileName = strings.TrimRight(fileName, ".tmpl")

		builder.Write(filepath.Join(command.ProjectDir, dirName, fileName), command)
	}
}

func (command *InitCommand) directoryInRoot(path string) string {
	return filepath.Join(command.ProjectDir, path)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func (command *InitCommand) createLayout() {
	for _, dirName := range dirsToCreate {
		path := command.directoryInRoot(dirName)
		must(os.MkdirAll(path, 00755))
	}
}
