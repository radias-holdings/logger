package logger

import (
	"bytes"
	"context"
	"os"
	"testing"

	"log/slog"

	c "github.com/smartystreets/goconvey/convey"
)

const (
	Debug = -4
)

func TestCreateLogger(t *testing.T) {
	ctx := context.Background()

	c.Convey("When the logger is created", t, func() {
		var buf bytes.Buffer
		logger := NewLogger(&buf, Debug) // Set to debug level

		c.Convey("Then the logger should not be nil", func() {
			c.So(logger, c.ShouldNotBeNil)
		})

		c.Convey("Then the log should be ERROR if in production", func() {
			previous := os.Getenv("ENV")
			os.Setenv("ENV", "prod")

			logger.Debug("This log shouldn't appear")
			logger.Error("This log should appear")
			output := buf.String()

			c.So(output, c.ShouldContainSubstring, "ERROR")
			os.Setenv("ENV", previous)
		})

		c.Convey("Then the log should be DEBUG if in development", func() {
			previous := os.Getenv("ENV")
			os.Setenv("ENV", "dev")
			logger.Debug("This log should appear")
			output := buf.String()

			c.So(output, c.ShouldContainSubstring, "DEBUG")
			os.Setenv("ENV", previous)
		})

		c.Convey("When a fatal log is sent", func() {
			logger.Log(ctx, LevelFatal, "test")

			c.Convey("Then the log should be written properly", func() {
				output := buf.String()

				c.So(output, c.ShouldContainSubstring, "FATAL")
				c.So(output, c.ShouldContainSubstring, "test")
			})
		})

		c.Convey("When a debug log is sent", func() {
			logger.Debug("debugging")

			c.Convey("Then the log should be written properly", func() {
				output := buf.String()

				c.So(output, c.ShouldContainSubstring, "DEBUG")
				c.So(output, c.ShouldContainSubstring, "debugging")
			})
		})

		c.Convey("When an unknown log level is sent", func() {
			logger.Log(ctx, slog.Level(100), "unknown")

			c.Convey("Then the log should default to the error level", func() {
				output := buf.String()

				c.So(output, c.ShouldContainSubstring, "92") // Default error level is 8
				c.So(output, c.ShouldContainSubstring, "unknown")
			})
		})
	})
}
