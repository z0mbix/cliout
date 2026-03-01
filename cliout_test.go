package cliout

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

// newTestOutput creates an Output that writes to a buffer with color enabled
// and level set to the lowest (LevelTrace) so all messages are captured.
func newTestOutput() (*Output, *bytes.Buffer) {
	var buf bytes.Buffer
	o := &Output{
		writer:       &buf,
		level:        LevelTrace,
		prefix:       defaultPrefix,
		hasPrefix:    true,
		theme:        ThemeDefault,
		colorEnabled: false, // disabled by default in tests for easy string matching
		exitFunc:     os.Exit,
	}
	return o, &buf
}

// --- Color tests ---

func TestHexValid(t *testing.T) {
	c := Hex("#FF5733")
	if !c.isTrueColor {
		t.Fatal("expected true color")
	}
	if c.r != 0xFF || c.g != 0x57 || c.b != 0x33 {
		t.Fatalf("expected RGB(255,87,51), got RGB(%d,%d,%d)", c.r, c.g, c.b)
	}
}

func TestHexWithoutHash(t *testing.T) {
	c := Hex("FF5733")
	if !c.isTrueColor {
		t.Fatal("expected true color")
	}
	if c.r != 0xFF || c.g != 0x57 || c.b != 0x33 {
		t.Fatalf("expected RGB(255,87,51), got RGB(%d,%d,%d)", c.r, c.g, c.b)
	}
}

func TestHexInvalid(t *testing.T) {
	cases := []string{"", "#", "GG0000", "#12345", "1234567", "#ZZZZZZ"}
	for _, hex := range cases {
		c := Hex(hex)
		if !c.isDefault() {
			t.Errorf("Hex(%q) should return ColorDefault, got %+v", hex, c)
		}
	}
}

func TestRGB(t *testing.T) {
	c := RGB(10, 20, 30)
	if !c.isTrueColor || c.r != 10 || c.g != 20 || c.b != 30 {
		t.Fatalf("unexpected RGB result: %+v", c)
	}
}

func TestColorIsDefault(t *testing.T) {
	if !ColorDefault.isDefault() {
		t.Fatal("ColorDefault should be default")
	}
	if ColorRed.isDefault() {
		t.Fatal("ColorRed should not be default")
	}
	if RGB(0, 0, 0).isDefault() {
		t.Fatal("RGB(0,0,0) should not be default (it's true color black)")
	}
}

func TestColorApplyDisabled(t *testing.T) {
	result := ColorRed.apply("hello", false)
	if result != "hello" {
		t.Fatalf("expected plain text, got %q", result)
	}
}

func TestColorApplyDefault(t *testing.T) {
	result := ColorDefault.apply("hello", true)
	if result != "hello" {
		t.Fatalf("expected plain text for default color, got %q", result)
	}
}

func TestColorApplyANSI(t *testing.T) {
	result := ColorRed.apply("hello", true)
	expected := "\033[31mhello\033[0m"
	if result != expected {
		t.Fatalf("expected %q, got %q", expected, result)
	}
}

func TestColorApplyTrueColor(t *testing.T) {
	c := RGB(255, 128, 0)
	result := c.apply("hello", true)
	expected := "\033[38;2;255;128;0mhello\033[0m"
	if result != expected {
		t.Fatalf("expected %q, got %q", expected, result)
	}
}

// --- Level tests ---

func TestLevelString(t *testing.T) {
	tests := []struct {
		level Level
		want  string
	}{
		{LevelTrace, "trace"},
		{LevelDebug, "debug"},
		{LevelInfo, "info"},
		{LevelWarn, "warn"},
		{LevelError, "error"},
		{LevelSilent, "silent"},
		{Level(99), "unknown"},
	}
	for _, tt := range tests {
		if got := tt.level.String(); got != tt.want {
			t.Errorf("Level(%d).String() = %q, want %q", tt.level, got, tt.want)
		}
	}
}

func TestLevelOrdering(t *testing.T) {
	if LevelTrace >= LevelDebug {
		t.Fatal("Trace should be less than Debug")
	}
	if LevelDebug >= LevelInfo {
		t.Fatal("Debug should be less than Info")
	}
	if LevelInfo >= LevelWarn {
		t.Fatal("Info should be less than Warn")
	}
	if LevelWarn >= LevelError {
		t.Fatal("Warn should be less than Error")
	}
	if LevelError >= LevelSilent {
		t.Fatal("Error should be less than Silent")
	}
}

// --- Output tests ---

func TestInfoOutput(t *testing.T) {
	o, buf := newTestOutput()
	o.Info("hello world")
	got := buf.String()
	if !strings.Contains(got, "hello world") {
		t.Fatalf("expected output to contain 'hello world', got %q", got)
	}
	if !strings.Contains(got, defaultPrefix) {
		t.Fatalf("expected output to contain prefix %q, got %q", defaultPrefix, got)
	}
}

func TestInfofOutput(t *testing.T) {
	o, buf := newTestOutput()
	o.Infof("count: %d, name: %s", 42, "test")
	got := buf.String()
	if !strings.Contains(got, "count: 42, name: test") {
		t.Fatalf("expected formatted output, got %q", got)
	}
}

func TestDebugOutput(t *testing.T) {
	o, buf := newTestOutput()
	o.Debug("debug message")
	got := buf.String()
	if !strings.Contains(got, "debug message") {
		t.Fatalf("expected debug output, got %q", got)
	}
}

