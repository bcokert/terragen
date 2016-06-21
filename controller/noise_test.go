package controller_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/bcokert/terragen/controller"
	"github.com/bcokert/terragen/errors"
	"github.com/bcokert/terragen/model"
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
		"unsupported mutlidimension from": {
			Url:          "/noise?from=52,55,62&to=12&resolution=14&noiseFunction=white",
			ExpectedBody: fmt.Sprintf(`{"error": "Noise - Invalid 'from' param: (%s)"}`, errors.UnsupportedError("Multiple dimensions").Error()),
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
		"unsupported mutlidimension to": {
			Url:          "/noise?from=15&to=52,52&resolution=14&noiseFunction=white",
			ExpectedBody: fmt.Sprintf(`{"error": "Noise - Invalid 'to' param: (%s)"}`, errors.UnsupportedError("Multiple dimensions").Error()),
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
		Url              string
		ExpectedResponse model.Noise
		ExpectedCode     int
	}{
		"violet1d simple": {
			Url: "/noise?from=0&to=10&resolution=2&noiseFunction=violet:1d",
			ExpectedResponse: model.Noise{
				RawNoise: map[string][]float64{
					"x": []float64{0, 0.5, 1, 1.5, 2, 2.5, 3, 3.5, 4, 4.5, 5, 5.5, 6, 6.5, 7, 7.5, 8, 8.5, 9, 9.5},
				},
				From:          []float64{0},
				To:            []float64{10},
				Resolution:    2,
				NoiseFunction: "violet:1d",
			},
			ExpectedCode: http.StatusOK,
		},
		"violet1d empty range": {
			Url: "/noise?from=0&to=0&resolution=2&noiseFunction=violet:1d",
			ExpectedResponse: model.Noise{
				RawNoise: map[string][]float64{
					"x": []float64{},
				},
				From:          []float64{0},
				To:            []float64{0},
				Resolution:    2,
				NoiseFunction: "violet:1d",
			},
			ExpectedCode: http.StatusOK,
		},
		"violet1d large resolution": {
			Url: "/noise?from=0&to=1&resolution=50&noiseFunction=violet:1d",
			ExpectedResponse: model.Noise{
				RawNoise: map[string][]float64{
					"x": testutils.FloatSlice(0, 1, 50),
				},
				From:          []float64{0},
				To:            []float64{1},
				Resolution:    50,
				NoiseFunction: "violet:1d",
			},
			ExpectedCode: http.StatusOK,
		},
		"blue1d simple": {
			Url: "/noise?from=0&to=10&resolution=2&noiseFunction=blue:1d",
			ExpectedResponse: model.Noise{
				RawNoise: map[string][]float64{
					"x": []float64{0, 0.5, 1, 1.5, 2, 2.5, 3, 3.5, 4, 4.5, 5, 5.5, 6, 6.5, 7, 7.5, 8, 8.5, 9, 9.5},
				},
				From:          []float64{0},
				To:            []float64{10},
				Resolution:    2,
				NoiseFunction: "blue:1d",
			},
			ExpectedCode: http.StatusOK,
		},
		"blue1d empty range": {
			Url: "/noise?from=0&to=0&resolution=2&noiseFunction=blue:1d",
			ExpectedResponse: model.Noise{
				RawNoise: map[string][]float64{
					"x": []float64{},
				},
				From:          []float64{0},
				To:            []float64{0},
				Resolution:    2,
				NoiseFunction: "blue:1d",
			},
			ExpectedCode: http.StatusOK,
		},
		"blue1d large resolution": {
			Url: "/noise?from=0&to=1&resolution=50&noiseFunction=blue:1d",
			ExpectedResponse: model.Noise{
				RawNoise: map[string][]float64{
					"x": testutils.FloatSlice(0, 1, 50),
				},
				From:          []float64{0},
				To:            []float64{1},
				Resolution:    50,
				NoiseFunction: "blue:1d",
			},
			ExpectedCode: http.StatusOK,
		},
		"white1d simple": {
			Url: "/noise?from=0&to=10&resolution=2&noiseFunction=white:1d",
			ExpectedResponse: model.Noise{
				RawNoise: map[string][]float64{
					"x": []float64{0, 0.5, 1, 1.5, 2, 2.5, 3, 3.5, 4, 4.5, 5, 5.5, 6, 6.5, 7, 7.5, 8, 8.5, 9, 9.5},
				},
				From:          []float64{0},
				To:            []float64{10},
				Resolution:    2,
				NoiseFunction: "white:1d",
			},
			ExpectedCode: http.StatusOK,
		},
		"white1d empty range": {
			Url: "/noise?from=0&to=0&resolution=2&noiseFunction=white:1d",
			ExpectedResponse: model.Noise{
				RawNoise: map[string][]float64{
					"x": []float64{},
				},
				From:          []float64{0},
				To:            []float64{0},
				Resolution:    2,
				NoiseFunction: "white:1d",
			},
			ExpectedCode: http.StatusOK,
		},
		"white1d large resolution": {
			Url: "/noise?from=0&to=1&resolution=50&noiseFunction=white:1d",
			ExpectedResponse: model.Noise{
				RawNoise: map[string][]float64{
					"x": testutils.FloatSlice(0, 1, 50),
				},
				From:          []float64{0},
				To:            []float64{1},
				Resolution:    50,
				NoiseFunction: "white:1d",
			},
			ExpectedCode: http.StatusOK,
		},
		"pink1d simple": {
			Url: "/noise?from=0&to=10&resolution=2&noiseFunction=pink:1d",
			ExpectedResponse: model.Noise{
				RawNoise: map[string][]float64{
					"x": []float64{0, 0.5, 1, 1.5, 2, 2.5, 3, 3.5, 4, 4.5, 5, 5.5, 6, 6.5, 7, 7.5, 8, 8.5, 9, 9.5},
				},
				From:          []float64{0},
				To:            []float64{10},
				Resolution:    2,
				NoiseFunction: "pink:1d",
			},
			ExpectedCode: http.StatusOK,
		},
		"pink1d empty range": {
			Url: "/noise?from=0&to=0&resolution=2&noiseFunction=pink:1d",
			ExpectedResponse: model.Noise{
				RawNoise: map[string][]float64{
					"x": []float64{},
				},
				From:          []float64{0},
				To:            []float64{0},
				Resolution:    2,
				NoiseFunction: "pink:1d",
			},
			ExpectedCode: http.StatusOK,
		},
		"pink1d large resolution": {
			Url: "/noise?from=0&to=1&resolution=50&noiseFunction=pink:1d",
			ExpectedResponse: model.Noise{
				RawNoise: map[string][]float64{
					"x": testutils.FloatSlice(0, 1, 50),
				},
				From:          []float64{0},
				To:            []float64{1},
				Resolution:    50,
				NoiseFunction: "pink:1d",
			},
			ExpectedCode: http.StatusOK,
		},
		"red1d simple": {
			Url: "/noise?from=0&to=10&resolution=2&noiseFunction=red:1d",
			ExpectedResponse: model.Noise{
				RawNoise: map[string][]float64{
					"x": []float64{0, 0.5, 1, 1.5, 2, 2.5, 3, 3.5, 4, 4.5, 5, 5.5, 6, 6.5, 7, 7.5, 8, 8.5, 9, 9.5},
				},
				From:          []float64{0},
				To:            []float64{10},
				Resolution:    2,
				NoiseFunction: "red:1d",
			},
			ExpectedCode: http.StatusOK,
		},
		"red1d empty range": {
			Url: "/noise?from=0&to=0&resolution=2&noiseFunction=red:1d",
			ExpectedResponse: model.Noise{
				RawNoise: map[string][]float64{
					"x": []float64{},
				},
				From:          []float64{0},
				To:            []float64{0},
				Resolution:    2,
				NoiseFunction: "red:1d",
			},
			ExpectedCode: http.StatusOK,
		},
		"red1d large resolution": {
			Url: "/noise?from=0&to=1&resolution=50&noiseFunction=red:1d",
			ExpectedResponse: model.Noise{
				RawNoise: map[string][]float64{
					"x": testutils.FloatSlice(0, 1, 50),
				},
				From:          []float64{0},
				To:            []float64{1},
				Resolution:    50,
				NoiseFunction: "red:1d",
			},
			ExpectedCode: http.StatusOK,
		},
	}

	for name, testCase := range testCases {
		server := &controller.Server{
			Seed:    42,
			Marshal: json.Marshal,
		}
		response := testutils.ExecuteTestRequest(router.CreateDefaultRouter(server), http.MethodGet, testCase.Url, nil)

		expectedNoiseFunction := presets.SpectralPresets[testCase.ExpectedResponse.NoiseFunction](42, []float64{1, 2, 4, 8})
		testCase.ExpectedResponse.RawNoise["value"] = []float64{}
		for _, x := range testCase.ExpectedResponse.RawNoise["x"] {
			testCase.ExpectedResponse.RawNoise["value"] = append(testCase.ExpectedResponse.RawNoise["value"], expectedNoiseFunction(x))
		}

		responseObject := model.Noise{}
		if err := json.NewDecoder(response.Body).Decode(&responseObject); err != nil {
			t.Errorf("%s failed. Failed to decode response: %s", name, response.Body.String())
		}

		if !responseObject.Equals(&testCase.ExpectedResponse) {
			t.Errorf("%s failed. Expected response '%#v', received '%#v'", name, testCase.ExpectedResponse, responseObject)
		}

		if response.Code != testCase.ExpectedCode {
			t.Errorf("%s failed. Expected code %d, received %d", name, testCase.ExpectedCode, response.Code)
		}
	}
}
