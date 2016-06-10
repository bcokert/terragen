package controller_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/bcokert/terragen/controller"
	"github.com/bcokert/terragen/errors"
	"github.com/bcokert/terragen/model"
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
		"white1d simple": {
			Url: "/noise?from=0&to=10&resolution=2&noiseFunction=white:1d",
			ExpectedResponse: model.Noise{
				RawNoise: map[string][]float64{
					"x":     []float64{0, 0.5, 1, 1.5, 2, 2.5, 3, 3.5, 4, 4.5, 5, 5.5, 6, 6.5, 7, 7.5, 8, 8.5, 9, 9.5},
					"value": []float64{0.3014276817564264, 0.35784827213896003, 0.46582953436262, 0.2082178777481407, 0.7359492079171576, 0.18061309800849304, 0.3946204513547035, 0.1984237256140557, 0.6289535023525084, 0.17185905231453527, 0.41797792696401115, 0.2879486980769115, 0.3999150923931667, 0.16147958604346457, 0.25354180360948003, 0.16761227230475298, 0.300315781312509, 0.21839282146232972, 0.5814827564981273, 0.22354643872406454},
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
					"x":     []float64{},
					"value": []float64{},
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
					"x":     testutils.FloatSlice(0, 1, 50),
					"value": []float64{0.3014276817564264, 0.678179234683142, 0.8204902749746891, 0.48291649010815363, 0.2882752272436985, 0.2894118401109828, 0.28677910572593346, 0.5526153639287348, 0.16609337039480054, -0.08603802431961347, -0.1732427064499742, -0.01624342535907583, -0.029985056954251077, 0.2805788365160484, 0.4699346333572204, 0.06557995870505971, 0.004470850695027145, -0.27227818630911216, -0.20497543035354893, -0.08834726146462168, -0.11055492965091696, -0.12544697775371016, -0.410389603022137, -0.5475388285634145, -0.015206326797833403, 0.32337533701770915, 0.47146601758199613, 0.37813534460786524, 0.1215104258239379, -0.006930447659418609, -0.14049749761827368, 0.0034414475543465273, -0.06390663665592747, -0.06842299055356799, -0.2378556629974401, -0.624342486889552, -0.5460878129153299, -0.04576759240490117, -0.018797368547241816, 0.10365145571522502, -0.11988235991019325, -0.3329207810440134, -0.41267191662202324, -0.35227493973581486, -0.24775122704726776, -0.16307633284587592, -0.3097352819774509, -0.5322705959799938, -0.4905921769848842, 0.014290153691196239},
				},
				From:          []float64{0},
				To:            []float64{1},
				Resolution:    50,
				NoiseFunction: "white:1d",
			},
			ExpectedCode: http.StatusOK,
		},
		"red1d simple": {
			Url: "/noise?from=0&to=10&resolution=2&noiseFunction=red:1d",
			ExpectedResponse: model.Noise{
				RawNoise: map[string][]float64{
					"x":     []float64{0, 0.5, 1, 1.5, 2, 2.5, 3, 3.5, 4, 4.5, 5, 5.5, 6, 6.5, 7, 7.5, 8, 8.5, 9, 9.5},
					"value": []float64{0.10491637639740707, 0.02522960371433485, 0.14240998620874543, -0.06951808844985238, 0.22028277620639095, -0.0036129201768815805, 0.16743035692290176, -0.08762397550026342, 0.25150755724035145, -0.1495588243161205, 0.21620418268684216, -0.11448528972805566, 0.18006882300120994, -0.1482459144529171, 0.12864632754084612, -0.15321647422470572, 0.040347426422610716, -0.14143833187807403, 0.16947245271105693, -0.09200993524802667},
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
					"x":     []float64{},
					"value": []float64{},
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
					"x":     testutils.FloatSlice(0, 1, 50),
					"value": []float64{0.10491637639740707, 0.09817856900932152, 0.22368158928951132, 0.22788808530931254, 0.2953241106825218, 0.23434663149763404, 0.29742780078138165, 0.2936903092842745, 0.2339142999144696, 0.2046673252326682, 0.20375269561653472, 0.2033993886280379, 0.17937033255408175, 0.09881476227942661, 0.17681578863666733, 0.038697001985125774, 0.17217199766771107, -0.02432421302904917, 0.05070744594905816, 0.008152800186307083, -0.09734163459175989, -0.015477935982265663, 0.02478857698209882, -0.1692852386543585, -0.013525596591508756, 0.054348503542155834, -0.14154570749415085, -0.1353363868575534, -0.06455723176743061, -0.12187174843847698, -0.12658962924570688, -0.16708784662327408, -0.19061436810581003, -0.19186499348563668, -0.21609116591051764, -0.2522841912003114, -0.283098315391673, -0.17653847765489233, -0.2665050974334386, -0.2423171235862273, -0.2120397259104831, -0.10699513608109999, -0.15335652541548386, -0.13164116902491907, -0.14829813101641812, -0.03985039577885589, -0.05839776352155067, -0.12066971734324261, -0.040775087365306044, 0.1878225546212253},
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
