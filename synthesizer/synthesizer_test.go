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
		Generator       synthesizer.NoiseFunctionGenerator1D
		WeightFunction  synthesizer.WeightFunction
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
				4*6 + 4*6,
				4*6 + 4*6,
				4*6 + 4*6,
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
				(-1+5)*5 + (-1+10)*10 + (-1+100)*100,
				(0+5)*5 + (0+10)*10 + (0+100)*100,
				(500+5)*5 + (500+10)*10 + (500+100)*100,
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

func TestOctave2D(t *testing.T) {
	testCases := map[string]struct {
		InputParams    [][2]float64
		Generator      synthesizer.NoiseFunctionGenerator2D
		WeightFunction synthesizer.WeightFunction
		Frequencies    []float64
		ExpectedFn     noise.Function2D
	}{
		"constant function and weight, 2 frequencies": {
			InputParams: [][2]float64{
				[2]float64{-1, -1},
				[2]float64{-1, 0},
				[2]float64{0, -1},
				[2]float64{0, 0},
				[2]float64{1, 1},
				[2]float64{0.999, 1.001},
				[2]float64{0, 500},
				[2]float64{9999, 2244},
				[2]float64{-33234, 0.0001},
			},
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
			InputParams: [][2]float64{
				[2]float64{-1, -1},
				[2]float64{-1, 0},
				[2]float64{0, -1},
				[2]float64{0, 0},
				[2]float64{1, 1},
				[2]float64{0.999, 1.001},
				[2]float64{0, 500},
				[2]float64{9999, 2244},
				[2]float64{-33234, 0.0001},
			},
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
		for i, params := range testCase.InputParams {
			result := noiseFunction(params[0], params[1])
			expected := testCase.ExpectedFn(params[0], params[1])
			if math.Abs(result-expected) > 0.00000000001 {
				t.Errorf("%s failed. Expected result %d to be %v, received %v", name, i, expected, result)
			}
		}
	}
}
