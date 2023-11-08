package assert

import (
	tt "github.com/stretchr/testify/assert"
	"testing"
)

func Equal[T any](t testing.TB, expected T, actual T, errorMsg ...any) bool {
	t.Helper()
	return tt.Equal(t, expected, actual, errorMsg...)
}

func Zero[T any](t testing.TB, value T, errorMsg ...any) bool {
	t.Helper()
	return tt.Zero(t, value, errorMsg...)
}

func Len[T ~[]any | ~[...]any | ~*[...]any | ~map[any]any | string | chan any](
	t testing.TB,
	expectedLen int,
	value T,
	errorMsg ...any) bool {
	t.Helper()
	return tt.Len(t, value, expectedLen, errorMsg...)
}
