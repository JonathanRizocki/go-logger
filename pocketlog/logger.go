package pocketlog

import (
	"encoding/json"
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

	contents := fmt.Sprintf(format, args...)
	//message = limitString(message, l.maxChars)

	// check the trimming is activated, and that we should apply it to this message
	// checking the length in runes, as this won't print unexpected characters
	if l.maxChars != 0 && uint(len([]rune(contents))) > l.maxChars {
		contents = string([]rune(contents)[:l.maxChars]) + "[TRIMMED]"
	}

	msg := message{
		Level:   lvl.String(),
		Message: contents,
	}

	// encode the message
	formattedMessage, err := json.Marshal(msg)
	if err != nil {
		_, _ = fmt.Fprintf(l.output, "unable to format message for %v\n", contents)
		return
	}

	_, _ = fmt.Fprintln(l.output, string(formattedMessage))
}

// message represents the JSON structure of the logged messages
type message struct {
	Level   string `json:"level"`
	Message string `json:"message"`
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
