package city

import "context"

type MemoryCityRepository struct {
	cities []City
}

func NewMemoryCityRepository(cities []City) *MemoryCityRepository {
	return &MemoryCityRepository{
		cities: cities,
	}
}

func (r *MemoryCityRepository) GetCities(ctx context.Context) ([]City, error) {
	return r.cities, nil
}
