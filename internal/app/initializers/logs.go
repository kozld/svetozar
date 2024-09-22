package initializers

import (
	"go.uber.org/zap"
)

// InitializeLogs setups zap logger
func InitializeLogs() *zap.Logger {
	logger, _ := zap.NewProduction()
	return logger
}
