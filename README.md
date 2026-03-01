# cliout

A Go package for printing formatted, coloured output to the command line. Not a logger -- just clean, user-facing CLI output with themes, levels, and customisable prefixes.

```
» configuring somoflange...
» updated settings
» done
```

## Features

- Levelled output: Trace, Debug, Info, Warn, Error, Success
- 32 built-in colour themes (Dracula, Nord, Monokai Pro, Catppuccino, Tokyo Night, and more)
- Customisable prefix character and colours
- True colour (24-bit RGB) and standard ANSI colour support
- Respects [`NO_COLOR`](https://no-color.org/) and `CLI_THEME` environment variables, auto-detects TTY
- Format string variants (`Infof`, `Debugf`, etc.)
- Zero external dependencies
- Go 1.20+

## Install

```sh
go get github.com/z0mbix/cliout
```

## Quick Start

```go
package main

import "github.com/z0mbix/cliout"

func main() {
    cliout.Info("configuring somoflange...")
    cliout.Info("updated settings")
    cliout.Info("done")
}
```

Output:

```
» configuring somoflange...
» updated settings
» done
```

## Output Levels

Messages are printed only if their level is at or above the configured minimum level. The default level is `LevelInfo`.

```
LevelTrace < LevelDebug < LevelInfo < LevelWarn < LevelError < LevelSilent
```

```go
package main

import "github.com/z0mbix/cliout"

func main() {
    cliout.Info("this is visible by default")
    cliout.Debug("this is hidden by default")

    // Lower the level to see debug messages
    cliout.SetLevel(cliout.LevelDebug)
    cliout.Debug("now this is visible")

    // Show everything
    cliout.SetLevel(cliout.LevelTrace)
    cliout.Trace("detailed trace output")

    // Suppress all output
    cliout.SetLevel(cliout.LevelSilent)
    cliout.Error("even errors are hidden")
}
```

### All Output Methods

Each level has a plain and a format-string variant:

```go
cliout.Trace("message")
cliout.Tracef("value: %d", 42)

cliout.Debug("message")
cliout.Debugf("value: %d", 42)

cliout.Info("message")
cliout.Infof("count: %d, name: %s", 100, "test")

cliout.Warn("message")
cliout.Warnf("disk usage: %d%%", 95)

cliout.Error("message")
cliout.Errorf("failed to connect: %v", err)

cliout.Success("message")
cliout.Successf("deployed %d services", 3)
```

`Success` prints at the Info level but uses the theme's dedicated success colour.

## Prefix Customisation

The default prefix is `»`. You can change it, or remove it entirely.

### Custom Prefix

```go
package main

import "github.com/z0mbix/cliout"

func main() {
    cliout.SetPrefix("->")
    cliout.Info("using arrow prefix")

    cliout.SetPrefix("*")
    cliout.Info("using asterisk prefix")

    cliout.SetPrefix("[myapp]")
    cliout.Info("using app name prefix")
}
```

Output:

```
-> using arrow prefix
* using asterisk prefix
[myapp] using app name prefix
```

### No Prefix

```go
package main

import "github.com/z0mbix/cliout"

func main() {
    cliout.ClearPrefix()
    cliout.Info("no prefix at all")
}
```

Output:

```
no prefix at all
```

### Restore Prefix After Clearing

```go
cliout.ClearPrefix()
cliout.Info("no prefix")

cliout.SetPrefix("»")
cliout.Info("prefix is back")
```

## Colours

### Built-in ANSI Colours

Standard and bright variants are available as package-level constants:

| Standard | Bright |
|---|---|
| `ColorBlack` | `ColorBrightBlack` |
| `ColorRed` | `ColorBrightRed` |
| `ColorGreen` | `ColorBrightGreen` |
| `ColorYellow` | `ColorBrightYellow` |
| `ColorBlue` | `ColorBrightBlue` |
| `ColorMagenta` | `ColorBrightMagenta` |
| `ColorCyan` | `ColorBrightCyan` |
| `ColorWhite` | `ColorBrightWhite` |

`ColorDefault` leaves the terminal's default colour unchanged.

### Custom Colours

Create true colours (24-bit) from hex strings or RGB values:

```go
// From hex (with or without #)
coral := cliout.Hex("#FF7F50")
teal := cliout.Hex("008080")

// From RGB components
purple := cliout.RGB(128, 0, 255)
```

### Setting Prefix Colour

```go
package main

import "github.com/z0mbix/cliout"

func main() {
    cliout.SetPrefixColor(cliout.ColorGreen)
    cliout.Info("green prefix")

    cliout.SetPrefixColor(cliout.Hex("#FF5733"))
    cliout.Info("custom orange prefix")
}
```

### Setting Message Colour

When set, this overrides the per-level theme colours for all messages:

```go
package main

import "github.com/z0mbix/cliout"

func main() {
    cliout.SetPrefixColor(cliout.ColorGreen)
    cliout.SetMessageColor(cliout.ColorWhite)
    cliout.Info("green prefix, white message")
    cliout.Warn("also white, even though it's a warning")
}
```

### Disabling Colour

Colour is automatically disabled when:

1. The `NO_COLOR` environment variable is set (any value)
2. stdout is not a TTY (e.g., piped to a file)

You can also disable it programmatically:

```go
cliout.SetColorEnabled(false)
cliout.Info("plain text, no ANSI codes")
```

### `CLI_THEME` Environment Variable

Users can set the `CLI_THEME` environment variable to select a theme across all tools that use cliout. The value is matched case-insensitively against the built-in theme names:

```sh
export CLI_THEME=dracula
```

Any tool using cliout will now default to Dracula without any code changes. The lookup happens in `New()`, so both the package-level default instance and any instances created with `cliout.New()` pick it up.

If `CLI_THEME` is unset, empty, or doesn't match a built-in theme, the Default theme is used. `SetTheme()` still overrides `CLI_THEME` if the application needs to.

`CLI_THEME` and `NO_COLOR` work together -- the theme is applied but colour output is disabled:

```sh
CLI_THEME=dracula NO_COLOR=1 ./mytool   # Dracula theme colours, but no ANSI codes emitted
```

### Multi-Colour Lines

Use `Colorize` to wrap individual text segments with colour, then compose them with `Infof`, `Errorf`, etc.:

```go
package main

import "github.com/z0mbix/cliout"

func main() {
    cliout.Infof("%s %s - %s",
        cliout.Colorize("ok", cliout.ColorGreen),
        cliout.Colorize("some blue text", cliout.ColorBlue),
        cliout.Colorize("some red text", cliout.ColorRed),
    )
}
```

Output (with colour):

```
» ok some blue text - some red text
```

`Colorize` respects the output's colour-enabled setting -- when colour is off (e.g., `NO_COLOR`, non-TTY, or `SetColorEnabled(false)`), it returns the text unchanged. It also works on `*Output` instances:

```go
out := cliout.New()
out.Infof("status: %s", out.Colorize("healthy", cliout.Hex("#00FF00")))
```

## Themes

Themes set the prefix colour and per-level message colours in one call. Setting a theme does not override explicit colour overrides from `SetPrefixColor` or `SetMessageColor`.

### Using a Theme

```go
package main

import "github.com/z0mbix/cliout"

func main() {
    cliout.SetTheme(cliout.ThemeDracula)
    cliout.Info("styled with Dracula colours")
    cliout.Warn("Dracula warning colour")
    cliout.Error("Dracula error colour")
    cliout.Success("Dracula success colour")
}
```

### Available Themes

| Theme | Variable |
|---|---|
| Default | `ThemeDefault` |
| Ayu | `ThemeAyu` |
| Ayu Light | `ThemeAyuLight` |
| Ayu Mirage | `ThemeAyuMirage` |
| Dracula | `ThemeDracula` |
| One Dark | `ThemeOneDark` |
| Solarized Dark | `ThemeSolarizedDark` |
| Solarized Light | `ThemeSolarizedLight` |
| Nord | `ThemeNord` |
| Gruvbox Dark | `ThemeGruvboxDark` |
| Gruvbox Light | `ThemeGruvboxLight` |
| Monokai | `ThemeMonokai` |
| Monokai Pro | `ThemeMonokaiPro` |
| Monokai Pro Classic | `ThemeMonokaiProClassic` |
| Monokai Pro Machine | `ThemeMonokaiProMachine` |
| Monokai Pro Octagon | `ThemeMonokaiProOctagon` |
| Monokai Pro Ristretto | `ThemeMonokaiProRistretto` |
| Monokai Pro Spectrum | `ThemeMonokaiProSpectrum` |
| Monokai Pro Light | `ThemeMonokaiProLight` |
| Material Dark | `ThemeMaterialDark` |
| Material Light | `ThemeMaterialLight` |
| Palenight | `ThemePalenight` |
| Catppuccino Frappe | `ThemeCatppuccinoFrappe` |
| Catppuccino Latte | `ThemeCatppuccinoLatte` |
| Catppuccino Macchiato | `ThemeCatppuccinoMacchiato` |
| Catppuccino Mocha | `ThemeCatppuccinoMocha` |
| Rose Pine | `ThemeRosePine` |
| Rose Pine Dawn | `ThemeRosePineDawn` |
| Rose Pine Moon | `ThemeRosePineMoon` |
| Tokyo Night Storm | `ThemeTokyoNightStorm` |
| Tokyo Night Day | `ThemeTokyoNightDay` |
| Tokyo Night Night | `ThemeTokyoNightNight` |

### Listing and Looking Up Themes

Use `Themes()` to get all built-in themes, and `ThemeByName()` to look one up by name (case-insensitive):

```go
package main

import (
    "fmt"

    "github.com/z0mbix/cliout"
)

func main() {
    // List all available theme names
    for _, t := range cliout.Themes() {
        fmt.Println(t.Name)
    }

    // Look up a theme by name (case-insensitive)
    theme, ok := cliout.ThemeByName("dracula")
    if ok {
        cliout.SetTheme(theme)
        cliout.Info("using Dracula theme")
    }
}
```

This is useful for letting users pick a theme via a flag or config file:

```go
themeName := flag.String("theme", "Default", "colour theme")
flag.Parse()

theme, ok := cliout.ThemeByName(*themeName)
if !ok {
    cliout.Errorf("unknown theme: %s", *themeName)
    os.Exit(1)
}
cliout.SetTheme(theme)
```

### Custom Themes

Create your own theme by defining a `Theme` struct:

```go
package main

import "github.com/z0mbix/cliout"

func main() {
    myTheme := cliout.Theme{
        Name:         "My Custom Theme",
        PrefixColor:  cliout.Hex("#E06C75"),
        InfoColor:    cliout.Hex("#ABB2BF"),
        DebugColor:   cliout.Hex("#5C6370"),
        TraceColor:   cliout.Hex("#5C6370"),
        WarnColor:    cliout.Hex("#E5C07B"),
        ErrorColor:   cliout.Hex("#FF0000"),
        SuccessColor: cliout.Hex("#98C379"),
    }

    cliout.SetTheme(myTheme)
    cliout.Info("custom theme in action")
    cliout.Error("custom error colour")
}
```

### Theme + Colour Overrides

Explicit colour overrides take priority over themes:

```go
cliout.SetTheme(cliout.ThemeNord)

// Override just the prefix colour, keep Nord's level colours
cliout.SetPrefixColor(cliout.ColorRed)
cliout.Info("red prefix, Nord info colour")
cliout.Error("red prefix, Nord error colour")
```

## Using Separate Instances

The package-level functions use a shared default instance. For independent configurations (e.g., in libraries or concurrent tools), create your own:

```go
package main

import (
    "os"

    "github.com/z0mbix/cliout"
)

func main() {
    // Create a custom instance
    out := cliout.New()
    out.SetTheme(cliout.ThemeTokyoNightStorm)
    out.SetPrefix(">>>")
    out.SetLevel(cliout.LevelDebug)
    out.Info("custom instance")
    out.Debug("debug is visible here")

    // Write to stderr instead of stdout
    errOut := cliout.New()
    errOut.SetWriter(os.Stderr)
    errOut.SetPrefix("!!")
    errOut.Error("this goes to stderr")

    // Access the default instance directly
    d := cliout.Default()
    d.SetTheme(cliout.ThemeDracula)
    d.Info("same as calling cliout.Info()")
}
```

## Custom Output Writer

By default, output goes to `os.Stdout`. You can redirect to any `io.Writer`:

```go
package main

import (
    "bytes"
    "fmt"
    "os"

    "github.com/z0mbix/cliout"
)

func main() {
    out := cliout.New()

    // Write to a buffer (useful for testing)
    var buf bytes.Buffer
    out.SetWriter(&buf)
    out.SetColorEnabled(false)
    out.Info("captured output")
    fmt.Print(buf.String())

    // Write to stderr
    out.SetWriter(os.Stderr)
    out.Error("this goes to stderr")
}
```

## Complete Example

```go
package main

import "github.com/z0mbix/cliout"

func main() {
    // Configure output
    cliout.SetTheme(cliout.ThemeCatppuccinoMocha)
    cliout.SetPrefix(">>>")
    cliout.SetLevel(cliout.LevelDebug)

    // Informational output
    cliout.Info("starting deployment")
    cliout.Infof("deploying %d services to %s", 3, "production")

    // Debug details (visible because level is set to Debug)
    cliout.Debug("connecting to cluster")
    cliout.Debugf("using endpoint %s", "https://k8s.example.com")

    // Warnings and errors
    cliout.Warn("service 'cache' has high memory usage")
    cliout.Warnf("memory at %d%% capacity", 92)
    cliout.Error("service 'auth' failed health check")
    cliout.Errorf("connection timeout after %ds", 30)

    // Success
    cliout.Success("deployment complete")
    cliout.Successf("all %d services healthy", 3)
}
```

## Colour Priority

When determining which colour to use for output, the following priority applies (highest first):

1. **Explicit message colour** -- set via `SetMessageColor()`
2. **Success colour** -- theme's `SuccessColor`, used only by `Success()`/`Successf()`
3. **Theme level colour** -- theme's per-level colour (InfoColor, WarnColor, etc.)

For the prefix:

1. **Explicit prefix colour** -- set via `SetPrefixColor()`
2. **Theme prefix colour** -- theme's `PrefixColor`

## API Reference

### Types

| Type | Description |
|---|---|
| `Output` | Holds all configuration and provides output methods |
| `Level` | Output verbosity level (`LevelTrace` through `LevelSilent`) |
| `Color` | Terminal colour (ANSI or 24-bit true colour) |
| `Theme` | Colour definitions for prefix and each output level |

### Constructors

| Function | Description |
|---|---|
| `New()` | Create a new `Output` with default settings |
| `Default()` | Get the package-level default `Output` instance |
| `RGB(r, g, b)` | Create a true colour from RGB components (0-255) |
| `Hex(hex)` | Create a true colour from a hex string (`"#FF5733"` or `"FF5733"`) |
| `Themes()` | Return a slice of all built-in themes |
| `ThemeByName(name)` | Look up a built-in theme by name (case-insensitive) |

### Configuration

All configuration methods are available on both `*Output` instances and as package-level functions:

| Method | Description |
|---|---|
| `SetLevel(Level)` | Set the minimum output level |
| `SetPrefix(string)` | Set the prefix string |
| `ClearPrefix()` | Remove the prefix |
| `SetPrefixColor(Color)` | Set the prefix colour (overrides theme) |
| `SetMessageColor(Color)` | Set the message colour for all levels (overrides theme) |
| `SetTheme(Theme)` | Set the colour theme |
| `SetColorEnabled(bool)` | Enable or disable colour output |
| `SetWriter(io.Writer)` | Set the output destination (instance method only) |

### Output Methods

| Method | Level | Description |
|---|---|---|
| `Trace(msg)` / `Tracef(fmt, ...)` | Trace | Most verbose, for detailed tracing |
| `Debug(msg)` / `Debugf(fmt, ...)` | Debug | Debugging information |
| `Info(msg)` / `Infof(fmt, ...)` | Info | General informational messages |
| `Warn(msg)` / `Warnf(fmt, ...)` | Warn | Warning messages |
| `Error(msg)` / `Errorf(fmt, ...)` | Error | Error messages |
| `Success(msg)` / `Successf(fmt, ...)` | Info | Success messages (uses theme's SuccessColor) |

### Colour Helpers

| Method | Description |
|---|---|
| `Colorize(text, Color)` | Wrap text with colour codes (respects colour-enabled setting) |

### Environment Variables

| Variable | Description |
|---|---|
| `NO_COLOR` | Disable all colour output when set (any value). See [no-color.org](https://no-color.org/) |
| `CLI_THEME` | Set the default theme by name (case-insensitive). Ignored if unset, empty, or unrecognised |

## License

MIT
