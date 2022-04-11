package logs

import "go.uber.org/zap"

var Logger *zap.Logger = GetLogs()
func GetLogs() *zap.Logger{
	Logger ,_ := zap.NewProduction()
	return Logger
}