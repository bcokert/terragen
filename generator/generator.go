package generator

import (
	"math/rand"

	"github.com/bcokert/terragen/noise"
)

// A Generator1D is a 1D noise function generator
type Generator1D func(seed int64) noise.Function1D

// Random1D Builds a 1d noise function that returns random floats in the range [0, 1]
func Random1D(seed int64) noise.Function1D {
	random := rand.New(rand.NewSource(seed))
	return func(t float64) float64 {
		return random.Float64()
	}
}

// A Generator2D is a 2D noise function generator
type Generator2D func(seed int64) noise.Function2D

// Random2D Builds a 2d noise function that returns random floats in the rang [0, 1]
func Random2D(seed int64) noise.Function2D {
	random := rand.New(rand.NewSource(seed))
	return func(tx, ty float64) float64 {
		return random.Float64()
	}
}
