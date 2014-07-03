package main

import (
	"github.com/glendc/cgreader"
)

// program

func CreateAndRunProgram(programFile []byte) {
	if main, result := ParseStaticProgram(programFile); result {
		cgreader.RunProgram(
			func(output chan string) {
				outputChannel = output
				outputIsAvailable = true
				main.excecute()
			})
	}
}

// static

func CreateStaticFunction(main *Command) cgreader.ProgramMain {
	return func(input <-chan string, output chan string) {
		inputChannel, outputChannel = input, output
		inputIsAvailable, outputIsAvailable = true, true
		main.excecute()
	}
}

func CreateAndRunStaticProgram(programFile []byte, programInputFile, programOutputFile string) {
	if main, result := ParseStaticProgram(programFile); result {
		cgreader.RunStaticProgram(
			programInputFile,
			programOutputFile,
			isVerbose,
			CreateStaticFunction(main))
	}
}

func CreateAndRunStaticPrograms(programFile []byte, programInputFiles, programOutputFiles []string) {
	if main, result := ParseStaticProgram(programFile); result {
		cgreader.RunStaticPrograms(
			programInputFiles,
			programOutputFiles,
			isVerbose,
			CreateStaticFunction(main))
	}
}

// interactive

func CreateInteractiveFunctions(initial, update *Command) (cgreader.UserInitializeFunction, cgreader.UserUpdateFunction) {
	initialFunction := func(input <-chan string) {
		inputChannel = input
		inputIsAvailable, outputIsAvailable = true, false
		initial.excecute()
	}

	updateFunction := func(input <-chan string, output chan string) {
		inputChannel, outputChannel = input, output
		inputIsAvailable, outputIsAvailable = true, true
		update.excecute()
	}

	return initialFunction, updateFunction
}

func CreateAndRunInteractiveProgram(programFile []byte, programType, programInputFile string) {
	if initial, update, result := ParseInteractiveProgram(programFile); result {
		initialFunction, updateFunction := CreateInteractiveFunctions(initial, update)
		cgreader.RunInteractiveProgram(programType, programInputFile, isVerbose, initialFunction, updateFunction)
	}
}

func CreateAndRunInteractivePrograms(programFile []byte, programType string, programInputFiles []string) {
	if initial, update, result := ParseInteractiveProgram(programFile); result {
		initialFunction, updateFunction := CreateInteractiveFunctions(initial, update)
		cgreader.RunInteractivePrograms(programType, programInputFiles, isVerbose, initialFunction, updateFunction)
	}
}
