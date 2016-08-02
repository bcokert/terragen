package controller

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/bcokert/terragen/errors"
	"github.com/bcokert/terragen/log"
	"github.com/bcokert/terragen/model"
	"github.com/bcokert/terragen/presets"
	"github.com/bcokert/terragen/random"
)

// Noise endpoint
// GET: Applies the given noise function (or creates one from the given params), and returns the noise applied to the given range
func (server *Server) Noise(response http.ResponseWriter, request *http.Request) {
	queryParams := request.URL.Query()
	log.Info("Request Started: %s %s", request.Method, request.URL.String())

	response.Header().Set("Access-Control-Allow-Origin", "*")

	switch request.Method {
	case http.MethodGet:
		var from, to []float64
		var resolution int
		var noiseFunction string
		var preset presets.Preset
		var err error
		var jsonResponse []byte

		if from, err = validateFrom(queryParams); err != nil {
			errors.WriteError(response, errors.ErrorWithCause("Noise - Invalid 'from' param", err), http.StatusBadRequest)
			return
		}

		if to, err = validateTo(queryParams); err != nil {
			errors.WriteError(response, errors.ErrorWithCause("Noise - Invalid 'to' param", err), http.StatusBadRequest)
			return
		}

		if resolution, err = validateResolution(queryParams); err != nil {
			errors.WriteError(response, errors.ErrorWithCause("Noise - Invalid 'resolution' param", err), http.StatusBadRequest)
			return
		}

		if noiseFunction, preset, err = validateNoiseFunctionAndRetrievePreset(queryParams); err != nil {
			errors.WriteError(response, errors.ErrorWithCause("Noise - Invalid 'noiseFunction' param", err), http.StatusBadRequest)
			return
		}

		if jsonResponse, err = server.getNoise(from, to, resolution, noiseFunction, preset); err != nil {
			errors.WriteError(response, errors.ErrorWithCause("Noise - Failed to generate noise", err), http.StatusInternalServerError)
			return
		}

		response.Write(jsonResponse)
	default:
		errors.WriteError(response, fmt.Errorf("Noise - Unsupported http method '%s'", request.Method), http.StatusBadRequest)
	}
}

func (server *Server) getNoise(from, to []float64, resolution int, noiseFunction string, preset presets.Preset) ([]byte, error) {
	response := model.NewNoise(noiseFunction)
	noiseFn := preset(random.NewDefaultSource(server.Seed), []float64{1, 2, 4, 8, 16, 32, 64})

	response.Generate(from, to, resolution, noiseFn)
	return server.Marshal(response)
}

func validateFrom(queryParams url.Values) (from []float64, err error) {
	if from, err = ParseFloatArrayParam(queryParams, "from"); err != nil {
		return []float64{}, fmt.Errorf("Illegal. Expected a list of numbers")
	}

	if len(from) == 0 {
		return []float64{}, fmt.Errorf("Invalid. Must not be empty")
	}

	return from, nil
}

func validateTo(queryParams url.Values) (to []float64, err error) {
	if to, err = ParseFloatArrayParam(queryParams, "to"); err != nil {
		return []float64{}, fmt.Errorf("Illegal. Expected a list of numbers")
	}

	if len(to) == 0 {
		return []float64{}, fmt.Errorf("Invalid. Must not be empty")
	}

	return to, nil
}

func validateResolution(queryParams url.Values) (resolution int, err error) {
	if resolution, err = strconv.Atoi(queryParams.Get("resolution")); err != nil {
		return 0, fmt.Errorf("Illegal. Expected an integer")
	}

	if resolution < 1 {
		return 0, fmt.Errorf("Invalid. Must be greater than 0")
	}

	return resolution, nil
}

func validateNoiseFunctionAndRetrievePreset(queryParams url.Values) (noiseFunction string, preset presets.Preset, err error) {
	noiseFunction = queryParams.Get("noiseFunction")

	if noiseFunction == "" {
		return "", nil, fmt.Errorf("Invalid. Expected a noise function preset or id")
	}

	for noiseFn, preset := range presets.SpectralPresets {
		if noiseFunction == noiseFn {
			return noiseFn, preset, nil
		}
	}

	for noiseFn, preset := range presets.LatticePresets {
		if noiseFunction == noiseFn {
			return noiseFn, preset, nil
		}
	}

	return "", nil, errors.UnsupportedError("Loading Noise Functions by Id")
}
