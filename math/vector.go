package math

import (
	"fmt"
	"math"
)

// Vec2 Represents a 2d vector with simple operations
type Vec2 [2]float64

// RandomDirectionVec2 Creates a random normalized Vec2 with dimensions in the range [-1,1]
func RandomDirectionVec2(random Source) Vec2 {
	a, b := (random.Float64()*2)-1, (random.Float64()*2)-1

	// Ensure the vector can be normalized
	if a == 0 && b == 0 {
		a = 1
	}

	vec := Vec2{0: a, 1: b}
	vec.Normalize()
	return vec
}

// IsEqual returns true if the two vectors are equal
func (vec *Vec2) IsEqual(other Vec2) bool {
	for i := range vec {
		if math.Abs(vec[i]-other[i]) > 0.00000000000001 {
			return false
		}
	}

	return true
}

// Length returns the length of a Vec2
func (vec *Vec2) Length() float64 {
	return math.Sqrt(vec[0]*vec[0] + vec[1]*vec[1])
}

// Normalize mutates a Vec2 so that it is normalized. It will fail if the vector is not a direction vector (any component is zero)
func (vec *Vec2) Normalize() error {
	length := vec.Length()
	if length == 0.0 {
		return fmt.Errorf("Tried to normalize a vector with length zero: %v", *vec)
	}

	vec[0] = vec[0] / length
	vec[1] = vec[1] / length

	return nil
}

// Dot takes another vectors and returns the dot product between the two
func (vec *Vec2) Dot(other Vec2) float64 {
	return vec[0]*other[0] + vec[1]*other[1]
}
