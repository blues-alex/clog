package clog

import (
	"fmt"
	"os"
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

	logFile *os.File // Файл для логирования
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
		// Если имя пустое, закрываем текущий файл и возвращаем к стандартному выводу
		if logFile != nil {
			logFile.Close()
			logFile = nil
		}
		return nil
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
	if sStr[0] == '[' && sStr[len(sStr)-1] == ']' {
		return sStr[1 : len(sStr)-1]
	}
	return sStr
}

// Вспомогательная функция для логирования
func loggerFunc(consoleLine, fileLine string) {
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
