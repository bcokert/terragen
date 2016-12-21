package math

// An Interpolator interpolates between two values
type Interpolator func(percentage, a, b float64) float64

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

// NewInterpolator creates an interpolator with the given easing function
// Smaller percentages make the result closer to a, large percentages closer to b
func NewInterpolator(easingFunc EasingFunction) Interpolator {
	return func(percentage, a, b float64) float64 {
		if percentage > 1 {
			percentage = 1
		}
		if percentage < 0 {
			percentage = 0
		}
		delta := easingFunc(percentage)
		return a*(1-delta) + b*delta
	}
}
