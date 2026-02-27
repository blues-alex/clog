package clog

import "fmt"

type LogFunc func(s ...any)

type LogFuncWrapper struct {
	fn LogFunc
}

func (w LogFuncWrapper) Print(a ...any) {
	w.fn(a...)
}

func (w LogFuncWrapper) Printf(format string, a ...any) {
	w.fn(fmt.Sprintf(format, a...))
}

func (w LogFuncWrapper) Println(a ...any) {
	w.fn(fmt.Sprint(a...))
}

var (
	ErrorWrapper   = LogFuncWrapper{fn: func(s ...any) { Error(s...) }}
	WarningWrapper = LogFuncWrapper{fn: func(s ...any) { Warning(s...) }}
	InfoWrapper    = LogFuncWrapper{fn: func(s ...any) { Info(s...) }}
	SuccessWrapper = LogFuncWrapper{fn: func(s ...any) { Success(s...) }}
	DebugWrapper   = LogFuncWrapper{fn: func(s ...any) { Debug(s...) }}
)

func (w LogFuncWrapper) Progress(a ...any) {
	fmt.Print("\r")
	w.fn(a...)
}

func (w LogFuncWrapper) Progressf(format string, a ...any) {
	fmt.Print("\r")
	w.fn(fmt.Sprintf(format, a...))
}
