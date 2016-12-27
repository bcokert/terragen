package noise

import (
	"math"

	"github.com/bcokert/terragen/log"
)

// standard1DInputParams are used to verify equality of 1D noise functions
// They're helpful for various testing of noise functions
var standard1DInputParams = [][]float64{
	[]float64{-5400},
	[]float64{-342.612},
	[]float64{-1},
	[]float64{-0.0001},
	[]float64{0},
	[]float64{0.0001},
	[]float64{0.999},
	[]float64{1},
	[]float64{1.001},
	[]float64{42},
	[]float64{654.23452},
	[]float64{99875},
}

// standard2DInputParams are used to verify equality of 2D noise functions
// They're helpful for various testing of noise functions
var standard2DInputParams = [][]float64{
	[]float64{-52345, -13.45},
	[]float64{-1, -1},
	[]float64{-1, 0},
	[]float64{0, -1},
	[]float64{0, 0},
	[]float64{1, 1},
	[]float64{0.999, 1.001},
	[]float64{0, 500},
	[]float64{9999, 2244},
	[]float64{-33234, 0.0001},
}

// Function is an n dimensional noise function
type Function func(t []float64) float64

// IsEqual determines whether two noise functions are equivalent empirically, with the given number of dimensions
func (leftFn Function) IsEqual(rightFn Function, dimensions int) bool {
	var dimensionParams [][]float64
	switch dimensions {
	case 1:
		dimensionParams = standard1DInputParams
	case 2:
		dimensionParams = standard2DInputParams
	default:
		log.Debug("IsEqual called with an invalid dimension (%d). Returning false.", dimensions)
		return false
	}

	for _, params := range dimensionParams {
		left, right := leftFn(params), rightFn(params)

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
