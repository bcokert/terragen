package interpolation_test

import (
	"testing"

	"github.com/bcokert/terragen/interpolation"
	"github.com/bcokert/terragen/testutils"
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
		if result := interpolation.DampCubicEase(testCase.Input); !testutils.IsFloatEqual(testCase.Expected, result) {
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
		if result := interpolation.LinearEase(testCase.Input); !testutils.IsFloatEqual(testCase.Expected, result) {
			t.Errorf("%s failed. Expected %v, received %v.", name, testCase.Expected, result)
		}
	}
}
