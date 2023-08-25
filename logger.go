package logger

import (
	"io"
	"log/slog"
	"os"
	"strings"
)

const (
	Debug = slog.Level(-4)
	Info  = slog.Level(0)
	Warn  = slog.Level(4)
	Error = slog.Level(8)
	Fatal = slog.Level(12)
)

// NewLogger takes a writer and an optional ENV value. It returns a new logger.
// The default writer is os.Stdout if nil and the logging level is ERROR if not defined.
// Logging levels: DEBUG = -4, INFO = 0, WARN = 4, ERROR = 8, FATAL = 12.
func NewLogger(writer io.Writer, level ...string) (logger *slog.Logger) {
	logLevel := Error
	if len(level) > 0 {
		setLevel := strings.ToUpper(level[0])
		switch setLevel {
		case "DEBUG":
			logLevel = Debug
		case "INFO":
			logLevel = Info
		case "WARN":
			logLevel = Warn
		case "ERROR":
			logLevel = Error
		case "FATAL":
			logLevel = Fatal
		}
	}
	if writer == nil {
		writer = os.Stdout
	}

	opts := &slog.HandlerOptions{
		Level: logLevel,
		// Replaces "ERROR+4" with "FATAL" in log output
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.LevelKey {
				level := a.Value.Any().(slog.Level)
				levelLabel := level.String()
				if level == Fatal {
					levelLabel = "FATAL"
				}
				a.Value = slog.StringValue(levelLabel)
			}
			return a
		},
	}

	return slog.New(slog.NewTextHandler(writer, opts))
}
