package log

import (
	"go.uber.org/zap"
)

// Field key-value pairs
type Field = zap.Field

// String create field with value type is string
func String(key, value string) Field {
	return zap.String(key, value)
}

// Any create field with value type is any
func Any(key string, value interface{}) Field {
	return zap.Any(key, value)
}

// Error create field with value type is error
func Error(err error) Field {
	return zap.Error(err)
}
