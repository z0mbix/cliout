// Example: Output levels and level filtering
//
// Demonstrates all output levels (Trace, Debug, Info, Warn, Error, Success)
// and how to control which messages are visible by setting the minimum level.
//
// Run: go run ./examples/levels/
package main

import (
	"fmt"

	"github.com/z0mbix/cliout"
)

func main() {
	// The default level is LevelInfo, so only Info, Warn, Error, and
	// Success messages are visible.
	fmt.Println("=== Default level (Info) ===")
	cliout.Trace("this trace message is hidden")
	cliout.Debug("this debug message is hidden")
	cliout.Info("this info message is visible")
	cliout.Warn("this warning is visible")
	cliout.Error("this error is visible")
	cliout.Success("this success message is visible")

	// Lower the level to Debug to also see debug messages.
	fmt.Println("\n=== Debug level ===")
	cliout.SetLevel(cliout.LevelDebug)
	cliout.Trace("still hidden at debug level")
	cliout.Debug("now debug is visible")
	cliout.Info("info is still visible")

	// Lower to Trace to see everything.
	fmt.Println("\n=== Trace level (everything visible) ===")
	cliout.SetLevel(cliout.LevelTrace)
	cliout.Trace("trace is now visible")
	cliout.Debug("debug is visible")
	cliout.Info("info is visible")
	cliout.Warn("warn is visible")
	cliout.Error("error is visible")

	// Raise to Warn to only see warnings and errors.
	fmt.Println("\n=== Warn level (only warnings and errors) ===")
	cliout.SetLevel(cliout.LevelWarn)
	cliout.Info("this info message is hidden")
	cliout.Warn("this warning is visible")
	cliout.Error("this error is visible")

	// Raise to Error to only see errors.
	fmt.Println("\n=== Error level (only errors) ===")
	cliout.SetLevel(cliout.LevelError)
	cliout.Warn("this warning is hidden")
	cliout.Error("only errors are visible now")

	// Silent suppresses all output.
	fmt.Println("\n=== Silent level (nothing visible) ===")
	cliout.SetLevel(cliout.LevelSilent)
	cliout.Error("even this error is hidden")
	fmt.Println("(no output above)")
}
