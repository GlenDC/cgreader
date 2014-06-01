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

type ProgramValidation func(output string) bool

func ReportResult(result bool) {
	if result {
		fmt.Println("Program is correct!")
	} else {
		fmt.Println("Program is incorrect!")
	}
}

func RunManualProgram(in string, main ProgramMain) {
	output := main(GetManualInput(in))
	fmt.Println(output)
}

func RunAndValidateManualProgram(in, test string, echo bool, main ProgramMain) {
	output := main(GetManualInput(in))

	if echo {
		fmt.Println(output)
	}

	result := TestOutput(test, output)
	ReportResult(result)
}

func RunAndSelfValidateManualProgram(in string, echo bool, main ProgramMain, validation ProgramValidation) {
	output := main(GetManualInput(in))

	if echo {
		fmt.Println(output)
	}

	result := validation(output)
	ReportResult(result)
}

type FlowProgram interface {
	Update(string)
	GetOutput() string
}

func RunFlowProgram(in string, program FlowProgram) {
	ch, ok := GetFlowInput(in)

	for <-ok {
		program.Update(<-ch)
	}

	output := program.GetOutput()
	fmt.Println(output)
}

func RunAndValidateFlowProgram(in, test string, echo bool, program FlowProgram) {
	ch, ok := GetFlowInput(in)

	for <-ok {
		program.Update(<-ch)
	}

	output := program.GetOutput()
	if echo {
		fmt.Println(output)
	}

	result := TestOutput(test, output)
	ReportResult(result)
}

func RunAndSelfValidateFlowProgram(in string, echo bool, program FlowProgram, validation ProgramValidation) {
	ch, ok := GetFlowInput(in)

	for <-ok {
		program.Update(<-ch)
	}

	output := program.GetOutput()
	if echo {
		fmt.Println(output)
	}

	result := validation(output)
	ReportResult(result)
}

type TargetProgram interface {
	ParseInitialData(<-chan string)
	GetInput() chan string
	Update(<-chan string) string
	SetOutput(string) string
	LoseConditionCheck() bool
	WinConditionCheck() bool
}

func RunTargetProgram(in string, trace bool, program TargetProgram) {
	ch := GetManualInput(in)
	program.ParseInitialData(ch)

	for {
		input := program.GetInput()
		output := program.Update(input)
		result := program.SetOutput(output)

		if trace {
			fmt.Printf("%s\n%s\n\n", output, result)
		}

		if program.WinConditionCheck() {
			ReportResult(true)
			break
		}

		if program.LoseConditionCheck() {
			ReportResult(false)
			break
		}
	}
}

type MapObject interface {
	GetMapCoordinates() string // returns string in format x;y
	GetMapIcon() string        // return 1 character string
}

func DrawMap(width, height int, background string, objects ...MapObject) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			position := fmt.Sprintf("%d;%d", x, y)
			c := background
			for _, object := range objects {
				pos := object.GetMapCoordinates()
				if pos == position {
					c = object.GetMapIcon()
					break
				}
			}
			fmt.Printf("%s  ", c)
		}
		fmt.Println("")
	}
	fmt.Println("")
}
