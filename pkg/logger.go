package pkg

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func customLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	var color string
	switch level {
	case zapcore.DebugLevel:
		color = "\x1b[47m\x1b[30m" // White background, black text
	case zapcore.InfoLevel:
		color = "\x1b[42m\x1b[30m" // Green background, black text
	case zapcore.WarnLevel:
		color = "\x1b[43m\x1b[30m" // Yellow background, black text
	case zapcore.ErrorLevel:
		color = "\x1b[41m\x1b[37m" // Red background, white text
	case zapcore.DPanicLevel, zapcore.PanicLevel, zapcore.FatalLevel:
		color = "\x1b[45m\x1b[37m" // Magenta background, white text
	default:
		color = "\x1b[47m\x1b[30m" // White background, black text
	}
	enc.AppendString(fmt.Sprintf("%s %-5s \x1b[0m", color, level.CapitalString()))
}

func CustomLogger() (*zap.Logger, error) {
	config := zap.Config{
		Level: zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: true,
		Encoding: "console",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey: "time",
			LevelKey: "level",
			NameKey: "logger",
			CallerKey: "caller",
			MessageKey: "msg",
			StacktraceKey: "stacktrace",
			LineEnding: zapcore.DefaultLineEnding,
			EncodeLevel: customLevelEncoder,
			EncodeTime: zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller: zapcore.FullCallerEncoder,
		},
		OutputPaths: []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	return config.Build()
}