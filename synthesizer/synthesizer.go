package synthesizer

import "github.com/bcokert/terragen/noisefunction"

// Octave1D synthesizes noise functions by combining source noise functions at different frequencies, with weights corresponding to their frequencies
func Octave1D(noiseFnGenerator func(freq float64) noisefunction.Function1D, weightFn func(freq float64) float64, frequencies []float64) noisefunction.Function1D {
	var noiseFunctions []noisefunction.Function1D
	for _, freq := range frequencies {
		freqFn := noiseFnGenerator(freq)
		weight := weightFn(freq) / float64(len(frequencies))
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
