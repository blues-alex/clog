package clog

import (
	"fmt"
	"os"
	"strings"
)

type LogFuncWrapper struct {
	fn      func(s ...any)
	colorFn func(format string, a ...any) string
	last    string
}

var carriageReturn = "\r"

func (w *LogFuncWrapper) Print(a ...any) {
	w.clearAndPrint(Sprint(a...))
}

func (w *LogFuncWrapper) Printf(format string, a ...any) {
	w.clearAndPrint(fmt.Sprintf(format, a...))
}

func (w *LogFuncWrapper) Println(a ...any) {
	w.clearAndPrint(Sprint(a...))
}

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

var (
	ErrorWrapper   = &LogFuncWrapper{fn: func(s ...any) { Error(s...) }, colorFn: Red}
	WarningWrapper = &LogFuncWrapper{fn: func(s ...any) { Warning(s...) }, colorFn: Yellow}
	InfoWrapper    = &LogFuncWrapper{fn: func(s ...any) { Info(s...) }, colorFn: nil}
	SuccessWrapper = &LogFuncWrapper{fn: func(s ...any) { Success(s...) }, colorFn: Green}
	DebugWrapper   = &LogFuncWrapper{fn: func(s ...any) { Debug(s...) }, colorFn: Cyan}
)

func (w *LogFuncWrapper) Progress(a ...any) {
	w.clearAndPrint(Sprint(a...))
}

func (w *LogFuncWrapper) Progressf(format string, a ...any) {
	s := fmt.Sprintf(format, a...)
	w.clearAndPrint(s)
}

func (w *LogFuncWrapper) Progressln(a ...any) {
	w.clearAndPrint(Sprint(a...))
}
