package template

import (
	"bitbucket.org/pkg/inflect"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var (
	funcMap = template.FuncMap{
		"Pluralize":  inflect.Pluralize,
		"Underscore": inflect.Underscore,
		"ToUpper":    strings.ToUpper,
		"ToLower":    strings.ToLower,
	}
)

type Builder struct {
	TemplateName string
	TemplatePath string
}

func NewBuilder(templatePath string) *Builder {
	templateName := filepath.Base(templatePath)
	builder := &Builder{
		TemplateName: templateName,
		TemplatePath: templatePath,
	}

	return builder
}

func (builder *Builder) Template() *template.Template {
	contents := LoadTemplateFromFile(builder.TemplatePath)
	tmpl := template.Must(template.New(builder.TemplateName).Funcs(funcMap).Parse(contents))

	return tmpl
}

func (builder *Builder) Write(outputPath string, data interface{}) {
	if _, err := os.Stat(outputPath); err == nil {
		fmt.Printf("File `%s` already exists. Skipping.\n", outputPath)
		return
	}

	tmpl := builder.Template()
	file, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	tmpl.Execute(writer, data)
	writer.Flush()
}
