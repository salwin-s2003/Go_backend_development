package logger

import (
	"go.uber.org/zap"
	"time"
)

var Logger *zap.Logger

func init() {
	var err error
	Logger, err = zap.NewProduction() // or zap.NewDevelopment() for readable logs
	if err != nil {
		panic(err)
	}
}

// Helper functions for structured logging
func ZapString(key, value string) zap.Field {
	return zap.String(key, value)
}

func ZapInt(key string, value int) zap.Field {
	return zap.Int(key, value)
}

func ZapInt32(key string, value int32) zap.Field {
	return zap.Int32(key, value)
}

func ZapDuration(key string, value time.Duration) zap.Field {
	return zap.Duration(key, value)
}

func ZapError(err error) zap.Field {
	return zap.Error(err)
}