package utils

import (
	"fmt"
	"github.com/gookit/color"
)

type Log interface {
	Success(format string, a ...interface{})
	Error(format string, a ...interface{})
	Warning(format string, a ...interface{})
	Info(format string, a ...interface{})
	DD(format string, a ...interface{})
}

type log struct {
}

func NewLog() Log {
	return &log{}
}

func (l *log) Success(format string, a ...interface{}) {
	fmt.Println(color.Green.Sprintf(format, a...))
}

func (l *log) Error(format string, a ...interface{}) {
	fmt.Println(color.Red.Sprintf(format, a...))
}

func (l *log) Warning(format string, a ...interface{}) {
	fmt.Println(color.Yellow.Sprintf(format, a...))
}

func (l *log) Info(format string, a ...interface{}) {
	fmt.Println(color.Blue.Sprintf(format, a...))
}

func (l *log) DD(format string, a ...interface{}) {
	fmt.Println(color.FgGray.Sprintf(format, a...))
}
