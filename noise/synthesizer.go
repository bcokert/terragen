package noise

// A WeightFunction determines a weight for an octave given the frequency of the octave
type WeightFunction func(freq float64) float64

// A NoiseFunctionGenerator produces noise functions at the given frequency
// It should take pre-constructed noise functions who's only remaining variable is the frequency
type NoiseFunctionGenerator func(freq float64) Function

// A Synthesizer synthesizes noise functions from existing noise functions at different frequencies and weights
type Synthesizer func(fnGenerator NoiseFunctionGenerator, weightFn WeightFunction, frequencies []float64) Function

// Octave synthesizes noise functions by linearly combining source noise functions
func Octave(noiseFnGenerator NoiseFunctionGenerator, weightFn WeightFunction, frequencies []float64) Function {
	var noiseFunctions []Function
	for _, freq := range frequencies {
		freqFn := noiseFnGenerator(freq)
		weight := weightFn(freq)
		noiseFunctions = append(noiseFunctions, func(t []float64) float64 {
			return freqFn(t) * weight
		})
	}

	return func(t []float64) (sum float64) {
		for _, fn := range noiseFunctions {
			sum += fn(t)
		}
		return sum
	}
}
