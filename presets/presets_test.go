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

func testSpectra1D(t *testing.T, preset presets.Spectral1DPreset, weightExponent float64) {
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
		expectedGeneratorFn := generator.Random(42)
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
	testSpectra1D(t, presets.Violet1D, 2)
}

func TestBlue1D(t *testing.T) {
	testSpectra1D(t, presets.Blue1D, 1)
}

func TestWhite1D(t *testing.T) {
	testSpectra1D(t, presets.White1D, 0)
}

func TestPink1D(t *testing.T) {
	testSpectra1D(t, presets.Pink1D, -1)
}

func TestRed1D(t *testing.T) {
	testSpectra1D(t, presets.Red1D, -2)
}
