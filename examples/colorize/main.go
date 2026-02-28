// Example: Colorize
//
// Demonstrates using Colorize to compose multi-colour lines. Each call
// to Colorize wraps a text segment with colour codes that respect the
// output's colour-enabled setting.
//
// Run: go run ./examples/colorize/
package main

import "github.com/z0mbix/cliout"

func main() {
	// Basic multi-colour line using ANSI colours.
	cliout.Infof("%s %s - %s",
		cliout.Colorize("ok", cliout.ColorGreen),
		cliout.Colorize("some blue text", cliout.ColorBlue),
		cliout.Colorize("some red text", cliout.ColorRed),
	)

	// Mix coloured and uncoloured segments.
	cliout.Infof("status: %s, latency: %s, cache: %s",
		cliout.Colorize("healthy", cliout.ColorGreen),
		cliout.Colorize("12ms", cliout.ColorYellow),
		cliout.Colorize("HIT", cliout.ColorCyan),
	)

	// Use true colours (24-bit hex).
	cliout.Infof("deployed %s to %s",
		cliout.Colorize("v2.4.1", cliout.Hex("#A9DC76")),
		cliout.Colorize("production", cliout.Hex("#FF6188")),
	)

	// Works with all output levels.
	cliout.Warnf("disk usage at %s on %s",
		cliout.Colorize("92%", cliout.ColorRed),
		cliout.Colorize("/dev/sda1", cliout.ColorWhite),
	)

	cliout.Errorf("service %s returned %s",
		cliout.Colorize("auth", cliout.ColorCyan),
		cliout.Colorize("503", cliout.ColorBrightRed),
	)

	// Also works on separate Output instances.
	out := cliout.New()
	out.SetPrefix("->")
	out.Infof("%s %s",
		out.Colorize("custom instance", cliout.ColorMagenta),
		out.Colorize("works too", cliout.ColorBrightGreen),
	)

	// When colour is disabled, Colorize returns plain text.
	out.SetColorEnabled(false)
	out.Infof("no colour: %s %s",
		out.Colorize("this", cliout.ColorRed),
		out.Colorize("is plain", cliout.ColorBlue),
	)
}
