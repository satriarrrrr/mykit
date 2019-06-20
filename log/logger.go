package log

import (
	internallog "github.com/satriarrrrr/kit/internal/log"
)

type Logger interface {
	Stop() error
	Debug(msg string, fields ...internallog.Field)
	Info(msg string, fields ...internallog.Field)
	Warn(msg string, fields ...internallog.Field)
	Error(msg string, fields ...internallog.Field)
	Panic(msg string, fields ...internallog.Field)
	Fatal(msg string, fields ...internallog.Field)
}

func NewLogger(commandName string, debugMode bool) (Logger, error) {
	return internallog.NewZapLogger(commandName, debugMode)
}
