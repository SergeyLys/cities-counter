package city

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sergeylys/city-counter/backend/internal/city"
)

func TestHandlerCount(t *testing.T) {
	cities := []city.City{
		{Name: "Rio de Janeiro"},
		{Name: "Cairo"},
		{Name: "Chongqing"},
		{Name: "Chengdu"},
	}

	repository := city.NewMemoryCityRepository(cities)
	service := city.NewCityService(repository)
	handler := city.NewCityHandler(service)

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/cities/count?letter=C",
		nil,
	)

	recorder := httptest.NewRecorder()

	handler.Count(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusOK,
			recorder.Code,
		)
	}

	expected := `{"letter":"C","count":3}` + "\n"

	if recorder.Body.String() != expected {
		t.Errorf(
			"expected body %q, got %q",
			expected,
			recorder.Body.String(),
		)
	}
}

func TestHandlerCountInvalidLetter(t *testing.T) {
	repository := city.NewMemoryCityRepository(nil)
	service := city.NewCityService(repository)
	handler := city.NewCityHandler(service)

	tests := []string{
		"/api/cities/count",
		"/api/cities/count?letter=",
		"/api/cities/count?letter=CC",
	}

	for _, url := range tests {
		t.Run(url, func(t *testing.T) {
			req := httptest.NewRequest(
				http.MethodGet,
				url,
				nil,
			)

			recorder := httptest.NewRecorder()

			handler.Count(recorder, req)

			if recorder.Code != http.StatusBadRequest {
				t.Errorf(
					"expected status %d, got %d",
					http.StatusBadRequest,
					recorder.Code,
				)
			}
		})
	}
}
