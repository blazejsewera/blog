package assert

import (
	tt "github.com/stretchr/testify/assert"
	"testing"
)

func Equal[T any](t testing.TB, expected T, actual T, errorMsg ...any) bool {
	t.Helper()
	return tt.Equal(t, expected, actual, errorMsg...)
}

func Zero(t testing.TB, value any, errorMsg ...any) bool {
	t.Helper()
	return tt.Zero(t, value, errorMsg...)
}

func Len(t testing.TB, expectedLen int, value any, errorMsg ...any) bool {
	t.Helper()
	return tt.Len(t, value, expectedLen, errorMsg...)
}
