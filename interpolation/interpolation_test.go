package interpolation_test

import (
	"testing"

	"github.com/bcokert/terragen/interpolation"
)

func TestNewInterpolator(t *testing.T) {
	testCases := map[string]struct {
		Percentage, A, B float64
		EasingFunc       interpolation.EasingFunction
		Expected         float64
	}{
		"zero": {
			Percentage: 0,
			A:          4,
			B:          5,
			EasingFunc: interpolation.LinearEase,
			Expected:   4,
		},
		"one": {
			Percentage: 1,
			A:          4,
			B:          5,
			EasingFunc: interpolation.LinearEase,
			Expected:   5,
		},
		"greater than 1": {
			Percentage: 1.7,
			A:          4,
			B:          5,
			EasingFunc: interpolation.LinearEase,
			Expected:   5,
		},
		"less than 0": {
			Percentage: -1.7,
			A:          4,
			B:          5,
			EasingFunc: interpolation.LinearEase,
			Expected:   4,
		},
		"in between": {
			Percentage: 0.5,
			A:          4,
			B:          5,
			EasingFunc: interpolation.LinearEase,
			Expected:   4.5,
		},
		"close to 1": {
			Percentage: 0.95,
			A:          4,
			B:          5,
			EasingFunc: interpolation.LinearEase,
			Expected:   4.95,
		},
		"weird easing func": {
			Percentage: 0.5,
			A:          4,
			B:          5,
			EasingFunc: func(t float64) float64 {
				return 0
			},
			Expected: 4,
		},
	}

	for name, testCase := range testCases {
		testFunc := interpolation.NewInterpolator(testCase.EasingFunc)
		result := testFunc(testCase.Percentage, testCase.A, testCase.B)
		if result != testCase.Expected {
			t.Errorf("'%s' failed. Expected %v, received %v", name, testCase.Expected, result)
		}
	}
}
