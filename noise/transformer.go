package noise

import (
	"math"
)

// A Transformer transforms a noise function into a new one, at the given frequency
type Transformer func(fn Function, freq float64) Function

// Sinusoid transforms a noise function into a phase shifted sinusoid, at the specified frequency
func Sinusoid(phaseFn Function, freq float64) Function {
	phase := phaseFn([]float64{})
	return func(t []float64) float64 {
		product := 1.0
		for _, tx := range t {
			product = product * math.Sin(2*math.Pi*freq*tx+phase)
		}
		return product
	}
}
