package http

import (
	"net/http"
	"net/url"
	"strconv"

	errs "errors"
	"github.com/bcokert/terragen/log"
	"github.com/bcokert/terragen/model"
	"github.com/bcokert/terragen/presets"
	"github.com/bcokert/terragen/random"
	"github.com/julienschmidt/httprouter"
	"time"
	"fmt"
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
		noise := model.NewNoise(params.presetName)
		noiseFn := params.preset(random.NewDefaultSource(params.seed), []float64{1, 2, 4, 8, 16, 32, 64})
		noise.Generate(params.from, params.to, params.resolution, noiseFn)

		return noise, http.StatusOK
	})
}

type queryParams struct {
	from       []int
	to         []int
	resolution int
	presetName string
	preset     presets.Preset
	seed       int64
}

func validateNoiseParams(params url.Values) (response queryParams, err error) {
	from := params.Get("from")
	to := params.Get("to")
	resolution := params.Get("resolution")
	noiseFunction := params.Get("noiseFunction")
	seed := params.Get("seed")

	// Check for missing params
	if from == "" {
		return queryParams{}, errs.New("From must be an array of integers")
	}

	if to == "" {
		return queryParams{}, errs.New("To must be an array of integers")
	}

	if resolution == "" {
		return queryParams{}, errs.New("Resolution must be a positive integer")
	}

	if noiseFunction == "" {
		return queryParams{}, errs.New("NoiseFunction must be a valid preset")
	}

	// Validate from and to values
	response.from = parseIntArray(from)
	if len(response.from) == 0 {
		return queryParams{}, errs.New("From must be an array of integers")
	}

	response.to = parseIntArray(to)
	if len(response.to) == 0 {
		return queryParams{}, errs.New("To must be an array of integers")
	}

	if len(response.to) != len(response.from) {
		return queryParams{}, errs.New("From and To must be the same length")
	}

	for i := range response.from {
		if response.from[i] >= response.to[i] {
			return queryParams{}, errs.New("The value of To must be greater than the value of From in each dimension")
		}
	}

	// Validate resolution value
	if response.resolution, err = strconv.Atoi(resolution); err != nil {
		return queryParams{}, errs.New("Resolution must be a positive integer")
	}
	if response.resolution < 1 {
		return queryParams{}, errs.New("Resolution must be a positive integer")
	}

	// Validate noise params value
	response.presetName = noiseFunction
	response.preset = searchPresets(noiseFunction)
	if response.preset == nil {
		return queryParams{}, errs.New("NoiseFunction must be a valid preset")
	}

	// Validate seed, or generate if missing
	if seed == "" {
		response.seed = time.Now().Unix()
	} else {
		if response.seed, err = strconv.ParseInt(seed, 10, 0); err != nil {
			return queryParams{}, errs.New("Seed must be a positive integer")
		}
	}

	return response, nil
}

// search through each preset collection for the specified preset
func searchPresets(name string) presets.Preset {
	for noiseFn, preset := range presets.SpectralPresets {
		if name == noiseFn {
			return preset
		}
	}

	for noiseFn, preset := range presets.LatticePresets {
		if name == noiseFn {
			return preset
		}
	}

	return nil
}
