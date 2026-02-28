// Example: Custom themes
//
// Demonstrates how to create your own theme with custom colours for each
// output level, and how theme colours interact with explicit overrides.
//
// Run: go run ./examples/custom-theme/
package main

import (
	"fmt"

	"github.com/z0mbix/cliout"
)

func main() {
	cliout.SetLevel(cliout.LevelTrace)

	// --- Define a custom theme from scratch ---
	fmt.Println("=== Custom cyberpunk theme ===")

	cyberpunk := cliout.Theme{
		Name:         "Cyberpunk",
		PrefixColor:  cliout.Hex("#FF00FF"), // magenta
		InfoColor:    cliout.Hex("#00FFFF"), // cyan
		DebugColor:   cliout.Hex("#888888"), // gray
		TraceColor:   cliout.Hex("#555555"), // dark gray
		WarnColor:    cliout.Hex("#FFD700"), // gold
		ErrorColor:   cliout.Hex("#FF0040"), // hot pink
		SuccessColor: cliout.Hex("#00FF41"), // matrix green
	}

	cliout.SetTheme(cyberpunk)
	cliout.Trace("trace: scanning network...")
	cliout.Debug("debug: found 47 nodes")
	cliout.Info("connecting to mainframe")
	cliout.Warn("firewall detected")
	cliout.Error("intrusion countermeasure triggered")
	cliout.Success("access granted")

	// --- Earth tones theme ---
	fmt.Println("\n=== Custom earth tones theme ===")

	earthTones := cliout.Theme{
		Name:         "Earth Tones",
		PrefixColor:  cliout.Hex("#8B4513"), // saddle brown
		InfoColor:    cliout.Hex("#D2B48C"), // tan
		DebugColor:   cliout.Hex("#A0522D"), // sienna
		TraceColor:   cliout.Hex("#6B4226"), // dark brown
		WarnColor:    cliout.Hex("#DAA520"), // goldenrod
		ErrorColor:   cliout.Hex("#B22222"), // firebrick
		SuccessColor: cliout.Hex("#228B22"), // forest green
	}

	cliout.SetTheme(earthTones)
	cliout.Trace("trace: reading soil samples")
	cliout.Debug("debug: pH level 6.8")
	cliout.Info("planting season started")
	cliout.Warn("frost warning tonight")
	cliout.Error("crop failure in sector 7")
	cliout.Success("harvest complete")

	// --- Theme with explicit overrides ---
	fmt.Println("\n=== Theme with prefix colour override ===")

	cliout.SetTheme(cliout.ThemeNord)
	cliout.SetPrefixColor(cliout.Hex("#FF0000")) // red prefix, but Nord level colours
	cliout.Info("red prefix, Nord info colour")
	cliout.Warn("red prefix, Nord warn colour")
	cliout.Error("red prefix, Nord error colour")
	cliout.Success("red prefix, Nord success colour")

	// Reset the prefix colour override to let the theme control it again
	cliout.SetPrefixColor(cliout.ColorDefault)
	cliout.Info("back to Nord's prefix colour")
}
