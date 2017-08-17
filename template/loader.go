package template

import (
	"io/ioutil"
	"os"
	"path/filepath"
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
	wd = filepath.ToSlash(wd)
	for _, p := range filepath.SplitList(os.Getenv("GOPATH")) {
		p = filepath.ToSlash(p)
		if strings.HasPrefix(strings.ToLower(wd), strings.ToLower(p)) {
			return wd[len(p+"/src/"):]
		}
	}
	return ""
}

func ImportPath() string {
	return os.Getenv("GOPATH") + "/src/github.com/dcu/gin-scaffold"
}

func TemplatePath(name string) string {
	return ImportPath() + "/template/data/" + name
}
