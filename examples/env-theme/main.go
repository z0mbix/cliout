// Example: CLI_THEME and CLI_PREFIX environment variables
//
// Demonstrates that cliout automatically picks up the CLI_THEME and CLI_PREFIX
// environment variables. This lets users configure a consistent theme and
// prefix across all tools that use cliout.
//
// Run:
//
//	CLI_THEME=dracula go run ./examples/env-theme/
//	CLI_THEME="tokyo night storm" go run ./examples/env-theme/
//	CLI_PREFIX="->" go run ./examples/env-theme/
//	CLI_PREFIX="" go run ./examples/env-theme/
//	CLI_THEME=nord CLI_PREFIX="::" go run ./examples/env-theme/
//	CLI_THEME=nord NO_COLOR=1 go run ./examples/env-theme/
//	go run ./examples/env-theme/
package main

import (
	"os"

	"github.com/z0mbix/cliout"
)

func main() {
	// The theme and prefix are already set from CLI_THEME / CLI_PREFIX by the
	// time New() / the default instance is initialised. No application code needed.
	cliout.SetLevel(cliout.LevelTrace)

	theme := os.Getenv("CLI_THEME")
	if theme == "" {
		theme = "(not set, using Default)"
	}
	cliout.Infof("CLI_THEME  = %s", theme)

	if prefix, ok := os.LookupEnv("CLI_PREFIX"); ok {
		if prefix == "" {
			cliout.Infof("CLI_PREFIX = (empty, prefix cleared)")
		} else {
			cliout.Infof("CLI_PREFIX = %s", prefix)
		}
	} else {
		cliout.Infof("CLI_PREFIX = (not set, using default)")
	}

	cliout.Trace("trace message")
	cliout.Debug("debug message")
	cliout.Info("info message")
	cliout.Warn("warning message")
	cliout.Error("error message")
	cliout.Success("success message")

	// SetTheme / SetPrefix still override the env vars if the application wants to.
	cliout.SetTheme(cliout.ThemeNord)
	cliout.SetPrefix("!!!")
	cliout.Info("this line uses Nord with '!!!' prefix, regardless of env vars")
}
