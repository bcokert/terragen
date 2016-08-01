package interpolation

/* Easing functions map a linear percentage into an eased version */

// DampCubicEase does a cubic easing that
func DampCubicEase(t float64) float64 {
	return 3*t*t - 2*t*t*t
}
