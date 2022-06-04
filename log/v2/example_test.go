package log_test

import (
	"fmt"
	"os"

	logv2 "github.com/exfly/pkg/log/v2"
	zaplog "github.com/exfly/pkg/log/v2/zap"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/go-kit/log/term"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Example_basic() {
	logger := log.NewLogfmtLogger(os.Stdout)
	level.Debug(logger).Log("msg", "this message is at the debug level")
	level.Info(logger).Log("msg", "this message is at the info level")
	level.Warn(logger).Log("msg", "this message is at the warn level")
	level.Error(logger).Log("msg", "this message is at the error level")

	// Output:
	// level=debug msg="this message is at the debug level"
	// level=info msg="this message is at the info level"
	// level=warn msg="this message is at the warn level"
	// level=error msg="this message is at the error level"
}

func Example_filtered() {
	// Set up logger with level filter.
	logger := log.NewLogfmtLogger(os.Stdout)
	logger = level.NewFilter(logger, level.AllowInfo())

	logv2.With(logger).Log("key", "val")

	// Use level helpers to log at different levels.
	logv2.Error(logger).Log("err", errors.New("bad data"))
	logv2.Info(logger).Log("event", "data saved")
	logv2.Debug(logger).Log("next item", 17) // filtered

	// Output:
	// key=val
	// level=error err="bad data"
	// level=info event="data saved"
}

func Example_term() {
	colorFn := func(keyvals ...interface{}) term.FgBgColor {
		for i := 0; i < len(keyvals)-1; i += 2 {
			if keyvals[i] != "level" {
				continue
			}

			level, ok := (keyvals[i+1]).(fmt.Stringer)
			if !ok {
				continue
			}

			switch level.String() {
			case "debug":
				return term.FgBgColor{Fg: term.DarkGray}
			case "info":
				return term.FgBgColor{Fg: term.Gray}
			case "warn":
				return term.FgBgColor{Fg: term.Yellow}
			case "error":
				return term.FgBgColor{Fg: term.Red}
			case "crit":
				return term.FgBgColor{Fg: term.Gray, Bg: term.DarkRed}
			default:
				return term.FgBgColor{}
			}
		}
		return term.FgBgColor{}
	}

	logger := term.NewColorLogger(
		term.NewColorWriter(os.Stdout),
		log.NewLogfmtLogger,
		colorFn,
	)

	logger = logv2.With(
		level.NewFilter(logger, level.AllowInfo()),
	)

	logv2.Error(logger).Log("err", errors.New("bad data"))
	logv2.Info(logger).Log("event", "data saved")
	logv2.Debug(logger).Log("next item", 17) // filtered

	// Output:
	// [31;1mlevel=error err="bad data"
	// [39;49;22m[37mlevel=info event="data saved"
	// [39;49;22m
}

func Example_error() {
	logger := log.NewLogfmtLogger(os.Stdout)
	err := errors.New("noop error")
	errWrap := errors.Wrap(err, "")
	level.Debug(logger).Log("msg", "err", "err", errWrap)
	level.Debug(logger).Log("msg", "err with stacktrace", "err", fmt.Sprintf("%+v", errWrap))

	// level=debug msg=err err=": noop error"
	// level=debug msg="err with stacktrace" err="noop error\ngithub.com/exfly/pkg/log/v2_test.Example_error\n\t/Volumes/code/github.com/exfly/pkg/log/v2/example_test.go:105\ntesting.runExample\n\t/Users/zhf/.goenv/versions/1.16.0/src/testing/run_example.go:63\ntesting.runExamples\n\t/Users/zhf/.goenv/versions/1.16.0/src/testing/example.go:44\ntesting.(*M).Run\n\t/Users/zhf/.goenv/versions/1.16.0/src/testing/testing.go:1419\nmain.main\n\t_testmain.go:57\nruntime.main\n\t/Users/zhf/.goenv/versions/1.16.0/src/runtime/proc.go:225\nruntime.goexit\n\t/Users/zhf/.goenv/versions/1.16.0/src/runtime/asm_amd64.s:1371\n\ngithub.com/exfly/pkg/log/v2_test.Example_error\n\t/Volumes/code/github.com/exfly/pkg/log/v2/example_test.go:106\ntesting.runExample\n\t/Users/zhf/.goenv/versions/1.16.0/src/testing/run_example.go:63\ntesting.runExamples\n\t/Users/zhf/.goenv/versions/1.16.0/src/testing/example.go:44\ntesting.(*M).Run\n\t/Users/zhf/.goenv/versions/1.16.0/src/testing/testing.go:1419\nmain.main\n\t_testmain.go:57\nruntime.main\n\t/Users/zhf/.goenv/versions/1.16.0/src/runtime/proc.go:225\nruntime.goexit\n\t/Users/zhf/.goenv/versions/1.16.0/src/runtime/asm_amd64.s:1371"
}

func Example_zap() {
	var err error

	encoderConfig := zap.NewDevelopmentEncoderConfig()
	// omit zap level time msg
	encoderConfig.LevelKey = ""
	encoderConfig.TimeKey = ""
	encoderConfig.MessageKey = ""

	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	writer := os.Stdout
	zaplogger := zap.New(
		zapcore.NewCore(encoder, zapcore.AddSync(writer), zap.DebugLevel),
		zap.Development(),
		zap.WithCaller(false),
	)

	logger := level.NewFilter(
		zaplog.NewZapSugarLogger(zaplogger, zapcore.InfoLevel),
		level.AllowAll(),
	)

	err = errors.New("noop error")
	errWrap := errors.Wrap(err, "wrap")
	level.Info(logger).Log(logv2.MsgKey, "has_err", logv2.ErrKey, errWrap)
	level.Debug(logger).Log(logv2.MsgKey, "has_err", logv2.ErrKey, errWrap)
}