func TestTraceOutput(t *testing.T) {
	o, buf := newTestOutput()
	o.Trace("trace message")
	got := buf.String()
	if !strings.Contains(got, "trace message") {
		t.Fatalf("expected trace output, got %q", got)
	}
}

func TestWarnOutput(t *testing.T) {
	o, buf := newTestOutput()
	o.Warn("warning!")
	got := buf.String()
	if !strings.Contains(got, "warning!") {
		t.Fatalf("expected warn output, got %q", got)
	}
}

func TestErrorOutput(t *testing.T) {
	o, buf := newTestOutput()
	o.Error("error occurred")
	got := buf.String()
	if !strings.Contains(got, "error occurred") {
		t.Fatalf("expected error output, got %q", got)
	}
}

func TestSuccessOutput(t *testing.T) {
	o, buf := newTestOutput()
	o.Success("it worked")
	got := buf.String()
	if !strings.Contains(got, "it worked") {
		t.Fatalf("expected success output, got %q", got)
	}
}

func TestSuccessfOutput(t *testing.T) {
	o, buf := newTestOutput()
	o.Successf("completed %d tasks", 5)
	got := buf.String()
	if !strings.Contains(got, "completed 5 tasks") {
		t.Fatalf("expected formatted success output, got %q", got)
	}
}

// --- Level filtering tests ---

func TestLevelFilteringDebugHiddenAtInfo(t *testing.T) {
	o, buf := newTestOutput()
	o.SetLevel(LevelInfo)
	o.Debug("should be hidden")
	if buf.Len() != 0 {
		t.Fatalf("debug message should be suppressed at info level, got %q", buf.String())
	}
}

func TestLevelFilteringTraceHiddenAtInfo(t *testing.T) {
	o, buf := newTestOutput()
	o.SetLevel(LevelInfo)
	o.Trace("should be hidden")
	if buf.Len() != 0 {
		t.Fatalf("trace message should be suppressed at info level, got %q", buf.String())
	}
}

func TestLevelFilteringDebugVisibleAtDebug(t *testing.T) {
	o, buf := newTestOutput()
	o.SetLevel(LevelDebug)
	o.Debug("visible debug")
	if !strings.Contains(buf.String(), "visible debug") {
		t.Fatalf("debug message should be visible at debug level, got %q", buf.String())
	}
}

func TestLevelFilteringInfoVisibleAtDebug(t *testing.T) {
	o, buf := newTestOutput()
	o.SetLevel(LevelDebug)
	o.Info("visible info")
	if !strings.Contains(buf.String(), "visible info") {
		t.Fatalf("info message should be visible at debug level, got %q", buf.String())
	}
}

func TestLevelFilteringErrorVisibleAtWarn(t *testing.T) {
	o, buf := newTestOutput()
	o.SetLevel(LevelWarn)
	o.Error("visible error")
	if !strings.Contains(buf.String(), "visible error") {
		t.Fatalf("error should be visible at warn level, got %q", buf.String())
	}
}

func TestLevelFilteringInfoHiddenAtWarn(t *testing.T) {
	o, buf := newTestOutput()
	o.SetLevel(LevelWarn)
	o.Info("hidden info")
	if buf.Len() != 0 {
		t.Fatalf("info should be suppressed at warn level, got %q", buf.String())
	}
}

func TestLevelSilent(t *testing.T) {
	o, buf := newTestOutput()
	o.SetLevel(LevelSilent)
	o.Error("should be hidden")
	o.Warn("should be hidden")
	o.Info("should be hidden")
	o.Debug("should be hidden")
	o.Trace("should be hidden")
	if buf.Len() != 0 {
		t.Fatalf("all messages should be suppressed at silent level, got %q", buf.String())
	}
}

// --- Prefix tests ---

func TestDefaultPrefix(t *testing.T) {
	o, buf := newTestOutput()
	o.Info("test")
	got := buf.String()
	if !strings.HasPrefix(got, defaultPrefix+" ") {
		t.Fatalf("expected output to start with prefix %q, got %q", defaultPrefix, got)
	}
}

func TestCustomPrefix(t *testing.T) {
	o, buf := newTestOutput()
	o.SetPrefix("->")
	o.Info("test")
	got := buf.String()
	if !strings.HasPrefix(got, "-> ") {
		t.Fatalf("expected output to start with '-> ', got %q", got)
	}
}

func TestClearPrefix(t *testing.T) {
	o, buf := newTestOutput()
	o.ClearPrefix()
	o.Info("no prefix")
	got := strings.TrimSpace(buf.String())
	if got != "no prefix" {
		t.Fatalf("expected 'no prefix' without any prefix, got %q", got)
	}
}

func TestSetPrefixAfterClear(t *testing.T) {
	o, buf := newTestOutput()
	o.ClearPrefix()
	o.SetPrefix("*")
	o.Info("test")
	got := buf.String()
	if !strings.HasPrefix(got, "* ") {
		t.Fatalf("expected output to start with '* ', got %q", got)
	}
}

// --- Color enable/disable tests ---

func TestColorDisabledNoAnsiCodes(t *testing.T) {
	o, buf := newTestOutput()
	o.SetColorEnabled(false)
	o.Info("plain text")
	got := buf.String()
	if strings.Contains(got, "\033[") {
		t.Fatalf("expected no ANSI escape codes when color disabled, got %q", got)
	}
}

