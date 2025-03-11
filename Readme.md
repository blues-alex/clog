## Overview

### Key Features

- `SetEnableError`, `SetEnableWarning`, `SetEnableInfo`, `SetEnableSuccess`, `SetEnableDebug`: Control the level of logging by enabling or disabling specific log types.
- `SetEnableAll`, `SetDisableAll`: Enable all or disable all logging levels at once.
- `SetLogFile`: Specify a file to store logs in.
- `Err`, `Error`, `Warning`, `Info`, `Success`, `Debug`: Core functions for logging messages, each tailored to different log types and colors.

### Package Structure

The project is organized around one package: `clog`.

### Used Libraries and Frameworks

- `fmt`: For string formatting.
- `os`: To manage file operations.
- `time`: To capture the current timestamp.
- `github.com/fatih/color`: Provides colored output for console logs.

## Code Documentation

### clog.go

#### Variables

```go
var (
    // Color functions to format strings with different colors
    Blue, Cyan, Green, HiGreen, Yellow, HiYellow, Red, HiRed, White func(string) string

    // Boolean flags controlling the logging levels
    EnableError   bool = true
    EnableWarning bool = true
    EnableInfo    bool = false
    EnableSuccess bool = false
    EnableDebug   bool = false

    logFile *os.File // Pointer to the log file
)
```

- `Blue`, `Cyan`, etc.: Functions that colorize console outputs.
- `EnableError`, `EnableWarning`, etc.: Flags indicating whether each type of logging is active or not.
- `logFile`: Reference to the file where logs are written.

#### Setters for Logging Levels

```go
func SetEnableError(enable bool) { EnableError = enable }
func SetEnableWarning(enable bool) { EnableWarning = enable }
func SetEnableInfo(enable bool) { EnableInfo = enable }
func SetEnableSuccess(enable bool) { EnableSuccess = enable }
func SetEnableDebug(enable bool) { EnableDebug = enable }
```

- These functions allow setting the active status of each logging level.

#### Function to set the log file

```go
func SetLogFile(filename string) error {
    if filename == "" {
        // Close existing log file and reset it to standard output if no new file is specified
        if logFile != nil {
            logFile.Close()
            logFile = nil
        }
        return nil
    }
    var err error
    logFile, err = os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        // Return the error if opening the file fails
        return err
    }
    return nil // Successfully set up the log file
}
```

- This function sets a new log file or resets logging to standard output if no filename is provided.

#### Helper Function for Formatting Messages

```go
func formatMessage(s ...any) string {
    // Remove square brackets from the start and end of the message
    return fmt.Sprintf("%v", s...)[1 : len(fmt.Sprint(s...))-1]
}
```

- This helper function cleans up messages by removing any enclosing square brackets, making them more readable.

#### Helper Function for Logging

```go
func loggerFunc(consoleLine, fileLine string) {
    // Print the message to console and optionally to log file if one is set
    fmt.Println(consoleLine)
    if logFile != nil {
        logFile.WriteString(fileLine + "\n")
    }
}
```

- This function logs messages both in the console and, if applicable, to a specified log file.

#### Core Logging Functions

Each of these functions handles logging for different levels with appropriate colors:

```go
func Err(e error) {
    if e != nil && EnableError {
        // If an error occurs and Error level is enabled, call the Error function
        Error(e)
    }
}

func Error(s ...any) {
    if !EnableError { return } // Skip if logging of errors is disabled
    mess := formatMessage(s...)
    ts := time.Now().Format("02.01.06 15:04:05.000")
    // Construct the console and file lines with red text for "ERROR:"
    consoleLine := fmt.Sprintf("[%s] %s %s\n", Green(ts), Red("ERROR:"), HiRed(mess))
    fileLine := fmt.Sprintf("[%s] ERROR: %s\n", ts, mess)
    loggerFunc(consoleLine, fileLine)
}

// Similarly for other levels of logging (Warning, Info, Success, Debug)
```

- These functions ensure that logs are correctly formatted and colored according to their level.

## Usage Examples

### Example of Using Core Logging Functions

This example demonstrates how to use the clog package by enabling all logging levels and directing them to a file named "app.log". It then shows various types of log messages being logged.

```go
package main

import (
    "github.com/blues-alex/clog"
)

func main() {
    // Enable all logging levels
    clog.SetEnableAll()
    // Set the log file to "app.log"
    err := clog.SetLogFile("app.log")
    if err != nil {
        panic(err) // Handle any errors that occur during setting up logging
    }

    // Log different types of messages
    clog.Debug("This is a debug message.")
    clog.Info("This is an informational message.")
    clog.Warning("This is a warning.")
    clog.Error("This is an error.")
    clog.Success("This operation was successful.")
}
```

- This example shows how to configure the clog package for comprehensive logging, including setting up a log file and using various logger functions.

### Example of Using Setters

Here's another way to use the clog package by selectively enabling certain levels of logging and directing them to an "error.log" file.

```go
package main

import (
    "github.com/blues-alex/clog"
)

func main() {
    // Enable only error level logging
    clog.SetEnableError(true)
    clog.SetEnableWarning(false) // Disable warning level logging
    err := clog.SetLogFile("error.log")
    if err != nil {
        panic(err)
    }

    // Log only error messages to the specified file
    clog.Error("This message will be logged in 'error.log'.")
    clog.Warning("This warning won't appear in 'error.log'.") // Won't log due to disabled setting
}
```

- This example illustrates how to customize logging levels and the destination file, focusing on error level messages only.
