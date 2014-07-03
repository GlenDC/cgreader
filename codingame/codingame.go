package codingame

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func GetManualInput(input string) <-chan string {
	ch := make(chan string, buffer)
	file, err := ioutil.ReadFile(input)
	if err == nil {
		lines := strings.Split(string(file), "\n")
		go func() {
			for _, line := range lines {
				if line != "" {
					ch <- fmt.Sprintf("%s\n", line)
				}
			}
			close(ch)
		}()
	} else {
		Printf("Error: finding input file with name \"%s\"\n", input)
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
		Printf("Error finding output file with name \"%s\"\n", test)
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

	var raw_output []byte

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
		case user_output, ok := <-och:
			if ok {
				raw_output = append(raw_output, []byte(user_output)...)
			} else {
				active = false
			}
		}
	}

	close(exit)

	select {
	case t := <-ch:
		report(strings.Split(string(raw_output), "\n"), t)
	default:
	}
	return
}

func IsAmountOfInputAndTestFilesEqual(input, test []string) bool {
	if len(input) != len(test) {
		Println("Make sure you give an equal amount of input files as the amount of test files.")
		return false
	}
	return true
}

type SandboxProgramFunction func(chan string)

func RunSandboxProgram(main SandboxProgramFunction) {
	InitializeCGReader()

	if ResetProgram != nil {
		ResetProgram()
	}

	RunProgram(func(output chan string) {
		main(output)
		close(output)
	}, func(output []string, time float64) {
		for _, line := range output {
			Println(line)
		}

		Printf("Your program finished in %fs :)\n", time)
	})
}

func RunManualProgram(input string, main ProgramMain) {
	InitializeCGReader()

	if ResetProgram != nil {
		ResetProgram()
	}

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
			Println(line)
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

	if ResetProgram != nil {
		ResetProgram()
	}

	ch := GetManualInput(input)
	result = RunProgram(func(output chan string) {
		main(ch, output)
		close(output)
	}, func(output []string, time float64) {
		if echo {
			for _, line := range output {
				Println(line)
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

	if ResetProgram != nil {
		ResetProgram()
	}

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
						Println(line)
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
