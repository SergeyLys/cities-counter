package cityRepository

import (
	"context"
	"strings"

	cityEntity "github.com/sergeylys/city-counter/backend/internal/city/entity"
)

type MemoryCityRepository struct {
	cities []cityEntity.City
}

func NewMemoryCityRepository(cities []cityEntity.City) *MemoryCityRepository {
	return &MemoryCityRepository{
		cities: cities,
	}
}

func (repository *MemoryCityRepository) GetCities(ctx context.Context, maxRows int, nameStartsWith *string) (GetCitiesResponse, error) {

	// Simulate db filtering

	filtered := make([]cityEntity.City, 0, len(repository.cities))

	for _, city := range repository.cities {
		if nameStartsWith != nil {
			if !strings.HasPrefix(strings.ToLower(city.Name), strings.ToLower(*nameStartsWith)) {
				continue
			}
		}

		filtered = append(filtered, city)
	}

	response := GetCitiesResponse{
		TotalResultsCount: len(filtered),
		Cities:            filtered,
	}

	return response, nil
}
