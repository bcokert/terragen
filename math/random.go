package math

import "math/rand"

// A Source provides a Float64 method that should return random floats between 0 and 1
type Source interface {
	Float64() float64
}

// NewDefaultSource creates a simple source with the standard random implementation
func NewDefaultSource(seed int64) Source {
	return rand.New(rand.NewSource(seed))
}
