package clog

import (
	"fmt"
	"os"
)

func ExampleSetEnableAll() {
	SetEnableAll()

	Error("This is an error message")
	Warning("This is a warning message")
	Info("This is an info message")
	Success("This is a success message")
	Debug("This is a debug message")
}

func ExampleError() {
	SetEnableError(true)

	Error("Something went wrong")
}

func ExampleWarning() {
	SetEnableWarning(true)

	Warning("This is a warning")
}

func ExampleInfo() {
	SetEnableInfo(true)

	Info("Information message")
}

func ExampleSuccess() {
	SetEnableSuccess(true)

	Success("Operation completed")
}

func ExampleDebug() {
	SetEnableDebug(true)

	Debug("Debug information")
}

func ExampleErr() {
	SetEnableError(true)

	err := fmt.Errorf("connection refused")
	Err(err)
}

func ExampleSetDisableAll() {
	SetEnableAll()
	SetDisableAll()

	Error("This will not be printed")
	Warning("This will not be printed")
	Info("This will not be printed")
	Success("This will not be printed")
	Debug("This will not be printed")
}

func ExampleSetLogFile() {
	SetEnableAll()

	SetLogFile("example.log")
	defer func() {
		SetLogFile("")
		os.Remove("example.log")
	}()

	Error("Error to file")
	Warning("Warning to file")

	content, _ := os.ReadFile("example.log")
	fmt.Println("File contents:", string(content))
}
