package clog

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/fatih/color"
)

var (
	Blue     = color.New(color.FgBlue).SprintfFunc()
	Cyan     = color.New(color.FgCyan).SprintfFunc()
	Green    = color.New(color.FgGreen).SprintfFunc()
	HiGreen  = color.New(color.FgHiGreen).SprintfFunc()
	Yellow   = color.New(color.FgYellow).SprintfFunc()
	HiYellow = color.New(color.FgHiYellow).SprintfFunc()
	Red      = color.New(color.FgRed).SprintfFunc()
	HiRed    = color.New(color.FgHiRed).SprintfFunc()
	White    = color.New(color.FgWhite).SprintfFunc()

	Errorf = fmt.Errorf

	Print   = fmt.Print
	Printf  = fmt.Printf
	Println = fmt.Println

	Sprint   = fmt.Sprint
	Sprintf  = fmt.Sprintf
	Sprintln = fmt.Sprintln

	EnableError   bool = true
	EnableWarning bool = true
	EnableInfo    bool = false
	EnableSuccess bool = false
	EnableDebug   bool = false

	logFile  *os.File
	logMutex sync.Mutex
)

// Сеттеры для уровней логирования (остаются без изменений)
func SetEnableError(enable bool) {
	EnableError = enable
}

func SetEnableWarning(enable bool) {
	EnableWarning = enable
}

func SetEnableInfo(enable bool) {
	EnableInfo = enable
}

func SetEnableSuccess(enable bool) {
	EnableSuccess = enable
}

func SetEnableDebug(enable bool) {
	EnableDebug = enable
}

func SetEnableAll() {
	EnableError = true
	EnableWarning = true
	EnableInfo = true
	EnableSuccess = true
	EnableDebug = true
}

func SetDisableAll() {
	EnableError = false
	EnableWarning = false
	EnableInfo = false
	EnableSuccess = false
	EnableDebug = false
}

// Функция для установки файла логирования
func SetLogFile(filename string) error {
	if filename == "" {
		if logFile != nil {
			logFile.Close()
			logFile = nil
		}
		return nil
	}

	if logFile != nil {
		logFile.Close()
	}

	var err error
	logFile, err = os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	return nil
}

// Вспомогательная функция для форматирования сообщения из ...any
func formatMessage(s ...any) string {
	sStr := fmt.Sprint(s...)
	if len(sStr) >= 2 && sStr[0] == '[' && sStr[len(sStr)-1] == ']' {
		return sStr[1 : len(sStr)-1]
	}
	return sStr
}

// Вспомогательная функция для логирования
func loggerFunc(consoleLine, fileLine string) {
	logMutex.Lock()
	defer logMutex.Unlock()
	fmt.Fprint(os.Stdout, consoleLine)
	if logFile != nil {
		fmt.Fprint(logFile, fileLine)
	}
}

// Основные функции логирования

func Err(e error) {
	if e != nil && EnableError {
		Error(e)
	}
}

func Error(s ...any) {
	if !EnableError {
		return
	}
	mess := formatMessage(s...)

	ts := time.Now().Format("02.01.06 15:04:05.000")

	consoleLine := fmt.Sprintf("[%s] %s %s\n",
		Green(ts),
		Red("ERROR:"),
		HiRed(mess))

	fileLine := fmt.Sprintf("[%s] ERROR: %s\n", ts, mess)

	loggerFunc(consoleLine, fileLine)
}

func Warning(s ...any) {
	if !EnableWarning {
		return
	}
	mess := formatMessage(s...)

	ts := time.Now().Format("02.01.06 15:04:05.000")

	consoleLine := fmt.Sprintf("[%s] %s %s\n",
		Green(ts),
		Yellow("Warning:"),
		HiYellow(mess))

	fileLine := fmt.Sprintf("[%s] Warning: %s\n", ts, mess)

	loggerFunc(consoleLine, fileLine)
}

