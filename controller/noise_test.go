package controller_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"testing"

	"github.com/bcokert/terragen/controller"
	"github.com/bcokert/terragen/errors"
	"github.com/bcokert/terragen/model"
	"github.com/bcokert/terragen/presets"
	"github.com/bcokert/terragen/random"
	"github.com/bcokert/terragen/router"
	"github.com/bcokert/terragen/testutils"
)

func TestGetNoise_InputValidation(t *testing.T) {
	testCases := map[string]struct {
		Url          string
		ExpectedBody string
		ExpectedCode int
	}{
		"missing from": {
			Url:          "/noise?to=12&resolution=14&noiseFunction=white",
			ExpectedBody: `{"error": "Noise - Invalid 'from' param: (Invalid. Must not be empty)"}`,
			ExpectedCode: http.StatusBadRequest,
		},
		"illegal from": {
			Url:          "/noise?from=52,banana&to=12&resolution=14&noiseFunction=white",
			ExpectedBody: `{"error": "Noise - Invalid 'from' param: (Illegal. Expected a list of numbers)"}`,
			ExpectedCode: http.StatusBadRequest,
		},
		"missing to": {
			Url:          "/noise?from=14&resolution=14&noiseFunction=white",
			ExpectedBody: `{"error": "Noise - Invalid 'to' param: (Invalid. Must not be empty)"}`,
			ExpectedCode: http.StatusBadRequest,
		},
		"illegal to": {
			Url:          "/noise?from=15&to=52,banana&resolution=14&noiseFunction=white",
			ExpectedBody: `{"error": "Noise - Invalid 'to' param: (Illegal. Expected a list of numbers)"}`,
			ExpectedCode: http.StatusBadRequest,
		},
		"missing resolution": {
			Url:          "/noise?from=14&to=14&noiseFunction=white",
			ExpectedBody: `{"error": "Noise - Invalid 'resolution' param: (Illegal. Expected an integer)"}`,
			ExpectedCode: http.StatusBadRequest,
		},
		"illegal resolution": {
			Url:          "/noise?from=15&to=52&resolution=51x23&noiseFunction=white",
			ExpectedBody: `{"error": "Noise - Invalid 'resolution' param: (Illegal. Expected an integer)"}`,
			ExpectedCode: http.StatusBadRequest,
		},
		"resolution too small": {
			Url:          "/noise?from=15&to=52&resolution=0&noiseFunction=white",
			ExpectedBody: `{"error": "Noise - Invalid 'resolution' param: (Invalid. Must be greater than 0)"}`,
			ExpectedCode: http.StatusBadRequest,
		},
		"missing noiseFunction": {
			Url:          "/noise?from=14&to=14&resolution=42",
			ExpectedBody: `{"error": "Noise - Invalid 'noiseFunction' param: (Invalid. Expected a noise function preset or id)"}`,
			ExpectedCode: http.StatusBadRequest,
		},
		"non-preset noise function": {
			Url:          "/noise?from=15&to=52&resolution=51&noiseFunction=ae1234",
			ExpectedBody: fmt.Sprintf(`{"error": "Noise - Invalid 'noiseFunction' param: (%s)"}`, errors.UnsupportedError("Loading Noise Functions by Id")),
			ExpectedCode: http.StatusBadRequest,
		},
	}

	for name, testCase := range testCases {
		server := &controller.Server{
			Seed:    42,
			Marshal: json.Marshal,
		}
		response := testutils.ExecuteTestRequest(router.CreateDefaultRouter(server), http.MethodGet, testCase.Url, nil)

		if response.Body.String() != testCase.ExpectedBody {
			t.Errorf("%s failed. Expected body '%s', received '%s'", name, testCase.ExpectedBody, response.Body.String())
		}

		if response.Code != testCase.ExpectedCode {
			t.Errorf("%s failed. Expected code %d, received %d", name, testCase.ExpectedCode, response.Code)
		}
	}
}

func TestGetNoise_Failure(t *testing.T) {
	testCases := map[string]struct {
		Url          string
		ExpectedBody string
		ExpectedCode int
	}{
		"fails to marshal": {
			Url:          "/noise?from=0&to=10&resolution=2&noiseFunction=white",
			ExpectedBody: `{"error": "Noise - Failed to generate noise: (Failed to marshal man!)"}`,
			ExpectedCode: http.StatusInternalServerError,
		},
	}

	for name, testCase := range testCases {
		server := &controller.Server{
			Seed: 42,
			Marshal: func(v interface{}) ([]byte, error) {
				return []byte{}, fmt.Errorf("Failed to marshal man!")
			},
		}
		response := testutils.ExecuteTestRequest(router.CreateDefaultRouter(server), http.MethodGet, testCase.Url, nil)

		if response.Body.String() != testCase.ExpectedBody {
			t.Errorf("%s failed. Expected body '%s', received '%s'", name, testCase.ExpectedBody, response.Body.String())
		}

		if response.Code != testCase.ExpectedCode {
			t.Errorf("%s failed. Expected code %d, received %d", name, testCase.ExpectedCode, response.Code)
		}
	}
}

