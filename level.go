package cliout

// Level represents the verbosity level of output.
type Level int

const (
	// LevelTrace is the most verbose level, for detailed tracing.
	LevelTrace Level = iota
	// LevelDebug is for debug messages.
	LevelDebug
	// LevelInfo is the default level for informational messages.
	LevelInfo
	// LevelWarn is for warning messages.
	LevelWarn
	// LevelError is for error messages.
	LevelError
	// LevelSilent suppresses all output.
	LevelSilent
)

// String returns a human-readable name for the level.
func (l Level) String() string {
	switch l {
	case LevelTrace:
		return "trace"
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warn"
	case LevelError:
		return "error"
	case LevelSilent:
		return "silent"
	default:
		return "unknown"
	}
}
