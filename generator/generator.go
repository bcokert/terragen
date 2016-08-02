package generator

import (
	"github.com/bcokert/terragen/interpolation"
	"github.com/bcokert/terragen/noise"
	"github.com/bcokert/terragen/random"
	"github.com/bcokert/terragen/vector"
)

// Random Builds a noise function that returns random floats in the range [0, 1]
func Random(random random.Source) noise.Function {
	return func(t []float64) float64 {
		return random.Float64()
	}
}

// Perlin builds a noise function that returns Lattice Gradient noise values as described by Ken Perlin
// TODO: Handle other dimensions than 2
func Perlin(cache vector.GridCache, interpolator interpolation.Interpolator) noise.Function {
	return func(t []float64) float64 {
		// Find the surrounding grid points of the lattice
		topLeftX, topLeftY := int(t[0]), int(t[1])
		gridPoints := [4][2]int{
			[2]int{topLeftX, topLeftY},
			[2]int{topLeftX + 1, topLeftY},
			[2]int{topLeftX, topLeftY - 1},
			[2]int{topLeftX + 1, topLeftY - 1},
		}

		// Generate influences by computing the dot product of a random vector and a direction vector for each grid point
		influences := [4]float64{}
		for i, gridPoint := range gridPoints {
			directionVector := vector.NewVec2(t[0]-float64(gridPoint[0]), t[1]-float64(gridPoint[1]))
			gridVector := cache.Get(gridPoint[0], gridPoint[1])
			influences[i] = gridVector.Dot(directionVector)
		}

		// Interpolate the influence values to get the noise value
		avgAB := interpolator(t[0]-float64(topLeftX), influences[0], influences[1])
		avgCD := interpolator(t[0]-float64(topLeftX), influences[2], influences[3])
		return interpolator(t[1]-float64(topLeftY-1), avgCD, avgAB)
	}
}
