package synthesizer_test

import (
	"math"
	"testing"

	"github.com/bcokert/terragen/noise"
	"github.com/bcokert/terragen/synthesizer"
)

func TestOctave1D(t *testing.T) {
	testCases := map[string]struct {
		InputParams     []float64
		Generator       func(freq float64) noise.Function1D
		WeightFunction  func(freq float64) float64
		Frequencies     []float64
		ExpectedResults []float64
	}{
		"constant function and weight, 2 frequencies": {
			InputParams: []float64{-1, 0, 500},
			Generator: func(freq float64) noise.Function1D {
				return func(t float64) float64 {
					return 4
				}
			},
			WeightFunction: func(freq float64) float64 {
				return 6
			},
			Frequencies: []float64{5, 10},
			ExpectedResults: []float64{
				4*(6/2) + 4*(6/2),
				4*(6/2) + 4*(6/2),
				4*(6/2) + 4*(6/2),
			},
		},
		"basic function and weight, 3 frequencies": {
			InputParams: []float64{-1, 0, 500},
			Generator: func(freq float64) noise.Function1D {
				return func(t float64) float64 {
					return freq + t
				}
			},
			WeightFunction: func(freq float64) float64 {
				return freq
			},
			Frequencies: []float64{5, 10, 100},
			ExpectedResults: []float64{
				(-1+5)*(5/3.0) + (-1+10)*(10/3.0) + (-1+100)*(100/3.0),
				(0+5)*(5/3.0) + (0+10)*(10/3.0) + (0+100)*(100/3.0),
				(500+5)*(5/3.0) + (500+10)*(10/3.0) + (500+100)*(100/3.0),
			},
		},
	}

	for name, testCase := range testCases {
		noiseFunction := synthesizer.Octave1D(testCase.Generator, testCase.WeightFunction, testCase.Frequencies)
		for i, expected := range testCase.ExpectedResults {
			if result := noiseFunction(testCase.InputParams[i]); math.Abs(result-expected) > 0.00000000001 {
				t.Errorf("%s failed. Expected result %d to be %v, received %v", name, i, expected, result)
			}
		}
	}
}