func TestColorEnabledHasAnsiCodes(t *testing.T) {
	o, buf := newTestOutput()
	o.SetColorEnabled(true)
	o.SetTheme(ThemeDracula) // has non-default colors
	o.Info("colored text")
	got := buf.String()
	if !strings.Contains(got, "\033[") {
		t.Fatalf("expected ANSI escape codes when color enabled with theme, got %q", got)
	}
}

// --- Theme tests ---

func TestSetThemeChangesColors(t *testing.T) {
	o, buf := newTestOutput()
	o.SetColorEnabled(true)
	o.SetTheme(ThemeDracula)
	o.Info("dracula info")
	got := buf.String()
	// Dracula info color is #F8F8F2 = RGB(248,248,242)
	if !strings.Contains(got, "38;2;248;248;242") {
		t.Fatalf("expected Dracula info color in output, got %q", got)
	}
	// Dracula prefix color is #BD93F9 = RGB(189,147,249)
	if !strings.Contains(got, "38;2;189;147;249") {
		t.Fatalf("expected Dracula prefix color in output, got %q", got)
	}
}

func TestThemeErrorColor(t *testing.T) {
	o, buf := newTestOutput()
	o.SetColorEnabled(true)
	o.SetTheme(ThemeNord)
	o.Error("nord error")
	got := buf.String()
	// Nord error color is #BF616A = RGB(191,97,106)
	if !strings.Contains(got, "38;2;191;97;106") {
		t.Fatalf("expected Nord error color in output, got %q", got)
	}
}

func TestThemeWarnColor(t *testing.T) {
	o, buf := newTestOutput()
	o.SetColorEnabled(true)
	o.SetTheme(ThemeOneDark)
	o.Warn("warning")
	got := buf.String()
	// One Dark warn color is #E5C07B = RGB(229,192,123)
	if !strings.Contains(got, "38;2;229;192;123") {
		t.Fatalf("expected One Dark warn color in output, got %q", got)
	}
}

func TestSuccessUsesSuccessColor(t *testing.T) {
	o, buf := newTestOutput()
	o.SetColorEnabled(true)
	o.SetTheme(ThemeDracula)
	o.Success("done")
	got := buf.String()
	// Dracula success color is #50FA7B = RGB(80,250,123)
	if !strings.Contains(got, "38;2;80;250;123") {
		t.Fatalf("expected Dracula success color in output, got %q", got)
	}
}

// --- Custom color override tests ---

func TestSetPrefixColorOverridesTheme(t *testing.T) {
	o, buf := newTestOutput()
	o.SetColorEnabled(true)
	o.SetTheme(ThemeDracula)
	o.SetPrefixColor(ColorRed) // override Dracula's purple prefix
	o.Info("test")
	got := buf.String()
	// Should have ANSI red (31) for prefix, not Dracula purple
	if !strings.Contains(got, "\033[31m"+defaultPrefix+"\033[0m") {
		t.Fatalf("expected ANSI red prefix, got %q", got)
	}
}

func TestSetMessageColorOverridesTheme(t *testing.T) {
	o, buf := newTestOutput()
	o.SetColorEnabled(true)
	o.SetTheme(ThemeDracula)
	o.SetMessageColor(ColorGreen) // override Dracula's info color
	o.Info("test")
	got := buf.String()
	// Should have ANSI green (32) for message
	if !strings.Contains(got, "\033[32mtest\033[0m") {
		t.Fatalf("expected ANSI green message, got %q", got)
	}
}

func TestSetMessageColorOverridesSuccessColor(t *testing.T) {
	o, buf := newTestOutput()
	o.SetColorEnabled(true)
	o.SetTheme(ThemeDracula)
	o.SetMessageColor(ColorYellow) // should override even success color
	o.Success("test")
	got := buf.String()
	if !strings.Contains(got, "\033[33mtest\033[0m") {
		t.Fatalf("expected ANSI yellow message even for success, got %q", got)
	}
}

// --- Format string variants ---

func TestDebugf(t *testing.T) {
	o, buf := newTestOutput()
	o.Debugf("value: %d", 123)
	if !strings.Contains(buf.String(), "value: 123") {
		t.Fatalf("expected formatted debug output, got %q", buf.String())
	}
}

func TestTracef(t *testing.T) {
	o, buf := newTestOutput()
	o.Tracef("trace %s %d", "item", 7)
	if !strings.Contains(buf.String(), "trace item 7") {
		t.Fatalf("expected formatted trace output, got %q", buf.String())
	}
}

func TestWarnf(t *testing.T) {
	o, buf := newTestOutput()
	o.Warnf("warn: %s", "careful")
	if !strings.Contains(buf.String(), "warn: careful") {
		t.Fatalf("expected formatted warn output, got %q", buf.String())
	}
}

func TestErrorf(t *testing.T) {
	o, buf := newTestOutput()
	o.Errorf("error: %v", "something broke")
	if !strings.Contains(buf.String(), "error: something broke") {
		t.Fatalf("expected formatted error output, got %q", buf.String())
	}
}

// --- Fatal tests ---

func TestFatalOutput(t *testing.T) {
	o, buf := newTestOutput()
	var exitCode int
	o.exitFunc = func(code int) { exitCode = code }

	o.Fatal("fatal error")
	got := buf.String()
	if !strings.Contains(got, "fatal error") {
		t.Fatalf("expected fatal output, got %q", got)
	}
	if exitCode != 1 {
		t.Fatalf("expected exit code 1, got %d", exitCode)
	}
}

