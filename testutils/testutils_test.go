package testutils_test

import (
	"testing"

	"github.com/bcokert/terragen/testutils"
)

func TestIsFloatEqual(t *testing.T) {
	testCases := map[string]struct {
		A, B     float64
		Expected bool
	}{
		"zero": {
			A:        0,
			B:        0,
			Expected: true,
		},
		"large": {
			A:        623462362.23452345,
			B:        623462362.23452345,
			Expected: true,
		},
		"almost": {
			A:        623462362.2345234,
			B:        623462362.2345235,
			Expected: false,
		},
	}

	for name, testCase := range testCases {
		result := testutils.IsFloatEqual(testCase.A, testCase.B)
		if result != testCase.Expected {
			t.Errorf("'%s' failed. Expected %v, received %v", name, testCase.Expected, result)
		}
	}
}
