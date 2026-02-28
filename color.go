package cliout

import "fmt"

// Color represents a color that can be applied to terminal output.
// It supports standard ANSI colors and true color (24-bit RGB).
type Color struct {
	r, g, b     uint8
	isTrueColor bool
	ansiCode    int
}

// Standard ANSI color constants.
var (
	ColorDefault = Color{}
	ColorBlack   = Color{ansiCode: 30}
	ColorRed     = Color{ansiCode: 31}
	ColorGreen   = Color{ansiCode: 32}
	ColorYellow  = Color{ansiCode: 33}
	ColorBlue    = Color{ansiCode: 34}
	ColorMagenta = Color{ansiCode: 35}
	ColorCyan    = Color{ansiCode: 36}
	ColorWhite   = Color{ansiCode: 37}

	// Bright variants.
	ColorBrightBlack   = Color{ansiCode: 90}
	ColorBrightRed     = Color{ansiCode: 91}
	ColorBrightGreen   = Color{ansiCode: 92}
	ColorBrightYellow  = Color{ansiCode: 93}
	ColorBrightBlue    = Color{ansiCode: 94}
	ColorBrightMagenta = Color{ansiCode: 95}
	ColorBrightCyan    = Color{ansiCode: 96}
	ColorBrightWhite   = Color{ansiCode: 97}
)

// RGB creates a true color from red, green, and blue components (0-255).
func RGB(r, g, b uint8) Color {
	return Color{r: r, g: g, b: b, isTrueColor: true}
}

// Hex creates a true color from a hex color string (e.g., "#FF5733" or "FF5733").
// Returns ColorDefault if the input is invalid.
func Hex(hex string) Color {
	if len(hex) > 0 && hex[0] == '#' {
		hex = hex[1:]
	}
	if len(hex) != 6 {
		return ColorDefault
	}
	var r, g, b uint8
	_, err := fmt.Sscanf(hex, "%02x%02x%02x", &r, &g, &b)
	if err != nil {
		return ColorDefault
	}
	return RGB(r, g, b)
}

// isDefault returns true if this is the default (unset) color.
func (c Color) isDefault() bool {
	return !c.isTrueColor && c.ansiCode == 0
}

// apply wraps text with the appropriate ANSI escape codes.
// If the color is default or colorEnabled is false, the text is returned unchanged.
func (c Color) apply(text string, colorEnabled bool) string {
	if !colorEnabled || c.isDefault() {
		return text
	}
	if c.isTrueColor {
		return fmt.Sprintf("\033[38;2;%d;%d;%dm%s\033[0m", c.r, c.g, c.b, text)
	}
	return fmt.Sprintf("\033[%dm%s\033[0m", c.ansiCode, text)
}
