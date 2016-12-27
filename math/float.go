package math

import "math"

// IsFloatEqual compares two floats to see if they're close enough to be considered equal
func IsFloatEqual(a, b float64) bool {
	return math.Abs(a-b) <= 0.00000000000001
}
