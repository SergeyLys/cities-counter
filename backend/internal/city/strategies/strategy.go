package cityStrategies

import "context"

type CountStrategy interface {
	Count(ctx context.Context, letter string) (int, error)
}