func Info(s ...any) {
	if !EnableInfo {
		return
	}
	mess := formatMessage(s...)

	ts := time.Now().Format("02.01.06 15:04:05.000")

	consoleLine := fmt.Sprintf("[%s] %s %s\n",
		Green(ts),
		"Info:",
		mess)

	fileLine := fmt.Sprintf("[%s] Info: %s\n", ts, mess)

	loggerFunc(consoleLine, fileLine)
}

func Success(s ...any) {
	if !EnableSuccess {
		return
	}
	mess := formatMessage(s...)

	ts := time.Now().Format("02.01.06 15:04:05.000")

	consoleLine := fmt.Sprintf("[%s] %s %s\n",
		Green(ts),
		Green("Success:"),
		HiGreen(mess))

	fileLine := fmt.Sprintf("[%s] Success: %s\n", ts, mess)

	loggerFunc(consoleLine, fileLine)
}

func Debug(s ...any) {
	if !EnableDebug {
		return
	}
	mess := formatMessage(s...)

	ts := time.Now().Format("02.01.06 15:04:05.000")

	consoleLine := fmt.Sprintf("[%s] %s %s\n",
		Green(ts),
		Blue("Debug:"),
		Cyan(mess))

	fileLine := fmt.Sprintf("[%s] Debug: %s\n", ts, mess)

	loggerFunc(consoleLine, fileLine)
}

// Color wrappers for fmt.Print*

func PrintRed(a ...any) (n int, err error) {
	return fmt.Print(Red(fmt.Sprint(a...)))
}

func PrintfRed(format string, a ...any) (n int, err error) {
	return fmt.Printf(Red(fmt.Sprintf(format, a...)))
}

func PrintlnRed(a ...any) (n int, err error) {
	return fmt.Println(Red(fmt.Sprint(a...)))
}

func SprintRed(a ...any) string {
	return Red(fmt.Sprint(a...))
}

func SprintfRed(format string, a ...any) string {
	return Red(fmt.Sprintf(format, a...))
}

func SprintlnRed(a ...any) string {
	return Red(fmt.Sprintln(a...))
}

func PrintHiRed(a ...any) (n int, err error) {
	return fmt.Print(HiRed(fmt.Sprint(a...)))
}

func PrintfHiRed(format string, a ...any) (n int, err error) {
	return fmt.Printf(HiRed(fmt.Sprintf(format, a...)))
}

func PrintlnHiRed(a ...any) (n int, err error) {
	return fmt.Println(HiRed(fmt.Sprint(a...)))
}

func SprintHiRed(a ...any) string {
	return HiRed(fmt.Sprint(a...))
}

func SprintfHiRed(format string, a ...any) string {
	return HiRed(fmt.Sprintf(format, a...))
}

func SprintlnHiRed(a ...any) string {
	return HiRed(fmt.Sprintln(a...))
}

func PrintYellow(a ...any) (n int, err error) {
	return fmt.Print(Yellow(fmt.Sprint(a...)))
}

func PrintfYellow(format string, a ...any) (n int, err error) {
	return fmt.Printf(Yellow(fmt.Sprintf(format, a...)))
}

func PrintlnYellow(a ...any) (n int, err error) {
	return fmt.Println(Yellow(fmt.Sprint(a...)))
}

func SprintYellow(a ...any) string {
	return Yellow(fmt.Sprint(a...))
}

func SprintfYellow(format string, a ...any) string {
	return Yellow(fmt.Sprintf(format, a...))
}

func SprintlnYellow(a ...any) string {
	return Yellow(fmt.Sprintln(a...))
}

func PrintHiYellow(a ...any) (n int, err error) {
	return fmt.Print(HiYellow(fmt.Sprint(a...)))
}

func PrintfHiYellow(format string, a ...any) (n int, err error) {
	return fmt.Printf(HiYellow(fmt.Sprintf(format, a...)))
}

func PrintlnHiYellow(a ...any) (n int, err error) {
	return fmt.Println(HiYellow(fmt.Sprint(a...)))
}

