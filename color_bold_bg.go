// Package clog provides additional color variants (Bold and Background).
//
// This file contains Print/Sprint/Printf functions for Bold and Background colors.

package clog

import (
	"fmt"
)

// BoldBlack

func PrintBoldBlack(a ...any) (n int, err error) {
	return fmt.Print(BoldBlack(fmt.Sprint(a...)))
}

func PrintfBoldBlack(format string, a ...any) (n int, err error) {
	return fmt.Printf(BoldBlack(fmt.Sprintf(format, a...)))
}

func PrintlnBoldBlack(a ...any) (n int, err error) {
	return fmt.Println(BoldBlack(fmt.Sprint(a...)))
}

func SprintBoldBlack(a ...any) string {
	return BoldBlack(fmt.Sprint(a...))
}

func SprintfBoldBlack(format string, a ...any) string {
	return BoldBlack(fmt.Sprintf(format, a...))
}

func SprintlnBoldBlack(a ...any) string {
	return BoldBlack(fmt.Sprintln(a...))
}

// BoldBlue

func PrintBoldBlue(a ...any) (n int, err error) {
	return fmt.Print(BoldBlue(fmt.Sprint(a...)))
}

func PrintfBoldBlue(format string, a ...any) (n int, err error) {
	return fmt.Printf(BoldBlue(fmt.Sprintf(format, a...)))
}

func PrintlnBoldBlue(a ...any) (n int, err error) {
	return fmt.Println(BoldBlue(fmt.Sprint(a...)))
}

func SprintBoldBlue(a ...any) string {
	return BoldBlue(fmt.Sprint(a...))
}

func SprintfBoldBlue(format string, a ...any) string {
	return BoldBlue(fmt.Sprintf(format, a...))
}

func SprintlnBoldBlue(a ...any) string {
	return BoldBlue(fmt.Sprintln(a...))
}

// BoldCyan

func PrintBoldCyan(a ...any) (n int, err error) {
	return fmt.Print(BoldCyan(fmt.Sprint(a...)))
}

func PrintfBoldCyan(format string, a ...any) (n int, err error) {
	return fmt.Printf(BoldCyan(fmt.Sprintf(format, a...)))
}

func PrintlnBoldCyan(a ...any) (n int, err error) {
	return fmt.Println(BoldCyan(fmt.Sprint(a...)))
}

func SprintBoldCyan(a ...any) string {
	return BoldCyan(fmt.Sprint(a...))
}

func SprintfBoldCyan(format string, a ...any) string {
	return BoldCyan(fmt.Sprintf(format, a...))
}

func SprintlnBoldCyan(a ...any) string {
	return BoldCyan(fmt.Sprintln(a...))
}

// BoldGreen

func PrintBoldGreen(a ...any) (n int, err error) {
	return fmt.Print(BoldGreen(fmt.Sprint(a...)))
}

func PrintfBoldGreen(format string, a ...any) (n int, err error) {
	return fmt.Printf(BoldGreen(fmt.Sprintf(format, a...)))
}

func PrintlnBoldGreen(a ...any) (n int, err error) {
	return fmt.Println(BoldGreen(fmt.Sprint(a...)))
}

func SprintBoldGreen(a ...any) string {
	return BoldGreen(fmt.Sprint(a...))
}

func SprintfBoldGreen(format string, a ...any) string {
	return BoldGreen(fmt.Sprintf(format, a...))
}

func SprintlnBoldGreen(a ...any) string {
	return BoldGreen(fmt.Sprintln(a...))
}

// BoldMagenta

func PrintBoldMagenta(a ...any) (n int, err error) {
	return fmt.Print(BoldMagenta(fmt.Sprint(a...)))
}

func PrintfBoldMagenta(format string, a ...any) (n int, err error) {
	return fmt.Printf(BoldMagenta(fmt.Sprintf(format, a...)))
}

func PrintlnBoldMagenta(a ...any) (n int, err error) {
	return fmt.Println(BoldMagenta(fmt.Sprint(a...)))
}

func SprintBoldMagenta(a ...any) string {
	return BoldMagenta(fmt.Sprint(a...))
}

func SprintfBoldMagenta(format string, a ...any) string {
	return BoldMagenta(fmt.Sprintf(format, a...))
}

func SprintlnBoldMagenta(a ...any) string {
	return BoldMagenta(fmt.Sprintln(a...))
}

// BoldRed

