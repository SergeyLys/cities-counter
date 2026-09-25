package cityStrategies

import (
	"context"

	citycache "github.com/sergeylys/city-counter/backend/internal/city/cache"
	cityRepository "github.com/sergeylys/city-counter/backend/internal/city/repositories"
)

type StartsWithStrategy struct {
	repository cityRepository.CityRepository
	cache      *citycache.CountCache
}

func NewStartsWithStrategy(
	repository cityRepository.CityRepository,
	cache *citycache.CountCache,
) *StartsWithStrategy {
	return &StartsWithStrategy{
		repository: repository,
		cache:      cache,
	}
}

func (strategy *StartsWithStrategy) Count(
	ctx context.Context,
	letter string,
) (int, error) {
	if count, ok := strategy.cache.Get(letter); ok {
		return count, nil
	}

	maxRows := 1
	response, err := strategy.repository.GetCities(ctx, maxRows, &letter)
	if err != nil {
		return 0, err
	}

	count := response.TotalResultsCount

	strategy.cache.Set(letter, count)

	return count, nil
}
