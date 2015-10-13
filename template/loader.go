package template

import (
	"io/ioutil"
	"os"
	"strings"
)

func LoadTemplate(name string) string {
	return LoadTemplateFromFile(TemplatePath(name))
}

func LoadTemplateFromFile(filePath string) string {
	data, err := ioutil.ReadFile(filePath)

	if err != nil {
		panic(err)
	}

	return strings.Replace(string(data), "}\\\n", "}", -1)
}

func PackageName() string {
	wd, _ := os.Getwd()

	return strings.TrimPrefix(wd, os.Getenv("GOPATH")+"/src/")
}

func ImportPath() string {
	return os.Getenv("GOPATH") + "/src/github.com/dcu/gin-scaffold"
}

func TemplatePath(name string) string {
	return ImportPath() + "/template/data/" + name
}
