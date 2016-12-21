package http_test

import (
	"net/http"
	"testing"

	"encoding/json"
	"fmt"
	"net/http/httptest"
	"strconv"
	"strings"

	tghttp "github.com/bcokert/terragen/http"
	"github.com/bcokert/terragen/math"
	"github.com/bcokert/terragen/model"
	"github.com/bcokert/terragen/presets"
)

func TestHandleNoise(t *testing.T) {
	testCases := map[string]struct {
		From                     string
		To                       string
		Resolution               string
		Preset                   string
		Seed                     string
		ExpectedPresetCollection map[string]presets.Preset // TODO: remove when noise composition refactor is done
		ExpectedStatusCode       int
		ExpectedErrorBody        string
	}{
		"Default from": {
			From: "", To: "3,2", Resolution: "5", Preset: "red", Seed: "922",
			ExpectedPresetCollection: presets.SpectralPresets,
			ExpectedStatusCode:       http.StatusOK,
			ExpectedErrorBody:        "",
		},
		"Default to": {
			From: "1,2", To: "", Resolution: "3", Preset: "blue", Seed: "123",
			ExpectedPresetCollection: presets.SpectralPresets,
			ExpectedStatusCode:       http.StatusOK,
			ExpectedErrorBody:        "",
		},
		"Default resolution": {
			From: "1", To: "10", Resolution: "", Preset: "white", Seed: "34",
			ExpectedPresetCollection: presets.SpectralPresets,
			ExpectedStatusCode:       http.StatusOK,
			ExpectedErrorBody:        "",
		},
		"Default preset": {
			From: "1", To: "3", Resolution: "20", Preset: "", Seed: "9999",
			ExpectedPresetCollection: presets.SpectralPresets,
			ExpectedStatusCode:       http.StatusOK,
			ExpectedErrorBody:        "",
		},
		"All Defaults except seed": {
			From: "", To: "", Resolution: "", Preset: "", Seed: "64326",
			ExpectedPresetCollection: presets.SpectralPresets,
			ExpectedStatusCode:       http.StatusOK,
			ExpectedErrorBody:        "",
		},
		"All set": {
			From: "0,2", To: "4,5", Resolution: "12", Preset: "rawPerlin", Seed: "62346236236",
			ExpectedPresetCollection: presets.LatticePresets,
			ExpectedStatusCode:       http.StatusOK,
			ExpectedErrorBody:        "",
		},
		"Illegal from": {
			From: "52,banana", To: "12", Resolution: "14", Preset: "white", Seed: "162",
			ExpectedPresetCollection: presets.SpectralPresets,
			ExpectedStatusCode:       http.StatusBadRequest,
			ExpectedErrorBody:        `{"error": "Invalid param: (From must be an array of integers)"}`,
		},
		"Float from": {
			From: "52.32", To: "12", Resolution: "14", Preset: "white", Seed: "162",
			ExpectedPresetCollection: presets.SpectralPresets,
			ExpectedStatusCode:       http.StatusBadRequest,
			ExpectedErrorBody:        `{"error": "Invalid param: (From must be an array of integers)"}`,
		},
		"illegal to": {
			From: "15", To: "52,banana", Resolution: "14", Preset: "white", Seed: "56",
			ExpectedPresetCollection: presets.SpectralPresets,
			ExpectedStatusCode:       http.StatusBadRequest,
			ExpectedErrorBody:        `{"error": "Invalid param: (To must be an array of integers)"}`,
		},
		"float to": {
			From: "15", To: "52.6", Resolution: "14", Preset: "white", Seed: "56",
			ExpectedPresetCollection: presets.SpectralPresets,
			ExpectedStatusCode:       http.StatusBadRequest,
			ExpectedErrorBody:        `{"error": "Invalid param: (To must be an array of integers)"}`,
		},
		"from and to diff lengths": {
			From: "15", To: "52,77", Resolution: "14", Preset: "white", Seed: "56",
			ExpectedPresetCollection: presets.SpectralPresets,
			ExpectedStatusCode:       http.StatusBadRequest,
			ExpectedErrorBody:        `{"error": "Invalid param: (From and To must be the same length)"}`,
		},
		"from greater than to": {
			From: "15", To: "7", Resolution: "14", Preset: "white", Seed: "56",
			ExpectedPresetCollection: presets.SpectralPresets,
			ExpectedStatusCode:       http.StatusBadRequest,
			ExpectedErrorBody:        `{"error": "Invalid param: (The value of To must be greater than the value of From in each dimension)"}`,
		},
		"from equals to": {
			From: "7", To: "7", Resolution: "14", Preset: "white", Seed: "56",
			ExpectedPresetCollection: presets.SpectralPresets,
			ExpectedStatusCode:       http.StatusBadRequest,
			ExpectedErrorBody:        `{"error": "Invalid param: (The value of To must be greater than the value of From in each dimension)"}`,
		},
		"illegal resolution": {
			From: "7", To: "11", Resolution: "14x15", Preset: "white", Seed: "56",
			ExpectedPresetCollection: presets.SpectralPresets,
			ExpectedStatusCode:       http.StatusBadRequest,
			ExpectedErrorBody:        `{"error": "Invalid param: (Resolution must be a positive integer)"}`,
		},
		"resolution too small": {
			From: "7", To: "11", Resolution: "0", Preset: "white", Seed: "56",
			ExpectedPresetCollection: presets.SpectralPresets,
			ExpectedStatusCode:       http.StatusBadRequest,
			ExpectedErrorBody:        `{"error": "Invalid param: (Resolution must be a positive integer)"}`,
		},
		"invalid noise function": {
			From: "7", To: "11", Resolution: "5", Preset: "ae1234", Seed: "56",
			ExpectedPresetCollection: presets.SpectralPresets,
			ExpectedStatusCode:       http.StatusBadRequest,
			ExpectedErrorBody:        `{"error": "Invalid param: (NoiseFunction must be a valid preset)"}`,
		},
		"invalid seed": {
			From: "7", To: "11", Resolution: "5", Preset: "white", Seed: "banana",
			ExpectedPresetCollection: presets.SpectralPresets,
			ExpectedStatusCode:       http.StatusBadRequest,
			ExpectedErrorBody:        `{"error": "Invalid param: (Seed must be a positive integer)"}`,
		},
	}

	for name, tc := range testCases {
		url := fmt.Sprintf("/noise?from=%s&to=%s&resolution=%s&noiseFunction=%s&seed=%s", tc.From, tc.To, tc.Resolution, tc.Preset, tc.Seed)

		w := httptest.NewRecorder()
		r, _ := http.NewRequest(http.MethodGet, url, nil)
		tghttp.HandleNoise()(w, r, nil)

		if w.Code != tc.ExpectedStatusCode {
			t.Errorf("%s failed. Expected status code %d, received %d", name, tc.ExpectedStatusCode, w.Code)

			// If we got an error when we should have got a success, print the error message as well so we can debug faster
			if tc.ExpectedStatusCode == http.StatusOK {
				t.Logf("Response: %s", w.Body.String())
			}

			continue
		}

		// Handle expected errors
		if tc.ExpectedErrorBody != "" {
			if w.Body.String() != tc.ExpectedErrorBody {
				t.Errorf("'%s' failed. Expected error response '%s', received '%s'", name, tc.ExpectedErrorBody, w.Body.String())
			}
			continue
		}

		// Handle expected successes
		from := []int{0, 0}
		to := []int{5, 5}
		resolution := 20
		var seed int64
		presetName := "red"
		var preset presets.Preset
		var err error
		var ok bool

		if tc.From != "" {
			from = tghttp.ParseIntArray(tc.From)
		}
		if tc.To != "" {
			to = tghttp.ParseIntArray(tc.To)
		}
		if tc.Resolution != "" {
			if resolution, err = strconv.Atoi(tc.Resolution); err != nil {
				t.Errorf("'%s' failed. Expected a success but passed in an invalid resolution: %s", name, tc.Resolution)
				continue
			}
		}
		if tc.Seed == "" {
			t.Errorf("'%s' failed. Cannot test for success with a random seed (must specify)", name)
			continue
		}
		if seed, err = strconv.ParseInt(tc.Seed, 10, 0); err != nil {
			t.Errorf("'%s' failed. Expected a success but passed in an invalid seed: %s", name, tc.Seed)
			continue
		}
		if tc.Preset != "" {
			presetName = tc.Preset
		}
		if preset, ok = tc.ExpectedPresetCollection[presetName]; !ok {
			t.Errorf("'%s' failed. Expected a success but passed in an invalid preset: %s", name, presetName)
			continue
		}

		noiseFunction := preset(math.NewDefaultSource(seed), []float64{1, 2, 4, 8, 16, 32, 64})
		expectedResponse := model.NewNoise(presetName)
		expectedResponse.Generate(from, to, resolution, noiseFunction)

		responseObject := model.Noise{}
		if err := json.NewDecoder(w.Body).Decode(&responseObject); err != nil {
			t.Errorf("'%s' failed. Failed to decode response: %s", name, w.Body.String())
			continue
		}

		if !responseObject.IsEqual(expectedResponse) {
			t.Errorf("'%s' failed. Expected response '%#v', received '%#v'", name, expectedResponse, responseObject)
		}
	}
}

