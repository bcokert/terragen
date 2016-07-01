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
		sum := 0.0
		for _, tx := range t {
			sum += tx
		}
		return math.Sin(2*math.Pi*freq*sum + phaseFn(t))
	}
}
