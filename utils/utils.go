package utils

import (
	"fmt"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func LoadLogger(logLevel string) (*zap.Logger, error) {
	level := func() zap.AtomicLevel {
		switch strings.ToUpper(logLevel) {
		case zapcore.InfoLevel.CapitalString():
			return zap.NewAtomicLevelAt(zapcore.InfoLevel)
		case zapcore.WarnLevel.CapitalString():
			return zap.NewAtomicLevelAt(zapcore.WarnLevel)
		}
		// Default Debug
		return zap.NewAtomicLevelAt(zapcore.DebugLevel)
	}()

	cfg := zap.NewProductionConfig()
	cfg.Level = level

	logger, err := cfg.Build()
	if err != nil {
		return nil, fmt.Errorf("failed loading logger: %w", err)
	}

	return logger, nil
}
