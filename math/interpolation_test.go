package math_test

import (
	"testing"

	"github.com/bcokert/terragen/math"
)

func TestDampCubicEase(t *testing.T) {
	testCases := map[string]struct {
		Input    float64
		Expected float64
	}{
		"zero": {
			Input:    0,
			Expected: 0,
		},
		"one": {
			Input:    1,
			Expected: 1,
		},
		"0.5": {
			Input:    0.5,
			Expected: 3*0.25 - 2*0.125,
		},
		"0.1": {
			Input:    0.1,
			Expected: 3*0.01 - 2*0.001,
		},
	}

	for name, testCase := range testCases {
		if result := math.DampCubicEase(testCase.Input); !math.IsFloatEqual(testCase.Expected, result) {
			t.Errorf("%s failed. Expected %v, received %v.", name, testCase.Expected, result)
		}
	}
}

func TestLinearEase(t *testing.T) {
	testCases := map[string]struct {
		Input    float64
		Expected float64
	}{
		"zero": {
			Input:    0,
			Expected: 0,
		},
		"one": {
			Input:    1,
			Expected: 1,
		},
		"0.5": {
			Input:    0.5,
			Expected: 0.5,
		},
		"0.1": {
			Input:    0.1,
			Expected: 0.1,
		},
	}

	for name, testCase := range testCases {
		if result := math.LinearEase(testCase.Input); !math.IsFloatEqual(testCase.Expected, result) {
			t.Errorf("%s failed. Expected %v, received %v.", name, testCase.Expected, result)
		}
	}
}

func TestNewInterpolator(t *testing.T) {
	testCases := map[string]struct {
		Percentage, A, B float64
		EasingFunc       math.EasingFunction
		Expected         float64
	}{
		"zero": {
			Percentage: 0,
			A:          4,
			B:          5,
			EasingFunc: math.LinearEase,
			Expected:   4,
		},
		"one": {
			Percentage: 1,
			A:          4,
			B:          5,
			EasingFunc: math.LinearEase,
			Expected:   5,
		},
		"greater than 1": {
			Percentage: 1.7,
			A:          4,
			B:          5,
			EasingFunc: math.LinearEase,
			Expected:   5,
		},
		"less than 0": {
			Percentage: -1.7,
			A:          4,
			B:          5,
			EasingFunc: math.LinearEase,
			Expected:   4,
		},
		"in between": {
			Percentage: 0.5,
			A:          4,
			B:          5,
			EasingFunc: math.LinearEase,
			Expected:   4.5,
		},
		"close to 1": {
			Percentage: 0.95,
			A:          4,
			B:          5,
			EasingFunc: math.LinearEase,
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
		testFunc := math.NewInterpolator(testCase.EasingFunc)
		result := testFunc(testCase.Percentage, testCase.A, testCase.B)
		if result != testCase.Expected {
			t.Errorf("'%s' failed. Expected %v, received %v", name, testCase.Expected, result)
		}
	}
}
