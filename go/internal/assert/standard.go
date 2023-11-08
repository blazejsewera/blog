package assert

import (
	tt "github.com/stretchr/testify/assert"
	"testing"
)

func Equal(t testing.TB, expected any, actual any, errorMsg ...any) bool {
	return tt.Equal(t, expected, actual, errorMsg...)
}

func Zero(t testing.TB, value any, errorMsg ...any) bool {
	return tt.Zero(t, value, errorMsg...)
}

func Len(t testing.TB, expectedLen int, value any, errorMsg ...any) bool {
	return tt.Len(t, value, expectedLen, errorMsg...)
}
