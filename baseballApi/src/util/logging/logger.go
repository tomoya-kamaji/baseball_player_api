package logging

import (
	"context"

	"go.uber.org/zap"
)

type Logger interface {
	Debug(context.Context, string, ...zap.Field)
	Info(context.Context, string, ...zap.Field)
	Warn(context.Context, string, ...zap.Field)
	Error(context.Context, string, ...zap.Field)
	Panic(context.Context, string, ...zap.Field)
	Fatal(context.Context, string, ...zap.Field)
}
