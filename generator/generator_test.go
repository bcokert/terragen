package generator_test

import (
	"math/rand"
	"testing"

	"github.com/bcokert/terragen/generator"
	"github.com/bcokert/terragen/noise"
)

func TestRandom(t *testing.T) {
	testCases := map[string]struct {
		Seed              int64
		ExpectedFnCreator func() noise.Function
		ExpectedDimension int
	}{
		"basic 1d": {
			Seed: 42,
			ExpectedFnCreator: func() noise.Function {
				rand := rand.New(rand.NewSource(42)).Float64
				return func(t []float64) float64 {
					return rand()
				}
			},
			ExpectedDimension: 1,
		},
		"basic 2d": {
			Seed: 99,
			ExpectedFnCreator: func() noise.Function {
				rand := rand.New(rand.NewSource(99)).Float64
				return func(t []float64) float64 {
					return rand()
				}
			},
			ExpectedDimension: 2,
		},
	}

	for name, testCase := range testCases {
		noiseFunction := generator.Random(testCase.Seed)
		if !noiseFunction.IsEqual(testCase.ExpectedFnCreator(), testCase.ExpectedDimension) {
			t.Errorf("%s failed. Noise function did not equal expected function", name)
		}
	}
}