func TestHandleNoise_Permutations(t *testing.T) {
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

	for name, tc := range testCases {
		w := httptest.NewRecorder()
		handler := tghttp.HandleNoise()

		for _, dimension := range tc.Dimensions {
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

				url := fmt.Sprintf("/noise?from=%s&to=%s&resolution=%s&noiseFunction=%s&seed=42", fromString, toString, resolutionString, tc.PresetName)
				r, _ := http.NewRequest(http.MethodGet, url, nil)
				handler(w, r, nil)

				noiseFunction := tc.PresetCollection[tc.PresetName](math.NewDefaultSource(42), []float64{1, 2, 4, 8, 16, 32, 64})
				expectedResponse := model.NewNoise(tc.PresetName)
				expectedResponse.Generate(params.From, params.To, params.Resolution, noiseFunction)

				responseObject := model.Noise{}
				if w.Code != http.StatusOK {
					t.Errorf("%s failed with param set %d. Expected code %d, received %d", name, i, http.StatusOK, w.Code)
					t.Logf("Response: %s", w.Body.String())
					continue
				}

				if err := json.NewDecoder(w.Body).Decode(&responseObject); err != nil {
					t.Errorf("%s failed with param set %d. Failed to decode response: %s", name, i, w.Body.String())
					continue
				}

				if !responseObject.IsEqual(expectedResponse) {
					if len(responseObject.Values) > 5 {
						responseObject.Values = responseObject.Values[:5]
					}
					if len(expectedResponse.Values) > 5 {
						expectedResponse.Values = expectedResponse.Values[:5]
					}
					t.Errorf("%s failed with param set %d. Expected response (values trimmed to length 5) '%+v', received '%+v'", name, i, expectedResponse, responseObject)
				}
			}
		}
	}
}
