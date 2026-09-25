package geonames

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGeonamesClientSearch(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got := r.URL.Query().Get("featureClass"); got != "P" {
			t.Fatalf("expected featureClass=P, got %q", got)
		}

		if got := r.URL.Query().Get("username"); got != "demo" {
			t.Fatalf("expected username=demo, got %q", got)
		}

		if got := r.URL.Query().Get("maxRows"); got != "10" {
			t.Fatalf("expected maxRows=10, got %q", got)
		}

		if got := r.URL.Query().Get("name_startsWith"); got != "a" {
			t.Fatalf("expected name_startsWith=a, got %q", got)
		}

		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write([]byte(`{"totalResultsCount":2,"geonames":[{"name":"Aachen"},{"name":"Athens"}]}`))
		if err != nil {
			t.Fatalf("write response: %v", err)
		}
	}))
	defer server.Close()

	client := NewGeonamesClient(server.URL, "demo", server.Client())
	letter := "a"
	params := client.WithBaseParams(10, &letter)

	result, err := client.Search(context.Background(), params)
	if err != nil {
		t.Fatalf("Search returned error: %v", err)
	}

	if result.TotalResultsCount != 2 {
		t.Fatalf("expected 2 results, got %d", result.TotalResultsCount)
	}

	if len(result.Geonames) != 2 {
		t.Fatalf("expected 2 geonames, got %d", len(result.Geonames))
	}

	if result.Geonames[0].Name != "Aachen" {
		t.Fatalf("expected first city Aachen, got %q", result.Geonames[0].Name)
	}
}
