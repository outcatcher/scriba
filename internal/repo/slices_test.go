package repo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type testFrom[T testTo] struct {
	FieldFrom bool
}

type testTo struct {
	FieldTo bool
}

func (tf testFrom[T]) toEntity() *T {
	return &T{FieldTo: tf.FieldFrom}
}

func TestConvertEntitySlice(t *testing.T) {
	t.Parallel()

	src := []testFrom[testTo]{{FieldFrom: true}}
	expected := []testTo{{FieldTo: true}}

	result := convertEntitySlice(src)

	require.Equal(t, expected, result)
}
