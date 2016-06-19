package transformer_test

import (
	"math"
	"testing"

	"github.com/bcokert/terragen/noise"
	"github.com/bcokert/terragen/transformer"
)

func TestSinusoid1D(t *testing.T) {
	testCases := map[string]struct {
		PhaseFn     noise.Function1D
		Frequency   float64
		InputParams []float64
		Expected    []float64
	}{
		"constant phase function": {
			PhaseFn: func(t float64) float64 {
				return 1
			},
			Frequency:   3,
			InputParams: []float64{-2, 0, 0.1, 1, 5000},
			Expected: []float64{
				math.Sin(2*math.Pi*3*-2 + 1),
				math.Sin(2*math.Pi*3*0 + 1),
				math.Sin(2*math.Pi*3*0.1 + 1),
				math.Sin(2*math.Pi*3*1 + 1),
				math.Sin(2*math.Pi*3*5000 + 1),
			},
		},
		"linear phase function": {
			PhaseFn: func(t float64) float64 {
				return t * 2
			},
			Frequency:   17.442,
			InputParams: []float64{-2, 0, 0.1, 1, 5000},
			Expected: []float64{
				math.Sin(2*math.Pi*17.442*-2 - 4),
				math.Sin(2*math.Pi*17.442*0 + 0),
				math.Sin(2*math.Pi*17.442*0.1 + 0.2),
				math.Sin(2*math.Pi*17.442*1 + 2),
				math.Sin(2*math.Pi*17.442*5000 + 10000),
			},
		},
	}

	for name, testCase := range testCases {
		noiseFn := transformer.Sinusoid1D(testCase.PhaseFn, testCase.Frequency)
		for i, param := range testCase.InputParams {
			if result := noiseFn(param); math.Abs(result-testCase.Expected[i]) > 0.00000000000001 {
				t.Errorf("'%s' failed. Expected result %d to be %v, received %v", name, i, testCase.Expected, result)
			}
		}
	}
}
