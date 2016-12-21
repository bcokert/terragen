package generator_test

import (
	"testing"

	"math"

	"github.com/bcokert/terragen/generator"
	tgmath "github.com/bcokert/terragen/math"
	"github.com/bcokert/terragen/noise"
)

func TestRandom(t *testing.T) {
	testCases := map[string]struct {
		Source            tgmath.Source
		ExpectedFnCreator func() noise.Function
		ExpectedDimension int
	}{
		"basic 1d": {
			Source: tgmath.NewDefaultSource(42),
			ExpectedFnCreator: func() noise.Function {
				rand := tgmath.NewDefaultSource(42).Float64
				return func(t []float64) float64 {
					return rand()
				}
			},
			ExpectedDimension: 1,
		},
		"basic 2d": {
			Source: tgmath.NewDefaultSource(99),
			ExpectedFnCreator: func() noise.Function {
				rand := tgmath.NewDefaultSource(99).Float64
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
		Cache        tgmath.GridCache
		Interpolator tgmath.Interpolator
	}{
		"grid edges": {
			Cache:        &tgmath.MockGridCache{},
			Interpolator: tgmath.NewInterpolator(tgmath.LinearEase),
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

			tl := tgmath.Vec2{a - float64(topLeft[0]), b - float64(topLeft[1])}
			tr := tgmath.Vec2{a - float64(topRight[0]), b - float64(topRight[1])}
			bl := tgmath.Vec2{a - float64(botLeft[0]), b - float64(botLeft[1])}
			br := tgmath.Vec2{a - float64(botRight[0]), b - float64(botRight[1])}

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