func SprintHiYellow(a ...any) string {
	return HiYellow(fmt.Sprint(a...))
}

func SprintfHiYellow(format string, a ...any) string {
	return HiYellow(fmt.Sprintf(format, a...))
}

func SprintlnHiYellow(a ...any) string {
	return HiYellow(fmt.Sprintln(a...))
}

func PrintGreen(a ...any) (n int, err error) {
	return fmt.Print(Green(fmt.Sprint(a...)))
}

func PrintfGreen(format string, a ...any) (n int, err error) {
	return fmt.Printf(Green(fmt.Sprintf(format, a...)))
}

func PrintlnGreen(a ...any) (n int, err error) {
	return fmt.Println(Green(fmt.Sprint(a...)))
}

func SprintGreen(a ...any) string {
	return Green(fmt.Sprint(a...))
}

func SprintfGreen(format string, a ...any) string {
	return Green(fmt.Sprintf(format, a...))
}

func SprintlnGreen(a ...any) string {
	return Green(fmt.Sprintln(a...))
}

func PrintHiGreen(a ...any) (n int, err error) {
	return fmt.Print(HiGreen(fmt.Sprint(a...)))
}

func PrintfHiGreen(format string, a ...any) (n int, err error) {
	return fmt.Printf(HiGreen(fmt.Sprintf(format, a...)))
}

func PrintlnHiGreen(a ...any) (n int, err error) {
	return fmt.Println(HiGreen(fmt.Sprint(a...)))
}

func SprintHiGreen(a ...any) string {
	return HiGreen(fmt.Sprint(a...))
}

func SprintfHiGreen(format string, a ...any) string {
	return HiGreen(fmt.Sprintf(format, a...))
}

func SprintlnHiGreen(a ...any) string {
	return HiGreen(fmt.Sprintln(a...))
}

func PrintBlue(a ...any) (n int, err error) {
	return fmt.Print(Blue(fmt.Sprint(a...)))
}

func PrintfBlue(format string, a ...any) (n int, err error) {
	return fmt.Printf(Blue(fmt.Sprintf(format, a...)))
}

func PrintlnBlue(a ...any) (n int, err error) {
	return fmt.Println(Blue(fmt.Sprint(a...)))
}

func SprintBlue(a ...any) string {
	return Blue(fmt.Sprint(a...))
}

func SprintfBlue(format string, a ...any) string {
	return Blue(fmt.Sprintf(format, a...))
}

func SprintlnBlue(a ...any) string {
	return Blue(fmt.Sprintln(a...))
}

func PrintCyan(a ...any) (n int, err error) {
	return fmt.Print(Cyan(fmt.Sprint(a...)))
}

func PrintfCyan(format string, a ...any) (n int, err error) {
	return fmt.Printf(Cyan(fmt.Sprintf(format, a...)))
}

func PrintlnCyan(a ...any) (n int, err error) {
	return fmt.Println(Cyan(fmt.Sprint(a...)))
}

func SprintCyan(a ...any) string {
	return Cyan(fmt.Sprint(a...))
}

func SprintfCyan(format string, a ...any) string {
	return Cyan(fmt.Sprintf(format, a...))
}

func SprintlnCyan(a ...any) string {
	return Cyan(fmt.Sprintln(a...))
}

func PrintWhite(a ...any) (n int, err error) {
	return fmt.Print(White(fmt.Sprint(a...)))
}

func PrintfWhite(format string, a ...any) (n int, err error) {
	return fmt.Printf(White(fmt.Sprintf(format, a...)))
}

func PrintlnWhite(a ...any) (n int, err error) {
	return fmt.Println(White(fmt.Sprint(a...)))
}

func SprintWhite(a ...any) string {
	return White(fmt.Sprint(a...))
}

func SprintfWhite(format string, a ...any) string {
	return White(fmt.Sprintf(format, a...))
}

func SprintlnWhite(a ...any) string {
	return White(fmt.Sprintln(a...))
}
