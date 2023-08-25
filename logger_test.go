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

		c.Convey("For default (no level specified)", func() {
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
				logger.Log(ctx, Fatal, "test")
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

		c.Convey("For DEBUG log level", func() {
			buf := &bytes.Buffer{}
			logger := NewLogger(buf, "DEBUG")

			c.So(logger, c.ShouldNotBeNil)

			logger.Debug("Debug message")
			output := buf.String()
			c.So(output, c.ShouldContainSubstring, "DEBUG")

			buf.Reset()
		})

		c.Convey("For INFO log level", func() {
			buf := &bytes.Buffer{}
			logger := NewLogger(buf, "INFO")

			c.So(logger, c.ShouldNotBeNil)

			logger.Info("Info message")
			output := buf.String()
			c.So(output, c.ShouldContainSubstring, "INFO")

			buf.Reset()
		})

		c.Convey("For WARN log level", func() {
			buf := &bytes.Buffer{}
			logger := NewLogger(buf, "WARN")

			c.So(logger, c.ShouldNotBeNil)

			logger.Warn("Warn message")
			output := buf.String()
			c.So(output, c.ShouldContainSubstring, "WARN")

			buf.Reset()
		})

		c.Convey("For ERROR log level", func() {
			buf := &bytes.Buffer{}
			logger := NewLogger(buf, "ERROR")

			c.So(logger, c.ShouldNotBeNil)

			logger.Error("Error message")
			output := buf.String()
			c.So(output, c.ShouldContainSubstring, "ERROR")

			buf.Reset()
		})

		c.Convey("For FATAL log level", func() {
			buf := &bytes.Buffer{}
			logger := NewLogger(buf, "FATAL")

			c.So(logger, c.ShouldNotBeNil)

			logger.Log(ctx, Fatal, "Fatal message")
			output := buf.String()
			c.So(output, c.ShouldContainSubstring, "FATAL")

			buf.Reset()
		})
	})
}
