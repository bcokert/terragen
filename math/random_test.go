package math_test

import (
	"math/rand"
	"testing"

	tgmath "github.com/bcokert/terragen/math"
	"github.com/bcokert/terragen/testutils"
)

func TestNewDefaultSource(t *testing.T) {
	testCases := map[string]struct {
		Source   tgmath.Source
		Expected tgmath.Source
	}{
		"Equal Seed": {
			Source:   tgmath.NewDefaultSource(44),
			Expected: rand.New(rand.NewSource(44)),
		},
	}

	for name, testCase := range testCases {
		for i := 0; i < 10; i++ {
			if !testutils.IsFloatEqual(testCase.Source.Float64(), testCase.Expected.Float64()) {
				t.Errorf("'%s' failed on iteration %d.", name, i)
			}
		}
	}
}

func TestConstantSourceMock_Float64(t *testing.T) {
	testCases := map[string]struct {
		Mock            *tgmath.ConstantSourceMock
		ExpectedResults []float64
	}{
		"zero": {
			Mock:            &tgmath.ConstantSourceMock{ConstantResult: 0},
			ExpectedResults: []float64{0, 0, 0, 0, 0},
		},
		"random": {
			Mock:            &tgmath.ConstantSourceMock{ConstantResult: 17.34},
			ExpectedResults: []float64{17.34, 17.34, 17.34},
		},
	}

	for name, testCase := range testCases {
		for i, expected := range testCase.ExpectedResults {
			if result := testCase.Mock.Float64(); result != expected {
				t.Errorf("'%s' failed on iteration %d. Expected %v, received %v", name, i, expected, result)
			}
		}
	}
}

func TestIncrementingSourceMock_Float64(t *testing.T) {
	testCases := map[string]struct {
		Mock            *tgmath.IncrementingSourceMock
		ExpectedResults []float64
	}{
		"zero": {
			Mock:            &tgmath.IncrementingSourceMock{IncrementingResult: 0},
			ExpectedResults: []float64{1, 2, 3, 4, 5},
		},
		"negative": {
			Mock:            &tgmath.IncrementingSourceMock{IncrementingResult: -5},
			ExpectedResults: []float64{-4, -3, -2, -1, 0},
		},
		"random": {
			Mock:            &tgmath.IncrementingSourceMock{IncrementingResult: 16.34},
			ExpectedResults: []float64{17.34, 18.34, 19.34},
		},
	}

	for name, testCase := range testCases {
		for i, expected := range testCase.ExpectedResults {
			if result := testCase.Mock.Float64(); result != expected {
				t.Errorf("'%s' failed on iteration %d. Expected %v, received %v", name, i, expected, result)
			}
		}
	}
}
