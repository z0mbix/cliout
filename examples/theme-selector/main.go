// Example: Theme selector
//
// Demonstrates using Themes() and ThemeByName() to let users pick a theme
// via a command-line flag. Run with --list to see all available themes,
// or --theme <name> to use a specific one.
//
// Run:
//
//	go run ./examples/theme-selector/
//	go run ./examples/theme-selector/ --list
//	go run ./examples/theme-selector/ --theme dracula
//	go run ./examples/theme-selector/ --theme "tokyo night storm"
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/z0mbix/cliout"
)

func main() {
	themeName := flag.String("theme", "Default", "colour theme to use")
	listThemes := flag.Bool("list", false, "list all available themes and exit")
	flag.Parse()

	// List all built-in themes and exit.
	if *listThemes {
		fmt.Println("Available themes:")
		for _, t := range cliout.Themes() {
			fmt.Printf("  %s\n", t.Name)
		}
		return
	}

	// Look up the requested theme by name (case-insensitive).
	theme, ok := cliout.ThemeByName(*themeName)
	if !ok {
		cliout.Errorf("unknown theme: %s", *themeName)
		fmt.Fprintln(os.Stderr, "Run with --list to see available themes.")
		os.Exit(1)
	}

	cliout.SetTheme(theme)
	cliout.SetLevel(cliout.LevelTrace)

	cliout.Infof("using theme: %s", theme.Name)
	cliout.Trace("trace message")
	cliout.Debug("debug message")
	cliout.Info("info message")
	cliout.Warn("warning message")
	cliout.Error("error message")
	cliout.Success("success message")
}
