package logger

import "go.uber.org/zap"

// SetLogger is the setter for log variable, it should be the only way to assign value to log.
func SetLogger(cfg *LoggingConfig) *zap.Logger {
	zapLogger := New(cfg)
	zap.ReplaceGlobals(zapLogger)
	return zapLogger
}
