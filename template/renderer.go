package template

import (
	"bitbucket.org/pkg/inflect"
	"os"
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

type Renderer struct {
	PackageName        string
	ModelName          string
	ModelNamePlural    string
	InstanceName       string
	InstanceNamePlural string
	TemplateName       string
	Data               map[string]string
}

func NewRenderer(templateName string, modelName string, data map[string]string) *Renderer {
	instanceName := inflect.CamelizeDownFirst(modelName)
	instanceNamePlural := inflect.Pluralize(instanceName)

	renderer := &Renderer{
		ModelName:          modelName,
		PackageName:        PackageName(),
		ModelNamePlural:    inflect.Pluralize(modelName),
		InstanceName:       instanceName,
		InstanceNamePlural: instanceNamePlural,
		TemplateName:       templateName,
		Data:               data,
	}

	return renderer
}

func (renderer *Renderer) Render() {
	contents := LoadTemplate(renderer.TemplateName)
	tmpl := template.Must(template.New(renderer.ModelName).Funcs(funcMap).Parse(contents))

	tmpl.Execute(os.Stdout, renderer)
}
