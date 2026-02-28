package cliout

// defaultOutput is the package-level Output instance used by the convenience functions.
var defaultOutput = New()

// Default returns the package-level default Output instance.
// This can be used to access the full Output API without creating a new instance.
func Default() *Output {
	return defaultOutput
}

// --- Package-level configuration functions ---

// SetLevel sets the minimum output level on the default output.
func SetLevel(l Level) {
	defaultOutput.SetLevel(l)
}

// SetPrefix sets the prefix string on the default output.
func SetPrefix(p string) {
	defaultOutput.SetPrefix(p)
}

// ClearPrefix removes the prefix from the default output.
func ClearPrefix() {
	defaultOutput.ClearPrefix()
}

// SetPrefixColor sets the prefix color on the default output.
func SetPrefixColor(c Color) {
	defaultOutput.SetPrefixColor(c)
}

// SetMessageColor sets the message color on the default output.
func SetMessageColor(c Color) {
	defaultOutput.SetMessageColor(c)
}

// SetTheme sets the theme on the default output.
func SetTheme(t Theme) {
	defaultOutput.SetTheme(t)
}

// SetColorEnabled enables or disables color on the default output.
func SetColorEnabled(enabled bool) {
	defaultOutput.SetColorEnabled(enabled)
}

// Colorize wraps text with the given color, respecting the default output's
// color-enabled setting.
func Colorize(text string, c Color) string {
	return defaultOutput.Colorize(text, c)
}

// --- Package-level output functions ---

// Info prints an info-level message.
func Info(msg string) {
	defaultOutput.Info(msg)
}

// Infof prints a formatted info-level message.
func Infof(format string, a ...any) {
	defaultOutput.Infof(format, a...)
}

// Debug prints a debug-level message.
func Debug(msg string) {
	defaultOutput.Debug(msg)
}

// Debugf prints a formatted debug-level message.
func Debugf(format string, a ...any) {
	defaultOutput.Debugf(format, a...)
}

// Trace prints a trace-level message.
func Trace(msg string) {
	defaultOutput.Trace(msg)
}

// Tracef prints a formatted trace-level message.
func Tracef(format string, a ...any) {
	defaultOutput.Tracef(format, a...)
}

// Warn prints a warn-level message.
func Warn(msg string) {
	defaultOutput.Warn(msg)
}

// Warnf prints a formatted warn-level message.
func Warnf(format string, a ...any) {
	defaultOutput.Warnf(format, a...)
}

// Error prints an error-level message.
func Error(msg string) {
	defaultOutput.Error(msg)
}

// Errorf prints a formatted error-level message.
func Errorf(format string, a ...any) {
	defaultOutput.Errorf(format, a...)
}

// Success prints a success message at info level using the theme's SuccessColor.
func Success(msg string) {
	defaultOutput.Success(msg)
}

// Successf prints a formatted success message at info level.
func Successf(format string, a ...any) {
	defaultOutput.Successf(format, a...)
}