func TestFatalfOutput(t *testing.T) {
	o, buf := newTestOutput()
	var exitCode int
	o.exitFunc = func(code int) { exitCode = code }

	o.Fatalf("fatal: %s", "disk full")
	got := buf.String()
	if !strings.Contains(got, "fatal: disk full") {
		t.Fatalf("expected formatted fatal output, got %q", got)
	}
	if exitCode != 1 {
		t.Fatalf("expected exit code 1, got %d", exitCode)
	}
}

func TestFatalUsesErrorColor(t *testing.T) {
	o, buf := newTestOutput()
	o.SetColorEnabled(true)
	o.SetTheme(ThemeDracula)
	o.exitFunc = func(code int) {}

	o.Fatal("crash")
	got := buf.String()
	// Dracula error color is #FF5555 = RGB(255,85,85)
	if !strings.Contains(got, "38;2;255;85;85") {
		t.Fatalf("expected Dracula error color in fatal output, got %q", got)
	}
}

func TestFatalSuppressedBySilent(t *testing.T) {
	o, buf := newTestOutput()
	var exitCode int
	o.exitFunc = func(code int) { exitCode = code }
	o.SetLevel(LevelSilent)

	o.Fatal("should be hidden")
	if buf.Len() != 0 {
		t.Fatalf("fatal message should be suppressed at silent level, got %q", buf.String())
	}
	// Exit must still be called even when output is suppressed.
	if exitCode != 1 {
		t.Fatalf("expected exit code 1 even at silent level, got %d", exitCode)
	}
}

func TestFatalfSuppressedBySilent(t *testing.T) {
	o, buf := newTestOutput()
	var exitCode int
	o.exitFunc = func(code int) { exitCode = code }
	o.SetLevel(LevelSilent)

	o.Fatalf("should be hidden: %d", 42)
	if buf.Len() != 0 {
		t.Fatalf("fatalf message should be suppressed at silent level, got %q", buf.String())
	}
	if exitCode != 1 {
		t.Fatalf("expected exit code 1 even at silent level, got %d", exitCode)
	}
}

// --- Output ends with newline ---

func TestOutputEndsWithNewline(t *testing.T) {
	o, buf := newTestOutput()
	o.Info("test")
	got := buf.String()
	if !strings.HasSuffix(got, "\n") {
		t.Fatalf("expected output to end with newline, got %q", got)
	}
}

// --- Writer tests ---

func TestSetWriter(t *testing.T) {
	o, _ := newTestOutput()
	var newBuf bytes.Buffer
	o.SetWriter(&newBuf)
	o.Info("new writer")
	if !strings.Contains(newBuf.String(), "new writer") {
		t.Fatalf("expected output in new writer, got %q", newBuf.String())
	}
}

// --- Default instance tests ---

func TestDefaultInstance(t *testing.T) {
	d := Default()
	if d == nil {
		t.Fatal("Default() should not return nil")
	}
	if d != defaultOutput {
		t.Fatal("Default() should return the package-level instance")
	}
}

// setupDefaultForTest redirects the package-level default instance to a
// buffer and enables all levels so we can capture output. It returns the
// buffer and a cleanup function that restores the original state.
func setupDefaultForTest() (*bytes.Buffer, func()) {
	d := Default()
	origWriter := d.writer
	origLevel := d.level
	origTheme := d.theme
	origPrefix := d.prefix
	origHasPrefix := d.hasPrefix
	origColor := d.colorEnabled
	origPrefixColor := d.prefixColor
	origMessageColor := d.messageColor
	origNoColorEnv := d.noColorEnv
	origExitFunc := d.exitFunc

	var buf bytes.Buffer
	d.writer = &buf
	d.level = LevelTrace
	d.colorEnabled = false
	d.noColorEnv = false
	d.theme = ThemeDefault
	d.prefix = defaultPrefix
	d.hasPrefix = true
	d.prefixColor = ColorDefault
	d.messageColor = ColorDefault
	d.exitFunc = os.Exit

	cleanup := func() {
		d.writer = origWriter
		d.level = origLevel
		d.theme = origTheme
		d.prefix = origPrefix
		d.hasPrefix = origHasPrefix
		d.colorEnabled = origColor
		d.noColorEnv = origNoColorEnv
		d.prefixColor = origPrefixColor
		d.messageColor = origMessageColor
		d.exitFunc = origExitFunc
	}
	return &buf, cleanup
}

func TestPackageLevelSetLevelAndOutput(t *testing.T) {
	buf, cleanup := setupDefaultForTest()
	defer cleanup()

	SetLevel(LevelWarn)
	Info("hidden")
	if buf.Len() != 0 {
		t.Fatalf("info should be hidden at warn level, got %q", buf.String())
	}
	Warn("visible warn")
	if !strings.Contains(buf.String(), "visible warn") {
		t.Fatalf("expected warn output, got %q", buf.String())
	}
}

func TestPackageLevelInfo(t *testing.T) {
	buf, cleanup := setupDefaultForTest()
	defer cleanup()

	Info("pkg info")
	if !strings.Contains(buf.String(), "pkg info") {
		t.Fatalf("expected 'pkg info', got %q", buf.String())
	}
}

