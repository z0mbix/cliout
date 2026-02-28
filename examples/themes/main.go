// Example: Built-in themes
//
// Demonstrates each of the built-in themes by printing sample output
// with each one. Best viewed in a terminal that supports true colour.
//
// Run: go run ./examples/themes/
package main

import (
	"fmt"

	"github.com/z0mbix/cliout"
)

func main() {
	cliout.SetLevel(cliout.LevelTrace)

	for _, t := range cliout.Themes() {
		fmt.Printf("\n--- %s ---\n", t.Name)
		cliout.SetTheme(t)
		cliout.Trace("trace message")
		cliout.Debug("debug message")
		cliout.Info("info message")
		cliout.Warn("warning message")
		cliout.Error("error message")
		cliout.Success("success message")
	}
}