func PrintBoldRed(a ...any) (n int, err error) {
	return fmt.Print(BoldRed(fmt.Sprint(a...)))
}

func PrintfBoldRed(format string, a ...any) (n int, err error) {
	return fmt.Printf(BoldRed(fmt.Sprintf(format, a...)))
}

func PrintlnBoldRed(a ...any) (n int, err error) {
	return fmt.Println(BoldRed(fmt.Sprint(a...)))
}

func SprintBoldRed(a ...any) string {
	return BoldRed(fmt.Sprint(a...))
}

func SprintfBoldRed(format string, a ...any) string {
	return BoldRed(fmt.Sprintf(format, a...))
}

func SprintlnBoldRed(a ...any) string {
	return BoldRed(fmt.Sprintln(a...))
}

// BoldWhite

func PrintBoldWhite(a ...any) (n int, err error) {
	return fmt.Print(BoldWhite(fmt.Sprint(a...)))
}

func PrintfBoldWhite(format string, a ...any) (n int, err error) {
	return fmt.Printf(BoldWhite(fmt.Sprintf(format, a...)))
}

func PrintlnBoldWhite(a ...any) (n int, err error) {
	return fmt.Println(BoldWhite(fmt.Sprint(a...)))
}

func SprintBoldWhite(a ...any) string {
	return BoldWhite(fmt.Sprint(a...))
}

func SprintfBoldWhite(format string, a ...any) string {
	return BoldWhite(fmt.Sprintf(format, a...))
}

func SprintlnBoldWhite(a ...any) string {
	return BoldWhite(fmt.Sprintln(a...))
}

// BoldYellow

func PrintBoldYellow(a ...any) (n int, err error) {
	return fmt.Print(BoldYellow(fmt.Sprint(a...)))
}

func PrintfBoldYellow(format string, a ...any) (n int, err error) {
	return fmt.Printf(BoldYellow(fmt.Sprintf(format, a...)))
}

func PrintlnBoldYellow(a ...any) (n int, err error) {
	return fmt.Println(BoldYellow(fmt.Sprint(a...)))
}

func SprintBoldYellow(a ...any) string {
	return BoldYellow(fmt.Sprint(a...))
}

func SprintfBoldYellow(format string, a ...any) string {
	return BoldYellow(fmt.Sprintf(format, a...))
}

func SprintlnBoldYellow(a ...any) string {
	return BoldYellow(fmt.Sprintln(a...))
}

// BgBlack

func PrintBgBlack(a ...any) (n int, err error) {
	return fmt.Print(BgBlack(fmt.Sprint(a...)))
}

func PrintfBgBlack(format string, a ...any) (n int, err error) {
	return fmt.Printf(BgBlack(fmt.Sprintf(format, a...)))
}

func PrintlnBgBlack(a ...any) (n int, err error) {
	return fmt.Println(BgBlack(fmt.Sprint(a...)))
}

func SprintBgBlack(a ...any) string {
	return BgBlack(fmt.Sprint(a...))
}

func SprintfBgBlack(format string, a ...any) string {
	return BgBlack(fmt.Sprintf(format, a...))
}

func SprintlnBgBlack(a ...any) string {
	return BgBlack(fmt.Sprintln(a...))
}

// BgBlue

func PrintBgBlue(a ...any) (n int, err error) {
	return fmt.Print(BgBlue(fmt.Sprint(a...)))
}

func PrintfBgBlue(format string, a ...any) (n int, err error) {
	return fmt.Printf(BgBlue(fmt.Sprintf(format, a...)))
}

func PrintlnBgBlue(a ...any) (n int, err error) {
	return fmt.Println(BgBlue(fmt.Sprint(a...)))
}

func SprintBgBlue(a ...any) string {
	return BgBlue(fmt.Sprint(a...))
}

func SprintfBgBlue(format string, a ...any) string {
	return BgBlue(fmt.Sprintf(format, a...))
}

func SprintlnBgBlue(a ...any) string {
	return BgBlue(fmt.Sprintln(a...))
}

// BgCyan

func PrintBgCyan(a ...any) (n int, err error) {
	return fmt.Print(BgCyan(fmt.Sprint(a...)))
}

func PrintfBgCyan(format string, a ...any) (n int, err error) {
	return fmt.Printf(BgCyan(fmt.Sprintf(format, a...)))
}

