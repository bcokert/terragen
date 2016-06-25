package noise

// Function1D represents a 1d noise function
type Function1D func(t float64) float64

// Function2D represents a 2d noise function
type Function2D func(tx, ty float64) float64
