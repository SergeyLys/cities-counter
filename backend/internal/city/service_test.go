package city

import (
	"context"
	"testing"

	cityCache "github.com/sergeylys/city-counter/backend/internal/city/cache"
	cityEntity "github.com/sergeylys/city-counter/backend/internal/city/entity"
	cityRepository "github.com/sergeylys/city-counter/backend/internal/city/repositories"
	cityStrategies "github.com/sergeylys/city-counter/backend/internal/city/strategies"
)

func TestCountByLetter(t *testing.T) {
	repository := cityRepository.NewMemoryCityRepository([]cityEntity.City{
		{Name: "Rio de Janeiro"},
		{Name: "Cairo"},
		{Name: "Chongqing"},
		{Name: "Chengdu"},
	})

	service := NewCityService(map[string]cityStrategies.CountStrategy{
		"startswith": cityStrategies.NewStartsWithStrategy(repository, cityCache.NewCountCache()),
		"bruteforce": cityStrategies.NewBruteforceStrategy(repository, cityCache.NewCountCache()),
	})

	tests := []struct {
		name     string
		letter   string
		strategy string
		expected int
	}{
		{name: "counts cities starting with C", letter: "C", strategy: "startswith", expected: 3},
		{name: "counts cities starting with R", letter: "R", strategy: "startswith", expected: 1},
		{name: "is case insensitive", letter: "c", strategy: "startswith", expected: 3},
		{name: "returns zero when there are no matches", letter: "X", strategy: "startswith", expected: 0},
		{name: "supports bruteforce strategy", letter: "c", strategy: "bruteforce", expected: 3},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := service.CountByLetter(context.Background(), tc.letter, tc.strategy)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if result != tc.expected {
				t.Fatalf("CountByLetter(%q,%q) = %d, expected %d", tc.letter, tc.strategy, result, tc.expected)
			}
		})
	}
}