func TestPackageLevelInfof(t *testing.T) {
	buf, cleanup := setupDefaultForTest()
	defer cleanup()

	Infof("count: %d", 42)
	if !strings.Contains(buf.String(), "count: 42") {
		t.Fatalf("expected formatted output, got %q", buf.String())
	}
}

func TestPackageLevelDebug(t *testing.T) {
	buf, cleanup := setupDefaultForTest()
	defer cleanup()

	Debug("pkg debug")
	if !strings.Contains(buf.String(), "pkg debug") {
		t.Fatalf("expected 'pkg debug', got %q", buf.String())
	}
}

func TestPackageLevelDebugf(t *testing.T) {
	buf, cleanup := setupDefaultForTest()
	defer cleanup()

	Debugf("val: %d", 7)
	if !strings.Contains(buf.String(), "val: 7") {
		t.Fatalf("expected formatted output, got %q", buf.String())
	}
}

func TestPackageLevelTrace(t *testing.T) {
	buf, cleanup := setupDefaultForTest()
	defer cleanup()

	Trace("pkg trace")
	if !strings.Contains(buf.String(), "pkg trace") {
		t.Fatalf("expected 'pkg trace', got %q", buf.String())
	}
}

func TestPackageLevelTracef(t *testing.T) {
	buf, cleanup := setupDefaultForTest()
	defer cleanup()

	Tracef("t: %s", "x")
	if !strings.Contains(buf.String(), "t: x") {
		t.Fatalf("expected formatted output, got %q", buf.String())
	}
}

func TestPackageLevelWarn(t *testing.T) {
	buf, cleanup := setupDefaultForTest()
	defer cleanup()

	Warn("pkg warn")
	if !strings.Contains(buf.String(), "pkg warn") {
		t.Fatalf("expected 'pkg warn', got %q", buf.String())
	}
}

func TestPackageLevelWarnf(t *testing.T) {
	buf, cleanup := setupDefaultForTest()
	defer cleanup()

	Warnf("w: %d", 1)
	if !strings.Contains(buf.String(), "w: 1") {
		t.Fatalf("expected formatted output, got %q", buf.String())
	}
}

func TestPackageLevelError(t *testing.T) {
	buf, cleanup := setupDefaultForTest()
	defer cleanup()

	Error("pkg error")
	if !strings.Contains(buf.String(), "pkg error") {
		t.Fatalf("expected 'pkg error', got %q", buf.String())
	}
}

func TestPackageLevelErrorf(t *testing.T) {
	buf, cleanup := setupDefaultForTest()
	defer cleanup()

	Errorf("e: %s", "fail")
	if !strings.Contains(buf.String(), "e: fail") {
		t.Fatalf("expected formatted output, got %q", buf.String())
	}
}

func TestPackageLevelFatal(t *testing.T) {
	buf, cleanup := setupDefaultForTest()
	defer cleanup()

	var exitCode int
	Default().exitFunc = func(code int) { exitCode = code }

	Fatal("pkg fatal")
	if !strings.Contains(buf.String(), "pkg fatal") {
		t.Fatalf("expected 'pkg fatal', got %q", buf.String())
	}
	if exitCode != 1 {
		t.Fatalf("expected exit code 1, got %d", exitCode)
	}
}

func TestPackageLevelFatalf(t *testing.T) {
	buf, cleanup := setupDefaultForTest()
	defer cleanup()

	var exitCode int
	Default().exitFunc = func(code int) { exitCode = code }

	Fatalf("fatal: %d", 99)
	if !strings.Contains(buf.String(), "fatal: 99") {
		t.Fatalf("expected formatted output, got %q", buf.String())
	}
	if exitCode != 1 {
		t.Fatalf("expected exit code 1, got %d", exitCode)
	}
}

func TestPackageLevelSuccess(t *testing.T) {
	buf, cleanup := setupDefaultForTest()
	defer cleanup()

	Success("pkg success")
	if !strings.Contains(buf.String(), "pkg success") {
		t.Fatalf("expected 'pkg success', got %q", buf.String())
	}
}

func TestPackageLevelSuccessf(t *testing.T) {
	buf, cleanup := setupDefaultForTest()
	defer cleanup()

	Successf("done: %d", 3)
	if !strings.Contains(buf.String(), "done: 3") {
		t.Fatalf("expected formatted output, got %q", buf.String())
	}
}

func TestPackageLevelSetPrefix(t *testing.T) {
	buf, cleanup := setupDefaultForTest()
	defer cleanup()

	SetPrefix("->")
	Info("test")
	if !strings.HasPrefix(buf.String(), "-> ") {
		t.Fatalf("expected '-> ' prefix, got %q", buf.String())
	}
}

func TestPackageLevelClearPrefix(t *testing.T) {
	buf, cleanup := setupDefaultForTest()
	defer cleanup()

	ClearPrefix()
	Info("no prefix")
	got := strings.TrimSpace(buf.String())
	if got != "no prefix" {
		t.Fatalf("expected 'no prefix' without prefix, got %q", got)
	}
}

func TestPackageLevelSetTheme(t *testing.T) {
	buf, cleanup := setupDefaultForTest()
	defer cleanup()

	Default().SetColorEnabled(true)
	SetTheme(ThemeDracula)
	Info("themed")
	// Dracula info color is #F8F8F2 = RGB(248,248,242)
	if !strings.Contains(buf.String(), "38;2;248;248;242") {
		t.Fatalf("expected Dracula info color, got %q", buf.String())
	}
}

