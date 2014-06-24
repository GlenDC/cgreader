package cgreader

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

// conf

var buffer int = 2048
var delay time.Duration
var timeout time.Duration = time.Second

func SetBuffer(size int) {
	buffer = size
}

func SetFrameRate(fps int) {
	if fps == 0 {
		SetDelay(0)
	} else {
		SetDelay(1000 / fps)
	}
}

func SetDelay(ms int) {
	t := fmt.Sprintf("%dms", ms)
	d, err := time.ParseDuration(t)
	if err == nil {
		delay = d
	}
}

func SetTimeout(seconds float64) {
	dur, err := time.ParseDuration(fmt.Sprintf("%fs", seconds))
	if err == nil {
		timeout = dur
	}
}

func InitializeCGReader() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// output

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

// debug

func Trace(msg string) {
	Print(msg)
}

func Traceln(msg string) {
	Println(msg)
}

func Tracef(format string, a ...interface{}) {
	Printf("%s\n", fmt.Sprintf(format, a...))
}

// help

func GetFileList(format string, n int) (files []string) {
	files = make([]string, n)
	for i := range files {
		files[i] = fmt.Sprintf(format, i+1)
	}
	return
}

// src

func GetManualInput(input string) <-chan string {
	ch := make(chan string, buffer)
	file, err := ioutil.ReadFile(input)
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
		Printf("Error finding input file with name \"%s\"", input)
		close(ch)
	}
	return ch
}

func TestOutput(test string, output []string) bool {
	if len(output) == 0 {
		return false
	}

	file, err := ioutil.ReadFile(test)
	if err == nil {
		test := strings.Split(string(file), "\n")

		for i, line := range output {
			if line != test[i] {
				return false
			}
		}

		return true
	} else {
		Printf("Error finding output file with name \"%s\"", test)
	}
	return false
}

type ProgramMain func(<-chan string, chan string)

func ReportResult(result bool, s float64) {
	if result {
		Printf("Your program finished in %fs and is correct! :)\n", s)
	} else {
		Printf("Your program finished in %fs and is incorrect. :(\n", s)
	}
}

func CheckProgramConditions(t time.Time) float64 {
	duration := time.Since(t)
	if duration.Seconds() > timeout.Seconds() {
		Printf("Your program timed out after %fs! :(\n", timeout.Seconds())
	}
	return duration.Seconds()
}

type Function func()
type Execute func(chan string)
type Report func([]string, float64)

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
			if CheckProgramConditions(start) > timeout.Seconds() {
				result = false
				return
			}
		}
	}
	return
}

func RunProgram(execute Execute, report Report) (result bool) {
	ch := make(chan float64)
	och := make(chan string, buffer)
	exit := make(chan struct{})
	error := make(chan struct{})

	output := make([]string, 0)

	result = true
	start := time.Now()
	go func() {
		execute(och)
		ch <- time.Since(start).Seconds()
		close(ch)
	}()

	go func() {
		for {
			select {
			case <-exit:
				return
			default:
				if CheckProgramConditions(start) > timeout.Seconds() {
					close(error)
				}
				time.Sleep(timeout)

			}

		}
	}()

	for active := true; active; {
		select {
		case <-error:
			active, result = false, false
		case line, ok := <-och:
			if ok {
				output = append(output, line)
			} else {
				active = false
			}
		}
	}

	close(exit)

	select {
	case t := <-ch:
		report(output, t)
	default:
	}
	return
}

func IsAmountOfInputAndTestFilesEqual(input, test []string) bool {
	if len(input) != len(test) {
		Print("Make sure you give an equal amount of input files as the amount of test files.")
		return false
	}
	return true
}

func RunManualProgram(input string, main ProgramMain) {
	InitializeCGReader()

	output := make(chan string, buffer)
	exit := make(chan struct{})

	go func() {
		main(GetManualInput(input), output)
		close(output)
		close(exit)
	}()

	for {
		select {
		case <-exit:
			return
		case line := <-output:
			Print(line)
		}
	}
}

func RunManualPrograms(input []string, main ProgramMain) {
	for i := range input {
		RunManualProgram(input[i], main)
		Println("")
	}
}

func RunAndValidateManualProgram(input, test string, echo bool, main ProgramMain) (result bool) {
	InitializeCGReader()

	ch := GetManualInput(input)
	result = RunProgram(func(output chan string) {
		main(ch, output)
		close(output)
	}, func(output []string, time float64) {
		if echo {
			for _, line := range output {
				Print(line)
			}
		}

		result = TestOutput(test, output)
		ReportResult(result, time)
	}) && result
	return
}

func ReportTotalResult(correct, total int) {
	emoji := ":)"
	if correct != total {
		emoji = ":("
	}
	Printf("All programs finished. %d/%d programs succeeded %s\n", correct, total, emoji)
}

func RunAndValidateManualPrograms(input, test []string, echo bool, main ProgramMain) {
	if IsAmountOfInputAndTestFilesEqual(input, test) {
		var counter int
		for i := range input {
			if RunAndValidateManualProgram(input[i], test[i], echo, main) {
				counter++
			}
			Println("")
		}
		ReportTotalResult(counter, len(input))
	}
}

type TargetProgram interface {
	ParseInitialData(<-chan string)
	GetInput() chan string
	Update(<-chan string, chan string)
	SetOutput([]string) string
	LoseConditionCheck() bool
	WinConditionCheck() bool
}

func RunTargetProgram(input string, trace bool, program TargetProgram) (isOK bool) {
	InitializeCGReader()

	ch := GetManualInput(input)

	if RunFunction(func() { program.ParseInitialData(ch) }) {
		for active := true; active; {
			input := program.GetInput()
			if RunProgram(func(output chan string) {
				program.Update(input, output)
				close(output)
			}, func(output []string, duration float64) {
				result := program.SetOutput(output)

				if trace {
					for _, line := range output {
						Print(line)
					}
					Printf("\n%s\n\n", result)
				}

				duration += duration

				if program.WinConditionCheck() {
					ReportResult(true, duration)
					active, isOK = false, true
				} else if program.LoseConditionCheck() {
					ReportResult(false, duration)
					active, isOK = false, false
				}

				time.Sleep(delay)
			}) == false {
				isOK = false
			}
		}
	}
	return
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
			Printf("%s ", c)
		}
		Println("")
	}
	Println("")
}
