// Package clog provides color Print wrappers.
//
// This file contains Print/Sprint/Printf functions for each color.
// Each color (Red, HiRed, Yellow, HiYellow, Green, HiGreen, Blue, Cyan, White)
// has 6 functions: Print, Printf, Println, Sprint, Sprintf, Sprintln.

package clog

import (
	"fmt"
)

// PrintRed outputs arguments in red color.
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
