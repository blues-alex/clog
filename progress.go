package clog

import (
	"fmt"
	"os"
	"strings"
)

// LogFuncWrapper provides methods for logging with progress bar support.
// It wraps a log function with additional methods for updating output in place.
type LogFuncWrapper struct {
	fn      func(s ...any)
	colorFn func(format string, a ...any) string
	last    string
}

var carriageReturn = "\r"

// Print outputs arguments without a newline, using the wrapper's color.
func (w *LogFuncWrapper) Print(a ...any) {
	w.clearAndPrint(Sprint(a...))
}

// Printf outputs a formatted string without a newline, using the wrapper's color.
func (w *LogFuncWrapper) Printf(format string, a ...any) {
	w.clearAndPrint(fmt.Sprintf(format, a...))
}

// Println outputs arguments with a newline, using the wrapper's color.
func (w *LogFuncWrapper) Println(a ...any) {
	w.clearAndPrint(Sprint(a...))
}

// clearAndPrint clears the previous line and prints the new message.
// Uses carriage return and spaces to overwrite previous output.
func (w *LogFuncWrapper) clearAndPrint(msg string) {
	if w.last != "" {
		spaces := strings.Repeat(" ", len(w.last))
		fmt.Fprint(os.Stdout, carriageReturn+spaces+carriageReturn)
	}
	coloredMsg := msg
	if w.colorFn != nil {
		coloredMsg = w.colorFn("%s", msg)
	}
	w.last = msg
	fmt.Fprint(os.Stdout, coloredMsg)
}

// Pre-configured wrappers for each log level.
// Each wrapper uses the same color as the corresponding log function.
var (
	// ErrorWrapper logs at ERROR level (red) and supports progress methods.
	ErrorWrapper = &LogFuncWrapper{fn: func(s ...any) { Error(s...) }, colorFn: Red}
	// WarningWrapper logs at WARNING level (yellow) and supports progress methods.
	WarningWrapper = &LogFuncWrapper{fn: func(s ...any) { Warning(s...) }, colorFn: Yellow}
	// InfoWrapper logs at INFO level (no color) and supports progress methods.
	InfoWrapper = &LogFuncWrapper{fn: func(s ...any) { Info(s...) }, colorFn: nil}
	// SuccessWrapper logs at SUCCESS level (green) and supports progress methods.
	SuccessWrapper = &LogFuncWrapper{fn: func(s ...any) { Success(s...) }, colorFn: Green}
	// DebugWrapper logs at DEBUG level (cyan) and supports progress methods.
	DebugWrapper = &LogFuncWrapper{fn: func(s ...any) { Debug(s...) }, colorFn: Cyan}
)

// Progress outputs arguments without a newline, overwriting the previous output.
// Useful for progress bars - displays only the latest message.
func (w *LogFuncWrapper) Progress(a ...any) {
	w.clearAndPrint(Sprint(a...))
}

// Progressf outputs a formatted string without a newline, overwriting the previous output.
// Useful for progress bars - displays only the latest message.
func (w *LogFuncWrapper) Progressf(format string, a ...any) {
	s := fmt.Sprintf(format, a...)
	w.clearAndPrint(s)
}

// Progressln outputs arguments with a newline, overwriting the previous output.
// Useful for progress bars - displays only the latest message.
func (w *LogFuncWrapper) Progressln(a ...any) {
	w.clearAndPrint(Sprint(a...))
}
