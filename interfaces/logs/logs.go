package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

func getEncoder() *zap.Logger {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	logger, _ := config.Build()
	return logger
}

var Logger *zap.Logger = GetLogs()

func GetLogs() *zap.Logger {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	config.EncoderConfig.TimeKey = "timestamp"
	logger, err := config.Build()
	if err != nil {
		logger, _ = zap.NewProduction()
		return logger
	}
	return logger
}
func NewLogger() *zap.Logger {
	return GetLogs()
}
