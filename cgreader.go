package cgreader

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func GetManualInput(in string) <-chan string {
	ch := make(chan string)
	go func() {
		file, err := ioutil.ReadFile(in)
		if err == nil {
			lines := strings.Split(string(file), "\n")
			for _, line := range lines {
				if line != "" {
					ch <- line
				}
			}
		}
		close(ch)
	}()
	return ch
}

func GetFlowInput(in string) (<-chan string, <-chan bool) {
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

func TestOutput(test, out string) bool {
	file, err := ioutil.ReadFile(test)
	if err == nil {
		test = fmt.Sprintf("%s", string(file))
		return out == test
	}
	return false
}

type ProgramMain func(<-chan string) string

func RunAndValidateManualProgram(in string, test string, echo bool, main ProgramMain) {
	output := main(GetManualInput(in))

	if echo {
		fmt.Println(output)
	}

	result := TestOutput(test, output)
	if result {
		fmt.Println("Program is correct!")
	} else {
		fmt.Println("Program is incorrect!")
	}
}

type FlowProgram interface {
	Update(string)
	GetOutput() string
}

func RunAndValidateFlowProgram(in string, test string, echo bool, program FlowProgram) {
	ch, ok := GetFlowInput(in)

	for <-ok {
		program.Update(<-ch)
	}

	output := program.GetOutput()
	if echo {
		fmt.Println(output)
	}

	result := TestOutput(test, output)
	if result {
		fmt.Println("Program is correct!")
	} else {
		fmt.Println("Program is incorrect!")
	}
}
