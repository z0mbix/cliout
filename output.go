package cliout

import (
	"fmt"
	"io"
	"os"
)

const defaultPrefix = "»"

// Output holds all configuration for CLI output rendering.
type Output struct {
	writer       io.Writer
	level        Level
	prefix       string
	hasPrefix    bool
	prefixColor  Color
	messageColor Color
	theme        Theme
	colorEnabled bool
	noColorEnv   bool // true when NO_COLOR was detected at construction time
}

// New creates an Output with sensible defaults:
//   - Writer: os.Stdout
//   - Level: LevelInfo
//   - Prefix: "»" (or the value of the CLI_PREFIX environment variable)
//   - Theme: ThemeDefault (or the theme named by the CLI_THEME environment variable)
//   - Color: auto-detected (disabled if NO_COLOR is set or stdout is not a TTY)
func New() *Output {
	theme := ThemeDefault
	if name, ok := os.LookupEnv("CLI_THEME"); ok && name != "" {
		if t, found := ThemeByName(name); found {
			theme = t
		}
	}

	prefix := defaultPrefix
	hasPrefix := true
	if v, ok := os.LookupEnv("CLI_PREFIX"); ok {
		prefix = v
		if v == "" {
			hasPrefix = false
		}
	}

	o := &Output{
		writer:       os.Stdout,
		level:        LevelInfo,
		prefix:       prefix,
		hasPrefix:    hasPrefix,
		theme:        theme,
		colorEnabled: true,
	}

	// Respect NO_COLOR environment variable.
	// See https://no-color.org/
	if _, ok := os.LookupEnv("NO_COLOR"); ok {
		o.colorEnabled = false
		o.noColorEnv = true
	}

	// Disable color if stdout is not a terminal.
	if f, ok := o.writer.(*os.File); ok {
		if !isTerminal(f) {
			o.colorEnabled = false
		}
	}

	return o
}

// isTerminal reports whether f is a terminal (character device).
func isTerminal(f *os.File) bool {
	fi, err := f.Stat()
	if err != nil {
		return false
	}
	return (fi.Mode() & os.ModeCharDevice) != 0
}

// --- Configuration methods ---

// SetLevel sets the minimum output level. Messages below this level are suppressed.
func (o *Output) SetLevel(l Level) {
	o.level = l
}

// SetPrefix sets the prefix string prepended to each output line.
func (o *Output) SetPrefix(p string) {
	o.prefix = p
	o.hasPrefix = true
}

// ClearPrefix removes the prefix from output lines.
func (o *Output) ClearPrefix() {
	o.prefix = ""
	o.hasPrefix = false
}

// SetPrefixColor sets the color of the prefix, overriding the theme's prefix color.
func (o *Output) SetPrefixColor(c Color) {
	o.prefixColor = c
}

// SetMessageColor sets the color for all messages, overriding per-level theme colors.
func (o *Output) SetMessageColor(c Color) {
	o.messageColor = c
}

// SetTheme sets the color theme. This affects prefix and per-level message colors
// unless explicitly overridden with SetPrefixColor or SetMessageColor.
func (o *Output) SetTheme(t Theme) {
	o.theme = t
}

// SetWriter sets the output destination.
func (o *Output) SetWriter(w io.Writer) {
	o.writer = w
}

// SetColorEnabled explicitly enables or disables color output.
// If the NO_COLOR environment variable was set when this Output was created,
// color remains disabled regardless of the value passed here.
// See https://no-color.org/
func (o *Output) SetColorEnabled(enabled bool) {
	if o.noColorEnv {
		return
	}
	o.colorEnabled = enabled
}

// --- Output methods ---

// Info prints an info-level message.
func (o *Output) Info(msg string) {
	o.print(LevelInfo, msg, false)
}

// Infof prints a formatted info-level message.
func (o *Output) Infof(format string, a ...any) {
	o.print(LevelInfo, fmt.Sprintf(format, a...), false)
}

// Debug prints a debug-level message.
func (o *Output) Debug(msg string) {
	o.print(LevelDebug, msg, false)
}

// Debugf prints a formatted debug-level message.
func (o *Output) Debugf(format string, a ...any) {
	o.print(LevelDebug, fmt.Sprintf(format, a...), false)
}

// Trace prints a trace-level message.
func (o *Output) Trace(msg string) {
	o.print(LevelTrace, msg, false)
}

// Tracef prints a formatted trace-level message.
func (o *Output) Tracef(format string, a ...any) {
	o.print(LevelTrace, fmt.Sprintf(format, a...), false)
}

// Warn prints a warn-level message.
func (o *Output) Warn(msg string) {
	o.print(LevelWarn, msg, false)
}

// Warnf prints a formatted warn-level message.
func (o *Output) Warnf(format string, a ...any) {
	o.print(LevelWarn, fmt.Sprintf(format, a...), false)
}

// Error prints an error-level message.
func (o *Output) Error(msg string) {
	o.print(LevelError, msg, false)
}

// Errorf prints a formatted error-level message.
func (o *Output) Errorf(format string, a ...any) {
	o.print(LevelError, fmt.Sprintf(format, a...), false)
}

// Success prints a success message at info level, using the theme's SuccessColor.
func (o *Output) Success(msg string) {
	o.print(LevelInfo, msg, true)
}

// Successf prints a formatted success message at info level.
func (o *Output) Successf(format string, a ...any) {
	o.print(LevelInfo, fmt.Sprintf(format, a...), true)
}

// --- Inline color helpers ---

// Colorize wraps text with the given color's ANSI escape codes. The result
// respects this Output's color-enabled setting: when color is disabled the
// text is returned unchanged. Use this to compose multi-color messages with
// Infof, Errorf, etc.
func (o *Output) Colorize(text string, c Color) string {
	return c.apply(text, o.colorEnabled)
}

// --- Internal rendering ---

// print is the core rendering method. It handles level filtering, color application,
// prefix rendering, and writing the final output line.
func (o *Output) print(level Level, msg string, isSuccess bool) {
	if level < o.level {
		return
	}

	// Determine message color: explicit override > success color > theme per-level color.
	var msgColor Color
	if !o.messageColor.isDefault() {
		msgColor = o.messageColor
	} else if isSuccess {
		msgColor = o.theme.SuccessColor
	} else {
		msgColor = o.colorForLevel(level)
	}

	// Determine prefix color: explicit override > theme prefix color.
	prefixColor := o.theme.PrefixColor
	if !o.prefixColor.isDefault() {
		prefixColor = o.prefixColor
	}

	// Build the output line.
	var line string
	if o.hasPrefix && o.prefix != "" {
		coloredPrefix := prefixColor.apply(o.prefix, o.colorEnabled)
		coloredMsg := msgColor.apply(msg, o.colorEnabled)
		line = coloredPrefix + " " + coloredMsg
	} else {
		line = msgColor.apply(msg, o.colorEnabled)
	}

	_, _ = fmt.Fprintln(o.writer, line)
}

// colorForLevel returns the theme color for a given output level.
func (o *Output) colorForLevel(level Level) Color {
	switch level {
	case LevelTrace:
		return o.theme.TraceColor
	case LevelDebug:
		return o.theme.DebugColor
	case LevelInfo:
		return o.theme.InfoColor
	case LevelWarn:
		return o.theme.WarnColor
	case LevelError:
		return o.theme.ErrorColor
	default:
		return o.theme.InfoColor
	}
}
