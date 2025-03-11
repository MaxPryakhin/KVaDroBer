package storage

import (
	"context"
	"errors"

	"go.uber.org/zap"
)

var (
	ErrorNotFound = errors.New("not found")
)

type Engine interface {
	Set(context.Context, string, string)
	Get(context.Context, string) (string, bool)
	Del(context.Context, string)
}

type Storage struct {
	engine Engine
	logger *zap.Logger
}

func NewStorage(engine Engine, logger *zap.Logger) (*Storage, error) {
	if engine == nil {
		return nil, errors.New("engine is invalid")
	}

	if logger == nil {
		return nil, errors.New("logger is invalid")
	}

	return &Storage{
		engine: engine,
		logger: logger,
	}, nil
}

func (s *Storage) Set(ctx context.Context, key, value string) {
	s.engine.Set(ctx, key, value)
}

func (s *Storage) Get(ctx context.Context, key string) (string, error) {
	value, ok := s.engine.Get(ctx, key)
	if !ok {
		return "", ErrorNotFound
	}

	return value, nil
}

func (s *Storage) Del(ctx context.Context, key string) {
	s.engine.Del(ctx, key)
}
