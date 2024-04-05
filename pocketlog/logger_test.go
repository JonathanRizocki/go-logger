package pocketlog_test

import (
	"learngo-pockets/logger/pocketlog"
	"testing"
)

func ExampleLogger_Debugf() {
	debugLogger := pocketlog.New(pocketlog.LevelDebug)
	debugLogger.LogF(pocketlog.LevelDebug, "Hello, %s", "world")
	// Output: [DEBUG] Hello, world
}

const (
	debugMessage = "Why write I still all one, ever the same"
	infoMessage  = "And keep invention in a noted weed,"
	errorMessage = "That every word doth almost tell my name,"
)

const (
	debugLevelMessage = "[DEBUG] " + debugMessage
	infoLevelMessage  = "[INFO] " + infoMessage
	errorLevelMessage = "[ERROR] " + errorMessage
)

func TestLogger_DebugInfoErrorf(t *testing.T) {
	type testCase struct {
		level    pocketlog.Level
		expected string
	}

	tt := map[string]testCase{
		"debug": {
			level:    pocketlog.LevelDebug,
			expected: debugLevelMessage + "\n" + infoLevelMessage + "\n" + errorLevelMessage + "\n",
		},
		"info": {
			level:    pocketlog.LevelInfo,
			expected: infoLevelMessage + "\n" + errorLevelMessage + "\n",
		},
		"error": {
			level:    pocketlog.LevelError,
			expected: errorLevelMessage + "\n",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			tw := &testWriter{}

			testedLogger := pocketlog.New(tc.level, pocketlog.WithOutput(tw))

			testedLogger.LogF(pocketlog.LevelDebug, debugMessage)
			testedLogger.LogF(pocketlog.LevelInfo, infoMessage)
			testedLogger.LogF(pocketlog.LevelError, errorMessage)

			if tw.contents != tc.expected {
				t.Errorf("Invalid contents, expected \n%q, got \n%q", tc.expected, tw.contents)
			}
		})
	}
}

// testWriter is a struct that implements io.Writer.
// We use it to validate that we can write to a specific output.
type testWriter struct {
	contents string
}

// Write implements the io.Writer interface.
func (tw *testWriter) Write(p []byte) (n int, err error) {
	tw.contents = tw.contents + string(p)
	return len(p), nil
}