func TestPackageLevelSetPrefixColor(t *testing.T) {
	buf, cleanup := setupDefaultForTest()
	defer cleanup()

	Default().SetColorEnabled(true)
	SetPrefixColor(ColorRed)
	Info("test")
	if !strings.Contains(buf.String(), "\033[31m") {
		t.Fatalf("expected red ANSI code for prefix, got %q", buf.String())
	}
}

func TestPackageLevelSetMessageColor(t *testing.T) {
	buf, cleanup := setupDefaultForTest()
	defer cleanup()

	Default().SetColorEnabled(true)
	SetMessageColor(ColorGreen)
	Info("test")
	if !strings.Contains(buf.String(), "\033[32mtest\033[0m") {
		t.Fatalf("expected green ANSI for message, got %q", buf.String())
	}
}

func TestPackageLevelSetColorEnabled(t *testing.T) {
	buf, cleanup := setupDefaultForTest()
	defer cleanup()

	SetColorEnabled(true)
	SetTheme(ThemeDracula)
	Info("colored")
	if !strings.Contains(buf.String(), "\033[") {
		t.Fatalf("expected ANSI codes when color enabled, got %q", buf.String())
	}

	buf.Reset()
	SetColorEnabled(false)
	Info("plain")
	if strings.Contains(buf.String(), "\033[") {
		t.Fatalf("expected no ANSI codes when color disabled, got %q", buf.String())
	}
}

// --- Theme existence tests (ensure all themes are defined and have non-empty names) ---

func TestAllThemesHaveNames(t *testing.T) {
	themes := Themes()
	if len(themes) != 32 {
		t.Fatalf("expected 32 built-in themes, got %d", len(themes))
	}
	for _, theme := range themes {
		if theme.Name == "" {
			t.Error("theme has empty name")
		}
		// All themes should have non-default error and success colors
		if theme.ErrorColor.isDefault() {
			t.Errorf("theme %q has default error color", theme.Name)
		}
		if theme.SuccessColor.isDefault() {
			t.Errorf("theme %q has default success color", theme.Name)
		}
	}
}

func TestThemesReturnsDistinctNames(t *testing.T) {
	themes := Themes()
	seen := make(map[string]bool)
	for _, theme := range themes {
		if seen[theme.Name] {
			t.Errorf("duplicate theme name: %q", theme.Name)
		}
		seen[theme.Name] = true
	}
}

func TestThemesReturnsFreshSlice(t *testing.T) {
	a := Themes()
	b := Themes()
	a[0] = Theme{Name: "mutated"}
	if b[0].Name == "mutated" {
		t.Fatal("Themes() should return a fresh slice each time")
	}
}

func TestThemeByNameFound(t *testing.T) {
	theme, ok := ThemeByName("Dracula")
	if !ok {
		t.Fatal("expected to find Dracula theme")
	}
	if theme.Name != "Dracula" {
		t.Fatalf("expected name 'Dracula', got %q", theme.Name)
	}
}

func TestThemeByNameCaseInsensitive(t *testing.T) {
	cases := []string{"dracula", "DRACULA", "DrAcUlA", "monokai pro", "MONOKAI PRO"}
	for _, name := range cases {
		_, ok := ThemeByName(name)
		if !ok {
			t.Errorf("ThemeByName(%q) should find a match", name)
		}
	}
}

func TestThemeByNameNotFound(t *testing.T) {
	_, ok := ThemeByName("nonexistent")
	if ok {
		t.Fatal("expected no match for 'nonexistent'")
	}
}

// --- colorForLevel default branch ---

func TestColorForLevelUnknownFallsBackToInfo(t *testing.T) {
	o, buf := newTestOutput()
	o.SetColorEnabled(true)
	o.SetTheme(ThemeDracula)
	// Call print directly with an invalid level value to hit the default branch.
	// Level(99) is above Silent so it won't be filtered, and colorForLevel will
	// hit the default case.
	o.print(Level(99), "fallback", false)
	got := buf.String()
	// Should use InfoColor from Dracula: #F8F8F2 = RGB(248,248,242)
	if !strings.Contains(got, "38;2;248;248;242") {
		t.Fatalf("expected Dracula info color as fallback, got %q", got)
	}
}

// --- Multiple messages test ---

func TestMultipleMessages(t *testing.T) {
	o, buf := newTestOutput()
	o.Info("line 1")
	o.Info("line 2")
	o.Info("line 3")
	lines := strings.Split(strings.TrimSpace(buf.String()), "\n")
	if len(lines) != 3 {
		t.Fatalf("expected 3 lines, got %d: %q", len(lines), buf.String())
	}
}

// --- Colorize tests ---

func TestColorizeWithColorEnabled(t *testing.T) {
	o, _ := newTestOutput()
	o.SetColorEnabled(true)
	got := o.Colorize("hello", ColorRed)
	expected := "\033[31mhello\033[0m"
	if got != expected {
		t.Fatalf("expected %q, got %q", expected, got)
	}
}

func TestColorizeWithColorDisabled(t *testing.T) {
	o, _ := newTestOutput()
	o.SetColorEnabled(false)
	got := o.Colorize("hello", ColorRed)
	if got != "hello" {
		t.Fatalf("expected plain text, got %q", got)
	}
}

func TestColorizeTrueColor(t *testing.T) {
	o, _ := newTestOutput()
	o.SetColorEnabled(true)
	got := o.Colorize("text", Hex("FF5733"))
	expected := "\033[38;2;255;87;51mtext\033[0m"
	if got != expected {
		t.Fatalf("expected %q, got %q", expected, got)
	}
}

