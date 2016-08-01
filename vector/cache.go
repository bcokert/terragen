package vector

import (
	"fmt"
)

// RandomSource is an interface that provides a Float64 method that should return random floats between 0 and 1
type RandomSource interface {
	Float64() float64
}

// A RandomGridCache is a cache for random vectors that guarantees a result even on misses
type RandomGridCache interface {
	Get(x, y int) Vec2
}

// DefaultRandomGridCache is a standard implementation of RandomGridCache
type DefaultRandomGridCache struct {
	grid   map[string]Vec2
	random RandomSource
}

// NewDefaultRandomGridCache creates a new DefaultRandomGridCache with the given random number generator
func NewDefaultRandomGridCache(random RandomSource) RandomGridCache {
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
