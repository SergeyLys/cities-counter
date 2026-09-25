package city_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	citypkg "github.com/sergeylys/city-counter/backend/internal/city"
	cityCache "github.com/sergeylys/city-counter/backend/internal/city/cache"
	cityEntity "github.com/sergeylys/city-counter/backend/internal/city/entity"
	cityRepository "github.com/sergeylys/city-counter/backend/internal/city/repositories"
	cityStrategies "github.com/sergeylys/city-counter/backend/internal/city/strategies"
)

func TestHandlerCount(t *testing.T) {
	repository := cityRepository.NewMemoryCityRepository([]cityEntity.City{
		{Name: "Rio de Janeiro"},
		{Name: "Cairo"},
		{Name: "Chongqing"},
		{Name: "Chengdu"},
	})

	service := citypkg.NewCityService(map[string]cityStrategies.CountStrategy{
		"startswith": cityStrategies.NewStartsWithStrategy(repository, cityCache.NewCountCache()),
		"bruteforce": cityStrategies.NewBruteforceStrategy(repository, cityCache.NewCountCache()),
	})

	handler := citypkg.NewCityHandler(service)

	req := httptest.NewRequest(http.MethodGet, "/api/cities/count?letter=C&strategy=startswith", nil)
	recorder := httptest.NewRecorder()

	handler.Count(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, recorder.Code)
	}

	expected := `{"letter":"C","count":3}` + "\n"
	if recorder.Body.String() != expected {
		t.Fatalf("expected body %q, got %q", expected, recorder.Body.String())
	}
}

func TestHandlerCountInvalidLetter(t *testing.T) {
	repository := cityRepository.NewMemoryCityRepository(nil)
	service := citypkg.NewCityService(map[string]cityStrategies.CountStrategy{
		"startswith": cityStrategies.NewStartsWithStrategy(repository, cityCache.NewCountCache()),
		"bruteforce": cityStrategies.NewBruteforceStrategy(repository, cityCache.NewCountCache()),
	})
	handler := citypkg.NewCityHandler(service)

	tests := []string{
		"/api/cities/count",
		"/api/cities/count?letter=",
		"/api/cities/count?letter=CC",
	}

	for _, url := range tests {
		t.Run(url, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, url, nil)
			recorder := httptest.NewRecorder()

			handler.Count(recorder, req)

			if recorder.Code != http.StatusOK && recorder.Code != http.StatusBadRequest {
				t.Fatalf("expected status 200 or 400, got %d", recorder.Code)
			}
		})
	}
}
