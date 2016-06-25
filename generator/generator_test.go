package generator_test

import (
	"math/rand"
	"testing"

	"github.com/bcokert/terragen/generator"
)

func TestRandom1D(t *testing.T) {
	testCases := map[string]struct {
		ExpectedFn func() float64
	}{
		"multiple functions with same seed": {
			ExpectedFn: rand.New(rand.NewSource(27)).Float64,
		},
	}

	for name, testCase := range testCases {
		randomFunction1 := generator.Random1D(27)
		randomFunction2 := generator.Random1D(27)
		for i := 0; i < 10; i++ {
			expected := testCase.ExpectedFn()
			if result := randomFunction1(1); result != expected {
				t.Errorf("%s failed. Expected function 1 iteration %d to be %v, received %v", name, i, expected, result)
			}
			if result := randomFunction2(1); result != expected {
				t.Errorf("%s failed. Expected function 1 iteration %d to be %v, received %v", name, i, expected, result)
			}
		}
	}
}

func TestRandom2D(t *testing.T) {
	testCases := map[string]struct {
		ExpectedFn func() float64
	}{
		"multiple functions with same seed": {
			ExpectedFn: rand.New(rand.NewSource(42)).Float64,
		},
	}

	for name, testCase := range testCases {

		random1 := generator.Random2D(42)
		random2 := generator.Random2D(42)

		for i := 0; i < 10; i++ {
			expected := testCase.ExpectedFn()
			if result := random1(1, 2); result != expected {
				t.Errorf("'%s' failed. Expected function 1 iteration %d to be %v, received %v", name, i, expected, result)
			}

			if result := random2(1, 2); result != expected {
				t.Errorf("'%s' failed. Expected function 1 iteration %d to be %v, received %v", name, i, expected, result)
			}
		}
	}
}
