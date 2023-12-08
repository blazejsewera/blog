package assert

import (
	"fmt"
	"github.com/yosssi/gohtml"
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
	normalizedExpected := formatHTML(string(expected))
	normalizedActual := formatHTML(string(actual))
	return Equal(t, normalizedExpected, normalizedActual)
}

func formatHTML(input string) string {
	return gohtml.Format(input)
}
