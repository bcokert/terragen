package vector

// RandomSourceMock is a mock RandomSource who's Float64 method turns an incrementing value on each call, which can be reset
type RandomSourceMock struct {
	IncrementingResult float64
}

// Float64 is a mock random number generator that returns the last number returned + 1
func (source *RandomSourceMock) Float64() float64 {
	source.IncrementingResult++
	return source.IncrementingResult
}

// RandomSourceConstantMock is a mock RandomSource who's Float64 always returns the ConstantResult
type RandomSourceConstantMock struct {
	ConstantResult float64
}

// Float64 is a mock random number generator that returns the last number returned + 1
func (source *RandomSourceConstantMock) Float64() float64 {
	return source.ConstantResult
}
