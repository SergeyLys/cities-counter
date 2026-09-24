package city

import (
	"context"
	"testing"

	"github.com/sergeylys/city-counter/backend/internal/city"
)

func TestCountByLetter(t *testing.T) {
	cities := []city.City{
		{Name: "Rio de Janeiro"},
		{Name: "Cairo"},
		{Name: "Chongqing"},
		{Name: "Chengdu"},
	}

	repository := city.NewMemoryCityRepository(cities)
	service := city.NewCityService(repository)

	tests := []struct {
		letter   rune
		name     string
		expected int
	}{
		{
			name:     "counts cities starting with C",
			letter:   'C',
			expected: 3,
		},
		{
			name:     "counts cities starting with R",
			letter:   'R',
			expected: 1,
		},
		{
			name:     "is case insensitive",
			letter:   'c',
			expected: 3,
		},
		{
			name:     "returns zero when there are no matches",
			letter:   'X',
			expected: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := service.CountByLetter(context.Background(), tc.letter)

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if result != tc.expected {
				t.Errorf("CountByLetter(%q) = %d, expected %d",
					tc.letter, result, tc.expected)
			}
		})
	}
}
