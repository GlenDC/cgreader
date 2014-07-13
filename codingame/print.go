package codingame

import (
	"fmt"
)

type PrintfCallback func(format string, a ...interface{})

var Printf PrintfCallback = func(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

type PrintCallback func(text string)

var Print PrintCallback = func(text string) {
	Printf("%s", text)
}

type PrintlnCallback func(text string)

var Println PrintCallback = func(text string) {
	Printf("%s\n", text)
}

func SetPrintfCallback(callback PrintfCallback) {
	Printf = callback
}
