package assert

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"golang.org/x/net/html"
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

func EqualHTMLTree[T, R []byte | string](t testing.TB, expected T, actual R) {
	t.Helper()
	nodesExpected := htmlTreeToList(t, string(expected))
	nodesActual := htmlTreeToList(t, string(actual))

	Equal(t, nodesExpected, nodesActual)
}

func htmlTreeToList(t testing.TB, actual string) []string {
	node, err := parseHTML(actual)
	if err != nil {
		t.Fatal(err)
	}
	var nodes []string
	for n := range node.Descendants() {
		stringData := strings.TrimSpace(n.Data)
		if stringData != "" {
			nodes = append(nodes, n.Data)
		}
	}
	return nodes
}

func parseHTML(input string) (*html.Node, error) {
	return html.Parse(strings.NewReader(input))
}
