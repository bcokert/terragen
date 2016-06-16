package transformer

import (
	"math"

	"github.com/bcokert/terragen/noisefunction"
)

// Sinusoid1D transforms a 1D noise function into a phase shifted sinusoid, at the specified frequency
func Sinusoid1D(phaseFn noisefunction.Function1D, freq float64) noisefunction.Function1D {
	return func(t float64) float64 {
		return math.Sin(2*math.Pi*freq*t + phaseFn(t))
	}
}
