package assert

import (
	"fmt"
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
