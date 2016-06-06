package random

import "math/rand"

// A NormalFunction is a random function that returns a number between 0 and 1
type NormalFunction func() float64

// SeededNormal Builds a NormalFunction with the given seed, for predictability
func SeededNormal(seed int64) NormalFunction {
	random := rand.New(rand.NewSource(seed))
	return func() float64 {
		return random.Float64()
	}
}
