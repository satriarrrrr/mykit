package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ZapLogger a logger that use Zap
type ZapLogger struct {
	logger *zap.Logger
}

func zapConfig(lvl zapcore.Level) zap.Config {
	return zap.Config{
		Level: zap.NewAtomicLevelAt(lvl),
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "message",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stdout"},
	}
}

// NewZapLogger initiate a new zap logger
func NewZapLogger(cmdName string, debugMode bool) (*ZapLogger, error) {
	level := zap.InfoLevel
	if debugMode {
		level = zap.DebugLevel
	}

	withFields := zap.Fields(zap.String("command", cmdName))

	logger, err := zapConfig(level).Build(withFields)
	if err != nil {
		return nil, err
	}

	return &ZapLogger{logger: logger}, nil
}

// Stop call this function before exiting program
func (zl *ZapLogger) Stop() error {
	return zl.logger.Sync()
}

// Debug write log with DEBUG level
func (zl *ZapLogger) Debug(msg string, fields ...Field) {
	zl.logger.Debug(msg, fields...)
}

// Info write log with INFO level
func (zl *ZapLogger) Info(msg string, fields ...Field) {
	zl.logger.Info(msg, fields...)
}

// Warn write log with WARN level
func (zl *ZapLogger) Warn(msg string, fields ...Field) {
	zl.logger.Warn(msg, fields...)
}

// Error write log with ERROR level
func (zl *ZapLogger) Error(msg string, fields ...Field) {
	zl.logger.Error(msg, fields...)
}

// Panic write log with PANIC level
func (zl *ZapLogger) Panic(msg string, fields ...Field) {
	zl.logger.Panic(msg, fields...)
}

// Fatal write log with FATAL level
func (zl *ZapLogger) Fatal(msg string, fields ...Field) {
	zl.logger.Fatal(msg, fields...)
}
