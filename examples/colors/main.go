// Example: Colours
//
// Demonstrates using built-in ANSI colours, custom RGB/Hex colours, and
// overriding prefix and message colours independently.
//
// Run: go run ./examples/colors/
package main

import (
	"fmt"

	"github.com/z0mbix/cliout"
)

func main() {
	// --- Built-in ANSI colours ---
	fmt.Println("=== ANSI prefix colours ===")

	cliout.SetPrefixColor(cliout.ColorRed)
	cliout.Info("red prefix")

	cliout.SetPrefixColor(cliout.ColorGreen)
	cliout.Info("green prefix")

	cliout.SetPrefixColor(cliout.ColorBlue)
	cliout.Info("blue prefix")

	cliout.SetPrefixColor(cliout.ColorMagenta)
	cliout.Info("magenta prefix")

	cliout.SetPrefixColor(cliout.ColorBrightCyan)
	cliout.Info("bright cyan prefix")

	// --- Custom hex colours ---
	fmt.Println("\n=== Hex prefix colours ===")

	cliout.SetPrefixColor(cliout.Hex("#FF5733"))
	cliout.Info("coral prefix via hex")

	cliout.SetPrefixColor(cliout.Hex("00CED1"))
	cliout.Info("dark turquoise prefix (no # needed)")

	// --- Custom RGB colours ---
	fmt.Println("\n=== RGB prefix colours ===")

	cliout.SetPrefixColor(cliout.RGB(255, 165, 0))
	cliout.Info("orange prefix via RGB")

	cliout.SetPrefixColor(cliout.RGB(128, 0, 255))
	cliout.Info("purple prefix via RGB")

	// --- Message colour override ---
	fmt.Println("\n=== Message colour override ===")

	// Reset prefix colour to default theme colour
	cliout.SetPrefixColor(cliout.ColorDefault)

	cliout.SetMessageColor(cliout.ColorWhite)
	cliout.Info("all messages are now white")
	cliout.Warn("even warnings are white")
	cliout.Error("and errors too")

	cliout.SetMessageColor(cliout.Hex("#A6E3A1"))
	cliout.Info("custom green message colour")

	// --- Both prefix and message colours ---
	fmt.Println("\n=== Combined prefix + message colours ===")

	cliout.SetPrefixColor(cliout.Hex("#FF79C6"))
	cliout.SetMessageColor(cliout.Hex("#F8F8F2"))
	cliout.Info("pink prefix, light message")

	cliout.SetPrefixColor(cliout.ColorYellow)
	cliout.SetMessageColor(cliout.ColorCyan)
	cliout.Info("yellow prefix, cyan message")

	// --- Reset to defaults ---
	fmt.Println("\n=== Reset to default colours ===")

	cliout.SetPrefixColor(cliout.ColorDefault)
	cliout.SetMessageColor(cliout.ColorDefault)
	cliout.Info("back to theme defaults")
}
