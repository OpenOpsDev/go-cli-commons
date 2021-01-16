package logger

import (
	"fmt"

	"github.com/gookit/color"
)

func Info(s string) {
	fmt.Println(color.Blue.Sprint(s))
}

func Warning(s string) {
	fmt.Println(color.Yellow.Sprint(s))
}

func Success(s string) {
	fmt.Println(color.Green.Sprint(s))
}

func Error(s string) {
	fmt.Println(color.Red.Sprint(s))
}

func Fatal(s string) {
	Error(s)
	os.Exit(1)
}
