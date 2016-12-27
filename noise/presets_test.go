package noise_test

import (
	"math"
	"testing"

	tgmath "github.com/bcokert/terragen/math"
	"github.com/bcokert/terragen/noise"
)

func testSpectralPreset(t *testing.T, preset noise.Preset, weightExponent float64) {
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
		expectedGeneratorFn := noise.Random(tgmath.NewDefaultSource(42))
		expectedNoiseFnGenerator := func(freq float64) noise.Function {
			return noise.Sinusoid(expectedGeneratorFn, freq)
		}
		expectedWeightFn := func(freq float64) float64 {
			return math.Pow(freq, weightExponent)
		}
		expectedSynthesizerFn := noise.Octave(expectedNoiseFnGenerator, expectedWeightFn, testCase.Frequencies)

		noiseFunction := preset(tgmath.NewDefaultSource(42), testCase.Frequencies)

		for _, dimension := range testCase.Dimensions {
			if !noiseFunction.IsEqual(expectedSynthesizerFn, dimension) {
				t.Errorf("%s failed in dimension %d. Noise function did not equal expected function", name, dimension)
			}
		}
	}
}

func TestViolet(t *testing.T) {
	testSpectralPreset(t, noise.Violet, 2)
}

func TestBlue(t *testing.T) {
	testSpectralPreset(t, noise.Blue, 1)
}

func TestWhite(t *testing.T) {
	testSpectralPreset(t, noise.White, 0)
}

func TestPink(t *testing.T) {
	testSpectralPreset(t, noise.Pink, -1)
}

func TestRed(t *testing.T) {
	testSpectralPreset(t, noise.Red, -2)
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
		expectedGeneratorFn := noise.Perlin(expectedCache, expectedInterpolator)

		noiseFunction := noise.RawPerlin(tgmath.NewDefaultSource(42), testCase.Frequencies)

		if !noiseFunction.IsEqual(expectedGeneratorFn, 2) {
			t.Errorf("%s failed in dimension %d. Noise function did not equal expected function", name, 2)
		}
	}
}
