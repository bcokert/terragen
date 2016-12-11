package generator

import (
	"github.com/bcokert/terragen/interpolation"
	"github.com/bcokert/terragen/noise"
	"github.com/bcokert/terragen/random"
	"github.com/bcokert/terragen/vector"
	"github.com/bcokert/terragen/log"
	"math"
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
		log.Debug("Calculating perlin noise at %v %v", t[0], t[1])

		// Find the surrounding grid points of the lattice
		topLeftX, topLeftY := int(t[0]), int(t[1])
		gridPoints := [4][2]int{
			[2]int{topLeftX, topLeftY},
			[2]int{topLeftX + 1, topLeftY},
			[2]int{topLeftX, topLeftY + 1},
			[2]int{topLeftX + 1, topLeftY + 1},
		}
		log.Debug("Top left is %v %v", topLeftX, topLeftY)
		log.Debug("Calculated grid points are: %v", gridPoints)

		// Generate influences by computing the dot product of a random vector and a direction vector for each grid point
		influences := [4]float64{}
		for i, gridPoint := range gridPoints {
			directionVector := vector.NewVec2(t[0]-float64(gridPoint[0]), t[1]-float64(gridPoint[1]))
			gridVector := cache.Get(gridPoint[0], gridPoint[1])
			influences[i] = gridVector.Dot(directionVector)

			// TODO: Need a system for different log levels
			//log.Debug("Direction vector for grid point %v (%v) is %v", i, gridPoint, directionVector)
			//log.Debug("Cached grid vector for grid point %v (%v) is %v", i, gridPoint, gridVector)
			//log.Debug("Influence (dot product) for grid point %v (%v) is %v", i, gridPoint, influences[i])
		}

		// Interpolate the influence values to get the noise value
		xBias := math.Abs(t[0] - float64(topLeftX))
		yBias := math.Abs(t[1] - float64(topLeftY))
		avgAB := interpolator(xBias, influences[0], influences[1])
		log.Debug("interpolation of grid points 0 and 1 at bias %v is %v", xBias, avgAB)
		avgCD := interpolator(xBias, influences[2], influences[3])
		log.Debug("interpolation of grid points 2 and 3 at bias %v is %v", xBias, avgCD)
		noiseValue := interpolator(yBias, avgAB, avgCD)
		log.Debug("interpolation of interpolations 01 and 23 at bias %v is %v", yBias, noiseValue)
		return noiseValue
	}
}
