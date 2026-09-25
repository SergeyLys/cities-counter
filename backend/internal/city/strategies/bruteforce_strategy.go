package cityStrategies

import (
	"context"
	"strings"

	citycache "github.com/sergeylys/city-counter/backend/internal/city/cache"
	cityRepository "github.com/sergeylys/city-counter/backend/internal/city/repositories"
)

type BruteforceStrategy struct {
	repository cityRepository.CityRepository
	cache      *citycache.CountCache
}

func NewBruteforceStrategy(
	repository cityRepository.CityRepository,
	cache *citycache.CountCache,
) *BruteforceStrategy {
	return &BruteforceStrategy{
		repository: repository,
		cache:      cache,
	}
}

func (strategy *BruteforceStrategy) Count(
	ctx context.Context,
	letter string,
) (int, error) {
	if count, ok := strategy.cache.Get(letter); ok {
		return count, nil
	}

	maxRows := 1000
	response, err := strategy.repository.GetCities(ctx, maxRows, nil)

	if err != nil {
		return 0, err
	}

	count := 0

	for _, city := range response.Cities {
		if strings.HasPrefix(strings.ToLower(city.Name), strings.ToLower(letter)) {
			count++
		}
	}

	strategy.cache.Set(letter, count)

	return count, nil
}
