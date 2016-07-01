package noise_test

import (
	"testing"

	"math"

	"github.com/bcokert/terragen/noise"
)

func TestIsEqual(t *testing.T) {
	testCases := map[string]struct {
		NoiseFn1  noise.Function
		NoiseFn2  noise.Function
		Dimension int
		Expected  bool
	}{
		"constants": {
			NoiseFn1: func(t []float64) float64 {
				return 1
			},
			NoiseFn2: func(t []float64) float64 {
				return 1
			},
			Dimension: 1,
			Expected:  true,
		},
		"linear": {
			NoiseFn1: func(t []float64) float64 {
				return t[0]
			},
			NoiseFn2: func(t []float64) float64 {
				return t[0]
			},
			Dimension: 1,
			Expected:  true,
		},
		"complex": {
			NoiseFn1: func(t []float64) float64 {
				return 2 * t[0] / (48 + t[0]) * math.Sin(t[0]) / 3 * t[0]
			},
			NoiseFn2: func(t []float64) float64 {
				return 2 * t[0] / (48 + t[0]) * math.Sin(t[0]) / 3 * t[0]
			},
			Dimension: 1,
			Expected:  true,
		},
		"different formulations": {
			NoiseFn1: func(t []float64) float64 {
				return 2*t[0]*6 + (4*t[0])/2
			},
			NoiseFn2: func(t []float64) float64 {
				return 14 * t[0]
			},
			Dimension: 1,
			Expected:  true,
		},
		"almost equal": {
			NoiseFn1: func(t []float64) float64 {
				return 2*t[0]*6 + (4*t[0])/2
			},
			NoiseFn2: func(t []float64) float64 {
				return 14*t[0] + 0.00000000001
			},
			Dimension: 1,
			Expected:  false,
		},
		"same magnitude, different sign": {
			NoiseFn1: func(t []float64) float64 {
				return -1
			},
			NoiseFn2: func(t []float64) float64 {
				return 1
			},
			Dimension: 1,
			Expected:  false,
		},
		"same but for one": {
			NoiseFn1: func(t []float64) float64 {
				if t[0] == -1 {
					return 9
				}
				return t[0]
			},
			NoiseFn2: func(t []float64) float64 {
				return t[0]
			},
			Dimension: 1,
			Expected:  false,
		},
		"differently signed zero": {
			NoiseFn1: func(t []float64) float64 {
				return 0
			},
			NoiseFn2: func(t []float64) float64 {
				return -0
			},
			Dimension: 1,
			Expected:  true,
		},
		"constants 2d": {
			NoiseFn1: func(t []float64) float64 {
				return 1
			},
			NoiseFn2: func(t []float64) float64 {
				return 1
			},
			Dimension: 2,
			Expected:  true,
		},
		"linear 2d": {
			NoiseFn1: func(t []float64) float64 {
				return t[0] + t[1]
			},
			NoiseFn2: func(t []float64) float64 {
				return t[0] + t[1]
			},
			Dimension: 2,
			Expected:  true,
		},
		"complex 2d": {
			NoiseFn1: func(t []float64) float64 {
				return 2 * t[0] / (48 + t[1]) * math.Sin(t[1]) / 3 * t[0]
			},
			NoiseFn2: func(t []float64) float64 {
				return 2 * t[0] / (48 + t[1]) * math.Sin(t[1]) / 3 * t[0]
			},
			Dimension: 2,
			Expected:  true,
		},
		"different formulations 2d": {
			NoiseFn1: func(t []float64) float64 {
				return 2*t[0]*6 + (4*t[1])/2
			},
			NoiseFn2: func(t []float64) float64 {
				return 12*t[0] + 2*t[1]
			},
			Dimension: 2,
			Expected:  true,
		},
		"almost equal 2d": {
			NoiseFn1: func(t []float64) float64 {
				return 2*t[0]*6 + (4*t[1])/2
			},
			NoiseFn2: func(t []float64) float64 {
				return 12*t[0] + 2*t[1] + 0.00000000001
			},
			Dimension: 2,
			Expected:  false,
		},
		"same magnitude, different sign 2d": {
			NoiseFn1: func(t []float64) float64 {
				return -1
			},
			NoiseFn2: func(t []float64) float64 {
				return 1
			},
			Dimension: 2,
			Expected:  false,
		},
		"same but for one 2d": {
			NoiseFn1: func(t []float64) float64 {
				if t[0] == -1 && t[1] == -1 {
					return 9.12
				}
				return t[0] + t[1] + 100
			},
			NoiseFn2: func(t []float64) float64 {
				return t[0] + t[1] + 100
			},
			Dimension: 2,
			Expected:  false,
		},
		"differently signed zero 2d": {
			NoiseFn1: func(t []float64) float64 {
				return 0
			},
			NoiseFn2: func(t []float64) float64 {
				return -0
			},
			Dimension: 2,
			Expected:  true,
		},
	}

	for name, testCase := range testCases {
		result := testCase.NoiseFn1.IsEqual(testCase.NoiseFn2, testCase.Dimension)
		if result != testCase.Expected {
			t.Errorf("'%s' failed. Expected %v, received %v", name, testCase.Expected, result)
		}
	}
}
