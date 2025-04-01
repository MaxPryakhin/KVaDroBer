package initialization

import (
	"context"
	"errors"
	"fmt"
	"kvadrober/internal/configuration"
	"kvadrober/internal/database"
	"kvadrober/internal/database/compute"
	"kvadrober/internal/database/storage"
	"kvadrober/internal/database/storage/engine/in_memory"
	"kvadrober/internal/network"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

type Initializer struct {
	logger *zap.Logger
	server *network.TCPServer
	engine *in_memory.Engine
}

func NewInitializer(cfg *configuration.Config) (*Initializer, error) {
	if cfg == nil {
		return nil, errors.New("failed to initialize: config is invalid")
	}

	logger, err := CreateLogger(cfg.Logging)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize logger: %w", err)
	}

	server, err := CreateNetwork(cfg.Network, logger)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize network: %w", err)
	}

	engine, err := CreateEngine(cfg.Engine, logger)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize engine: %w", err)
	}

	return &Initializer{
		logger: logger,
		server: server,
		engine: engine,
	}, nil
}

func (i *Initializer) StartDatabase(ctx context.Context) error {
	compute, err := compute.NewCompute(i.logger)
	if err != nil {
		return err
	}

	storage, err := storage.NewStorage(i.engine, i.logger)
	if err != nil {
		return err
	}

	database, err := database.NewDatabase(compute, storage, i.logger)
	if err != nil {
		return err
	}

	group, groupCtx := errgroup.WithContext(ctx)

	group.Go(func() error {
		i.server.HandleQueries(groupCtx, func(ctx context.Context, query []byte) []byte {
			response := database.HandleQuery(ctx, string(query))
			return []byte(response)
		})

		return nil
	})

	err = group.Wait()
	i.logger.Sync()
	return err
}
