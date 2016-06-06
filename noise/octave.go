package noise

import (
	"math"

	"github.com/bcokert/terragen/random"
)

// Octave1D creates Parametric1D functions with custom noise functions and weights
func Octave1D(generator OctaveParametric1D, weightFn OctaveWeight1D, frequencies []float64) Parametric1D {
	var noiseFunctions []Parametric1D
	for _, freq := range frequencies {
		freqFn := generator(freq)
		weight := weightFn(freq) / float64(len(frequencies))
		noiseFunctions = append(noiseFunctions, func(t float64) float64 {
			return freqFn(t) * weight
		})
	}

	return func(t float64) (sum float64) {
		for _, fn := range noiseFunctions {
			sum += fn(t)
		}
		return sum
	}
}

// SineOctave1D is an Octave1D generator that combines sine waves, using random phases, and custom weights
func SineOctave1D(phaseFn random.NormalFunction, weightFn OctaveWeight1D, frequencies []float64) Parametric1D {
	noiseFunctionGenerator := func(freq float64) Parametric1D {
		return func(t float64) float64 {
			return math.Sin(2*math.Pi*freq*t + phaseFn())
		}
	}

	return Octave1D(noiseFunctionGenerator, weightFn, frequencies)
}

// WhiteNoise1D is a SinceOctave1D generator that has equal weights for all frequencies
func WhiteNoise1D(phaseFn random.NormalFunction, frequencies []float64) Parametric1D {
	weightFn := func(frequency float64) float64 {
		return 1
	}

	return SineOctave1D(phaseFn, weightFn, frequencies)
}

// RedNoise1D is a SinceOctave1D generator that has larger weights for lower frequencies
// The ratio is 1/f*f
func RedNoise1D(phaseFn random.NormalFunction, frequencies []float64) Parametric1D {
	weightFn := func(frequency float64) float64 {
		return 1 / (frequency * frequency)
	}

	return SineOctave1D(phaseFn, weightFn, frequencies)
}
