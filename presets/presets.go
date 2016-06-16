package presets

import (
	"github.com/bcokert/terragen/noisefunction"
	"github.com/bcokert/terragen/synthesizer"
	"github.com/bcokert/terragen/transformer"
)

// WhiteNoise1D is a SinceOctave1D generator that has equal weights for all frequencies
func WhiteNoise1D(phaseFn noisefunction.Function1D, frequencies []float64) noisefunction.Function1D {
	weightFn := func(frequency float64) float64 {
		return 1
	}

	noiseFunctionGenerator := func(freq float64) noisefunction.Function1D {
		return transformer.Sinusoid1D(phaseFn, freq)
	}

	return synthesizer.Octave1D(noiseFunctionGenerator, weightFn, frequencies)
}

// RedNoise1D is a SinceOctave1D generator that has larger weights for lower frequencies
// The ratio is 1/f*f
func RedNoise1D(phaseFn noisefunction.Function1D, frequencies []float64) noisefunction.Function1D {
	weightFn := func(frequency float64) float64 {
		return 1 / (frequency * frequency)
	}

	noiseFunctionGenerator := func(freq float64) noisefunction.Function1D {
		return transformer.Sinusoid1D(phaseFn, freq)
	}

	return synthesizer.Octave1D(noiseFunctionGenerator, weightFn, frequencies)
}
