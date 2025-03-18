package initialization

import (
	"errors"
	"fmt"
	"kvadrober/internal/configuration"

	"go.uber.org/zap"
)

type Initializer struct {
	logger *zap.Logger
}

func NewInitializer(cfg *configuration.Config) (*Initializer, error) {
	if cfg == nil {
		return nil, errors.New("failed to initialize: config is invalid")
	}

	logger, err := CreateLogger(cfg.Logging)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize logger: %w", err)
	}

	return &Initializer{
		logger: logger,
	}, nil
}
