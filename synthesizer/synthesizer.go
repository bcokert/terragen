package synthesizer

import "github.com/bcokert/terragen/noise"

type WeightFunction func(freq float64) float64
type NoiseFunctionGenerator func(freq float64) noise.Function1D
type Frequencies []float64

type Synthesizer1D func(fnGenerator NoiseFunctionGenerator, weightFn WeightFunction, frequencies Frequencies) noise.Function1D

// Octave1D synthesizes noise functions by combining source noise functions at different frequencies, with weights corresponding to their frequencies
func Octave1D(noiseFnGenerator NoiseFunctionGenerator, weightFn WeightFunction, frequencies Frequencies) noise.Function1D {
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
