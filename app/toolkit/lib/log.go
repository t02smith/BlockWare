package lib

import (
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Logger *zap.SugaredLogger
	once   sync.Once
)

func InitLogger() *zap.SugaredLogger {
	once.Do(func() {
		config := zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		config.EncoderConfig.ConsoleSeparator = "\t"

		l, _ := config.Build()
		Logger = l.Sugar()

		defer Logger.Sync()
	})

	return Logger
}
