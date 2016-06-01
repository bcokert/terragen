package log_test

import (
	"testing"

	"github.com/bcokert/terragen/log"
)

func TestInfo(t *testing.T) {
	testCases := map[string]struct {
		Message  string
		Args     []interface{}
		Expected string
	}{
		"Simple Log": {
			Message:  "Hello World",
			Args:     []interface{}{},
			Expected: "I: Hello World",
		},
		"Format Log": {
			Message:  "Hello %s",
			Args:     []interface{}{"Wallalla"},
			Expected: "I: Hello Wallalla",
		},
	}

	log.UseTestLogger()

	for name, testCase := range testCases {
		log.Info(testCase.Message, testCase.Args...)
		output := log.FlushTestLogger()
		if output != (testCase.Expected + "\n") {
			t.Errorf("%s failed. Expected %v, received %v", name, testCase.Expected, output)
		}
	}
}

func TestError(t *testing.T) {
	testCases := map[string]struct {
		Message  string
		Args     []interface{}
		Expected string
	}{
		"Simple Log": {
			Message:  "Hello World",
			Args:     []interface{}{},
			Expected: "ERROR: Hello World",
		},
		"Format Log": {
			Message:  "Hello %s",
			Args:     []interface{}{"Wallalla"},
			Expected: "ERROR: Hello Wallalla",
		},
	}

	log.UseTestLogger()

	for name, testCase := range testCases {
		log.Error(testCase.Message, testCase.Args...)
		output := log.FlushTestLogger()
		if output != (testCase.Expected + "\n") {
			t.Errorf("%s failed. Expected %v, received %v", name, testCase.Expected, output)
		}
	}
}
