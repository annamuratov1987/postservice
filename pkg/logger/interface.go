package logger

import "go.uber.org/zap"

type ILogger interface {
	Sync() error
	With(fields ...zap.Field) ILogger

	Info(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	Fatal(msg string, fields ...zap.Field)
	Debug(msg string, fields ...zap.Field)
}
