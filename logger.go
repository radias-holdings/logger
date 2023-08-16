package logger

import (
	"io"
	"log/slog"
	"os"
)

const (
	LevelFatal = slog.Level(12)
)

func init() {
	NewLogger(nil)
}

// NewLogger takes a writer and an optional logging level. It returns a new logger.
// The default writer is os.Stdout if nil and the logging level is ERROR if not defined.
// Logging levels: DEBUG = -4, INFO = 0, WARN = 4, ERROR = 8, FATAL = 12.
func NewLogger(writer io.Writer, level ...slog.Level) (logger *slog.Logger) {
	logLevel := slog.Level(8)

	if writer == nil {
		writer = os.Stdout
	}
	if len(level) > 0 {
		logLevel = level[0]
	}

	opts := &slog.HandlerOptions{
		Level: logLevel,
		// Replaces "ERROR+4" with "FATAL" in log output
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.LevelKey {
				level := a.Value.Any().(slog.Level)
				levelLabel := level.String()
				if level == LevelFatal {
					levelLabel = "FATAL"
				}
				a.Value = slog.StringValue(levelLabel)
			}
			return a
		},
	}

	return slog.New(slog.NewTextHandler(writer, opts))
}
