package city

import (
	"context"
	"unicode"
)

type CityService struct {
	repository CityRepository
}

func NewCityService(repository CityRepository) *CityService {
	return &CityService{
		repository: repository,
	}
}

func (service *CityService) CountByLetter(ctx context.Context, letter rune) (int, error) {
	count := 0

	cities, err := service.repository.GetCities(ctx)

	if err != nil {
		return 0, err
	}

	for _, city := range cities {
		runes := []rune(city.Name)

		if len(runes) == 0 {
			continue
		}

		if unicode.ToLower(runes[0]) == unicode.ToLower(letter) {
			count++
		}
	}

	return count, nil
}
