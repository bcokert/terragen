package synthesizer_test

import (
	"testing"

	"github.com/bcokert/terragen/noise"
	"github.com/bcokert/terragen/synthesizer"
)

func TestOctave(t *testing.T) {
	testCases := map[string]struct {
		Generator      synthesizer.NoiseFunctionGenerator
		WeightFunction synthesizer.WeightFunction
		Frequencies    []float64
		Dimension      int
		ExpectedFn     noise.Function
	}{
		"constant function and weight, 2 frequencies 1d": {
			Generator: func(freq float64) noise.Function {
				return func(t []float64) float64 {
					return 4
				}
			},
			WeightFunction: func(freq float64) float64 {
				return 6
			},
			Frequencies: []float64{5, 10},
			Dimension:   1,
			ExpectedFn: func(t []float64) float64 {
				return 4*6 + 4*6
			},
		},
		"basic function and weight, 3 frequencies 1d": {
			Generator: func(freq float64) noise.Function {
				return func(t []float64) float64 {
					return freq + t[0]
				}
			},
			WeightFunction: func(freq float64) float64 {
				return freq
			},
			Frequencies: []float64{5, 10, 100},
			Dimension:   1,
			ExpectedFn: func(t []float64) float64 {
				return (5+t[0])*5 + (10+t[0])*10 + (100+t[0])*100
			},
		},
		"constant function and weight, 2 frequencies 2d": {
			Generator: func(freq float64) noise.Function {
				return func(t []float64) float64 {
					return 4
				}
			},
			WeightFunction: func(freq float64) float64 {
				return 6
			},
			Frequencies: []float64{5, 10},
			Dimension:   2,
			ExpectedFn: func(t []float64) float64 {
				return 4*6 + 4*6
			},
		},
		"basic function and weight, 3 frequencies 2d": {
			Generator: func(freq float64) noise.Function {
				return func(t []float64) float64 {
					return freq + 2*t[0] + 3*t[1]
				}
			},
			WeightFunction: func(freq float64) float64 {
				return freq
			},
			Frequencies: []float64{5, 10, 100},
			Dimension:   2,
			ExpectedFn: func(t []float64) float64 {
				return (5+2*t[0]+3*t[1])*5 + (10+2*t[0]+3*t[1])*10 + (100+2*t[0]+3*t[1])*100
			},
		},
	}

	for name, testCase := range testCases {
		noiseFunction := synthesizer.Octave(testCase.Generator, testCase.WeightFunction, testCase.Frequencies)
		if !noiseFunction.IsEqual(testCase.ExpectedFn, testCase.Dimension) {
			t.Errorf("%s failed. Noise function did not equal expected function", name)
		}
	}
}
