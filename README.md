# Logger Package

The logger package provides a simplified composite literal for creating structured loggers, leveraging the `slog` package.

## NewLogger

 This is the primary function in this package. The `NewLogger` function takes an `io.Writer` (optional), and a logging level (optional). It returns a new logger instance from the slog package.

```go
func NewLogger(writer io.Writer, level ...slog.Level) (logger *slog.Logger)
```

## Parameters
- `writer`: An io.Writer where the log output will be written. This is optional, and if not provided, the logger will default to writing to os.Stdout.
- `level`: A slog.Level which sets the logging level for the logger. This is optional. If not provided, it defaults to `slog.LevelError`.

## Logging Levels
The logging levels used are:

- DEBUG = -4
- INFO = 0
- WARN = 4
- ERROR = 8
- FATAL = 12 (custom log level)

If a message is logged at a level of ERROR+4, it is output as "FATAL" in the log.

## Usage
Use of the logging package should be used at the top level function such as within `main`. Functions which return errors should remain

To create a logger with default parameters (writing to `os.Stdout` at ERROR level and above):
```go
log := logger.NewLogger(nil)
...
// Standard error
_, err := exampleFunction()
	if err != nil {
		log.Error("describe action here", "rsp", err)
	}

// Fatal
_, err := exampleFunction2()
	if err != nil {
		log.Log(context.Background(), logger.LevelFatal, "describe action here", "rsp", err)
		os.Exit(1)
	}
```
Outputs:
```text
time=2023-07-20T16:54:40.558+10:00 level=ERROR msg="describe action here" rsp="error: something in exampleFunction went wrong"
time=2023-07-20T16:54:40.558+10:00 level=FATAL msg="describe action here" rsp="error: something in exampleFunction2 went very wrong"
exit status 1
```

To create a logger writing to a provided `io.Writer` at the ERROR level (this is mainly for testing the output stream):
```go
buf := new(bytes.Buffer)
logger := NewLogger(buf)
```

