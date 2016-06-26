package presets

import (
	"math"

	"github.com/bcokert/terragen/generator"
	"github.com/bcokert/terragen/noise"
	"github.com/bcokert/terragen/synthesizer"
	"github.com/bcokert/terragen/transformer"
)

// Spectral1DPresets is a map from preset names to 1D spectral noise functions
var Spectral1DPresets map[string]func(int64, []float64) noise.Function1D = map[string]func(int64, []float64) noise.Function1D{
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

func spectral1D(seed int64, frequencies []float64, weightExponent float64) noise.Function1D {
	phaseFn := generator.Random1D(seed)

	weightFn := func(frequency float64) float64 {
		return math.Pow(frequency, weightExponent)
	}

	noiseFunctionGenerator := func(freq float64) noise.Function1D {
		return transformer.Sinusoid1D(phaseFn, freq)
	}

	return synthesizer.Octave1D(noiseFunctionGenerator, weightFn, frequencies)
}

// Spectral2DPresets is a map from preset names to 2D spectral noise functions
var Spectral2DPresets map[string]func(int64, []float64) noise.Function2D = map[string]func(int64, []float64) noise.Function2D{
	"violet:2d": Violet2D,
	"blue:2d":   Blue2D,
	"white:2d":  White2D,
	"pink:2d":   Pink2D,
	"red:2d":    Red2D,
}

// A Spectral2DPreset creates octave noise functions with randomly phased 2D sinusoidal noise functions.
// It combines octaves proportional to their frequencies, using a function f^X, where X is the weightExponent and corresponds to a normalized electromagnetic spectrum
// 2 = violet, 1 = blue, 0 = white, -1 = pink, -2 = red
type Spectral2DPreset func(seed int64, frequencies []float64) noise.Function2D

// Violet2D is a Spectral2DPreset with heavy emphasis on high frequencies
func Violet2D(seed int64, frequencies []float64) noise.Function2D {
	return spectral2D(seed, frequencies, 2)
}

// Blue2D is a Spectral2DPreset with light emphasis on high frequencies
func Blue2D(seed int64, frequencies []float64) noise.Function2D {
	return spectral2D(seed, frequencies, 1)
}

// White2D is a Spectral2DPreset with equal emphasis on all frequencies
func White2D(seed int64, frequencies []float64) noise.Function2D {
	return spectral2D(seed, frequencies, 0)
}

// Pink2D is a Spectral2DPreset with light emphasis on low frequencies
func Pink2D(seed int64, frequencies []float64) noise.Function2D {
	return spectral2D(seed, frequencies, -1)
}

// Red2D is a Spectral2DPreset with heavy emphasis on low frequencies
func Red2D(seed int64, frequencies []float64) noise.Function2D {
	return spectral2D(seed, frequencies, -2)
}

func spectral2D(seed int64, frequencies []float64, weightExponent float64) noise.Function2D {
	phaseFn := generator.Random2D(seed)

	weightFn := func(frequency float64) float64 {
		return math.Pow(frequency, weightExponent)
	}

	noiseFunctionGenerator := func(freq float64) noise.Function2D {
		return transformer.Sinusoid2D(phaseFn, freq)
	}

	return synthesizer.Octave2D(noiseFunctionGenerator, weightFn, frequencies)
}
