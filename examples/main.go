package main

import (
	"fmt"
	"os"
	"time"

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

	fmt.Println("\n=== Color Print wrappers ===")
	clog.PrintRed("This is red text\n")
	clog.PrintfRed("This is %s text\n", "red")
	clog.PrintlnRed("This is red line")

	clog.PrintHiRed("This is hi-red text\n")
	clog.PrintlnHiRed("This is hi-red line")

	clog.PrintYellow("This is yellow text\n")
	clog.PrintlnYellow("This is yellow line")

	clog.PrintHiYellow("This is hi-yellow text\n")
	clog.PrintlnHiYellow("This is hi-yellow line")

	clog.PrintGreen("This is green text\n")
	clog.PrintlnGreen("This is green line")

	clog.PrintHiGreen("This is hi-green text\n")
	clog.PrintlnHiGreen("This is hi-green line")

	clog.PrintBlue("This is blue text\n")
	clog.PrintlnBlue("This is blue line")

	clog.PrintCyan("This is cyan text\n")
	clog.PrintlnCyan("This is cyan line")

	clog.PrintWhite("This is white text\n")
	clog.PrintlnWhite("This is white line")

	clog.PrintfYellow("New Error: %#v\n", clog.Errorf("New error"))

	fmt.Println("\n=== Color functions (sprint) ===")
	s := clog.SprintfRed("Value: %d\n", 42)
	fmt.Print(s)

	s = clog.SprintfHiYellow("Warning: %s\n", "check this")
	fmt.Print(s)

	s = clog.SprintfGreen("Success: %s\n", "done")
	fmt.Print(s)

	s = clog.SprintfBlue("Info: %d items\n", 10)
	fmt.Print(s)

	fmt.Println("\n=== Progress (\\r) ===")
	clog.SetEnableAll()

	for i := 0; i <= 100; i += 20 {
		clog.DebugWrapper.Progressf("Progress: %d%%", i)
		time.Sleep(50 * time.Millisecond)
	}
	fmt.Println()

	for i := 0; i <= 100; i += 20 {
		clog.WarningWrapper.Progress("Step ", i, "/5")
		time.Sleep(50 * time.Millisecond)
	}
	fmt.Println()

	fmt.Println("\n=== Log rotation ===")
	clog.SetMaxLogFileSize(1024)   // 1KB max
	clog.SetLogFileTrimPercent(30) // Trim 30%
	clog.SetLogFile("rotation.log")
	clog.SetEnableAll()

	for i := 0; i < 50; i++ {
		clog.Info("Log entry number ", i)
	}
	clog.SetLogFile("")

	content, _ = os.ReadFile("rotation.log")
	fmt.Println("Log after rotation (last entries):")
	fmt.Println(string(content))
	os.Remove("rotation.log")
}
