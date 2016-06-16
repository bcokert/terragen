package model_test

import (
	"testing"

	"github.com/bcokert/terragen/model"
	"github.com/bcokert/terragen/noisefunction"
	"github.com/bcokert/terragen/testutils"
)

func TestNewNoise(t *testing.T) {
	testCases := map[string]struct {
		NoiseFunction string
		Expected      model.Noise
	}{
		"basic": {
			NoiseFunction: "white:1d",
			Expected: model.Noise{
				NoiseFunction: "white:1d",
			},
		},
	}

	for name, testCase := range testCases {
		result := model.NewNoise(testCase.NoiseFunction)
		if !result.Equals(&testCase.Expected) {
			t.Errorf("%s failed. Expected %#v, received %#v", name, testCase.Expected, result)
		}
	}
}

func TestGenerate(t *testing.T) {
	testCases := map[string]struct {
		From          []float64
		To            []float64
		Resolution    int
		NoiseFunction noisefunction.Function1D
		Expected      model.Noise
	}{
		"basic 1d": {
			From:       []float64{1},
			To:         []float64{3},
			Resolution: 4,
			NoiseFunction: func(t float64) float64 {
				return t
			},
			Expected: model.Noise{
				RawNoise: map[string][]float64{
					"x":     testutils.FloatSlice(1, 3, 4),
					"value": testutils.FloatSlice(1, 3, 4),
				},
				From:       []float64{1},
				To:         []float64{3},
				Resolution: 4,
			},
		},
	}

	for name, testCase := range testCases {
		noise := &model.Noise{}
		noise.Generate(testCase.From, testCase.To, testCase.Resolution, testCase.NoiseFunction)

		if !noise.Equals(&testCase.Expected) {
			t.Errorf("%s failed. Expected %#v, receinved %#v", name, testCase.Expected, noise)
		}
	}
}

