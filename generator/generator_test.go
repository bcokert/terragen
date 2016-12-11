package generator_test

import (
	"testing"

	"github.com/bcokert/terragen/generator"
	"github.com/bcokert/terragen/interpolation"
	"github.com/bcokert/terragen/noise"
	"github.com/bcokert/terragen/random"
	"github.com/bcokert/terragen/vector"
	"math"
)

func TestRandom(t *testing.T) {
	testCases := map[string]struct {
		Source            random.Source
		ExpectedFnCreator func() noise.Function
		ExpectedDimension int
	}{
		"basic 1d": {
			Source: random.NewDefaultSource(42),
			ExpectedFnCreator: func() noise.Function {
				rand := random.NewDefaultSource(42).Float64
				return func(t []float64) float64 {
					return rand()
				}
			},
			ExpectedDimension: 1,
		},
		"basic 2d": {
			Source: random.NewDefaultSource(99),
			ExpectedFnCreator: func() noise.Function {
				rand := random.NewDefaultSource(99).Float64
				return func(t []float64) float64 {
					return rand()
				}
			},
			ExpectedDimension: 2,
		},
	}

	for name, testCase := range testCases {
		noiseFunction := generator.Random(testCase.Source)
		if !noiseFunction.IsEqual(testCase.ExpectedFnCreator(), testCase.ExpectedDimension) {
			t.Errorf("%s failed. Noise function did not equal expected function", name)
		}
	}
}

func TestPerlin(t *testing.T) {
	testCases := map[string]struct {
		Cache        vector.GridCache
		Interpolator interpolation.Interpolator
	}{
		"grid edges": {
			Cache:        &vector.MockGridCache{},
			Interpolator: interpolation.NewInterpolator(interpolation.LinearEase),
		},
	}

	for name, testCase := range testCases {

		testNoiseFunction := generator.Perlin(testCase.Cache, testCase.Interpolator)
		expectedNoiseFunction := func(input []float64) float64 {
			a, b := input[0], input[1]

			topLeft := [2]int{int(a), int(b)}
			topRight := [2]int{int(a) + 1, int(b)}
			botLeft := [2]int{int(a), int(b) + 1}
			botRight := [2]int{int(a) + 1, int(b) + 1}

			tl := vector.NewVec2(a-float64(topLeft[0]), b-float64(topLeft[1]))
			tr := vector.NewVec2(a-float64(topRight[0]), b-float64(topRight[1]))
			bl := vector.NewVec2(a-float64(botLeft[0]), b-float64(botLeft[1]))
			br := vector.NewVec2(a-float64(botRight[0]), b-float64(botRight[1]))

			itl := tl.Dot(testCase.Cache.Get(topLeft[0], topLeft[1]))
			itr := tr.Dot(testCase.Cache.Get(topRight[0], topRight[1]))
			ibl := bl.Dot(testCase.Cache.Get(botLeft[0], botLeft[1]))
			ibr := br.Dot(testCase.Cache.Get(botRight[0], botRight[1]))

			avgAB := testCase.Interpolator(math.Abs(a-float64(topLeft[0])), itl, itr)
			avgCD := testCase.Interpolator(math.Abs(a-float64(topLeft[0])), ibl, ibr)
			return testCase.Interpolator(math.Abs(b-float64(topLeft[1])), avgAB, avgCD)
		}

		if !testNoiseFunction.IsEqual(expectedNoiseFunction, 2) {
			t.Errorf("'%s' failed.", name)
		}
	}
}
