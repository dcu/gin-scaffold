package command

import (
//"github.com/dcu/gin-scaffold/template"
//"os"
)

type Base interface {
	Execute(args []string)
	Help()
}
