package presets

import (
	"math"

	"github.com/bcokert/terragen/generator"
	"github.com/bcokert/terragen/noise"
	"github.com/bcokert/terragen/synthesizer"
	"github.com/bcokert/terragen/transformer"
)

// A Preset is a function that takes a seed and frequencies and produces a pre-constructed noise function
type Preset func(seed int64, frequencies []float64) noise.Function

// SpectralPresets is a map from preset names to Spectral Presets
// A Spectral Preset creates octave noise functions with randomly phased sinusoidal noise functions.
// It combines octaves proportional to their frequencies, using a function f^X, where X is the weightExponent and corresponds to a normalized electromagnetic spectrum
// 2 = violet, 1 = blue, 0 = white, -1 = pink, -2 = red
var SpectralPresets map[string]Preset = map[string]Preset{
	"violet": Violet,
	"blue":   Blue,
	"white":  White,
	"pink":   Pink,
	"red":    Red,
}

// Violet is a Preset with heavy emphasis on high frequencies
func Violet(seed int64, frequencies []float64) noise.Function {
	return spectral(seed, frequencies, 2)
}

// Blue is a Preset with light emphasis on high frequencies
func Blue(seed int64, frequencies []float64) noise.Function {
	return spectral(seed, frequencies, 1)
}

// White is a Preset with equal emphasis on all frequencies
func White(seed int64, frequencies []float64) noise.Function {
	return spectral(seed, frequencies, 0)
}

// Pink is a Preset with light emphasis on low frequencies
func Pink(seed int64, frequencies []float64) noise.Function {
	return spectral(seed, frequencies, -1)
}

// Red is a Preset with heavy emphasis on low frequencies
func Red(seed int64, frequencies []float64) noise.Function {
	return spectral(seed, frequencies, -2)
}

func spectral(seed int64, frequencies []float64, weightExponent float64) noise.Function {
	phaseFn := generator.Random(seed)

	weightFn := func(frequency float64) float64 {
		return math.Pow(frequency, weightExponent)
	}

	noiseFunctionGenerator := func(freq float64) noise.Function {
		return transformer.Sinusoid(phaseFn, freq)
	}

	return synthesizer.Octave(noiseFunctionGenerator, weightFn, frequencies)
}