func TestEquals(t *testing.T) {
	testCases := map[string]struct {
		Left     model.Noise
		Right    model.Noise
		Expected bool
	}{
		"equal empty": {
			Left:     model.Noise{},
			Right:    model.Noise{},
			Expected: true,
		},
		"equal full": {
			Left: model.Noise{
				RawNoise: map[string][]float64{
					"x":     []float64{0, 0.5, 1, 1.5, 2, 2.5, 3, 3.5, 4, 4.5, 5, 5.5, 6, 6.5, 7, 7.5, 8, 8.5, 9, 9.5},
					"value": []float64{0.10491637639740707, 0.02522960371433485, 0.14240998620874543, -0.06951808844985238, 0.22028277620639095, -0.0036129201768815805, 0.16743035692290176, -0.08762397550026342, 0.25150755724035145, -0.1495588243161205, 0.21620418268684216, -0.11448528972805566, 0.18006882300120994, -0.1482459144529171, 0.12864632754084612, -0.15321647422470572, 0.040347426422610716, -0.14143833187807403, 0.16947245271105693, -0.09200993524802667},
				},
				From:          []float64{0},
				To:            []float64{10},
				Resolution:    2,
				NoiseFunction: "red:1d",
			},
			Right: model.Noise{
				RawNoise: map[string][]float64{
					"x":     []float64{0, 0.5, 1, 1.5, 2, 2.5, 3, 3.5, 4, 4.5, 5, 5.5, 6, 6.5, 7, 7.5, 8, 8.5, 9, 9.5},
					"value": []float64{0.10491637639740707, 0.02522960371433485, 0.14240998620874543, -0.06951808844985238, 0.22028277620639095, -0.0036129201768815805, 0.16743035692290176, -0.08762397550026342, 0.25150755724035145, -0.1495588243161205, 0.21620418268684216, -0.11448528972805566, 0.18006882300120994, -0.1482459144529171, 0.12864632754084612, -0.15321647422470572, 0.040347426422610716, -0.14143833187807403, 0.16947245271105693, -0.09200993524802667},
				},
				From:          []float64{0},
				To:            []float64{10},
				Resolution:    2,
				NoiseFunction: "red:1d",
			},
			Expected: true,
		},
		"almost equal full": {
			Left: model.Noise{
				RawNoise: map[string][]float64{
					"x":     []float64{0, 0.5, 1, 1.5, 2, 2.5, 3, 3.5, 4, 4.5, 5, 5.5, 6, 6.5, 7, 7.5, 8, 8.5, 9, 9.5},
					"value": []float64{0.10491637639740707, 0.02522960371433485, 0.14240998620874543, -0.06951808844985238, 0.22028277620639095, -0.0036129201768815805, 0.16743035692590176, -0.08762397550026342, 0.25150755724035145, -0.1495588243161205, 0.21620418268684216, -0.11448528972805566, 0.18006882300120994, -0.1482459144529171, 0.12864632754084612, -0.15321647422470572, 0.040347426422610716, -0.14143833187807403, 0.16947245271105693, -0.09200993524802667},
				},
				From:          []float64{0},
				To:            []float64{10},
				Resolution:    2,
				NoiseFunction: "red:1d",
			},
			Right: model.Noise{
				RawNoise: map[string][]float64{
					"x":     []float64{0, 0.5, 1, 1.5, 2, 2.5, 3, 3.5, 4, 4.5, 5, 5.5, 6, 6.5, 7, 7.5, 8, 8.5, 9, 9.5},
					"value": []float64{0.10491637639740707, 0.02522960371433485, 0.14240998620874543, -0.06951808844985238, 0.22028277620639095, -0.0036129201768815805, 0.16743035692290176, -0.08762397550026342, 0.25150755724035145, -0.1495588243161205, 0.21620418268684216, -0.11448528972805566, 0.18006882300120994, -0.1482459144529171, 0.12864632754084612, -0.15321647422470572, 0.040347426422610716, -0.14143833187807403, 0.16947245271105693, -0.09200993524802667},
				},
				From:          []float64{0},
				To:            []float64{10},
				Resolution:    2,
				NoiseFunction: "red:1d",
			},
			Expected: false,
		},
		"different from": {
			Left: model.Noise{
				From: []float64{1, 2, 3},
			},
			Right: model.Noise{
				From: []float64{1, 3, 3},
			},
			Expected: false,
		},
		"different to": {
			Left: model.Noise{
				To: []float64{1, 2, 3},
			},
			Right: model.Noise{
				To: []float64{1, 3, 3},
			},
			Expected: false,
		},
		"different dimensions": {
			Left: model.Noise{
				RawNoise: map[string][]float64{
					"x":     []float64{1, 2},
					"value": []float64{1, 2, 3},
				},
			},
			Right: model.Noise{
				RawNoise: map[string][]float64{
					"y":     []float64{1, 2},
					"value": []float64{1, 2, 3},
				},
			},
			Expected: false,
		},
		"different number of dimensions": {
			Left: model.Noise{
				RawNoise: map[string][]float64{
					"x":     []float64{1, 2},
					"y":     []float64{1, 2},
					"value": []float64{1, 2, 3},
				},
			},
			Right: model.Noise{
				RawNoise: map[string][]float64{
					"x":     []float64{1, 2},
					"value": []float64{1, 2, 3},
				},
			},
			Expected: false,
		},
		"different length of dimensions": {
			Left: model.Noise{
				RawNoise: map[string][]float64{
					"x":     []float64{1, 2, 3},
					"value": []float64{1, 2, 3},
				},
			},
			Right: model.Noise{
				RawNoise: map[string][]float64{
					"x":     []float64{1, 2},
					"value": []float64{1, 2, 3},
				},
			},
			Expected: false,
		},
	}

	for name, testCase := range testCases {
		if testCase.Left.Equals(&testCase.Right) != testCase.Expected {
			t.Errorf("%s failed. Expected %#v == %#v to be %v", name, testCase.Left, testCase.Right, testCase.Expected)
		}
	}
}
