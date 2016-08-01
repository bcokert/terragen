package vector_test

import (
	"testing"

	"math"

	"github.com/bcokert/terragen/random"
	"github.com/bcokert/terragen/vector"
)

func TestDefaultRandomGridCache_Implements(t *testing.T) {
	var _ vector.GridCache = &vector.DefaultRandomGridCache{}
}

func TestDefaultRandomGridCache_Get(t *testing.T) {
	testCases := map[string]struct {
		Cache    vector.GridCache
		Expected []struct {
			X, Y           int
			ExpectedVector vector.Vec2
		}
	}{
		"cache hits": {
			Cache: vector.NewDefaultRandomGridCache(&random.IncrementingSourceMock{IncrementingResult: 0}),
			Expected: []struct {
				X, Y           int
				ExpectedVector vector.Vec2
			}{
				{X: -9, Y: -9, ExpectedVector: vector.NewVec2(1, 3)},
				{X: -9, Y: -8, ExpectedVector: vector.NewVec2(5, 7)},
				{X: 0, Y: 0, ExpectedVector: vector.NewVec2(379*2-1, 380*2-1)},
				{X: 10, Y: 10, ExpectedVector: vector.NewVec2(799*2-1, 800*2-1)},
			},
		},
		"cache misses": {
			Cache: vector.NewDefaultRandomGridCache(&random.IncrementingSourceMock{IncrementingResult: 0}),
			Expected: []struct {
				X, Y           int
				ExpectedVector vector.Vec2
			}{
				{X: -10, Y: -10, ExpectedVector: vector.NewVec2(801*2-1, 802*2-1)},
				{X: 12, Y: 3, ExpectedVector: vector.NewVec2(803*2-1, 804*2-1)},
			},
		},
		"misses then hits": {
			Cache: vector.NewDefaultRandomGridCache(&random.IncrementingSourceMock{IncrementingResult: 0}),
			Expected: []struct {
				X, Y           int
				ExpectedVector vector.Vec2
			}{
				{X: -10, Y: -10, ExpectedVector: vector.NewVec2(801*2-1, 802*2-1)},
				{X: -10, Y: -10, ExpectedVector: vector.NewVec2(801*2-1, 802*2-1)},
				{X: 12, Y: 3, ExpectedVector: vector.NewVec2(803*2-1, 804*2-1)},
				{X: -10, Y: -10, ExpectedVector: vector.NewVec2(801*2-1, 802*2-1)},
				{X: 12, Y: 3, ExpectedVector: vector.NewVec2(803*2-1, 804*2-1)},
				{X: 12, Y: 3, ExpectedVector: vector.NewVec2(803*2-1, 804*2-1)},
				{X: 12, Y: 3, ExpectedVector: vector.NewVec2(803*2-1, 804*2-1)},
				{X: -10, Y: -10, ExpectedVector: vector.NewVec2(801*2-1, 802*2-1)},
				{X: -12, Y: 0, ExpectedVector: vector.NewVec2(805*2-1, 806*2-1)},
			},
		},
	}

	for name, testCase := range testCases {
		for _, inputs := range testCase.Expected {
			result := testCase.Cache.Get(inputs.X, inputs.Y)
			inputs.ExpectedVector.Normalize()
			if !result.IsEqual(inputs.ExpectedVector) {
				t.Errorf("'%s' failed on inputs %v,%v. Expected %v, received %v", name, inputs.X, inputs.Y, inputs.ExpectedVector, result)
			}
		}
	}
}

func TestMockGridCache_Get(t *testing.T) {
	testCases := map[string]struct {
		X, Y     int
		Cache    *vector.MockGridCache
		Expected vector.Vec2
	}{
		"basic": {
			X:        1,
			Y:        0,
			Cache:    &vector.MockGridCache{},
			Expected: vector.Vec2{1, 0},
		},
		"random": {
			X:        52,
			Y:        23,
			Cache:    &vector.MockGridCache{},
			Expected: vector.Vec2{52 / math.Sqrt(52*52+23*23), 23 / math.Sqrt(52*52+23*23)},
		},
		"zero": {
			X:        0,
			Y:        0,
			Cache:    &vector.MockGridCache{},
			Expected: vector.Vec2{1, 0},
		},
	}

	for name, testCase := range testCases {
		result := testCase.Cache.Get(testCase.X, testCase.Y)
		if !testCase.Expected.IsEqual(result) {
			t.Errorf("'%s' failed. Expected %v, received %v", name, testCase.Expected, result)
		}
	}
}
