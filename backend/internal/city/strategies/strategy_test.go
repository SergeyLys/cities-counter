package cityStrategies

import (
	"context"
	"testing"

	cityCache "github.com/sergeylys/city-counter/backend/internal/city/cache"
	cityEntity "github.com/sergeylys/city-counter/backend/internal/city/entity"
	cityRepository "github.com/sergeylys/city-counter/backend/internal/city/repositories"
)

func TestStartsWithStrategyCountUsesMemoryRepository(t *testing.T) {
	repository := cityRepository.NewMemoryCityRepository([]cityEntity.City{
		{Name: "Aachen"},
		{Name: "Berlin"},
		{Name: "Athens"},
		{Name: "Auckland"},
	})

	strategy := NewStartsWithStrategy(repository, cityCache.NewCountCache())

	count, err := strategy.Count(context.Background(), "a")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if count != 3 {
		t.Fatalf("expected 3 cities starting with A, got %d", count)
	}
}

func TestBruteforceStrategyCountUsesMemoryRepository(t *testing.T) {
	repository := cityRepository.NewMemoryCityRepository([]cityEntity.City{
		{Name: "Aachen"},
		{Name: "Berlin"},
		{Name: "Athens"},
		{Name: "Auckland"},
	})

	strategy := NewBruteforceStrategy(repository, cityCache.NewCountCache())

	count, err := strategy.Count(context.Background(), "a")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if count != 3 {
		t.Fatalf("expected 3 cities starting with A, got %d", count)
	}
}
