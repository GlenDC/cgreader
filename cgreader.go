package cgreader

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func RunProgram(in string) (<-chan string, <-chan bool) {
	ch := make(chan string)
	ok := make(chan bool)
	go func() {
		file, err := ioutil.ReadFile(in)
		if err == nil {
			lines := strings.Split(string(file), "\n")
			for _, line := range lines {
				if line != "" {
					ok <- true
					ch <- line
				}
			}
			ok <- false
		}
		close(ch)
		close(ok)
	}()
	return ch, ok
}

func TestProgram(out, test string) bool {
	output, err := ioutil.ReadFile(out)
	if err == nil {
		out = fmt.Sprintf("%s", string(output))
		return test == out
	}
	return false
}
