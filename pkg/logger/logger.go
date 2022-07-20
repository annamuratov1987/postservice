package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(
	outputPath string,
) (ILogger, error) {

	config := zap.NewProductionConfig()
	config.DisableCaller = true
	config.DisableStacktrace = true
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	if len(outputPath) > 0 {
		config.OutputPaths = append(config.OutputPaths, outputPath)
		config.ErrorOutputPaths = append(config.ErrorOutputPaths, outputPath)
	}

	core, err := config.Build()
	if err != nil {
		return nil, err
	}

	log := &logger{core: core}

	return log, nil
}

type logger struct {
	core *zap.Logger
}

func (l *logger) With(fields ...zap.Field) ILogger {
	copy := l.clone()
	copy.core = copy.core.With(fields...)
	return copy
}

func (l *logger) Sync() error {
	return l.core.Sync()
}

func (l *logger) Info(msg string, fields ...zap.Field) {
	l.core.Info(msg, fields...)
}

func (l *logger) Warn(msg string, fields ...zap.Field) {
	l.core.Warn(msg, fields...)
}

func (l *logger) Debug(msg string, fields ...zap.Field) {
	l.core.Debug(msg, fields...)
}

func (l *logger) Error(msg string, fields ...zap.Field) {
	l.core.Error(msg, fields...)
}

func (l *logger) Fatal(msg string, fields ...zap.Field) {
	l.core.Fatal(msg, fields...)
}

func (l *logger) clone() *logger {
	copy := *l
	return &copy
}
