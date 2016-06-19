package generator

import (
	"math/rand"

	"github.com/bcokert/terragen/noise"
)

// Generator1D is a basic 1D noise function generator
type Generator1D func(seed int64) noise.Function1D

// Random Builds a noise function that returns random floats in the range [0, 1]
func Random(seed int64) noise.Function1D {
	random := rand.New(rand.NewSource(seed))
	return func(t float64) float64 {
		return random.Float64()
	}
}
