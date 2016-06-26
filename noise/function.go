package noise

import (
	"math"

	"github.com/bcokert/terragen/log"
)

// Standard1DInputParams are used to verify equality of 1D noise functions
// They're helpful for various testing of noise functions
var Standard1DInputParams = []float64{
	-5400,
	-342.612,
	-1,
	-0.0001,
	0,
	0.0001,
	0.999,
	1,
	1.001,
	42,
	654.23452,
	99875,
}

// Standard2DInputParams are used to verify equality of 2D noise functions
// They're helpful for various testing of noise functions
var Standard2DInputParams = [][2]float64{
	[2]float64{-52345, -13.45},
	[2]float64{-1, -1},
	[2]float64{-1, 0},
	[2]float64{0, -1},
	[2]float64{0, 0},
	[2]float64{1, 1},
	[2]float64{0.999, 1.001},
	[2]float64{0, 500},
	[2]float64{9999, 2244},
	[2]float64{-33234, 0.0001},
}

// Function is an n dimensional noise function
type Function func(t []float64) float64

// Function1D represents a 1d noise function
type Function1D func(t float64) float64

// Function2D represents a 2d noise function
type Function2D func(tx, ty float64) float64

func (fn Function1D) IsEqual(other Function1D) bool {
	for _, param := range Standard1DInputParams {
		left, right := fn(param), other(param)

		if (left < 0) != (right < 0) {
			log.Debug("IsEqual failed on inputs %v: %v != %v", param, left, right)
			return false
		}

		if math.Abs(left-right) > 0.00000000000001 {
			log.Debug("IsEqual failed on inputs %v: %v != %v", param, left, right)
			return false
		}
	}

	return true
}

func (fn Function2D) IsEqual(other Function2D) bool {
	for _, params := range Standard2DInputParams {
		left, right := fn(params[0], params[1]), other(params[0], params[1])

		if (left < 0) != (right < 0) {
			log.Debug("IsEqual failed on inputs %v: %v != %v", params, left, right)
			return false
		}

		if math.Abs(left-right) > 0.00000000000001 {
			log.Debug("IsEqual failed on inputs %v: %v != %v", params, left, right)
			return false
		}
	}

	return true
}