func TestGetNoise_MissingHTTPMethods(t *testing.T) {
	testCases := map[string]struct {
		Method       string
		ExpectedBody string
		ExpectedCode int
	}{
		"POST": {
			Method:       http.MethodPost,
			ExpectedBody: fmt.Sprintf(`{"error": "Noise - Unsupported http method '%s'"}`, http.MethodPost),
			ExpectedCode: http.StatusBadRequest,
		},
		"DELETE": {
			Method:       http.MethodDelete,
			ExpectedBody: fmt.Sprintf(`{"error": "Noise - Unsupported http method '%s'"}`, http.MethodDelete),
			ExpectedCode: http.StatusBadRequest,
		},
		"PUT": {
			Method:       http.MethodPut,
			ExpectedBody: fmt.Sprintf(`{"error": "Noise - Unsupported http method '%s'"}`, http.MethodPut),
			ExpectedCode: http.StatusBadRequest,
		},
	}

	for name, testCase := range testCases {
		server := &controller.Server{
			Seed:    42,
			Marshal: json.Marshal,
		}
		response := testutils.ExecuteTestRequest(router.CreateDefaultRouter(server), testCase.Method, "/noise", nil)

		if response.Body.String() != testCase.ExpectedBody {
			t.Errorf("%s failed. Expected body '%s', received '%s'", name, testCase.ExpectedBody, response.Body.String())
		}

		if response.Code != testCase.ExpectedCode {
			t.Errorf("%s failed. Expected code %d, received %d", name, testCase.ExpectedCode, response.Code)
		}
	}
}

func TestGetNoise_Success(t *testing.T) {
	testCases := map[string]struct {
		PresetName string
		Dimensions []int
	}{
		"violet": {PresetName: "violet", Dimensions: []int{1, 2}},
		"blue":   {PresetName: "blue", Dimensions: []int{1, 2}},
		"white":  {PresetName: "white", Dimensions: []int{1, 2}},
		"pink":   {PresetName: "pink", Dimensions: []int{1, 2}},
		"red":    {PresetName: "red", Dimensions: []int{1, 2}},
	}

	// For each preset, for each dimension, we have a set of test cases (aka sets of params)
	testCaseParams := []struct {
		From       []float64
		To         []float64
		Resolution int
	}{
		{From: []float64{-3}, To: []float64{-1}, Resolution: 4},
		{From: []float64{-0}, To: []float64{6}, Resolution: 50},
		{From: []float64{-12}, To: []float64{55}, Resolution: 1},
		{From: []float64{-5}, To: []float64{-5}, Resolution: 2},
		{From: []float64{-3, 0}, To: []float64{0, 5}, Resolution: 4},
		{From: []float64{-1, -1}, To: []float64{0, 0}, Resolution: 50},
		{From: []float64{0, 21}, To: []float64{1, 23}, Resolution: 1},
		{From: []float64{55, 91}, To: []float64{55, 999}, Resolution: 2},
		{From: []float64{-4, 7}, To: []float64{-2, 7}, Resolution: 2},
	}

	for name, testCase := range testCases {
		server := &controller.Server{
			Seed:    42,
			Marshal: json.Marshal,
		}

		for i, params := range testCaseParams {
			froms := []string{}
			for _, from := range params.From {
				froms = append(froms, fmt.Sprintf("%v", from))
			}
			fromString := strings.Join(froms, ",")
			tos := []string{}
			for _, to := range params.To {
				tos = append(tos, fmt.Sprintf("%v", to))
			}
			toString := strings.Join(tos, ",")
			resolutionString := strconv.Itoa(params.Resolution)

			url := fmt.Sprintf("/noise?from=%s&to=%s&resolution=%s&noiseFunction=%s", fromString, toString, resolutionString, testCase.PresetName)
			noiseFunction := presets.SpectralPresets[testCase.PresetName](random.NewDefaultSource(42), []float64{1, 2, 4, 8, 16, 32, 64})
			response := testutils.ExecuteTestRequest(router.CreateDefaultRouter(server), http.MethodGet, url, nil)

			expectedResponse := model.NewNoise(testCase.PresetName)
			expectedResponse.Generate(params.From, params.To, params.Resolution, noiseFunction)

			responseObject := model.Noise{}
			if response.Code != http.StatusOK {
				t.Errorf("%s failed with param set %d. Expected code %d, received %d", name, i, http.StatusOK, response.Code)
				t.Logf("Response: %s", response.Body.String())
				continue
			}

			if err := json.NewDecoder(response.Body).Decode(&responseObject); err != nil {
				t.Errorf("%s failed with param set %d. Failed to decode response: %s", name, i, response.Body.String())
				continue
			}

			if !responseObject.IsEqual(expectedResponse) {
				t.Errorf("%s failed with param set %d. Expected response '%#v', received '%#v'", name, i, expectedResponse, responseObject)
			}
		}
	}
}
