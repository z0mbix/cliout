// Example: Separate Output instances
//
// Demonstrates creating independent Output instances with their own
// configuration, using custom writers, and writing to stderr.
//
// Run: go run ./examples/instance/
package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/z0mbix/cliout"
)

func main() {
	// --- Two independent instances with different themes ---
	fmt.Println("=== Independent instances ===")

	app := cliout.New()
	app.SetTheme(cliout.ThemeDracula)
	app.SetPrefix("[app]")
	app.Info("application started")
	app.Info("loading configuration")

	db := cliout.New()
	db.SetTheme(cliout.ThemeNord)
	db.SetPrefix("[db]")
	db.Info("connected to database")
	db.Info("running migrations")

	// They don't interfere with each other
	app.Success("ready to serve requests")
	db.Success("migrations complete")

	// --- Instance with debug level enabled ---
	fmt.Println("\n=== Instance with debug level ===")

	verbose := cliout.New()
	verbose.SetLevel(cliout.LevelDebug)
	verbose.SetTheme(cliout.ThemeTokyoNightStorm)
	verbose.SetPrefix(">>>")
	verbose.Debug("this debug message is visible")
	verbose.Info("info is visible too")

	// The default instance is unaffected
	cliout.Debug("this is still hidden (default is Info level)")
	cliout.Info("default instance is independent")

	// --- Writing to stderr ---
	fmt.Println("\n=== Writing to stderr ===")

	errOut := cliout.New()
	errOut.SetWriter(os.Stderr)
	errOut.SetPrefix("!!")
	errOut.SetTheme(cliout.ThemeGruvboxDark)
	errOut.Error("this goes to stderr")
	errOut.Warn("this warning also goes to stderr")

	// --- Writing to a buffer (useful for testing or capturing output) ---
	fmt.Println("\n=== Writing to a buffer ===")

	var buf bytes.Buffer
	captured := cliout.New()
	captured.SetWriter(&buf)
	captured.SetColorEnabled(false)
	captured.Info("first captured line")
	captured.Info("second captured line")
	captured.Warn("captured warning")

	fmt.Printf("Captured %d bytes:\n", buf.Len())
	fmt.Print(buf.String())

	// --- Accessing the default instance ---
	fmt.Println("\n=== Default instance ===")

	d := cliout.Default()
	d.SetTheme(cliout.ThemeCatppuccinoMocha)
	d.Info("via Default() - same as cliout.Info()")
	cliout.Info("via package function - same instance")
}
