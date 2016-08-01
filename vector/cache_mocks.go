package vector

// A MockGridCache always misses, and simply returns a normalized vector whos initial components were the grid coordinates
type MockGridCache struct{}

// Get returns a vector with components equal the input grid coordinates, which is then normalized
func (cache *MockGridCache) Get(x, y int) Vec2 {
	vec := NewVec2(float64(x), float64(y))
	if x == 0 && y == 0 {
		vec[0] = 1
	}
	vec.Normalize()
	return vec
}
