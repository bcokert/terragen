package math

// IncrementingSourceMock is a mock Source who's Float64 method turns an incrementing value on each call, which can be reset
type IncrementingSourceMock struct {
	IncrementingResult float64
}

// Float64 is a mock random number generator that returns the last number returned + 1
func (source *IncrementingSourceMock) Float64() float64 {
	source.IncrementingResult++
	return source.IncrementingResult
}

// ConstantSourceMock is a mock Source who's Float64 always returns the ConstantResult
type ConstantSourceMock struct {
	ConstantResult float64
}

// Float64 is a mock random number generator that returns the last number returned + 1
func (source *ConstantSourceMock) Float64() float64 {
	return source.ConstantResult
}
