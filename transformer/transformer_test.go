package transformer_test

import (
	"math"
	"testing"

	"github.com/bcokert/terragen/noise"
	"github.com/bcokert/terragen/transformer"
)

func TestSinusoid(t *testing.T) {
	testCases := map[string]struct {
		PhaseFn    noise.Function
		Frequency  float64
		Dimension  int
		ExpectedFn noise.Function
	}{
		"constant phase function": {
			PhaseFn: func(t []float64) float64 {
				return 1
			},
			Frequency: 3,
			Dimension: 1,
			ExpectedFn: func(t []float64) float64 {
				return math.Sin(2*math.Pi*3*t[0] + 1)
			},
		},
		"constant phase function 2d": {
			PhaseFn: func(t []float64) float64 {
				return 7
			},
			Frequency: 3,
			Dimension: 2,
			ExpectedFn: func(t []float64) float64 {
				return math.Sin(2*math.Pi*3*t[0]+7) * math.Sin(2*math.Pi*3*t[1]+7)
			},
		},
	}

	for name, testCase := range testCases {
		noiseFunction := transformer.Sinusoid(testCase.PhaseFn, testCase.Frequency)
		if !noiseFunction.IsEqual(testCase.ExpectedFn, testCase.Dimension) {
			t.Errorf("%s failed. Noise function did not equal expected function", name)
		}
	}
}
