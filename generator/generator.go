package generator

import (
	"math/rand"

	"github.com/bcokert/terragen/noise"
)

// A Generator creates a specifically seeded noise function
type Generator func(seed int64) noise.Function

// Random Builds a noise function that returns random floats in the range [0, 1]
func Random(seed int64) noise.Function {
	random := rand.New(rand.NewSource(seed))
	return func(t []float64) float64 {
		return random.Float64()
	}
}
