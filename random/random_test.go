package random_test

import (
	"testing"

	"github.com/bcokert/terragen/random"
)

func TestSeededNormal(t *testing.T) {
	testCases := map[string]struct {
		Seed            int64
		ExpectedResults [10]float64
	}{
		"simple golden test": {
			Seed: 27,
			ExpectedResults: [10]float64{
				0.011887685710215895,
				0.12999249953770575,
				0.5154010770958206,
				0.2514663236893434,
				0.9583614835153325,
				0.9265658146275855,
				0.1010051413537622,
				0.7275874447787803,
				0.9718979257366778,
				0.6703870049799081,
			},
		},
	}

	for name, testCase := range testCases {
		randomFunction := random.SeededNormal(testCase.Seed)
		for i, expected := range testCase.ExpectedResults {
			if result := randomFunction(); result != expected {
				t.Errorf("%s failed. Expected case %d to be %v, received %v", name, i, expected, result)
			}
		}
	}
}
