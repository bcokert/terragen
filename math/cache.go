package math

import (
	"fmt"
)

// A GridCache is a cache for vectors keyed by 2d grid coordinates that guarantees a result even on misses
type GridCache interface {
	Get(x, y int) Vec2
}

// DefaultRandomGridCache is a standard implementation of GridCache
type DefaultRandomGridCache struct {
	grid   map[string]Vec2
	random Source
}

// NewDefaultRandomGridCache creates a new DefaultRandomGridCache with the given random number generator
func NewDefaultRandomGridCache(random Source) GridCache {
	cache := &DefaultRandomGridCache{random: random}

	// Pre-populate the cache with 20x20 grid points, from -9 to 10 inclusive
	cache.grid = make(map[string]Vec2, 400)
	for x := -9; x <= 10; x++ {
		for y := -9; y <= 10; y++ {
			cache.grid[fmt.Sprintf("%d,%d", x, y)] = RandomDirectionVec2(random)
		}
	}

	return cache
}

// Get returns the cached value for the default cache, or generates one, caching it and returning it
func (cache *DefaultRandomGridCache) Get(x, y int) Vec2 {
	hashKey := fmt.Sprintf("%d,%d", x, y)
	vector, ok := cache.grid[hashKey]

	if !ok {
		vector = RandomDirectionVec2(cache.random)
		cache.grid[hashKey] = vector
	}

	return vector
}
