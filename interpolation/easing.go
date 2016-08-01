package interpolation

// EasingFunctions map a linear percentage into an eased percentage
type EasingFunction func(t float64) float64

// DampCubicEase does a cubic easing that favors either endpoint
func DampCubicEase(t float64) float64 {
	return 3*t*t - 2*t*t*t
}

// LinearEase simply returns the input percentage. It's also good for mocking easing functions
func LinearEase(t float64) float64 {
	return t
}
