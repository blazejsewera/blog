package page_test

import (
	"bytes"
	"testing"

	"github.com/blazejsewera/blog/renderer/internal/templates"
)

type renderedPage struct {
	html []byte
	err  error
}

func assertRenderPageContains(t testing.TB, rendered renderedPage, expectedData map[string]string) {
	t.Helper()
	if rendered.err != nil {
		t.Fatalf("unexpected error: %s", rendered.err)
	}

	if bytes.Contains(rendered.html, templates.ErrorBytes) {
		t.Errorf(templates.ErrorMessage)
	}

	for _, v := range expectedData {
		value := []byte(v)
		if !bytes.Contains(rendered.html, value) {
			t.Errorf("rendered does not contain '%s'", value)
		}
	}
}
