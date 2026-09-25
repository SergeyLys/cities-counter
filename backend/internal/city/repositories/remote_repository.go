package cityRepository

import (
	"context"

	cityEntity "github.com/sergeylys/city-counter/backend/internal/city/entity"
	"github.com/sergeylys/city-counter/backend/internal/geonames"
)

type RemoteCityRepository struct {
	client *geonames.GeonamesClient
}

func NewRemoteCityRepository(client *geonames.GeonamesClient) *RemoteCityRepository {
	return &RemoteCityRepository{
		client: client,
	}
}

func (repository *RemoteCityRepository) GetCities(ctx context.Context, maxRows int, nameStartsWith *string) (GetCitiesResponse, error) {
	params := repository.client.WithBaseParams(
		maxRows,
		nameStartsWith,
	)

	response, err := repository.client.Search(ctx, params)
	if err != nil {
		return GetCitiesResponse{}, err
	}

	cities := make([]cityEntity.City, 0, len(response.Geonames))

	for _, place := range response.Geonames {
		cities = append(cities, cityEntity.City{
			Name: place.Name,
		})
	}

	return GetCitiesResponse{
		TotalResultsCount: response.TotalResultsCount,
		Cities:            cities,
	}, nil
}
