package clog

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetters(t *testing.T) {
	tests := []struct {
		name   string
		setter func(bool)
		flag   *bool
	}{
		{"EnableError", SetEnableError, &EnableError},
		{"EnableWarning", SetEnableWarning, &EnableWarning},
		{"EnableInfo", SetEnableInfo, &EnableInfo},
		{"EnableSuccess", SetEnableSuccess, &EnableSuccess},
		{"EnableDebug", SetEnableDebug, &EnableDebug},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prevVal := *tt.flag
			defer func() { *tt.flag = prevVal }()
			newVal := !prevVal

			tt.setter(newVal)
			assert.Equal(t, newVal, *tt.flag, "Флаг не установлен")
		})
	}
}

func TestSetLogFile(t *testing.T) {
	tests := []struct {
		name        string
		filename    string
		expectError bool
	}{
		{"ValidFile", "test.log", false},
		{"EmptyFilename", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := SetLogFile(tt.filename)
			if tt.expectError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)

			if tt.filename == "" {
				assert.Nil(t, logFile)
			} else {
				assert.NotNil(t, logFile)
				defer func() { _ = os.Remove(tt.filename) }()
			}

			// Восстанавливаем состояние
			SetLogFile("")
		})
	}
}

func TestFormatMessage(t *testing.T) {
	tests := []struct {
		input    any
		expected string
	}{
		{[]string{"a", "b"}, "a b"},
		{"test", "test"},
		{123, "123"},
		{nil, "<nil>"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.input), func(t *testing.T) {
			actual := formatMessage(tt.input)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestErrorOutput(t *testing.T) {
	testCases := []struct {
		name         string
		enable       bool
		expected     string
		expectedFile string
	}{
		{"Enabled", true, "\x1b[31mERROR:\x1b[0m \x1b[91mTest error message\x1b[0m\n", "ERROR: Test error message"},
		{"Disabled", false, "", ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			prevEnable := EnableError
			defer func() { SetEnableError(prevEnable) }()
			SetEnableError(tc.enable)

			// Перехват stdout
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			// Тестовый файл для лога
			tmpFile, err := os.CreateTemp("", "test_log_*.log")
			assert.NoError(t, err)
			defer func() {
				SetLogFile("")
				os.Remove(tmpFile.Name())
			}()
			err = SetLogFile(tmpFile.Name())
			assert.NoError(t, err)

			// Вызываем функцию
			Error("Test error message")

			w.Close()
			buf := new(bytes.Buffer)
			_, _ = buf.ReadFrom(r)
			os.Stdout = oldStdout

			if tc.expected != "" {
				assert.True(t, buf.String()[33:] == tc.expected, "Unexpected console output: %s != %s", buf.String(), tc.expected)

				fileContent, err := os.ReadFile(tmpFile.Name())
				assert.NoError(t, err)
				assert.Contains(t, string(fileContent), tc.expectedFile)
			} else {
				assert.Empty(t, buf.String(), "Output should be empty when disabled")
				assert.FileExists(t, tmpFile.Name()) // Файл создан, но запись не происходит
			}
		})
	}
}

func TestWarningOutput(t *testing.T) {
	testCases := []struct {
		name         string
		enable       bool
		expected     string
		expectedFile string
	}{
		{"Enabled", true, "\x1b[33mWarning:\x1b[0m \x1b[93mTest warning message\x1b[0m\n", "Warning: Test warning message"},
		{"Disabled", false, "", ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			prevEnable := EnableWarning
			defer func() { SetEnableWarning(prevEnable) }()
			SetEnableWarning(tc.enable)

			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			tmpFile, err := os.CreateTemp("", "test_log_*.log")
			assert.NoError(t, err)
			defer func() {
				SetLogFile("")
				os.Remove(tmpFile.Name())
			}()
			err = SetLogFile(tmpFile.Name())
			assert.NoError(t, err)

			Warning("Test warning message")

			w.Close()
			buf := new(bytes.Buffer)
			_, _ = buf.ReadFrom(r)
			os.Stdout = oldStdout

			if tc.expected != "" {
				assert.True(t, buf.String()[33:] == tc.expected, "Unexpected console output: %s", buf.String())

				fileContent, err := os.ReadFile(tmpFile.Name())
				assert.NoError(t, err)
				assert.Contains(t, string(fileContent), tc.expectedFile)
			} else {
				assert.Empty(t, buf.String(), "Output should be empty when disabled")
			}
		})
	}
}
