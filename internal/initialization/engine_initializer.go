package initialization

import (
	"errors"

	"go.uber.org/zap"

	"kvadrober/internal/configuration"
	"kvadrober/internal/database/storage/engine/in_memory"
)

func CreateEngine(cfg *configuration.EngineConfig, logger *zap.Logger) (*in_memory.Engine, error) {
	if logger == nil {
		return nil, errors.New("logger is invalid")
	}

	if cfg == nil {
		return in_memory.NewEngine(logger)
	}

	if cfg.Type != "" {
		supportedTypes := map[string]struct{}{
			"in_memory": {},
		}

		if _, found := supportedTypes[cfg.Type]; !found {
			return nil, errors.New("engine type is incorrect")
		}
	}

	return in_memory.NewEngine(logger)
}
