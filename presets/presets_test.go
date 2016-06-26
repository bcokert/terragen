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
		Frequencies []float64
	}{
		"one frequency": {
			Frequencies: []float64{1},
		},
		"multi frequency": {
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

		if !noiseFunction.IsEqual(expectedSynthesizerFn) {
			t.Errorf("%s failed. Noise function did not equal expected function", name)
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
		Frequencies []float64
	}{
		"one frequency": {
			Frequencies: []float64{1},
		},
		"multi frequency": {
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

		if !noiseFunction.IsEqual(expectedSynthesizerFn) {
			t.Errorf("%s failed. Noise function did not equal expected function", name)
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
