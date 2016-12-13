package controller_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"testing"

	"github.com/bcokert/terragen/controller"
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
			ExpectedBody: `{"error": "Invalid param: (From must be an array of integers)"}`,
			ExpectedCode: http.StatusBadRequest,
		},
		"illegal from": {
			Url:          "/noise?from=52,banana&to=12&resolution=14&noiseFunction=white",
			ExpectedBody: `{"error": "Invalid param: (From must be an array of integers)"}`,
			ExpectedCode: http.StatusBadRequest,
		},
		"float from": {
			Url:          "/noise?from=52.32&to=12&resolution=14&noiseFunction=white",
			ExpectedBody: `{"error": "Invalid param: (From must be an array of integers)"}`,
			ExpectedCode: http.StatusBadRequest,
		},
		"missing to": {
			Url:          "/noise?from=14&resolution=14&noiseFunction=white",
			ExpectedBody: `{"error": "Invalid param: (To must be an array of integers)"}`,
			ExpectedCode: http.StatusBadRequest,
		},
		"illegal to": {
			Url:          "/noise?from=15&to=52,banana&resolution=14&noiseFunction=white",
			ExpectedBody: `{"error": "Invalid param: (To must be an array of integers)"}`,
			ExpectedCode: http.StatusBadRequest,
		},
		"float to": {
			Url:          "/noise?from=15&to=52.55&resolution=14&noiseFunction=white",
			ExpectedBody: `{"error": "Invalid param: (To must be an array of integers)"}`,
			ExpectedCode: http.StatusBadRequest,
		},
		"from and to diff lengths": {
			Url:          "/noise?from=15&to=52,77&resolution=14&noiseFunction=white",
			ExpectedBody: `{"error": "Invalid param: (From and To must be the same length)"}`,
			ExpectedCode: http.StatusBadRequest,
		},
		"from greater than to": {
			Url:          "/noise?from=15&to=10&resolution=14&noiseFunction=white",
			ExpectedBody: `{"error": "Invalid param: (The value of To must be greater than the value of From in each dimension)"}`,
			ExpectedCode: http.StatusBadRequest,
		},
		"from equals to": {
			Url:          "/noise?from=15&to=15&resolution=14&noiseFunction=white",
			ExpectedBody: `{"error": "Invalid param: (The value of To must be greater than the value of From in each dimension)"}`,
			ExpectedCode: http.StatusBadRequest,
		},
		"missing resolution": {
			Url:          "/noise?from=14&to=16&noiseFunction=white",
			ExpectedBody: `{"error": "Invalid param: (Resolution must be a positive integer)"}`,
			ExpectedCode: http.StatusBadRequest,
		},
		"illegal resolution": {
			Url:          "/noise?from=15&to=52&resolution=51x23&noiseFunction=white",
			ExpectedBody: `{"error": "Invalid param: (Resolution must be a positive integer)"}`,
			ExpectedCode: http.StatusBadRequest,
		},
		"resolution too small": {
			Url:          "/noise?from=15&to=52&resolution=0&noiseFunction=white",
			ExpectedBody: `{"error": "Invalid param: (Resolution must be a positive integer)"}`,
			ExpectedCode: http.StatusBadRequest,
		},
		"missing noiseFunction": {
			Url:          "/noise?from=14&to=15&resolution=42",
			ExpectedBody: `{"error": "Invalid param: (NoiseFunction must be a valid preset)"}`,
			ExpectedCode: http.StatusBadRequest,
		},
		"invalid noise function": {
			Url:          "/noise?from=15&to=52&resolution=51&noiseFunction=ae1234",
			ExpectedBody: `{"error": "Invalid param: (NoiseFunction must be a valid preset)"}`,
			ExpectedCode: http.StatusBadRequest,
		},
		"invalid seed": {
			Url:          "/noise?from=15&to=52&resolution=51&noiseFunction=white&seed=banana",
			ExpectedBody: `{"error": "Invalid param: (Seed must be a positive integer)"}`,
			ExpectedCode: http.StatusBadRequest,
		},
	}

	for name, testCase := range testCases {
		server := &controller.Server{
			Marshal: json.Marshal,
		}
		response := testutils.ExecuteTestRequest(router.CreateDefaultRouter(server, "."), http.MethodGet, testCase.Url, nil)

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
			ExpectedBody: `{"error": "Failed to generate noise: (Failed to marshal man!)"}`,
			ExpectedCode: http.StatusInternalServerError,
		},
	}

	for name, testCase := range testCases {
		server := &controller.Server{
			Marshal: func(v interface{}) ([]byte, error) {
				return []byte{}, fmt.Errorf("Failed to marshal man!")
			},
		}
		response := testutils.ExecuteTestRequest(router.CreateDefaultRouter(server, "."), http.MethodGet, testCase.Url, nil)

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
		PresetName       string
		Dimensions       []int
		PresetCollection map[string]presets.Preset
	}{
		"violet":    {PresetName: "violet", Dimensions: []int{1, 2}, PresetCollection: presets.SpectralPresets},
		"blue":      {PresetName: "blue", Dimensions: []int{1, 2}, PresetCollection: presets.SpectralPresets},
		"white":     {PresetName: "white", Dimensions: []int{1, 2}, PresetCollection: presets.SpectralPresets},
		"pink":      {PresetName: "pink", Dimensions: []int{1, 2}, PresetCollection: presets.SpectralPresets},
		"red":       {PresetName: "red", Dimensions: []int{1, 2}, PresetCollection: presets.SpectralPresets},
		"rawPerlin": {PresetName: "rawPerlin", Dimensions: []int{2}, PresetCollection: presets.LatticePresets},
	}

	// For each preset, for each dimension, we have a set of test cases (aka sets of params)
	testCaseParams := map[int][]struct {
		From       []int
		To         []int
		Resolution int
	}{
		1: {
			{From: []int{-3}, To: []int{-1}, Resolution: 4},
			{From: []int{-0}, To: []int{6}, Resolution: 50},
			{From: []int{-12}, To: []int{55}, Resolution: 1},
			{From: []int{-5}, To: []int{-4}, Resolution: 2},
		},
		2: {
			{From: []int{-3, 0}, To: []int{0, 5}, Resolution: 4},
			{From: []int{-1, -1}, To: []int{0, 0}, Resolution: 50},
			{From: []int{0, 21}, To: []int{1, 23}, Resolution: 1},
			{From: []int{55, 91}, To: []int{56, 999}, Resolution: 2},
			{From: []int{-4, 7}, To: []int{-2, 8}, Resolution: 2},
		},
	}

	for name, testCase := range testCases {
		server := &controller.Server{
			Marshal: json.Marshal,
		}

		for _, dimension := range testCase.Dimensions {
			for i, params := range testCaseParams[dimension] {
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

				url := fmt.Sprintf("/noise?from=%s&to=%s&resolution=%s&noiseFunction=%s&seed=42", fromString, toString, resolutionString, testCase.PresetName)
				noiseFunction := testCase.PresetCollection[testCase.PresetName](random.NewDefaultSource(42), []float64{1, 2, 4, 8, 16, 32, 64})
				response := testutils.ExecuteTestRequest(router.CreateDefaultRouter(server, "."), http.MethodGet, url, nil)

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
}
