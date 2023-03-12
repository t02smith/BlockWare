package util

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Logger *zap.SugaredLogger
)

func InitLogger() *zap.SugaredLogger {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.ConsoleSeparator = "\t"
	config.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)

	// config.OutputPaths = []string{
	// 	"./zap.log",
	// }

	l, _ := config.Build()
	Logger = l.Sugar()
	defer Logger.Sync()

	return Logger
}
