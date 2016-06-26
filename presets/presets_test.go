package presets_test

import (
	"math"
	"testing"

	"github.com/bcokert/terragen/generator"
	"github.com/bcokert/terragen/noise"
	"github.com/bcokert/terragen/presets"
	"github.com/bcokert/terragen/synthesizer"
	"github.com/bcokert/terragen/transformer"
)

func testSpectral1D(t *testing.T, preset presets.Spectral1DPreset, weightExponent float64) {
	testCases := map[string]struct {
		InputParams []float64
		Frequencies []float64
	}{
		"one frequency": {
			InputParams: []float64{-1, 0, 1, 1.5, 2, 50},
			Frequencies: []float64{1},
		},
		"multi frequency": {
			InputParams: []float64{-1, 0, 1},
			Frequencies: []float64{1, 2, 3},
		},
	}

	for name, testCase := range testCases {
		expectedGeneratorFn := generator.Random1D(42)
		expectedNoiseFnGenerator := func(freq float64) noise.Function1D {
			return transformer.Sinusoid1D(expectedGeneratorFn, freq)
		}
		expectedWeightFn := func(freq float64) float64 {
			return math.Pow(freq, weightExponent)
		}
		expectedSynthesizerFn := synthesizer.Octave1D(expectedNoiseFnGenerator, expectedWeightFn, testCase.Frequencies)

		noiseFunction := preset(42, testCase.Frequencies)

		for _, param := range testCase.InputParams {
			expected := expectedSynthesizerFn(param)
			result := noiseFunction(param)
			if math.Abs(result-expected) > 0.0000000000001 {
				t.Errorf("%s failed. Expected param %v to result in %v, received %v", name, param, expected, result)
			}
		}
	}
}

func TestViolet1D(t *testing.T) {
	testSpectral1D(t, presets.Violet1D, 2)
}

func TestBlue1D(t *testing.T) {
	testSpectral1D(t, presets.Blue1D, 1)
}

func TestWhite1D(t *testing.T) {
	testSpectral1D(t, presets.White1D, 0)
}

func TestPink1D(t *testing.T) {
	testSpectral1D(t, presets.Pink1D, -1)
}

func TestRed1D(t *testing.T) {
	testSpectral1D(t, presets.Red1D, -2)
}

func testSpectral2D(t *testing.T, preset presets.Spectral2DPreset, weightExponent float64) {
	testCases := map[string]struct {
		InputParams [][2]float64
		Frequencies []float64
	}{
		"one frequency": {
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
			Frequencies: []float64{1},
		},
		"multi frequency": {
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
			Frequencies: []float64{1, 2, 3},
		},
	}

	for name, testCase := range testCases {
		expectedGeneratorFn := generator.Random2D(42)
		expectedNoiseFnGenerator := func(freq float64) noise.Function2D {
			return transformer.Sinusoid2D(expectedGeneratorFn, freq)
		}
		expectedWeightFn := func(freq float64) float64 {
			return math.Pow(freq, weightExponent)
		}
		expectedSynthesizerFn := synthesizer.Octave2D(expectedNoiseFnGenerator, expectedWeightFn, testCase.Frequencies)

		noiseFunction := preset(42, testCase.Frequencies)

		for _, params := range testCase.InputParams {
			expected := expectedSynthesizerFn(params[0], params[1])
			result := noiseFunction(params[0], params[1])
			if math.Abs(result-expected) > 0.0000000000001 {
				t.Errorf("%s failed. Expected params %v to result in %v, received %v", name, params, expected, result)
			}
		}
	}
}

func TestViolet2D(t *testing.T) {
	testSpectral2D(t, presets.Violet2D, 2)
}

func TestBlue2D(t *testing.T) {
	testSpectral2D(t, presets.Blue2D, 1)
}

func TestWhite2D(t *testing.T) {
	testSpectral2D(t, presets.White2D, 0)
}

func TestPink2D(t *testing.T) {
	testSpectral2D(t, presets.Pink2D, -1)
}

func TestRed2D(t *testing.T) {
	testSpectral2D(t, presets.Red2D, -2)
}
