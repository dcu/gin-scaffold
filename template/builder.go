package template

import (
	"bitbucket.org/pkg/inflect"
	"fmt"
	"io"
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
	if templatePath[0:1] != "/" {
		templatePath = TemplatePath(templatePath)
	}

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

func (builder *Builder) Write(writer io.Writer, data interface{}) {
	tmpl := builder.Template()
	err := tmpl.Execute(writer, data)
	if err != nil {
		panic(err)
	}
}

func (builder *Builder) WriteToPath(outputPath string, data interface{}) {
	fmt.Printf("Creating file: %s\n", outputPath)
	if _, err := os.Stat(outputPath); err == nil {
		fmt.Printf("File `%s` already exists. Skipping.\n", outputPath)
		return
	}

	file, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	builder.Write(file, data)
}
