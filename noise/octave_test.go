package noise_test

import (
	"math"
	"testing"

	"github.com/bcokert/terragen/noise"
	"github.com/bcokert/terragen/random"
)

func TestWhiteNoise1D(t *testing.T) {
	testCases := map[string]struct {
		InputParams     []float64
		Frequencies     []float64
		ExpectedResults []float64
	}{
		"one frequency": {
			InputParams: []float64{-1, 0, 1, 1.5, 2, 50},
			Frequencies: []float64{1},
			ExpectedResults: []float64{
				math.Sin(2*math.Pi*1*-1 + 1),
				math.Sin(2*math.Pi*1*0 + 1),
				math.Sin(2*math.Pi*1*1 + 1),
				math.Sin(2*math.Pi*1*1.5 + 1),
				math.Sin(2*math.Pi*1*2 + 1),
				math.Sin(2*math.Pi*1*50 + 1),
			},
		},
		"multi frequency": {
			InputParams: []float64{-1, 0, 1},
			Frequencies: []float64{1, 2, 3},
			ExpectedResults: []float64{
				math.Sin(2*math.Pi*1*-1+1)/3 + math.Sin(2*math.Pi*1*-1+1)/3 + math.Sin(2*math.Pi*3*-1+1)/3,
				math.Sin(2*math.Pi*1*0+1)/3 + math.Sin(2*math.Pi*2*0+1)/3 + math.Sin(2*math.Pi*3*0+1)/3,
				math.Sin(2*math.Pi*1*1+1)/3 + math.Sin(2*math.Pi*2*1+1)/3 + math.Sin(2*math.Pi*3*1+1)/3,
			},
		},
	}

	for name, testCase := range testCases {
		noiseFunction := noise.WhiteNoise1D(func() float64 { return 1 }, testCase.Frequencies)
		for i, expected := range testCase.ExpectedResults {
			if result := noiseFunction(testCase.InputParams[i]); math.Abs(result-expected) > 0.0000000000001 {
				t.Errorf("%s failed. Expected result %d to be %v, received %v", name, i, expected, result)
			}
		}
	}
}

func TestRedNoise1D(t *testing.T) {
	testCases := map[string]struct {
		InputParams     []float64
		Frequencies     []float64
		ExpectedResults []float64
	}{
		"one frequency": {
			InputParams: []float64{-1, 0, 1, 1.5, 2, 50},
			Frequencies: []float64{1},
			ExpectedResults: []float64{
				math.Sin(2*math.Pi*1*-1 + 1),
				math.Sin(2*math.Pi*1*0 + 1),
				math.Sin(2*math.Pi*1*1 + 1),
				math.Sin(2*math.Pi*1*1.5 + 1),
				math.Sin(2*math.Pi*1*2 + 1),
				math.Sin(2*math.Pi*1*50 + 1),
			},
		},
		"multi frequency": {
			InputParams: []float64{-1, 0, 1},
			Frequencies: []float64{1, 2, 3},
			ExpectedResults: []float64{
				math.Sin(2*math.Pi*1*-1+1)/3 + math.Sin(2*math.Pi*1*-1+1)/12 + math.Sin(2*math.Pi*3*-1+1)/27,
				math.Sin(2*math.Pi*1*0+1)/3 + math.Sin(2*math.Pi*2*0+1)/12 + math.Sin(2*math.Pi*3*0+1)/27,
				math.Sin(2*math.Pi*1*1+1)/3 + math.Sin(2*math.Pi*2*1+1)/12 + math.Sin(2*math.Pi*3*1+1)/27,
			},
		},
	}

	for name, testCase := range testCases {
		noiseFunction := noise.RedNoise1D(func() float64 { return 1 }, testCase.Frequencies)
		for i, expected := range testCase.ExpectedResults {
			if result := noiseFunction(testCase.InputParams[i]); math.Abs(result-expected) > 0.0000000000001 {
				t.Errorf("%s failed. Expected result %d to be %v, received %v", name, i, expected, result)
			}
		}
	}
}

func TestSineOctave1D(t *testing.T) {
	testCases := map[string]struct {
		InputParams     []float64
		PhaseFunction   random.NormalFunction
		WeightFunction  noise.OctaveWeight1D
		Frequencies     []float64
		ExpectedResults []float64
	}{
		"constant phase and weight, 2 frequencies": {
			InputParams: []float64{-1, 0, 500},
			PhaseFunction: func() float64 {
				return 3
			},
			WeightFunction: func(freq float64) float64 {
				return 6
			},
			Frequencies: []float64{5, 10},
			ExpectedResults: []float64{
				math.Sin(2*math.Pi*5*-1+3)*(6/2) + math.Sin(2*math.Pi*10*-1+3)*(6/2),
				math.Sin(2*math.Pi*5*0+3)*(6/2) + math.Sin(2*math.Pi*10*0+3)*(6/2),
				math.Sin(2*math.Pi*5*500+3)*(6/2) + math.Sin(2*math.Pi*10*500+3)*(6/2),
			},
		},
	}

	for name, testCase := range testCases {
		noiseFunction := noise.SineOctave1D(testCase.PhaseFunction, testCase.WeightFunction, testCase.Frequencies)
		for i, expected := range testCase.ExpectedResults {
			if result := noiseFunction(testCase.InputParams[i]); math.Abs(result-expected) > 0.0000000000001 {
				t.Errorf("%s failed. Expected result %d to be %v, received %v", name, i, expected, result)
			}
		}
	}
}

func TestOctave1D(t *testing.T) {
	testCases := map[string]struct {
		InputParams     []float64
		Generator       noise.OctaveParametric1D
		WeightFunction  noise.OctaveWeight1D
		Frequencies     []float64
		ExpectedResults []float64
	}{
		"constant function and weight, 2 frequencies": {
			InputParams: []float64{-1, 0, 500},
			Generator: func(freq float64) noise.Parametric1D {
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
			Generator: func(freq float64) noise.Parametric1D {
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
		noiseFunction := noise.Octave1D(testCase.Generator, testCase.WeightFunction, testCase.Frequencies)
		for i, expected := range testCase.ExpectedResults {
			if result := noiseFunction(testCase.InputParams[i]); math.Abs(result-expected) > 0.00000000001 {
				t.Errorf("%s failed. Expected result %d to be %v, received %v", name, i, expected, result)
			}
		}
	}
}
