package cgreader

import (
	"io/ioutil"
	"strings"
)

func RunProgram(input string) (<-chan string, <-chan bool) {
	ch := make(chan string)
	ok := make(chan bool)
	go func() {
		file, err := ioutil.ReadFile(input)
		if err == nil {
			lines := strings.Split(string(file), "\n")
			for _, line := range lines {
				ok <- true
				ch <- line
			}
			ok <- false
		}
		close(ch)
		close(ok)
	}()
	return ch, ok
}
