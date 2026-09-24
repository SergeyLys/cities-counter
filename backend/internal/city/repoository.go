package city

import "context"

type CityRepository interface {
	GetCities(ctx context.Context) ([]City, error)
}