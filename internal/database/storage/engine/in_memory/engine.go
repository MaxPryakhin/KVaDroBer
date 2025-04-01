package in_memory

import (
	"context"
	"errors"

	"go.uber.org/zap"
)

type Engine struct {
	table  *HashTable
	logger *zap.Logger
}

func NewEngine(logger *zap.Logger) (*Engine, error) {
	if logger == nil {
		return nil, errors.New("logger is invalid")
	}

	return &Engine{
		table:  NewHashTable(),
		logger: logger,
	}, nil
}
func (e *Engine) Set(ctx context.Context, key, value string) {
	e.table.Set(key, value)
}

func (e *Engine) Get(ctx context.Context, key string) (string, bool) {
	return e.table.Get(key)
}

func (e *Engine) Del(ctx context.Context, key string) {
	e.table.Del(key)
}
