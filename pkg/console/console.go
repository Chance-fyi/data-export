package console

import (
	"github.com/gookit/color"
	"os"
)

func Log(a ...interface{}) {
	color.Note.Print(a...)
}

func Logp(format string, a ...interface{}) {
	color.Note.Prompt(format, a...)
}

func Logln(a ...interface{}) {
	color.Note.Println(a...)
}

func Logf(format string, a ...interface{}) {
	color.Note.Printf(format, a...)
}

func Success(a ...interface{}) {
	color.Success.Print(a...)
}

func Successp(format string, a ...interface{}) {
	color.Success.Prompt(format, a...)
}

func Successln(a ...interface{}) {
	color.Success.Println(a...)
}

func Successf(format string, a ...interface{}) {
	color.Success.Printf(format, a...)
}

func Error(a ...interface{}) {
	color.Errorp(a...)
}

func Errorp(format string, a ...interface{}) {
	color.Error.Prompt(format, a...)
}

func Errorln(a ...interface{}) {
	color.Errorln(a...)
}

func Errorf(format string, a ...interface{}) {
	color.Errorf(format, a...)
}

func Exit(msg string) {
	Errorp(msg)
	os.Exit(1)
}

func ExitIf(err error) {
	if err != nil {
		Exit(err.Error())
	}
}

func D(msg string) {
	Log(msg)
	os.Exit(1)
}
