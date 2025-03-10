package initialization

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	defaultEncoding   = "json"
	defaultLevel      = zapcore.DebugLevel
	defaultOutputPath = "kvadrober.log"
)

func CreateLogger() (*zap.Logger, error) {
	level := defaultLevel
	output := defaultOutputPath

	loggerCfg := zap.Config{
		Encoding:    defaultEncoding,
		Level:       zap.NewAtomicLevelAt(level),
		OutputPaths: []string{output},
	}

	return loggerCfg.Build()
}
