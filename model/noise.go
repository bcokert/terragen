package model

import (
	"strconv"

	"math"

	"github.com/bcokert/terragen/log"
	"github.com/bcokert/terragen/noise"
)

// Noise represents generated noise, typically from GetNoise
type Noise struct {
	RawNoise      map[string][]float64 `json:"rawNoise"`
	From          []float64            `json:"from"`
	To            []float64            `json:"to"`
	Resolution    int                  `json:"resolution"`
	NoiseFunction string               `json:"noiseFunction"`
}

// NewNoise creates a new Noise object
func NewNoise(noiseFunction string) *Noise {
	return &Noise{
		NoiseFunction: noiseFunction,
	}
}

// Generate populates the RawNoise and related fields of this noise, by iterating over the range and calling the given noise function
func (noise *Noise) Generate(from, to []float64, resolution int, noiseFunction noise.Function) {
	noise.RawNoise = map[string][]float64{}

	// initialize storage for parameters and values
	// for an n dimensional lattice, the number of points is range[dimension0]*resolution * range[dimension1]*resolution * ...
	numParams := len(from)
	numTotalSamples := 0
	for i := range from {
		numSamplesInDimension := int(to[i] - from[i])
		numTotalSamples *= numSamplesInDimension
	}
	numTotalSamples *= resolution
	for i := range from {
		noise.RawNoise["t"+strconv.Itoa(i+1)] = make([]float64, 0, numTotalSamples)
	}
	noise.RawNoise["value"] = make([]float64, 0, numTotalSamples)

	var eachSample func(point []float64, dimensionIndex int)
	eachSample = func(point []float64, dimensionIndex int) {
		if dimensionIndex == numParams {
			for i, component := range point {
				key := "t" + strconv.Itoa(i+1)
				noise.RawNoise[key] = append(noise.RawNoise[key], component)
			}
			noise.RawNoise["value"] = append(noise.RawNoise["value"], noiseFunction(point))
		} else {
			for i := from[dimensionIndex]; i < to[dimensionIndex]; i++ {
				for j := 0; j < resolution; j++ {
					eachSample(append(point[:], i+float64(j)/float64(resolution)), dimensionIndex+1)
				}
			}
		}
	}

	eachSample([]float64{}, 0)

	noise.From = from
	noise.To = to
	noise.Resolution = resolution
}

// IsEqual returns true if the other Noise is equal to this one
func (noise *Noise) IsEqual(other *Noise) bool {
	if len(noise.RawNoise) != len(other.RawNoise) {
		return false
	}

	for dimension, values := range other.RawNoise {
		if _, ok := noise.RawNoise[dimension]; !ok {
			log.Debug("The other did not have dimension %v", dimension)
			return false
		}
		if len(noise.RawNoise[dimension]) != len(values) {
			log.Debug("Dimension %v length: expected %v, found %v", dimension, len(noise.RawNoise[dimension]), len(values))
			return false
		}
		for i, value := range values {
			if math.Abs(noise.RawNoise[dimension][i]-value) > 0.00000000000001 {
				log.Debug("Dimension %v at index %d: expected %v, found %v", dimension, i, noise.RawNoise[dimension][i], value)
				return false
			}
		}
	}

	for i, v := range other.From {
		if noise.From[i] != v {
			return false
		}
	}

	for i, v := range other.To {
		if noise.To[i] != v {
			return false
		}
	}

	return noise.Resolution == other.Resolution && noise.NoiseFunction == other.NoiseFunction
}
