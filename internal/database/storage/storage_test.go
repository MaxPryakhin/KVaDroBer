package storage

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestNewStorage(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)

	tests := map[string]struct {
		engine Engine
		logger *zap.Logger

		expectedErr    error
		expectedNilObj bool
	}{
		"create storage without engine": {
			expectedErr:    errors.New("engine is invalid"),
			expectedNilObj: true,
		},
		"create storage without logger": {
			engine:         NewMockEngine(ctrl),
			expectedErr:    errors.New("logger is invalid"),
			expectedNilObj: true,
		},
		"create storage": {
			engine:      NewMockEngine(ctrl),
			logger:      zap.NewNop(),
			expectedErr: nil,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			storage, err := NewStorage(test.engine, test.logger)
			assert.Equal(t, test.expectedErr, err)
			if test.expectedNilObj {
				assert.Nil(t, storage)
			} else {
				assert.NotNil(t, storage)
			}
		})
	}
}

func TestStorageSet(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)

	tests := map[string]struct {
		engine func() Engine

		expectedErr error
	}{
		"set without wal": {
			engine: func() Engine {
				engine := NewMockEngine(ctrl)
				engine.EXPECT().
					Set(gomock.Any(), "key", "value")
				return engine
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			storage, err := NewStorage(test.engine(), zap.NewNop())
			require.NoError(t, err)

			storage.Set(context.Background(), "key", "value")
			assert.Equal(t, test.expectedErr, err)
		})
	}
}

func TestStorageDel(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)

	tests := map[string]struct {
		engine func() Engine

		expectedErr error
	}{
		"del existing element": {
			engine: func() Engine {
				engine := NewMockEngine(ctrl)
				engine.EXPECT().
					Del(gomock.Any(), "key")
				return engine
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			storage, err := NewStorage(test.engine(), zap.NewNop())
			require.NoError(t, err)

			storage.Del(context.Background(), "key")
			assert.Equal(t, test.expectedErr, err)
		})
	}
}

func TestStorageGet(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)

	tests := map[string]struct {
		engine func() Engine

		expectedValue string
		expectedErr   error
	}{
		"get with unexesiting element": {
			engine: func() Engine {
				engine := NewMockEngine(ctrl)
				engine.EXPECT().
					Get(gomock.Any(), "key").
					Return("", false)
				return engine
			},
			expectedErr: ErrorNotFound,
		},
		"get with exesiting element": {
			engine: func() Engine {
				engine := NewMockEngine(ctrl)
				engine.EXPECT().
					Get(gomock.Any(), "key").
					Return("value", true)
				return engine
			},
			expectedValue: "value",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			storage, err := NewStorage(test.engine(), zap.NewNop())
			require.NoError(t, err)

			value, err := storage.Get(context.Background(), "key")
			assert.Equal(t, test.expectedErr, err)
			assert.Equal(t, test.expectedValue, value)
		})
	}
}
