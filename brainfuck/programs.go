package main

import (
	"github.com/glendc/cgreader"
)

func CreateManualFunction(main *Command) cgreader.ProgramMain {
	return func(input <-chan string, output chan string) {
		inputChannel, outputChannel = input, output
		inputIsAvailable, outputIsAvailable = true, true
		main.excecute()
	}
}

func CreateAndRunManulProgram(programFile []byte, programInputFile, programOutputFile string) {
	if main, result := ParseManualProgram(programFile); result {
		cgreader.RunAndValidateManualProgram(
			programInputFile,
			programOutputFile,
			isVerbose,
			CreateManualFunction(main))
	}
}

func CreateAndRunManulPrograms(programFile []byte, programInputFiles, programOutputFiles []string) {
	if main, result := ParseManualProgram(programFile); result {
		cgreader.RunAndValidateManualPrograms(
			programInputFiles,
			programOutputFiles,
			isVerbose,
			CreateManualFunction(main))
	}
}

func CreateTargetFunctions(initial, update *Command) (cgreader.UserInitializeFunction, cgreader.UserUpdateFunction) {
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

func CreateAndRunTargetProgram(programFile []byte, programType, programInputFile string) {
	if initial, update, result := ParseTargetProgram(programFile); result {
		initialFunction, updateFunction := CreateTargetFunctions(initial, update)

		switch programType {
		case CMD_KIRK:
			cgreader.RunKirkProgram(programInputFile, isVerbose, initialFunction, updateFunction)
		case CMD_RAGNAROK:
			cgreader.RunRagnarokProgram(programInputFile, isVerbose, initialFunction, updateFunction)
		case CMD_RAGNAROK_GIANTS:
			cgreader.RunRagnarokGiantsProgram(programInputFile, isVerbose, initialFunction, updateFunction)
		}
	}
}

func CreateAndRunTargetPrograms(programFile []byte, programType string, programInputFiles []string) {
	if initial, update, result := ParseTargetProgram(programFile); result {
		initialFunction, updateFunction := CreateTargetFunctions(initial, update)

		switch programType {
		case CMD_KIRK:
			cgreader.RunKirkPrograms(programInputFiles, isVerbose, initialFunction, updateFunction)
		case CMD_RAGNAROK:
			cgreader.RunRagnarokPrograms(programInputFiles, isVerbose, initialFunction, updateFunction)
		case CMD_RAGNAROK_GIANTS:
			cgreader.RunRagnarokGiantsPrograms(programInputFiles, isVerbose, initialFunction, updateFunction)
		}
	}
}
