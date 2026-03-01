// Example: CLI_THEME environment variable
//
// Demonstrates that cliout automatically picks up the CLI_THEME environment
// variable to set the colour theme. This lets users configure a consistent
// theme across all tools that use cliout.
//
// Run:
//
//	CLI_THEME=dracula go run ./examples/env-theme/
//	CLI_THEME="tokyo night storm" go run ./examples/env-theme/
//	CLI_THEME=nord NO_COLOR=1 go run ./examples/env-theme/
//	go run ./examples/env-theme/
package main

import (
	"os"

	"github.com/z0mbix/cliout"
)

func main() {
	// The theme is already set from CLI_THEME by the time New() / the
	// default instance is initialised. No application code needed.
	cliout.SetLevel(cliout.LevelTrace)

	theme := os.Getenv("CLI_THEME")
	if theme == "" {
		theme = "(not set, using Default)"
	}
	cliout.Infof("CLI_THEME = %s", theme)

	cliout.Trace("trace message")
	cliout.Debug("debug message")
	cliout.Info("info message")
	cliout.Warn("warning message")
	cliout.Error("error message")
	cliout.Success("success message")

	// SetTheme still overrides CLI_THEME if the application wants to.
	cliout.SetTheme(cliout.ThemeNord)
	cliout.Info("this line uses Nord, regardless of CLI_THEME")
}
