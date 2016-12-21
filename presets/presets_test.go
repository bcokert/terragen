package presets_test

import (
	"math"
	"testing"

	"github.com/bcokert/terragen/generator"
	tgmath "github.com/bcokert/terragen/math"
	"github.com/bcokert/terragen/noise"
	"github.com/bcokert/terragen/presets"
	"github.com/bcokert/terragen/synthesizer"
	"github.com/bcokert/terragen/transformer"
)

func testSpectralPreset(t *testing.T, preset presets.Preset, weightExponent float64) {
	testCases := map[string]struct {
		Frequencies []float64
		Dimensions  []int
	}{
		"one frequency": {
			Frequencies: []float64{1},
			Dimensions:  []int{1, 2},
		},
		"multi frequency": {
			Frequencies: []float64{1, 2, 3},
			Dimensions:  []int{1, 2},
		},
	}

	for name, testCase := range testCases {
		expectedGeneratorFn := generator.Random(tgmath.NewDefaultSource(42))
		expectedNoiseFnGenerator := func(freq float64) noise.Function {
			return transformer.Sinusoid(expectedGeneratorFn, freq)
		}
		expectedWeightFn := func(freq float64) float64 {
			return math.Pow(freq, weightExponent)
		}
		expectedSynthesizerFn := synthesizer.Octave(expectedNoiseFnGenerator, expectedWeightFn, testCase.Frequencies)

		noiseFunction := preset(tgmath.NewDefaultSource(42), testCase.Frequencies)

		for _, dimension := range testCase.Dimensions {
			if !noiseFunction.IsEqual(expectedSynthesizerFn, dimension) {
				t.Errorf("%s failed in dimension %d. Noise function did not equal expected function", name, dimension)
			}
		}
	}
}

func TestViolet(t *testing.T) {
	testSpectralPreset(t, presets.Violet, 2)
}

func TestBlue(t *testing.T) {
	testSpectralPreset(t, presets.Blue, 1)
}

func TestWhite(t *testing.T) {
	testSpectralPreset(t, presets.White, 0)
}

func TestPink(t *testing.T) {
	testSpectralPreset(t, presets.Pink, -1)
}

func TestRed(t *testing.T) {
	testSpectralPreset(t, presets.Red, -2)
}

func TestRawPerlin(t *testing.T) {
	testCases := map[string]struct {
		Frequencies []float64
	}{
		"smoke": {
			Frequencies: []float64{1},
		},
	}

	for name, testCase := range testCases {
		expectedCache := tgmath.NewDefaultRandomGridCache(tgmath.NewDefaultSource(42))
		expectedInterpolator := tgmath.NewInterpolator(tgmath.DampCubicEase)
		expectedGeneratorFn := generator.Perlin(expectedCache, expectedInterpolator)

		noiseFunction := presets.RawPerlin(tgmath.NewDefaultSource(42), testCase.Frequencies)

		if !noiseFunction.IsEqual(expectedGeneratorFn, 2) {
			t.Errorf("%s failed in dimension %d. Noise function did not equal expected function", name, 2)
		}
	}
}
