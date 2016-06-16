package presets_test

import (
	"math"
	"testing"

	"github.com/bcokert/terragen/presets"
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
		noiseFunction := presets.WhiteNoise1D(func(t float64) float64 { return 1 }, testCase.Frequencies)
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
		noiseFunction := presets.RedNoise1D(func(t float64) float64 { return 1 }, testCase.Frequencies)
		for i, expected := range testCase.ExpectedResults {
			if result := noiseFunction(testCase.InputParams[i]); math.Abs(result-expected) > 0.0000000000001 {
				t.Errorf("%s failed. Expected result %d to be %v, received %v", name, i, expected, result)
			}
		}
	}
}
