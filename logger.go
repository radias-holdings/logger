package logger

import (
	"io"
	"log/slog"
	"os"
	"strings"
)

const (
	LevelFatal = slog.Level(12)
	LevelError = slog.Level(8)
)

func init() {
	NewLogger(nil)
}

// NewLogger takes a writer and an optional ENV value. It returns a new logger.
// The default writer is os.Stdout if nil and the logging level is ERROR if not defined.
// Logging levels: DEBUG = -4, INFO = 0, WARN = 4, ERROR = 8, FATAL = 12.
// Environment levels: development, test and github = DEBUG, production = ERROR (default)
func NewLogger(writer io.Writer, env ...string) (logger *slog.Logger) {
	debugEnvs := []string{"DEVELOPMENT", "TEST", "GITHUB"}
	logLevel := LevelError

	if writer == nil {
		writer = os.Stdout
	}
	if len(env) > 0 {
		setLevel := strings.ToUpper(env[0])
		for _, env := range debugEnvs {
			if setLevel == env {
				logLevel = slog.Level(-4)
				break
			}
		}
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
