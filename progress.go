package clog

import (
	"fmt"
	"os"
	"strings"
)

type LogFunc func(s ...any)

type LogFuncWrapper struct {
	fn   LogFunc
	last string
}

var clearLine = "\033[2K"
var carriageReturn = "\r"

func (w *LogFuncWrapper) Print(a ...any) {
	w.clearAndPrint(Sprint(a...))
}

func (w *LogFuncWrapper) Printf(format string, a ...any) {
	w.clearAndPrint(Sprintf(format, a...))
}

func (w *LogFuncWrapper) Println(a ...any) {
	w.clearAndPrint(Sprint(a...))
}

func (w *LogFuncWrapper) clearAndPrint(msg string) {
	if w.last != "" {
		spaces := strings.Repeat(" ", len(w.last))
		fmt.Fprint(os.Stdout, carriageReturn+spaces+carriageReturn)
	}
	w.last = msg
	fmt.Fprint(os.Stdout, msg)
}

var (
	ErrorWrapper   = &LogFuncWrapper{fn: func(s ...any) { Error(s...) }}
	WarningWrapper = &LogFuncWrapper{fn: func(s ...any) { Warning(s...) }}
	InfoWrapper    = &LogFuncWrapper{fn: func(s ...any) { Info(s...) }}
	SuccessWrapper = &LogFuncWrapper{fn: func(s ...any) { Success(s...) }}
	DebugWrapper   = &LogFuncWrapper{fn: func(s ...any) { Debug(s...) }}
)

func (w *LogFuncWrapper) Progress(a ...any) {
	w.clearAndPrint(Sprint(a...))
}

func (w *LogFuncWrapper) Progressf(format string, a ...any) {
	w.clearAndPrint(Sprintf(format, a...))
}

func (w *LogFuncWrapper) Progressln(a ...any) {
	w.clearAndPrint(Sprint(a...))
}
