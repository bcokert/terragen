package noise

//////// Noise Functions

// A Parametric1D is a Generic Noise function that returns some continuous noise
type Parametric1D func(t float64) float64

// A Parametric2D is a Generic Noise function of two parameters that returns some continuous noise
type Parametric2D func(x, y float64) float64

//////// Octave Functions

// OctaveWeight1D takes a frequency to generate a weight. Used to combine multi frequency noise functions
// A weight is equivalent to an amplitude for periodic Noise Functions
// For example, red noise is a combination of Noise Functions at different frequencies, where lower frequencies have higher weights
type OctaveWeight1D func(freq float64) float64

// OctaveParametric1D takes a frequency to generate Paramtetric Noise Functions, which will be combined along with weights
type OctaveParametric1D func(freq float64) Parametric1D
