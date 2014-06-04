package cgreader

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

const buffer = 8192

func GetManualInput(in string) <-chan string {
	ch := make(chan string, buffer)
	file, err := ioutil.ReadFile(in)
	if err == nil {
		lines := strings.Split(string(file), "\n")
		go func() {
			for _, line := range lines {
				if line != "" {
					ch <- line
				}
			}
			close(ch)
		}()
	} else {
		close(ch)
	}
	return ch
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

func ReportResult(result bool, s float64) {
	if result {
		fmt.Printf("Your program finished in %fs and is correct! :)\n", s)
	} else {
		fmt.Printf("Your program finished in %fs and is incorrect. :(\n", s)
	}
}

var timeout float64 = 1.0

func SetTimeout(f float64) {
	timeout = f
}

func CheckProgramConditions(t time.Time) (s float64) {
	duration := time.Since(t)
	if s = duration.Seconds(); s > timeout {
		fmt.Printf("Your program timed out after %fs! :(\n", timeout)
	}
	return
}

type Function func()
type Execute func() string
type Report func(string, float64)

type ProgramInformation struct {
	time   float64
	output string
}

func RunFunction(function Function) (result bool) {
	ch := make(chan struct{})
	start := time.Now()
	go func() {
		function()
		close(ch)
	}()

	for {
		select {
		case <-ch:
			result = true
			return
		default:
			if CheckProgramConditions(start) > timeout {
				result = false
				return
			}
		}
	}
}

func RunProgram(execute Execute, report Report) (result bool) {
	ch := make(chan ProgramInformation)

	start := time.Now()
	go func() {
		output := execute()
		ch <- ProgramInformation{time.Since(start).Seconds(), output}
		close(ch)
	}()

	var info ProgramInformation
	for {
		select {
		case info = <-ch:
			report(info.output, info.time)
			result = true
			return
		default:
			if CheckProgramConditions(start) > timeout {
				result = false
				return
			}
		}
	}
}

func RunManualProgram(in string, main ProgramMain) {
	output := main(GetManualInput(in))
	fmt.Println(output)
}

func RunAndValidateManualProgram(in, test string, echo bool, main ProgramMain) {
	RunAndSelfValidateManualProgram(in, echo, main, func(out string) bool {
		return TestOutput(test, out)
	})
}

func RunAndSelfValidateManualProgram(in string, echo bool, main ProgramMain, validation ProgramValidation) {
	input := GetManualInput(in)
	RunProgram(func() string {
		return main(input)
	}, func(output string, time float64) {
		if echo {
			fmt.Println(output)
		}

		result := validation(output)
		ReportResult(result, time)
	})
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

	if RunFunction(func() { program.ParseInitialData(ch) }) {
		var duration float64
		for active := true; active; {
			input := program.GetInput()
			if RunProgram(func() string {
				return program.Update(input)
			}, func(output string, time float64) {
				result := program.SetOutput(output)

				if trace {
					fmt.Printf("%s\n%s\n\n", output, result)
				}

				duration += time

				if program.WinConditionCheck() {
					ReportResult(true, duration)
					active = false
				} else if program.LoseConditionCheck() {
					ReportResult(false, duration)
					active = false
				}
			}) == false {
				return
			}
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
			fmt.Printf("%s ", c)
		}
		fmt.Println("")
	}
	fmt.Println("")
}
