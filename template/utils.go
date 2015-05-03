package template

import (
	"fmt"
	"github.com/mgutz/ansi"
)

func printAction(color string, action string, target string) {
	fmt.Printf("	%s  %s\n", ansi.Color(action, color), target)
}
