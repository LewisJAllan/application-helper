package zaphelper

import (
	"context"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var ZapLogger = zap.New(zapcore.NewCore(
	zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:       "msg",
		LevelKey:         "level",
		TimeKey:          "@timestamp",
		NameKey:          "logger",
		CallerKey:        "caller",
		FunctionKey:      "function",
		StacktraceKey:    "stacktrace",
		LineEnding:       zapcore.DefaultLineEnding,
		EncodeLevel:      zapcore.LowercaseLevelEncoder,
		EncodeTime:       zapcore.RFC3339NanoTimeEncoder,
		EncodeDuration:   zapcore.NanosDurationEncoder,
		EncodeCaller:     zapcore.ShortCallerEncoder,
		EncodeName:       zapcore.FullNameEncoder,
		ConsoleSeparator: "",
	}),
	zapcore.AddSync(os.Stdout),
	zap.NewAtomicLevelAt(zapcore.InfoLevel),
),
	zap.AddCaller(),
	zap.AddCallerSkip(1),
)

// FromContext will return the logger associated with the context if present, otherwise the ZapLogger
func FromContext(ctx context.Context) *zap.Logger {
	if l, ok := ctx.Value(struct{}{}).(*zap.Logger); ok {
		return l
	}
	return ZapLogger
}
