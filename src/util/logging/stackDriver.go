package logging

import (
	"context"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type CtxKey string

const (
	KeyUserAgent CtxKey = "userAgent"
)

// SetUserAgent sets UserAgent to given context
func SetUserAgent(ctx context.Context, userAgent string) context.Context {
	return context.WithValue(ctx, KeyUserAgent, userAgent)
}

// GetUserAgent returns UserAgent from given context
func GetUserAgent(ctx context.Context) string {
	val, ok := ctx.Value(KeyUserAgent).(string)
	if !ok {
		return ""
	}
	return val
}

type FuncCtxValueStr func(ctx context.Context) string

type stackDriverLogger struct {
	funcUserAgernt FuncCtxValueStr
	logger         *zap.Logger
	Logger
}

func NewStackDriverLoggerByName(name string) Logger {
	return NewStackDriverLogger(name, GetUserAgent)
}

func NewStackDriverLogger(name string, funcUserAgernt FuncCtxValueStr) Logger {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeLevel = stackDriverLevelEncoder()
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	config.Development = false

	// TODO：本番のみここを設定してもよいかも。warnLevel以上のログを出すように
	// config.Level = zap.NewAtomicLevelAt(zap.WarnLevel)

	logger, err := config.Build()
	if err != nil {
		panic(err)
	}
	return &stackDriverLogger{
		logger:         logger.Named(name),
		funcUserAgernt: funcUserAgernt,
	}
}

func (l *stackDriverLogger) Debug(ctx context.Context, msg string, fields ...zap.Field) {
	l.logger.Debug(msg, append(fields,
		zap.Object("timestamp", NewStackDriverTimestamp()),
		zap.String("user_agent", l.funcUserAgernt(ctx)),
	)...)

}

func (l *stackDriverLogger) Info(ctx context.Context, msg string, fields ...zap.Field) {
	l.logger.Info(msg, append(fields,
		zap.Object("timestamp", NewStackDriverTimestamp()),
		zap.String("user_agent", l.funcUserAgernt(ctx)),
	)...)
}

func (l *stackDriverLogger) Warn(ctx context.Context, msg string, fields ...zap.Field) {
	l.logger.Warn(msg, append(fields,
		zap.Object("timestamp", NewStackDriverTimestamp()),
		zap.String("user_agent", l.funcUserAgernt(ctx)),
	)...)
}

func (l *stackDriverLogger) Error(ctx context.Context, msg string, fields ...zap.Field) {
	l.logger.Error(msg, append(fields,
		zap.Object("timestamp", NewStackDriverTimestamp()),
		zap.String("user_agent", l.funcUserAgernt(ctx)),
	)...)
}

func stackDriverLevelEncoder() zapcore.LevelEncoder {
	var stackDriverLogLevelSeverity = map[zapcore.Level]string{
		zapcore.DebugLevel:  "DEBUG",
		zapcore.InfoLevel:   "INFO",
		zapcore.WarnLevel:   "WARNING",
		zapcore.ErrorLevel:  "ERROR",
		zapcore.DPanicLevel: "ALERT",
		zapcore.PanicLevel:  "ALERT",
		zapcore.FatalLevel:  "EMERGENCY",
	}

	return func(lv zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(stackDriverLogLevelSeverity[lv])
	}
}

type stackDriverTimestamp struct {
	nanoSeconds int64
}

func (ts stackDriverTimestamp) MarshalLogObject(e zapcore.ObjectEncoder) error {
	e.AddInt64("seconds", ts.nanoSeconds/int64(time.Second))
	e.AddInt64("nanos", ts.nanoSeconds%(int64(time.Second)))
	return nil
}

func NewStackDriverTimestamp() zapcore.ObjectMarshaler {
	return stackDriverTimestamp{nanoSeconds: time.Now().UnixNano()}
}
