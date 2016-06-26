package transformer_test

import (
	"math"
	"testing"

	"github.com/bcokert/terragen/noise"
	"github.com/bcokert/terragen/transformer"
)

func TestSinusoid1D(t *testing.T) {
	testCases := map[string]struct {
		PhaseFn    noise.Function1D
		Frequency  float64
		ExpectedFn noise.Function1D
	}{
		"constant phase function": {
			PhaseFn: func(t float64) float64 {
				return 1
			},
			Frequency: 3,
			ExpectedFn: func(t float64) float64 {
				return math.Sin(2*math.Pi*3*t + 1)
			},
		},
		"linear phase function": {
			PhaseFn: func(t float64) float64 {
				return t * 2
			},
			Frequency: 17.442,
			ExpectedFn: func(t float64) float64 {
				return math.Sin(2*math.Pi*17.442*t + (t * 2))
			},
		},
	}

	for name, testCase := range testCases {
		noiseFunction := transformer.Sinusoid1D(testCase.PhaseFn, testCase.Frequency)
		if !noiseFunction.IsEqual(testCase.ExpectedFn) {
			t.Errorf("%s failed. Noise function did not equal expected function", name)
		}
	}
}

func TestSinusoid2D(t *testing.T) {
	testCases := map[string]struct {
		PhaseFn    noise.Function2D
		Frequency  float64
		ExpectedFn noise.Function2D
	}{
		"constant phase function": {
			PhaseFn: func(tx, ty float64) float64 {
				return 1
			},
			Frequency: 3,
			ExpectedFn: func(tx, ty float64) float64 {
				return math.Sin(2*math.Pi*3*(tx+ty) + 1)
			},
		},
		"linear phase function": {
			PhaseFn: func(tx, ty float64) float64 {
				return tx * 2
			},
			Frequency: 17.442,
			ExpectedFn: func(tx, ty float64) float64 {
				return math.Sin(2*math.Pi*17.442*(tx+ty) + (tx * 2))
			},
		},
	}

	for name, testCase := range testCases {
		noiseFunction := transformer.Sinusoid2D(testCase.PhaseFn, testCase.Frequency)
		if !noiseFunction.IsEqual(testCase.ExpectedFn) {
			t.Errorf("%s failed. Noise function did not equal expected function", name)
		}
	}
}
