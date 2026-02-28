// Example: Format string variants
//
// Demonstrates using the printf-style format string methods (Infof, Debugf,
// Tracef, Warnf, Errorf, Successf) for dynamic output.
//
// Run: go run ./examples/formatting/
package main

import (
	"errors"

	"github.com/z0mbix/cliout"
)

func main() {
	cliout.SetLevel(cliout.LevelTrace)

	// String interpolation
	name := "somoflange"
	cliout.Infof("configuring %s...", name)

	// Integer formatting
	count := 42
	cliout.Infof("processed %d items", count)

	// Multiple arguments
	cliout.Infof("deployed %d services to %s", 3, "production")

	// Float formatting
	cliout.Infof("uptime: %.2f%%", 99.97)

	// Debug with context
	cliout.Debugf("connecting to %s:%d", "db.example.com", 5432)
	cliout.Debugf("query took %dms", 150)

	// Trace with detailed info
	cliout.Tracef("request headers: %s", "Content-Type: application/json")
	cliout.Tracef("response body size: %d bytes", 4096)

	// Warnings with values
	cliout.Warnf("disk usage at %d%% capacity", 92)
	cliout.Warnf("rate limit: %d/%d requests remaining", 12, 1000)

	// Errors with context
	err := errors.New("connection refused")
	cliout.Errorf("failed to connect to database: %v", err)
	cliout.Errorf("timeout after %ds on endpoint %s", 30, "/api/health")

	// Success with details
	cliout.Successf("all %d tests passed in %.1fs", 128, 3.7)
	cliout.Successf("backup complete: %d files, %d MB", 1500, 230)
}
