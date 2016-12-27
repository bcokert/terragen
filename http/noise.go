package http

import (
	"net/http"
	"net/url"
	"strconv"

	"errors"
	"fmt"
	"time"

	"github.com/bcokert/terragen/log"
	"github.com/bcokert/terragen/math"
	"github.com/bcokert/terragen/noise"
	"github.com/julienschmidt/httprouter"
)

// HandleNoise generates noise with the given params. It is an idempotent call
func HandleNoise() httprouter.Handle {
	return Handle(func(response http.ResponseWriter, request *http.Request, _ httprouter.Params) (interface{}, int) {
		log.Info("Request Started: %s %s", request.Method, request.URL.String())

		// Validate the params, and get the related data
		params, err := validateNoiseParams(request.URL.Query())
		if err != nil {
			return fmt.Errorf("Invalid param: (%s)", err.Error()), http.StatusBadRequest
		}

		// Generate noise from the given params and preset
		log.Info("Generating noise with the following params: %+v", params)
		noise := noise.NewNoise(params.presetName)
		noiseFn := params.preset(math.NewDefaultSource(params.seed), []float64{1, 2, 4, 8, 16, 32, 64})
		noise.Generate(params.from, params.to, params.resolution, noiseFn)

		return noise, http.StatusOK
	})
}

type queryParams struct {
	from       []int
	to         []int
	resolution int
	presetName string
	preset     noise.Preset
	seed       int64
}

func validateNoiseParams(params url.Values) (response queryParams, err error) {
	from := params.Get("from")
	to := params.Get("to")
	resolution := params.Get("resolution")
	noiseFunction := params.Get("noiseFunction")
	seed := params.Get("seed")

	// Validate from and to values
	response.from = []int{0, 0}
	if from != "" {
		response.from = ParseIntArray(from)
		if len(response.from) == 0 {
			return queryParams{}, errors.New("From must be an array of integers")
		}
	}

	response.to = []int{5, 5}
	if to != "" {
		response.to = ParseIntArray(to)
		if len(response.to) == 0 {
			return queryParams{}, errors.New("To must be an array of integers")
		}
	}

	if len(response.to) != len(response.from) {
		return queryParams{}, errors.New("From and To must be the same length")
	}

	for i := range response.from {
		if response.from[i] >= response.to[i] {
			return queryParams{}, errors.New("The value of To must be greater than the value of From in each dimension")
		}
	}

	// Validate resolution value
	response.resolution = 20
	if resolution != "" {
		if response.resolution, err = strconv.Atoi(resolution); err != nil {
			return queryParams{}, errors.New("Resolution must be a positive integer")
		}
		if response.resolution < 1 {
			return queryParams{}, errors.New("Resolution must be a positive integer")
		}
	}

	// Validate noise params value
	response.presetName = "red"
	if noiseFunction != "" {
		response.presetName = noiseFunction
	}
	response.preset = searchPresets(response.presetName)
	if response.preset == nil {
		return queryParams{}, errors.New("NoiseFunction must be a valid preset")
	}

	// Validate seed, or generate if missing
	response.seed = time.Now().Unix()
	if seed != "" {
		if response.seed, err = strconv.ParseInt(seed, 10, 0); err != nil {
			return queryParams{}, errors.New("Seed must be a positive integer")
		}
	}

	return response, nil
}

// search through each preset collection for the specified preset
func searchPresets(name string) noise.Preset {
	for noiseFn, preset := range noise.SpectralPresets {
		if name == noiseFn {
			return preset
		}
	}

	for noiseFn, preset := range noise.LatticePresets {
		if name == noiseFn {
			return preset
		}
	}

	return nil
}
