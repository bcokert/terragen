package interpolation

// An Interpolator interpolates between two values interpolate between two values
type Interpolator func(percentage, a, b float64) float64

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
