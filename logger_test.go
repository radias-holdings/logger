package logger

import (
	"bytes"
	"context"
	"testing"

	"log/slog"

	c "github.com/smartystreets/goconvey/convey"
)

func TestCreateLogger(t *testing.T) {
	ctx := context.Background()

	c.Convey("Logger initialization", t, func() {

		c.Convey("For default or production environment", func() {
			buf := &bytes.Buffer{}
			logger := NewLogger(buf)

			c.So(logger, c.ShouldNotBeNil)

			c.Convey("Should log at ERROR level", func() {
				logger.Debug("This log shouldn't appear")
				logger.Error("This log should appear")
				output := buf.String()

				c.So(output, c.ShouldContainSubstring, "ERROR")
			})

			c.Convey("Should log FATAL properly", func() {
				logger.Log(ctx, LevelFatal, "test")
				output := buf.String()

				c.So(output, c.ShouldContainSubstring, "FATAL")
				c.So(output, c.ShouldContainSubstring, "test")
			})

			c.Convey("Unknown log level should default to error level", func() {
				logger.Log(ctx, slog.Level(100), "unknown")
				output := buf.String()

				c.So(output, c.ShouldContainSubstring, "92")
				c.So(output, c.ShouldContainSubstring, "unknown")
			})

			buf.Reset()
		})

		c.Convey("For debug environments (development, test or github pipeline)", func() {
			for _, env := range []string{"development", "test", "github"} {

				buf := &bytes.Buffer{}
				logger := NewLogger(buf, env)

				c.So(logger, c.ShouldNotBeNil)

				c.Convey("Should log at DEBUG level for "+env, func() {
					logger.Debug("This log should appear")
					output := buf.String()

					c.So(output, c.ShouldContainSubstring, "DEBUG")
				})

				c.Convey("Debug log should be written properly for "+env, func() {
					logger.Debug("debugging")
					output := buf.String()

					c.So(output, c.ShouldContainSubstring, "DEBUG")
					c.So(output, c.ShouldContainSubstring, "debugging")
				})

				buf.Reset()
			}
		})
	})
}
