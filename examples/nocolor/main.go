// Example: Disabling colour
//
// Demonstrates the different ways colour can be disabled: programmatically,
// via NO_COLOR, and automatic TTY detection.
//
// Run: go run ./examples/nocolor/
//
// Try also:
//
//	NO_COLOR=1 go run ./examples/nocolor/
//	go run ./examples/nocolor/ | cat
package main

import (
	"fmt"

	"github.com/z0mbix/cliout"
)

func main() {
	// By default, colour is auto-detected:
	// - Disabled if NO_COLOR env var is set (any value)
	// - Disabled if stdout is not a TTY (e.g., piped to a file or another command)
	// - Enabled otherwise

	fmt.Println("=== Default (auto-detected) ===")
	cliout.SetTheme(cliout.ThemeDracula)
	cliout.Info("colour depends on your terminal")
	cliout.Warn("try piping this to 'cat' to see it without colour")

	// --- Programmatically disable colour ---
	fmt.Println("\n=== Colour disabled ===")
	cliout.SetColorEnabled(false)
	cliout.Info("this has no ANSI escape codes")
	cliout.Warn("plain text warning")
	cliout.Error("plain text error")
	cliout.Success("plain text success")

	// --- Re-enable colour ---
	fmt.Println("\n=== Colour re-enabled ===")
	cliout.SetColorEnabled(true)
	cliout.Info("colour is back")
	cliout.Error("coloured error")

	// --- Per-instance colour control ---
	fmt.Println("\n=== Per-instance colour control ===")
	coloured := cliout.New()
	coloured.SetColorEnabled(true)
	coloured.SetTheme(cliout.ThemeNord)
	coloured.SetPrefix("[colour]")
	coloured.Info("this instance has colour enabled")

	plain := cliout.New()
	plain.SetColorEnabled(false)
	plain.SetPrefix("[plain]")
	plain.Info("this instance has colour disabled")

	// Both work independently
	coloured.Success("coloured success")
	plain.Success("plain success")
}
