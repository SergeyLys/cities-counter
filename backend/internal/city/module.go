package city

import (
	"net/http"
	"time"

	cityCache "github.com/sergeylys/city-counter/backend/internal/city/cache"
	cityRepositories "github.com/sergeylys/city-counter/backend/internal/city/repositories"
	cityStrategies "github.com/sergeylys/city-counter/backend/internal/city/strategies"
	"github.com/sergeylys/city-counter/backend/internal/config"
	"github.com/sergeylys/city-counter/backend/internal/geonames"
)

func CityModule() *CityHandler {
	cfg := config.Load()
	client := geonames.NewGeonamesClient(
		cfg.GeoNamesURL,
		cfg.GeoNamesUsername,
		&http.Client{
			Timeout: 10 * time.Second,
		},
	)

	repository := cityRepositories.NewRemoteCityRepository(client)

	startsWithStrategy := cityStrategies.NewStartsWithStrategy(
		repository,
		cityCache.NewCountCache(),
	)

	bruteforceStrategy := cityStrategies.NewBruteforceStrategy(
		repository,
		cityCache.NewCountCache(),
	)

	service := NewCityService(
		map[string]cityStrategies.CountStrategy{
			"startswith": startsWithStrategy,
			"bruteforce": bruteforceStrategy,
		},
	)

	return NewCityHandler(service)
}
