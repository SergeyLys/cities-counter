package citycache

import "sync"

type CountCache struct {
	mu     sync.RWMutex
	values map[string]int
}

func NewCountCache() *CountCache {
	return &CountCache{
		values: make(map[string]int),
	}
}

func (c *CountCache) Get(letter string) (int, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	count, exists := c.values[letter]

	return count, exists
}

func (c *CountCache) Set(letter string, count int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.values[letter] = count
}
