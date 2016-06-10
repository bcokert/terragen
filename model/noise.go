package model

import "github.com/bcokert/terragen/noise"

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
func (noise *Noise) Generate(from, to []float64, resolution int, noiseFunction noise.Parametric1D) {
	noise.RawNoise = map[string][]float64{}

	numSamples := int(to[0]-from[0]+1) * resolution
	noise.RawNoise["x"] = make([]float64, 0, numSamples)
	noise.RawNoise["value"] = make([]float64, 0, numSamples)

	for lattice := from[0]; lattice < to[0]; lattice++ {
		for res := 0; res < resolution; res++ {
			x := lattice + float64(res)/float64(resolution)
			noise.RawNoise["x"] = append(noise.RawNoise["x"], x)
			noise.RawNoise["value"] = append(noise.RawNoise["value"], noiseFunction(x))
		}
	}

	noise.From = from
	noise.To = to
	noise.Resolution = resolution
}

// Equals returns true if the other Noise is equal
func (noise *Noise) Equals(o *Noise) bool {
	if len(noise.RawNoise) != len(o.RawNoise) {
		return false
	}

	for dimension, values := range o.RawNoise {
		if _, ok := noise.RawNoise[dimension]; !ok {
			return false
		}
		if len(noise.RawNoise[dimension]) != len(values) {
			return false
		}
		for i, value := range values {
			if noise.RawNoise[dimension][i] != value {
				return false
			}
		}
	}

	for i, v := range o.From {
		if noise.From[i] != v {
			return false
		}
	}

	for i, v := range o.To {
		if noise.To[i] != v {
			return false
		}
	}

	return noise.Resolution == o.Resolution && noise.NoiseFunction == o.NoiseFunction
}
