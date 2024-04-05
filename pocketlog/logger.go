package pocketlog

import (
	"fmt"
	"io"
	"os"
)

type Logger struct {
	threshold Level
	output    io.Writer
	maxChars  uint
}

// New returns you a logger, ready to log at the required threshold.
func New(threshold Level, opts ...Option) *Logger {
	lgr := &Logger{threshold: threshold, output: os.Stdout}

	for _, configFunc := range opts {
		configFunc(lgr)
	}

	return lgr
}

func (l *Logger) LogF(lvl Level, format string, args ...any) {
	if l.threshold > lvl {
		return
	}

	l.logf(lvl, format, args...)
}

// logf prints the message to the output.
// Add decorations here, if any.
func (l *Logger) logf(lvl Level, format string, args ...any) {

	message := fmt.Sprintf(format, args...)
	//message = limitString(message, l.maxChars)

	if l.maxChars != 0 && uint(len([]rune(message))) > l.maxChars {
		message = string([]rune(message)[:l.maxChars]) + "[TRIMMED]"
	}

	_, _ = fmt.Fprintf(l.output, "%s %s\n", lvl, message)
}

// // Function to limit the string to a certain number of characters
// func limitString(s string, maxLength int) string {
// 	// Check if the string length exceeds the limit
// 	if utf8.RuneCountInString(s) > maxLength {
// 		// Truncate the string to the specified length
// 		runes := []rune(s)
// 		return string(runes[:maxLength])
// 	}
// 	return s
// }
