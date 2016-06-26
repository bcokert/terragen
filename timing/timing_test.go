package timing_test

import (
	"regexp"
	"testing"
	"time"

	"github.com/bcokert/terragen/log"
	"github.com/bcokert/terragen/timing"
)

func TestTrack(t *testing.T) {
	testCases := map[string]struct {
		StartTime        time.Time
		ExpectedLogRegex string
	}{
		"12 Ms": {
			StartTime:        time.Now().Add(-12 * 1000000),
			ExpectedLogRegex: `INFO test track took 12\.\d*ms`,
		},
		"2 hours": {
			StartTime:        time.Now().Add(-2 * 3600000000000),
			ExpectedLogRegex: `INFO: test track took 2h0m0.0\d*s`,
		},
	}

	for name, testCase := range testCases {
		log.UseTestLogger()

		timing.Track(testCase.StartTime, "test track")

		output := log.FlushTestLogger()
		if matches, err := regexp.Match(testCase.ExpectedLogRegex, []byte(output)); !matches || err != nil {
			t.Errorf("%s failed. Expected %v, received %v", name, testCase.ExpectedLogRegex, output)
		}
	}
}
