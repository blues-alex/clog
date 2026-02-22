package main

import (
	"fmt"
	"os"

	"github.com/blues-alex/clog"
)

func main() {
	fmt.Println("=== SetEnableAll ===")
	clog.SetEnableAll()
	clog.Error("This is an error message")
	clog.Warning("This is a warning message")
	clog.Info("This is an info message")
	clog.Success("This is a success message")
	clog.Debug("This is a debug message")

	fmt.Println("\n=== SetDisableAll ===")
	clog.SetDisableAll()
	clog.Error("This will NOT be printed")

	fmt.Println("\n=== Individual enables ===")
	clog.SetDisableAll()
	clog.SetEnableError(true)
	clog.Error("Only Error enabled")

	fmt.Println("\n=== SetLogFile ===")
	clog.SetEnableAll()
	clog.SetLogFile("demo.log")
	clog.Error("Error to file")
	clog.Warning("Warning to file")
	clog.SetLogFile("")

	content, _ := os.ReadFile("demo.log")
	fmt.Println("Log file contents:")
	fmt.Println(string(content))
	os.Remove("demo.log")
}
