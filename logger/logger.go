package logger

import (
	tm "github.com/buger/goterm"
)

func Info(s string) {
	tm.Println(tm.Color(tm.Bold(s), tm.BLUE))
}

func Warning(s string) {
	tm.Println(tm.Color(tm.Bold(s), tm.YELLOW))
}

func Success(s string) {
	tm.Println(tm.Color(tm.Bold(s), tm.GREEN))
}

func Error(s string) {
	tm.Println(tm.Color(tm.Bold(s), tm.RED))
}
