package compute

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQuery(t *testing.T) {
	query := NewQuery(GetCommandID, []string{"GET", "key"})
	require.Equal(t, GetCommandID, query.CommandID())
	require.True(t, reflect.DeepEqual([]string{"GET", "key"}, query.Arguments()))
}
