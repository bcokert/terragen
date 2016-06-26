package noise_test

import (
	"testing"

	"math"

	"github.com/bcokert/terragen/noise"
)

func TestIsEqual1D(t *testing.T) {
	testCases := map[string]struct {
		NoiseFn1 noise.Function1D
		NoiseFn2 noise.Function1D
		Expected bool
	}{
		"constants": {
			NoiseFn1: func(t float64) float64 {
				return 1
			},
			NoiseFn2: func(t float64) float64 {
				return 1
			},
			Expected: true,
		},
		"linear": {
			NoiseFn1: func(t float64) float64 {
				return t
			},
			NoiseFn2: func(t float64) float64 {
				return t
			},
			Expected: true,
		},
		"complex": {
			NoiseFn1: func(t float64) float64 {
				return 2 * t / (48 + t) * math.Sin(t) / 3 * t
			},
			NoiseFn2: func(t float64) float64 {
				return 2 * t / (48 + t) * math.Sin(t) / 3 * t
			},
			Expected: true,
		},
		"different formulations": {
			NoiseFn1: func(t float64) float64 {
				return 2*t*6 + (4*t)/2
			},
			NoiseFn2: func(t float64) float64 {
				return 14 * t
			},
			Expected: true,
		},
		"almost equal": {
			NoiseFn1: func(t float64) float64 {
				return 2*t*6 + (4*t)/2
			},
			NoiseFn2: func(t float64) float64 {
				return 14*t + 0.00000000001
			},
			Expected: false,
		},
		"same magnitude, different sign": {
			NoiseFn1: func(t float64) float64 {
				return -1
			},
			NoiseFn2: func(t float64) float64 {
				return 1
			},
			Expected: false,
		},
		"same but for one": {
			NoiseFn1: func(t float64) float64 {
				if t == -1 {
					return 9
				}
				return t
			},
			NoiseFn2: func(t float64) float64 {
				return t
			},
			Expected: false,
		},
		"differently signed zero": {
			NoiseFn1: func(t float64) float64 {
				return 0
			},
			NoiseFn2: func(t float64) float64 {
				return -0
			},
			Expected: true,
		},
	}

	for name, testCase := range testCases {
		result := testCase.NoiseFn1.IsEqual(testCase.NoiseFn2)
		if result != testCase.Expected {
			t.Errorf("'%s' failed. Expected %v, received %v", name, testCase.Expected, result)
		}
	}
}

func TestIsEqual2D(t *testing.T) {
	testCases := map[string]struct {
		NoiseFn1 noise.Function2D
		NoiseFn2 noise.Function2D
		Expected bool
	}{
		"constants": {
			NoiseFn1: func(tx, ty float64) float64 {
				return 1
			},
			NoiseFn2: func(tx, ty float64) float64 {
				return 1
			},
			Expected: true,
		},
		"linear": {
			NoiseFn1: func(tx, ty float64) float64 {
				return tx + ty
			},
			NoiseFn2: func(tx, ty float64) float64 {
				return tx + ty
			},
			Expected: true,
		},
		"complex": {
			NoiseFn1: func(tx, ty float64) float64 {
				return 2 * tx / (48 + ty) * math.Sin(ty) / 3 * tx
			},
			NoiseFn2: func(tx, ty float64) float64 {
				return 2 * tx / (48 + ty) * math.Sin(ty) / 3 * tx
			},
			Expected: true,
		},
		"different formulations": {
			NoiseFn1: func(tx, ty float64) float64 {
				return 2*tx*6 + (4*ty)/2
			},
			NoiseFn2: func(tx, ty float64) float64 {
				return 12*tx + 2*ty
			},
			Expected: true,
		},
		"almost equal": {
			NoiseFn1: func(tx, ty float64) float64 {
				return 2*tx*6 + (4*ty)/2
			},
			NoiseFn2: func(tx, ty float64) float64 {
				return 12*tx + 2*ty + 0.00000000001
			},
			Expected: false,
		},
		"same magnitude, different sign": {
			NoiseFn1: func(tx, ty float64) float64 {
				return -1
			},
			NoiseFn2: func(tx, ty float64) float64 {
				return 1
			},
			Expected: false,
		},
		"same but for one": {
			NoiseFn1: func(tx, ty float64) float64 {
				if tx == -1 && ty == -1 {
					return 9.12
				}
				return tx + ty + 100
			},
			NoiseFn2: func(tx, ty float64) float64 {
				return tx + ty + 100
			},
			Expected: false,
		},
		"differently signed zero": {
			NoiseFn1: func(tx, ty float64) float64 {
				return 0
			},
			NoiseFn2: func(tx, ty float64) float64 {
				return -0
			},
			Expected: true,
		},
	}

	for name, testCase := range testCases {
		result := testCase.NoiseFn1.IsEqual(testCase.NoiseFn2)
		if result != testCase.Expected {
			t.Errorf("'%s' failed. Expected %v, received %v", name, testCase.Expected, result)
		}
	}
}
