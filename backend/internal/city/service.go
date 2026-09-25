package city

import (
	"context"
	"fmt"
	"strings"

	cityStrategies "github.com/sergeylys/city-counter/backend/internal/city/strategies"
)

type CityService struct {
	strategies map[string]cityStrategies.CountStrategy
}

func NewCityService(strategies map[string]cityStrategies.CountStrategy) *CityService {
	return &CityService{
		strategies: strategies,
	}
}

func (service *CityService) CountByLetter(ctx context.Context, letter string, strategyName string) (int, error) {
	strategy, exists := service.strategies[strategyName]

	if !exists {
		return 0, fmt.Errorf(
			"unknown strategy: %s",
			strategyName,
		)
	}

	return strategy.Count(
		ctx,
		strings.ToLower(letter),
	)
}