func TestColorizeDefaultColor(t *testing.T) {
	o, _ := newTestOutput()
	o.SetColorEnabled(true)
	got := o.Colorize("text", ColorDefault)
	if got != "text" {
		t.Fatalf("expected plain text for default color, got %q", got)
	}
}

func TestColorizeComposedInInfof(t *testing.T) {
	o, buf := newTestOutput()
	o.SetColorEnabled(true)
	o.Infof("%s %s %s",
		o.Colorize("green", ColorGreen),
		"plain",
		o.Colorize("red", ColorRed),
	)
	got := buf.String()
	if !strings.Contains(got, "\033[32mgreen\033[0m") {
		t.Fatalf("expected green ANSI in output, got %q", got)
	}
	if !strings.Contains(got, "plain") {
		t.Fatalf("expected plain text in output, got %q", got)
	}
	if !strings.Contains(got, "\033[31mred\033[0m") {
		t.Fatalf("expected red ANSI in output, got %q", got)
	}
}

func TestPackageLevelColorize(t *testing.T) {
	// Just verify the package-level function doesn't panic and returns a string.
	got := Colorize("test", ColorBlue)
	// The default output may or may not have color enabled depending on the
	// test environment, so just verify it contains the original text.
	if !strings.Contains(got, "test") {
		t.Fatalf("expected result to contain 'test', got %q", got)
	}
}

// --- CLI_THEME environment variable tests ---

func TestNewRespectsCliThemeEnv(t *testing.T) {
	t.Setenv("CLI_THEME", "Dracula")

	o := New()
	if o.theme.Name != "Dracula" {
		t.Fatalf("expected theme 'Dracula', got %q", o.theme.Name)
	}
}

func TestNewCliThemeEnvCaseInsensitive(t *testing.T) {
	t.Setenv("CLI_THEME", "tokyo night storm")

	o := New()
	if o.theme.Name != "Tokyo Night Storm" {
		t.Fatalf("expected theme 'Tokyo Night Storm', got %q", o.theme.Name)
	}
}

func TestNewCliThemeEnvUnknownFallsBackToDefault(t *testing.T) {
	t.Setenv("CLI_THEME", "nonexistent-theme")

	o := New()
	if o.theme.Name != ThemeDefault.Name {
		t.Fatalf("expected default theme, got %q", o.theme.Name)
	}
}

func TestNewCliThemeEnvEmptyFallsBackToDefault(t *testing.T) {
	t.Setenv("CLI_THEME", "")

	o := New()
	if o.theme.Name != ThemeDefault.Name {
		t.Fatalf("expected default theme, got %q", o.theme.Name)
	}
}

func TestNewCliThemeEnvUnsetUsesDefault(t *testing.T) {
	t.Setenv("CLI_THEME", "")
	os.Unsetenv("CLI_THEME") //nolint:errcheck // test cleanup handled by t.Setenv

	o := New()
	if o.theme.Name != ThemeDefault.Name {
		t.Fatalf("expected default theme, got %q", o.theme.Name)
	}
}

func TestNewCliThemeWithNoColorStillDisablesColor(t *testing.T) {
	t.Setenv("CLI_THEME", "Dracula")
	t.Setenv("NO_COLOR", "1")

	o := New()
	if o.theme.Name != "Dracula" {
		t.Fatalf("expected theme 'Dracula', got %q", o.theme.Name)
	}
	if o.colorEnabled {
		t.Fatal("expected color to be disabled when NO_COLOR is set")
	}
}

// --- CLI_PREFIX environment variable tests ---

func TestNewRespectsCliPrefixEnv(t *testing.T) {
	t.Setenv("CLI_PREFIX", "->")

	o := New()
	if o.prefix != "->" {
		t.Fatalf("expected prefix '->', got %q", o.prefix)
	}
	if !o.hasPrefix {
		t.Fatal("expected hasPrefix to be true")
	}
}

func TestNewCliPrefixEnvEmptyClearsPrefix(t *testing.T) {
	t.Setenv("CLI_PREFIX", "")

	o := New()
	if o.prefix != "" {
		t.Fatalf("expected empty prefix, got %q", o.prefix)
	}
	if o.hasPrefix {
		t.Fatal("expected hasPrefix to be false when CLI_PREFIX is empty")
	}
}

func TestNewCliPrefixEnvUnsetUsesDefault(t *testing.T) {
	t.Setenv("CLI_PREFIX", "")
	os.Unsetenv("CLI_PREFIX") //nolint:errcheck // test cleanup handled by t.Setenv

	o := New()
	if o.prefix != defaultPrefix {
		t.Fatalf("expected default prefix %q, got %q", defaultPrefix, o.prefix)
	}
	if !o.hasPrefix {
		t.Fatal("expected hasPrefix to be true")
	}
}

func TestNewSetPrefixOverridesCliPrefixEnv(t *testing.T) {
	t.Setenv("CLI_PREFIX", "->")

	o := New()
	if o.prefix != "->" {
		t.Fatalf("expected initial prefix '->', got %q", o.prefix)
	}

	o.SetPrefix("***")
	if o.prefix != "***" {
		t.Fatalf("expected prefix '***' after SetPrefix, got %q", o.prefix)
	}
}

