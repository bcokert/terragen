package transformer

import (
	"math"

	"github.com/bcokert/terragen/noise"
)

// A Transformer transforms a noise function into a new one, at the given frequency
type Transformer func(fn noise.Function, freq float64) noise.Function

// Sinusoid transforms a noise function into a phase shifted sinusoid, at the specified frequency
func Sinusoid(phaseFn noise.Function, freq float64) noise.Function {
	return func(t []float64) float64 {
		product := 1.0
		for _, tx := range t {
			product = product * math.Sin(2*math.Pi*freq*tx+phaseFn(t))
		}
		return product
	}
}
