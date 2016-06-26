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
	"github.com/bcokert/terragen/noise"
	"github.com/bcokert/terragen/presets"
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
			Url:          "/noise?from=0&to=10&resolution=2&noiseFunction=white:1d",
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
	}{
		"violet:1d": {PresetName: "violet:1d"},
		"blue:1d":   {PresetName: "blue:1d"},
		"white:1d":  {PresetName: "white:1d"},
		"pink:1d":   {PresetName: "pink:1d"},
		"red:1d":    {PresetName: "red:1d"},
		"violet:2d": {PresetName: "violet:2d"},
		"blue:2d":   {PresetName: "blue:2d"},
		"white:2d":  {PresetName: "white:2d"},
		"pink:2d":   {PresetName: "pink:2d"},
		"red:2d":    {PresetName: "red:2d"},
	}

	fromCases := map[string][][]float64{
		"1d": [][]float64{[]float64{-3}, []float64{0}, []float64{-12}, []float64{5}},
		"2d": [][]float64{[]float64{-3, 0}, []float64{-1, -1}, []float64{0, 21}, []float64{55, 91}, []float64{-4, 7}},
	}
	toCases := map[string][][]float64{
		"1d": [][]float64{[]float64{-1}, []float64{6}, []float64{55}, []float64{5}},
		"2d": [][]float64{[]float64{0, 5}, []float64{0, 0}, []float64{1, 23}, []float64{55, 999}, []float64{-2, 7}},
	}
	resolutionCases := map[string][]int{
		"1d": []int{4, 50, 1, 2},
		"2d": []int{4, 50, 1, 2, 2},
	}

	for name, testCase := range testCases {
		server := &controller.Server{
			Seed:    42,
			Marshal: json.Marshal,
		}

		dimension := testCase.PresetName[len(testCase.PresetName)-2:]

		for i := range fromCases[dimension] {
			froms := []string{}
			for _, from := range fromCases[dimension][i] {
				froms = append(froms, fmt.Sprintf("%v", from))
			}
			fromString := strings.Join(froms, ",")
			tos := []string{}
			for _, to := range toCases[dimension][i] {
				tos = append(tos, fmt.Sprintf("%v", to))
			}
			toString := strings.Join(tos, ",")
			resolutionString := strconv.Itoa(resolutionCases[dimension][i])
			url := fmt.Sprintf("/noise?from=%s&to=%s&resolution=%s&noiseFunction=%s", fromString, toString, resolutionString, testCase.PresetName)

			var noiseFunction noise.Function
			switch dimension {
			case "1d":
				func1d := presets.Spectral1DPresets[testCase.PresetName](42, []float64{1, 2, 4, 8, 16, 32, 64})
				noiseFunction = func(t []float64) float64 {
					return func1d(t[0])
				}
			case "2d":
				func2d := presets.Spectral2DPresets[testCase.PresetName](42, []float64{1, 2, 4, 8, 16, 32, 64})
				noiseFunction = func(t []float64) float64 {
					return func2d(t[0], t[1])
				}
			default:
				t.Fatalf("Unkown noise function dimension: %s", dimension)
			}

			response := testutils.ExecuteTestRequest(router.CreateDefaultRouter(server), http.MethodGet, url, nil)
			expectedResponse := model.NewNoise(testCase.PresetName)
			expectedResponse.Generate(fromCases[dimension][i], toCases[dimension][i], resolutionCases[dimension][i], noiseFunction)

			responseObject := model.Noise{}
			if response.Code != http.StatusOK {
				t.Errorf("%s failed. Expected code %d, received %d", name, http.StatusOK, response.Code)
				t.Logf("Response: %s", response.Body.String())
				continue
			}

			if err := json.NewDecoder(response.Body).Decode(&responseObject); err != nil {
				t.Errorf("%s failed. Failed to decode response: %s", name, response.Body.String())
				continue
			}

			if !responseObject.IsEqual(expectedResponse) {
				t.Errorf("%s with request '%s' failed. Expected response '%#v', received '%#v'", name, url, expectedResponse, responseObject)
			}
		}
	}
}