func TestNewCliPrefixEnvOutput(t *testing.T) {
	t.Setenv("CLI_PREFIX", "::")

	o := New()
	var buf bytes.Buffer
	o.SetWriter(&buf)
	o.SetColorEnabled(false)
	o.SetLevel(LevelTrace)
	o.Info("hello")

	got := buf.String()
	if !strings.HasPrefix(got, ":: ") {
		t.Fatalf("expected output to start with ':: ', got %q", got)
	}
}

func TestNewCliPrefixEnvEmptyOutput(t *testing.T) {
	t.Setenv("CLI_PREFIX", "")

	o := New()
	var buf bytes.Buffer
	o.SetWriter(&buf)
	o.SetColorEnabled(false)
	o.SetLevel(LevelTrace)
	o.Info("no prefix")

	got := strings.TrimSpace(buf.String())
	if got != "no prefix" {
		t.Fatalf("expected 'no prefix' without any prefix, got %q", got)
	}
}

func TestNewCliPrefixWithCliTheme(t *testing.T) {
	t.Setenv("CLI_PREFIX", "=>")
	t.Setenv("CLI_THEME", "Dracula")

	o := New()
	if o.prefix != "=>" {
		t.Fatalf("expected prefix '=>', got %q", o.prefix)
	}
	if o.theme.Name != "Dracula" {
		t.Fatalf("expected theme 'Dracula', got %q", o.theme.Name)
	}
}

func TestNewSetThemeOverridesCliThemeEnv(t *testing.T) {
	t.Setenv("CLI_THEME", "Dracula")

	o := New()
	if o.theme.Name != "Dracula" {
		t.Fatalf("expected initial theme 'Dracula', got %q", o.theme.Name)
	}

	o.SetTheme(ThemeNord)
	if o.theme.Name != "Nord" {
		t.Fatalf("expected theme 'Nord' after SetTheme, got %q", o.theme.Name)
	}
}

// --- NO_COLOR enforcement tests ---

func TestSetColorEnabledBlockedByNoColor(t *testing.T) {
	t.Setenv("NO_COLOR", "1")

	o := New()
	if o.colorEnabled {
		t.Fatal("expected color disabled when NO_COLOR is set")
	}

	// Attempt to re-enable color — should be blocked.
	o.SetColorEnabled(true)
	if o.colorEnabled {
		t.Fatal("SetColorEnabled(true) must not override NO_COLOR")
	}
}

func TestNoColorPreventsAnsiCodesAfterSetColorEnabled(t *testing.T) {
	t.Setenv("NO_COLOR", "1")

	o := New()
	var buf bytes.Buffer
	o.SetWriter(&buf)
	o.SetLevel(LevelTrace)
	o.SetTheme(ThemeDracula)

	// Try to force color on.
	o.SetColorEnabled(true)

	o.Info("should be plain")
	got := buf.String()
	if strings.Contains(got, "\033[") {
		t.Fatalf("expected no ANSI codes when NO_COLOR is set, got %q", got)
	}
}

func TestNoColorBlocksColorizeOutput(t *testing.T) {
	t.Setenv("NO_COLOR", "1")

	o := New()
	o.SetColorEnabled(true) // should be ignored

	got := o.Colorize("text", ColorRed)
	if strings.Contains(got, "\033[") {
		t.Fatalf("Colorize must not emit ANSI codes when NO_COLOR is set, got %q", got)
	}
	if got != "text" {
		t.Fatalf("expected plain 'text', got %q", got)
	}
}

func TestNoColorEmptyValueStillDisables(t *testing.T) {
	t.Setenv("NO_COLOR", "")

	o := New()
	o.SetColorEnabled(true)

	if o.colorEnabled {
		t.Fatal("NO_COLOR with empty value should still disable color")
	}
}

func TestNoColorWithThemeAndPrefixProducesPlainOutput(t *testing.T) {
	t.Setenv("NO_COLOR", "1")
	t.Setenv("CLI_THEME", "Dracula")
	t.Setenv("CLI_PREFIX", "::")

	o := New()
	var buf bytes.Buffer
	o.SetWriter(&buf)
	o.SetLevel(LevelTrace)

	// Try every output method.
	o.Trace("t")
	o.Debug("d")
	o.Info("i")
	o.Warn("w")
	o.Error("e")
	o.Success("s")
	o.Infof("f %d", 1)

	got := buf.String()
	if strings.Contains(got, "\033[") {
		t.Fatalf("expected no ANSI codes anywhere in output when NO_COLOR is set, got %q", got)
	}
	// Verify prefix is still applied.
	if !strings.Contains(got, ":: i") {
		t.Fatalf("expected CLI_PREFIX '::' in output, got %q", got)
	}
}

func TestSetColorEnabledWorksWithoutNoColor(t *testing.T) {
	// Verify that SetColorEnabled still works normally when NO_COLOR is not set.
	o, buf := newTestOutput()
	o.SetTheme(ThemeDracula)

	o.SetColorEnabled(true)
	o.Info("colored")
	got := buf.String()
	if !strings.Contains(got, "\033[") {
		t.Fatalf("expected ANSI codes when color is enabled and NO_COLOR is not set, got %q", got)
	}

	buf.Reset()
	o.SetColorEnabled(false)
	o.Info("plain")
	got = buf.String()
	if strings.Contains(got, "\033[") {
		t.Fatalf("expected no ANSI codes after disabling color, got %q", got)
	}
}
