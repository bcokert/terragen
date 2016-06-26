package generator_test

import (
	"math/rand"
	"testing"

	"github.com/bcokert/terragen/generator"
	"github.com/bcokert/terragen/noise"
)

func TestRandom1D(t *testing.T) {
	rand := rand.New(rand.NewSource(27)).Float64
	testCases := map[string]struct {
		ExpectedFn noise.Function1D
	}{
		"basic": {
			ExpectedFn: func(t float64) float64 {
				return rand()
			},
		},
	}

	for name, testCase := range testCases {
		noiseFunction := generator.Random1D(27)
		if !noiseFunction.IsEqual(testCase.ExpectedFn) {
			t.Errorf("%s failed. Noise function did not equal expected function", name)
		}
	}
}

func TestRandom2D(t *testing.T) {
	rand := rand.New(rand.NewSource(42)).Float64
	testCases := map[string]struct {
		ExpectedFn noise.Function2D
	}{
		"basic": {
			ExpectedFn: func(tx, ty float64) float64 {
				return rand()
			},
		},
	}

	for name, testCase := range testCases {
		noiseFunction := generator.Random2D(42)
		if !noiseFunction.IsEqual(testCase.ExpectedFn) {
			t.Errorf("%s failed. Noise function did not equal expected function", name)
		}
	}
}
