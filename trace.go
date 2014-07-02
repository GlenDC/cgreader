package cgreader

import (
	"fmt"
)

func Trace(msg string) {
	Print(msg)
}

func Traceln(msg string) {
	Println(msg)
}

func Tracef(format string, a ...interface{}) {
	Printf("%s", fmt.Sprintf(format, a...))
}
