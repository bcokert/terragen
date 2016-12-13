package controller

import (
	"net/http"
	"net/url"
	"strconv"

	errs "errors"
	"github.com/bcokert/terragen/errors"
	"github.com/bcokert/terragen/log"
	"github.com/bcokert/terragen/model"
	"github.com/bcokert/terragen/presets"
	"github.com/bcokert/terragen/random"
	"github.com/julienschmidt/httprouter"
)

// Noise endpoint
// Applies the given noise function (or creates one from the given params), and returns the noise applied to the given range
func (server *Server) Noise(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	log.Info("Request Started: %s %s", request.Method, request.URL.String())
	response.Header().Set("Access-Control-Allow-Origin", "*")

	// Validate the params, and get the related data
	params, err := validateNoiseParams(request.URL.Query())
	if err != nil {
		errors.WriteError(response, errors.ErrorWithCause("Invalid param", err), http.StatusBadRequest)
		return
	}

	// Generate noise from the given params and preset
	noise := model.NewNoise(params.PresetName)
	noiseFn := params.Preset(random.NewDefaultSource(server.Seed), []float64{1, 2, 4, 8, 16, 32, 64})
	noise.Generate(params.From, params.To, params.Resolution, noiseFn)

	// Write the result
	json, err := server.Marshal(noise)
	if err != nil {
		errors.WriteError(response, errors.ErrorWithCause("Failed to generate noise", err), http.StatusInternalServerError)
		return
	}

	response.Write(json)
}

type NoiseQueryParams struct {
	From       []int
	To         []int
	Resolution int
	PresetName string
	Preset     presets.Preset
}

func validateNoiseParams(params url.Values) (response NoiseQueryParams, err error) {
	from := params.Get("from")
	to := params.Get("to")
	resolution := params.Get("resolution")
	noiseFunction := params.Get("noiseFunction")

	// Check for missing params
	if from == "" {
		return NoiseQueryParams{}, errs.New("From must be an array of integers")
	}

	if to == "" {
		return NoiseQueryParams{}, errs.New("To must be an array of integers")
	}

	if resolution == "" {
		return NoiseQueryParams{}, errs.New("Resolution must be a positive integer")
	}

	if noiseFunction == "" {
		return NoiseQueryParams{}, errs.New("NoiseFunction must be a valid preset")
	}

	// Validate from and to values
	response.From = ParseIntArray(from)
	if len(response.From) == 0 {
		return NoiseQueryParams{}, errs.New("From must be an array of integers")
	}

	response.To = ParseIntArray(to)
	if len(response.To) == 0 {
		return NoiseQueryParams{}, errs.New("To must be an array of integers")
	}

	if len(response.To) != len(response.From) {
		return NoiseQueryParams{}, errs.New("From and To must be the same length")
	}

	for i := range response.From {
		if response.From[i] >= response.To[i] {
			return NoiseQueryParams{}, errs.New("The value of To must be greater than the value of From in each dimension")
		}
	}

	// Validate resolution value
	if response.Resolution, err = strconv.Atoi(resolution); err != nil {
		return NoiseQueryParams{}, errs.New("Resolution must be a positive integer")
	}
	if response.Resolution < 1 {
		return NoiseQueryParams{}, errs.New("Resolution must be a positive integer")
	}

	// Validate noise params value
	response.PresetName = noiseFunction
	response.Preset = searchPresets(noiseFunction)
	if response.Preset == nil {
		return NoiseQueryParams{}, errs.New("NoiseFunction must be a valid preset")
	}

	return response, nil
}

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