func PrintlnBgCyan(a ...any) (n int, err error) {
	return fmt.Println(BgCyan(fmt.Sprint(a...)))
}

func SprintBgCyan(a ...any) string {
	return BgCyan(fmt.Sprint(a...))
}

func SprintfBgCyan(format string, a ...any) string {
	return BgCyan(fmt.Sprintf(format, a...))
}

func SprintlnBgCyan(a ...any) string {
	return BgCyan(fmt.Sprintln(a...))
}

// BgGreen

func PrintBgGreen(a ...any) (n int, err error) {
	return fmt.Print(BgGreen(fmt.Sprint(a...)))
}

func PrintfBgGreen(format string, a ...any) (n int, err error) {
	return fmt.Printf(BgGreen(fmt.Sprintf(format, a...)))
}

func PrintlnBgGreen(a ...any) (n int, err error) {
	return fmt.Println(BgGreen(fmt.Sprint(a...)))
}

func SprintBgGreen(a ...any) string {
	return BgGreen(fmt.Sprint(a...))
}

func SprintfBgGreen(format string, a ...any) string {
	return BgGreen(fmt.Sprintf(format, a...))
}

func SprintlnBgGreen(a ...any) string {
	return BgGreen(fmt.Sprintln(a...))
}

// BgMagenta

func PrintBgMagenta(a ...any) (n int, err error) {
	return fmt.Print(BgMagenta(fmt.Sprint(a...)))
}

func PrintfBgMagenta(format string, a ...any) (n int, err error) {
	return fmt.Printf(BgMagenta(fmt.Sprintf(format, a...)))
}

func PrintlnBgMagenta(a ...any) (n int, err error) {
	return fmt.Println(BgMagenta(fmt.Sprint(a...)))
}

func SprintBgMagenta(a ...any) string {
	return BgMagenta(fmt.Sprint(a...))
}

func SprintfBgMagenta(format string, a ...any) string {
	return BgMagenta(fmt.Sprintf(format, a...))
}

func SprintlnBgMagenta(a ...any) string {
	return BgMagenta(fmt.Sprintln(a...))
}

// BgRed

func PrintBgRed(a ...any) (n int, err error) {
	return fmt.Print(BgRed(fmt.Sprint(a...)))
}

func PrintfBgRed(format string, a ...any) (n int, err error) {
	return fmt.Printf(BgRed(fmt.Sprintf(format, a...)))
}

func PrintlnBgRed(a ...any) (n int, err error) {
	return fmt.Println(BgRed(fmt.Sprint(a...)))
}

func SprintBgRed(a ...any) string {
	return BgRed(fmt.Sprint(a...))
}

func SprintfBgRed(format string, a ...any) string {
	return BgRed(fmt.Sprintf(format, a...))
}

func SprintlnBgRed(a ...any) string {
	return BgRed(fmt.Sprintln(a...))
}

// BgWhite

func PrintBgWhite(a ...any) (n int, err error) {
	return fmt.Print(BgWhite(fmt.Sprint(a...)))
}

func PrintfBgWhite(format string, a ...any) (n int, err error) {
	return fmt.Printf(BgWhite(fmt.Sprintf(format, a...)))
}

func PrintlnBgWhite(a ...any) (n int, err error) {
	return fmt.Println(BgWhite(fmt.Sprint(a...)))
}

func SprintBgWhite(a ...any) string {
	return BgWhite(fmt.Sprint(a...))
}

func SprintfBgWhite(format string, a ...any) string {
	return BgWhite(fmt.Sprintf(format, a...))
}

func SprintlnBgWhite(a ...any) string {
	return BgWhite(fmt.Sprintln(a...))
}

// BgYellow

func PrintBgYellow(a ...any) (n int, err error) {
	return fmt.Print(BgYellow(fmt.Sprint(a...)))
}

func PrintfBgYellow(format string, a ...any) (n int, err error) {
	return fmt.Printf(BgYellow(fmt.Sprintf(format, a...)))
}

func PrintlnBgYellow(a ...any) (n int, err error) {
	return fmt.Println(BgYellow(fmt.Sprint(a...)))
}

func SprintBgYellow(a ...any) string {
	return BgYellow(fmt.Sprint(a...))
}

func SprintfBgYellow(format string, a ...any) string {
	return BgYellow(fmt.Sprintf(format, a...))
}

func SprintlnBgYellow(a ...any) string {
	return BgYellow(fmt.Sprintln(a...))
}
