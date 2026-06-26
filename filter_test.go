package cloapi

import (
	"context"
	"net/http"
	"net/url"
	"testing"
)

// applyEditors builds a request and runs the editors, returning the resulting query.
func applyEditors(t *testing.T, editors ...RequestEditorFn) url.Values {
	t.Helper()
	req, err := http.NewRequest(http.MethodGet, "https://api.clo.ru/v2/things", nil)
	if err != nil {
		t.Fatal(err)
	}
	for _, e := range editors {
		if err := e(context.Background(), req); err != nil {
			t.Fatalf("editor error: %v", err)
		}
	}
	return req.URL.Query()
}

func TestWithPage(t *testing.T) {
	q := applyEditors(t, WithPage(50, 100))
	if q.Get("limit") != "50" || q.Get("offset") != "100" {
		t.Fatalf("limit=%q offset=%q", q.Get("limit"), q.Get("offset"))
	}
}

func TestWithOrderBy(t *testing.T) {
	q := applyEditors(t, WithOrderBy("-created"))
	if q.Get("ordering") != "-created" {
		t.Fatalf("ordering=%q", q.Get("ordering"))
	}
}

func TestWithQuery(t *testing.T) {
	q := applyEditors(t, WithQuery(map[string]string{"foo": "bar", "baz": "qux"}))
	if q.Get("foo") != "bar" || q.Get("baz") != "qux" {
		t.Fatalf("unexpected query: %v", q)
	}
}

func TestWithFilter(t *testing.T) {
	q := applyEditors(t, WithFilter(
		FilterField{Field: "size", Condition: "gte", Value: "10"},
		FilterField{Field: "size", Condition: "lt", Value: "100"},
		FilterField{Field: "name", Condition: "bogus", Value: "x"}, // skipped
	))
	if q.Get("size__gte") != "10" {
		t.Errorf("size__gte=%q", q.Get("size__gte"))
	}
	if q.Get("size__lt") != "100" {
		t.Errorf("size__lt=%q", q.Get("size__lt"))
	}
	if _, ok := q["name__bogus"]; ok {
		t.Errorf("unsupported condition should be skipped, got %v", q["name__bogus"])
	}
}

func TestEditorsCompose(t *testing.T) {
	q := applyEditors(t, WithPage(10, 0), WithFilter(FilterField{Field: "status", Condition: "in", Value: "active"}))
	if q.Get("limit") != "10" || q.Get("status__in") != "active" {
		t.Fatalf("composition failed: %v", q)
	}
}
