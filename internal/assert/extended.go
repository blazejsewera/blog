package assert

import (
	"fmt"
	"github.com/blazejsewera/blog/postprocess"
	"reflect"
	"testing"
)

func EqualFields[T any](t testing.TB, expected T, actual T, fields ...string) {
	t.Helper()
	ev := reflect.ValueOf(expected)
	av := reflect.ValueOf(actual)

	for _, field := range fields {
		Equal(t,
			ev.FieldByName(field).Interface(),
			av.FieldByName(field).Interface(),
			fmt.Sprintf("Field: %s", field))
	}
}

func EqualHTML[T, R []byte | string](t testing.TB, expected T, actual R) bool {
	t.Helper()
	normalizedExpected := string(postprocess.MinifyHTML([]byte(expected)))
	normalizedActual := string(postprocess.MinifyHTML([]byte(actual)))
	return Equal(t, normalizedExpected, normalizedActual)
}
