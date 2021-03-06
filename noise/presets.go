package noise

import (
	"math"

	tgmath "github.com/bcokert/terragen/math"
)

// A Preset is a function that takes a seed and frequencies and produces a pre-constructed noise function
type Preset func(source tgmath.Source, frequencies []float64) Function

// SpectralPresets is a map from preset names to Spectral Presets
// A Spectral Preset creates octave noise functions with randomly phased sinusoidal noise functions.
// It combines octaves proportional to their frequencies, using a function f^X, where X is the weightExponent and corresponds to a normalized electromagnetic spectrum
// 2 = violet, 1 = blue, 0 = white, -1 = pink, -2 = red
var SpectralPresets = map[string]Preset{
	"violet": Violet,
	"blue":   Blue,
	"white":  White,
	"pink":   Pink,
	"red":    Red,
}

// Violet is a Preset with heavy emphasis on high frequencies
func Violet(source tgmath.Source, frequencies []float64) Function {
	return spectral(source, frequencies, 2)
}

// Blue is a Preset with light emphasis on high frequencies
func Blue(source tgmath.Source, frequencies []float64) Function {
	return spectral(source, frequencies, 1)
}

// White is a Preset with equal emphasis on all frequencies
func White(source tgmath.Source, frequencies []float64) Function {
	return spectral(source, frequencies, 0)
}

// Pink is a Preset with light emphasis on low frequencies
func Pink(source tgmath.Source, frequencies []float64) Function {
	return spectral(source, frequencies, -1)
}

// Red is a Preset with heavy emphasis on low frequencies
func Red(source tgmath.Source, frequencies []float64) Function {
	return spectral(source, frequencies, -2)
}

func spectral(source tgmath.Source, frequencies []float64, weightExponent float64) Function {
	phaseFn := Random(source)

	weightFn := func(frequency float64) float64 {
		return math.Pow(frequency, weightExponent)
	}

	noiseFunctionGenerator := func(freq float64) Function {
		return Sinusoid(phaseFn, freq)
	}

	return Octave(noiseFunctionGenerator, weightFn, frequencies)
}

// LatticePresets is a map from preset names to Lattice Presets
// A Lattice Preset creates lattice gradient noise from a grid of random vectors
// It computes the influence of a given point at each of the surrounding grid coordinates, and then interpolates them to get a final result
var LatticePresets = map[string]Preset{
	"rawPerlin": RawPerlin,
}

// RawPerlin is a lattice preset that returns the output of a perlin generator, without modifying it
func RawPerlin(source tgmath.Source, frequencies []float64) Function {
	cache := tgmath.NewDefaultRandomGridCache(source)
	interpolator := tgmath.NewInterpolator(tgmath.DampCubicEase)
	return Perlin(cache, interpolator)
}
