// Package clog provides a colored logger for Go with multiple log levels,
// file logging with rotation, and progress bar support.
package clog

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/fatih/color"
)

// Color functions for console output.
// Each function returns a colored string using ANSI escape codes.
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

	// Re-export fmt functions for convenience
	Errorf = fmt.Errorf

	Print   = fmt.Print
	Printf  = fmt.Printf
	Println = fmt.Println

	Sprint   = fmt.Sprint
	Sprintf  = fmt.Sprintf
	Sprintln = fmt.Sprintln

	// Flags to enable/disable log levels.
	// By default, Error and Warning are enabled.
	EnableError   bool = true
	EnableWarning bool = true
	EnableInfo    bool = false
	EnableSuccess bool = false
	EnableDebug   bool = false

	// Internal state for file logging
	logFile            *os.File
	logMutex           sync.Mutex
	logFilePath        string
	MaxLogFileSize     int64 = 0  // 0 means no rotation
	LogFileTrimPercent int   = 20 // Default trim 20% when max size reached
)

// SetEnableError enables or disables Error level logging.
func SetEnableError(enable bool) {
	EnableError = enable
}

// SetEnableWarning enables or disables Warning level logging.
func SetEnableWarning(enable bool) {
	EnableWarning = enable
}

// SetEnableInfo enables or disables Info level logging.
func SetEnableInfo(enable bool) {
	EnableInfo = enable
}

// SetEnableSuccess enables or disables Success level logging.
func SetEnableSuccess(enable bool) {
	EnableSuccess = enable
}

// SetEnableDebug enables or disables Debug level logging.
func SetEnableDebug(enable bool) {
	EnableDebug = enable
}

// SetEnableAll enables all log levels.
func SetEnableAll() {
	EnableError = true
	EnableWarning = true
	EnableInfo = true
	EnableSuccess = true
	EnableDebug = true
}

// SetDisableAll disables all log levels.
func SetDisableAll() {
	EnableError = false
	EnableWarning = false
	EnableInfo = false
	EnableSuccess = false
	EnableDebug = false
}

// SetMaxLogFileSize sets the maximum size of the log file in bytes.
// When the file exceeds this size, it will be trimmed.
// Set to 0 to disable rotation (default).
func SetMaxLogFileSize(size int64) {
	MaxLogFileSize = size
}

// SetLogFileTrimPercent sets the percentage of the file to remove when max size is reached.
// Must be between 5 and 50. Default is 20.
func SetLogFileTrimPercent(percent int) {
	if percent < 5 {
		percent = 5
	}
	if percent > 50 {
		percent = 50
	}
	LogFileTrimPercent = percent
}

// SetLogFile sets the log file path.
// If filename is empty, logging to file is disabled.
// Messages are written to both console and file.
func SetLogFile(filename string) error {
	if filename == "" {
		if logFile != nil {
			logFile.Close()
			logFile = nil
		}
		logFilePath = ""
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
	logFilePath = filename
	return nil
}

// trimLogFile trims the log file when it exceeds MaxLogFileSize.
// It removes LogFileTrimPercent from the beginning of the file,
// ensuring to trim at line boundaries for readability.
func trimLogFile() {
	if logFile == nil || logFilePath == "" || MaxLogFileSize <= 0 {
		return
	}

	stat, err := logFile.Stat()
	if err != nil || stat.Size() < MaxLogFileSize {
		return
	}

	content, err := os.ReadFile(logFilePath)
	if err != nil {
		return
	}

	trimSize := int64(float64(len(content)) * float64(LogFileTrimPercent) / 100.0)
	if trimSize < int64(len(content))/2 {
		// Find next newline to avoid cutting in middle of line
		startIdx := trimSize
		for startIdx < int64(len(content)) && content[startIdx] != '\n' {
			startIdx++
		}
		if startIdx < int64(len(content)) {
			startIdx++
		}
		newContent := content[startIdx:]
		err := os.WriteFile(logFilePath, newContent, 0644)
		if err == nil {
			logFile.Seek(0, 2)
		}
	}
}

// formatMessage formats log message arguments.
// If the message starts and ends with brackets, they are removed.
func formatMessage(s ...any) string {
	sStr := fmt.Sprint(s...)
	if len(sStr) >= 2 && sStr[0] == '[' && sStr[len(sStr)-1] == ']' {
		return sStr[1 : len(sStr)-1]
	}
	return sStr
}

// loggerFunc writes the message to both console and file (if configured).
// Thread-safe with mutex.
func loggerFunc(consoleLine, fileLine string) {
	logMutex.Lock()
	defer logMutex.Unlock()
	fmt.Fprint(os.Stdout, consoleLine)
	if logFile != nil {
		fmt.Fprint(logFile, fileLine)
		trimLogFile()
	}
}

// Err logs an error if it is not nil and Error level is enabled.
func Err(e error) {
	if e != nil && EnableError {
		Error(e)
	}
}

// Error logs a message at ERROR level (red color).
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

// Warning logs a message at WARNING level (yellow color).
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

// Info logs a message at INFO level (white color).
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

// Success logs a message at SUCCESS level (green color).
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

// Debug logs a message at DEBUG level (blue color).
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
