package cityRepository

import (
	"context"

	cityEntity "github.com/sergeylys/city-counter/backend/internal/city/entity"
)

type GetCitiesResponse struct {
	TotalResultsCount int
	Cities            []cityEntity.City
}

type CityRepository interface {
	GetCities(ctx context.Context, maxRows int, nameStartsWith *string) (GetCitiesResponse, error)
}
