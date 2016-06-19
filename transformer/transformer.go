package transformer

import (
	"math"

	"github.com/bcokert/terragen/noise"
)

// A Transformer1D is a function that transforms a noise function into a new one, with the given frequency
type Transformer1D func(fn noise.Function1D, freq float64) noise.Function1D

// Sinusoid1D transforms a 1D noise function into a phase shifted sinusoid, at the specified frequency
func Sinusoid1D(phaseFn noise.Function1D, freq float64) noise.Function1D {
	return func(t float64) float64 {
		return math.Sin(2*math.Pi*freq*t + phaseFn(t))
	}
}
