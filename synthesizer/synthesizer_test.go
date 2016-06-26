package synthesizer_test

import (
	"testing"

	"github.com/bcokert/terragen/noise"
	"github.com/bcokert/terragen/synthesizer"
)

func TestOctave1D(t *testing.T) {
	testCases := map[string]struct {
		Generator      synthesizer.NoiseFunctionGenerator1D
		WeightFunction synthesizer.WeightFunction
		Frequencies    []float64
		ExpectedFn     noise.Function1D
	}{
		"constant function and weight, 2 frequencies": {
			Generator: func(freq float64) noise.Function1D {
				return func(t float64) float64 {
					return 4
				}
			},
			WeightFunction: func(freq float64) float64 {
				return 6
			},
			Frequencies: []float64{5, 10},
			ExpectedFn: func(t float64) float64 {
				return 4*6 + 4*6
			},
		},
		"basic function and weight, 3 frequencies": {
			Generator: func(freq float64) noise.Function1D {
				return func(t float64) float64 {
					return freq + t
				}
			},
			WeightFunction: func(freq float64) float64 {
				return freq
			},
			Frequencies: []float64{5, 10, 100},
			ExpectedFn: func(t float64) float64 {
				return (5+t)*5 + (10+t)*10 + (100+t)*100
			},
		},
	}

	for name, testCase := range testCases {
		noiseFunction := synthesizer.Octave1D(testCase.Generator, testCase.WeightFunction, testCase.Frequencies)
		if !noiseFunction.IsEqual(testCase.ExpectedFn) {
			t.Errorf("%s failed. Noise function did not equal expected function", name)
		}
	}
}

func TestOctave2D(t *testing.T) {
	testCases := map[string]struct {
		Generator      synthesizer.NoiseFunctionGenerator2D
		WeightFunction synthesizer.WeightFunction
		Frequencies    []float64
		ExpectedFn     noise.Function2D
	}{
		"constant function and weight, 2 frequencies": {
			Generator: func(freq float64) noise.Function2D {
				return func(tx, ty float64) float64 {
					return 4
				}
			},
			WeightFunction: func(freq float64) float64 {
				return 6
			},
			Frequencies: []float64{5, 10},
			ExpectedFn: func(tx, ty float64) float64 {
				return 4*6 + 4*6
			},
		},
		"basic function and weight, 3 frequencies": {
			Generator: func(freq float64) noise.Function2D {
				return func(tx, ty float64) float64 {
					return freq + 2*tx + 3*ty
				}
			},
			WeightFunction: func(freq float64) float64 {
				return freq
			},
			Frequencies: []float64{5, 10, 100},
			ExpectedFn: func(tx, ty float64) float64 {
				return (5+2*tx+3*ty)*5 + (10+2*tx+3*ty)*10 + (100+2*tx+3*ty)*100
			},
		},
	}

	for name, testCase := range testCases {
		noiseFunction := synthesizer.Octave2D(testCase.Generator, testCase.WeightFunction, testCase.Frequencies)
		if !noiseFunction.IsEqual(testCase.ExpectedFn) {
			t.Errorf("%s failed. Noise function did not equal expected function", name)
		}
	}
}
