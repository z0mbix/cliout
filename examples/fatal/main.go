// Example: Fatal output
//
// Demonstrates Fatal and Fatalf which print an error-level message and then
// call os.Exit(1). Useful for unrecoverable errors in CLI tools.
//
// Run:
//
//	go run ./examples/fatal/
package main

import "github.com/z0mbix/cliout"

func main() {
	cliout.SetTheme(cliout.ThemeDracula)

	cliout.Info("opening configuration file")
	cliout.Fatal("config.yaml: no such file or directory")

	// This line is never reached because Fatal calls os.Exit(1).
	cliout.Success("configuration loaded")
}
