package model

import (
	"strconv"

	"math"

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
	numParams := len(from)
	numTotalSamples := 0
	for i := range from {
		numSamples := int(to[0]-from[0]+1) * resolution
		numTotalSamples *= numSamples
		noise.RawNoise["t"+strconv.Itoa(i+1)] = make([]float64, 0, numSamples)
	}
	noise.RawNoise["value"] = make([]float64, 0, numTotalSamples)

	storeNoise := func(latPoints []float64) {
		for res := 0; res < resolution; res++ {
			params := make([]float64, 0, numParams)
			resParam := float64(res) / float64(resolution)
			for i, latPoint := range latPoints {
				key := "t" + strconv.Itoa(i+1)
				lat := latPoint + resParam
				params = append(params, lat)
				noise.RawNoise[key] = append(noise.RawNoise[key], lat)
			}
			noise.RawNoise["value"] = append(noise.RawNoise["value"], noiseFunction(params))
		}
	}

	var eachLatticePoint func(latticePoints []float64, dimensionIndex int)
	eachLatticePoint = func(latticePoints []float64, dimensionIndex int) {
		if dimensionIndex == numParams {
			storeNoise(latticePoints)
		} else {
			for lat := from[dimensionIndex]; lat < to[dimensionIndex]; lat++ {
				newLatticePoints := append(latticePoints[:], lat)
				eachLatticePoint(newLatticePoints, dimensionIndex+1)
			}
		}
	}

	eachLatticePoint([]float64{}, 0)

	noise.From = from
	noise.To = to
	noise.Resolution = resolution
}

// Equals returns true if the other Noise is equal
func (noise *Noise) Equals(other *Noise) bool {
	if len(noise.RawNoise) != len(other.RawNoise) {
		return false
	}

	for dimension, values := range other.RawNoise {
		if _, ok := noise.RawNoise[dimension]; !ok {
			return false
		}
		if len(noise.RawNoise[dimension]) != len(values) {
			return false
		}
		for i, value := range values {
			if math.Abs(noise.RawNoise[dimension][i]-value) > 0.00000000000001 {
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
