package transformer

import (
	"math"

	"github.com/bcokert/terragen/noise"
)

// A Transformer1D transforms a 1d noise function into a new one, at the given frequency
type Transformer1D func(fn noise.Function1D, freq float64) noise.Function1D

// Sinusoid1D transforms a 1D noise function into a phase shifted sinusoid, at the specified frequency
func Sinusoid1D(phaseFn noise.Function1D, freq float64) noise.Function1D {
	return func(t float64) float64 {
		return math.Sin(2*math.Pi*freq*t + phaseFn(t))
	}
}

// A Transformer2D transforms a 2d noise function into a new one, at the given frequency
type Transformer2D func(fn noise.Function2D, freq float64) noise.Function2D

// Sinusoid2D transforms a 2D noise function into a phase shifted sinusoid, at the specified frequency
func Sinusoid2D(phaseFn noise.Function2D, freq float64) noise.Function2D {
	return func(tx, ty float64) float64 {
		return math.Sin(2*math.Pi*freq*(tx+ty) + phaseFn(tx, ty))
	}
}
