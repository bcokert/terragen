package synthesizer

import "github.com/bcokert/terragen/noise"

// A WeightFunction determines a weight for an octave given the frequency of the octave
type WeightFunction func(freq float64) float64

// A NoiseFunctionGenerator1D produces 1D noise functions at the given frequency
type NoiseFunctionGenerator1D func(freq float64) noise.Function1D

// A Synthesizer1D synthesizes 1D noise functions from existing 1D noise functions at different frequencies and weights
// It should take pre-constructed noise functions who's only remaining variable is the frequency
type Synthesizer1D func(fnGenerator NoiseFunctionGenerator1D, weightFn WeightFunction, frequencies []float64) noise.Function1D

// Octave1D synthesizes 1D noise functions by linearly combining source 2D noise functions
func Octave1D(noiseFnGenerator NoiseFunctionGenerator1D, weightFn WeightFunction, frequencies []float64) noise.Function1D {
	var noiseFunctions []noise.Function1D
	for _, freq := range frequencies {
		freqFn := noiseFnGenerator(freq)
		weight := weightFn(freq)
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

// A NoiseFunctionGenerator2D produces 2D noise functions at the given frequency
// It should take pre-constructed noise functions who's only remaining variable is the frequency
type NoiseFunctionGenerator2D func(freq float64) noise.Function2D

// A Synthesizer2D synthesizes 2D noise functions from existing 2D noise functions at different frequencies and weights
type Synthesizer2D func(fnGenerator NoiseFunctionGenerator1D, weightFn WeightFunction, frequencies []float64) noise.Function2D

// Octave2D synthesizes 2D noise functions by linearly combining source 2D noise functions
func Octave2D(noiseFnGenerator NoiseFunctionGenerator2D, weightFn WeightFunction, frequencies []float64) noise.Function2D {
	var noiseFunctions []noise.Function2D
	for _, freq := range frequencies {
		freqFn := noiseFnGenerator(freq)
		weight := weightFn(freq)
		noiseFunctions = append(noiseFunctions, func(tx, ty float64) float64 {
			return freqFn(tx, ty) * weight
		})
	}

	return func(tx, ty float64) (sum float64) {
		for _, fn := range noiseFunctions {
			sum += fn(tx, ty)
		}
		return sum
	}
}
