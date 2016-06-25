package presets

import (
	"math"

	"github.com/bcokert/terragen/generator"
	"github.com/bcokert/terragen/noise"
	"github.com/bcokert/terragen/synthesizer"
	"github.com/bcokert/terragen/transformer"
)

// SpectralPresets is a map from preset names to spectral noise functions
var SpectralPresets map[string]func(int64, []float64) noise.Function1D = map[string]func(int64, []float64) noise.Function1D{
	"violet:1d": Violet1D,
	"blue:1d":   Blue1D,
	"white:1d":  White1D,
	"pink:1d":   Pink1D,
	"red:1d":    Red1D,
}

// A Spectral1DPreset creates octave noise functions with randomly phased sinusoidal noise functions.
// It combines octaves proportional to their frequencies, using a function f^X, where X is the weightExponent and corresponds to a normalized electromagnetic spectrum
// 2 = violet, 1 = blue, 0 = white, -1 = pink, -2 = red
type Spectral1DPreset func(seed int64, frequencies []float64) noise.Function1D

// Violet1D is a Spectral1DPreset with heavy emphasis on high frequencies
func Violet1D(seed int64, frequencies []float64) noise.Function1D {
	return spectral1D(seed, frequencies, 2)
}

// Blue1D is a Spectral1DPreset with light emphasis on high frequencies
func Blue1D(seed int64, frequencies []float64) noise.Function1D {
	return spectral1D(seed, frequencies, 1)
}

// White1D is a Spectral1DPreset with equal emphasis on all frequencies
func White1D(seed int64, frequencies []float64) noise.Function1D {
	return spectral1D(seed, frequencies, 0)
}

// Pink1D is a Spectral1DPreset with light emphasis on low frequencies
func Pink1D(seed int64, frequencies []float64) noise.Function1D {
	return spectral1D(seed, frequencies, -1)
}

// Red1D is a Spectral1DPreset with heavy emphasis on low frequencies
func Red1D(seed int64, frequencies []float64) noise.Function1D {
	return spectral1D(seed, frequencies, -2)
}

func spectral1D(seed int64, frequencies synthesizer.Frequencies, weightExponent float64) noise.Function1D {
	phaseFn := generator.Random1D(seed)

	weightFn := func(frequency float64) float64 {
		return math.Pow(frequency, weightExponent)
	}

	noiseFunctionGenerator := func(freq float64) noise.Function1D {
		return transformer.Sinusoid1D(phaseFn, freq)
	}

	return synthesizer.Octave1D(noiseFunctionGenerator, weightFn, frequencies)
}
