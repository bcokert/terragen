package vector_test

import (
	"math"
	"testing"

	"github.com/bcokert/terragen/random"
	"github.com/bcokert/terragen/vector"
)

func TestNewVec2(t *testing.T) {
	testCases := map[string]struct {
		X, Y     float64
		Expected vector.Vec2
	}{
		"zero": {
			X:        0,
			Y:        0,
			Expected: vector.Vec2{0: 0, 1: 0},
		},
		"large": {
			X:        12342142153.5123,
			Y:        -923512512,
			Expected: vector.Vec2{0: 12342142153.5123, 1: -923512512},
		},
	}

	for name, testCase := range testCases {
		result := vector.NewVec2(testCase.X, testCase.Y)
		if !result.IsEqual(testCase.Expected) {
			t.Errorf("'%s' failed. Expected %v, received %v", name, testCase.Expected, result)
		}
	}
}

func TestRandomDirectionVec2(t *testing.T) {
	testCases := map[string]struct {
		RandomSource random.Source
		Expected     vector.Vec2
	}{
		"one component": {
			RandomSource: &random.IncrementingSourceMock{IncrementingResult: -0.5},
			Expected:     vector.Vec2{0: 0, 1: 1},
		},
		"two component": {
			RandomSource: &random.IncrementingSourceMock{IncrementingResult: 1},
			Expected:     vector.Vec2{0: 3 / math.Sqrt(34), 1: 5 / math.Sqrt(34)},
		},
		"converts zero vector to a direction vector": {
			RandomSource: &random.ConstantSourceMock{ConstantResult: 0.5},
			Expected:     vector.Vec2{0: 1, 1: 0},
		},
	}

	for name, testCase := range testCases {
		result := vector.RandomDirectionVec2(testCase.RandomSource)
		if !result.IsEqual(testCase.Expected) {
			t.Errorf("'%s' failed. Expected %v, received %v", name, testCase.Expected, result)
		}
	}
}

func TestVec2_IsEqual(t *testing.T) {
	testCases := map[string]struct {
		Vec1, Vec2 vector.Vec2
		Expected   bool
	}{
		"zero": {
			Vec1:     vector.Vec2{0: 0, 1: 0},
			Vec2:     vector.Vec2{0: 0, 1: 0},
			Expected: true,
		},
		"equal": {
			Vec1:     vector.Vec2{0: 0.0956276, 1: -4532926.2364},
			Vec2:     vector.Vec2{0: 0.0956276, 1: -4532926.2364},
			Expected: true,
		},
		"not equal": {
			Vec1:     vector.Vec2{0: 3, 1: 7},
			Vec2:     vector.Vec2{0: 0, 1: 0},
			Expected: false,
		},
		"almost equal": {
			Vec1:     vector.Vec2{0: 0.0956276, 1: -4532926.2364},
			Vec2:     vector.Vec2{0: 0.0956277, 1: -4532926.2364},
			Expected: false,
		},
	}

	for name, testCase := range testCases {
		result := testCase.Vec1.IsEqual(testCase.Vec2)
		if result != testCase.Expected {
			t.Errorf("'%s' failed. Expected %v, received %v", name, testCase.Expected, result)
		}
	}
}

func TestVec2_Length(t *testing.T) {
	testCases := map[string]struct {
		Vec      vector.Vec2
		Expected float64
	}{
		"zero": {
			Vec:      vector.Vec2{0: 0, 1: 0},
			Expected: 0,
		},
		"orthogonal": {
			Vec:      vector.Vec2{0: 0, 1: 5},
			Expected: 5,
		},
		"negatives": {
			Vec:      vector.Vec2{0: -3, 1: 1.5},
			Expected: math.Sqrt(9 + 1.5*1.5),
		},
	}

	for name, testCase := range testCases {
		result := testCase.Vec.Length()
		if result != testCase.Expected {
			t.Errorf("'%s' failed. Expected %v, received %v", name, testCase.Expected, result)
		}
	}
}

func TestVec2_Normalize(t *testing.T) {
	testCases := map[string]struct {
		Vec              vector.Vec2
		Expected         vector.Vec2
		ExpectedErrorMsg string
	}{
		"zero": {
			Vec:              vector.Vec2{0: 0, 1: 0},
			Expected:         vector.Vec2{},
			ExpectedErrorMsg: "Tried to normalize a vector with length zero: [0 0]",
		},
		"orthogonal": {
			Vec:              vector.Vec2{0: 0, 1: 5},
			Expected:         vector.Vec2{0: 0, 1: 1},
			ExpectedErrorMsg: "",
		},
		"regular": {
			Vec:              vector.Vec2{0: -3, 1: 1.5},
			Expected:         vector.Vec2{0: -3 / math.Sqrt(9+1.5*1.5), 1: 1.5 / math.Sqrt(9+1.5*1.5)},
			ExpectedErrorMsg: "",
		},
	}

	for name, testCase := range testCases {
		err := testCase.Vec.Normalize()
		if testCase.ExpectedErrorMsg == "" {
			if err != nil {
				t.Errorf("'%s' failed. An unexpected error occurred: %v", name, err.Error())
			}
			if !testCase.Vec.IsEqual(testCase.Expected) {
				t.Errorf("'%s' failed. Expected %v, received %v", name, testCase.Expected, testCase.Vec)
			}
		} else {
			if testCase.ExpectedErrorMsg != err.Error() {
				t.Errorf("'%s' failed. Expected error '%v', received '%v'", name, testCase.ExpectedErrorMsg, err.Error())
			}
		}
	}
}

func TestVec2_Dot(t *testing.T) {
	testCases := map[string]struct {
		Vec1, Vec2 vector.Vec2
		Expected   float64
	}{
		"zero": {
			Vec1:     vector.Vec2{0: 0, 1: 0},
			Vec2:     vector.Vec2{0: 0, 1: 0},
			Expected: 0,
		},
		"orthogonal": {
			Vec1:     vector.Vec2{0: 1, 1: 0},
			Vec2:     vector.Vec2{0: 0, 1: 1},
			Expected: 0,
		},
		"single component": {
			Vec1:     vector.Vec2{0: 0, 1: 2},
			Vec2:     vector.Vec2{0: 0, 1: 3},
			Expected: 6,
		},
		"complex": {
			Vec1:     vector.Vec2{0: -3, 1: 4},
			Vec2:     vector.Vec2{0: 1, 1: 2},
			Expected: -3 + 8,
		},
	}

	for name, testCase := range testCases {
		result := testCase.Vec1.Dot(testCase.Vec2)
		if result != testCase.Expected {
			t.Errorf("'%s' failed. Expected %v, received %v", name, testCase.Expected, result)
		}
	}
}
