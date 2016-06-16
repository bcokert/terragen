package generator_test

import (
	"testing"

	"github.com/bcokert/terragen/generator"
)

func TestRandom(t *testing.T) {
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
		randomFunction1 := generator.Random(testCase.Seed)
		randomFunction2 := generator.Random(testCase.Seed)
		for i, expected := range testCase.ExpectedResults {
			if result := randomFunction1(1); result != expected {
				t.Errorf("%s failed. Expected randomFunciton1 case %d to be %v, received %v", name, i, expected, result)
			}
			if result := randomFunction2(1); result != expected {
				t.Errorf("%s failed. Expected randomFunciton2 case %d to be %v, received %v", name, i, expected, result)
			}
		}
	}
}
