package noise

import (
	"math"
)

// Noise represents generated noise, typically from GetNoise
type Noise struct {
	Values        []float64 `json:"values"`
	From          []int     `json:"from"`
	To            []int     `json:"to"`
	Resolution    int       `json:"resolution"`
	NoiseFunction string    `json:"noiseFunction"`
}

// NewNoise creates a new Noise object
func NewNoise(noiseFunction string) *Noise {
	return &Noise{
		NoiseFunction: noiseFunction,
	}
}

// Generate populates the RawNoise and related fields of this noise, by iterating over the range and calling the given noise function
func (noise *Noise) Generate(from, to []int, resolution int, noiseFunction Function) {
	noise.Values = []float64{}

	// initialize storage for parameters and values
	// for an n dimensional lattice, the number of points is range[dimension0]*resolution * range[dimension1]*resolution * ...
	numParams := len(from)
	numTotalSamples := 0
	for i := range from {
		numSamplesInDimension := to[i] - from[i]
		numTotalSamples *= numSamplesInDimension
	}
	numTotalSamples *= resolution
	noise.Values = make([]float64, 0, numTotalSamples)

	var eachSample func(point []float64, dimensionIndex int)
	eachSample = func(point []float64, dimensionIndex int) {
		if dimensionIndex == numParams {
			noise.Values = append(noise.Values, noiseFunction(point))
		} else {
			for i := from[dimensionIndex]; i < to[dimensionIndex]; i++ {
				for j := 0; j < resolution; j++ {
					eachSample(append(point[:], float64(i)+float64(j)/float64(resolution)), dimensionIndex+1)
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
	for i, v := range other.Values {
		if math.Abs(noise.Values[i]-v) > 0.00000000000001 {
			return false
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
