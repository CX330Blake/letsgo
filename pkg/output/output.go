package output

import "github.com/fatih/color"

func Good(format string, a ...interface{}) {
	color.Green(format, a...)
}

func Err(format string, a ...interface{}) {
	color.Red(format, a...)
}

func Warn(format string, a ...interface{}) {
	color.Yellow(format, a...)
}
