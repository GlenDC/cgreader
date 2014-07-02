package main

import (
	"fmt"
)

func ErrorTextParseError(message string, line, character int) {
	fmt.Printf("ERROR! %s at line %d (%d)\n", message, line, character)
}

func ErrorMissingProgramType() {
	fmt.Printf("ERROR! Please provide a program type\n%s\n", SYNOPSIS)
}

func ErrorMissingBrainfuckProgram() {
	fmt.Printf("ERROR! Please provide the path to the brainfuck program\n%s\n", SYNOPSIS)
}

func ErrorMissingInputFile() {
	fmt.Printf("ERROR! Please provide the path to an input file...\n%s\n", SYNOPSIS)
}

func ErrorMissingOutputFile() {
	fmt.Printf("ERROR! Please provide the path to an output file...\n%s\n", SYNOPSIS)
}

func ErrorIllegalEmbbedFormat() {
	fmt.Println("ERROR! Illegal Embedded Info Format.")
}

func ErrorIllegalProgramType(programType string) {
	fmt.Printf(
		"ERROR! \"%s\" is not recognized as a valid program type\nLegal program types: %s, %s, %s, %s\n",
		programType,
		CMD_MANUAL,
		CMD_KIRK,
		CMD_RAGNAROK,
		CMD_RAGNAROK_GIANTS)
}

func IllegalTargetProgramFormat() {
	fmt.Printf("ERROR! Please seperate your intial and update logic with \"%s\"\n", SEPERATOR)
}

func ErrorIllegalProgramFilePath(program string) {
	fmt.Printf("ERROR! \"%s\" is not recognized as a valid path\n", program)
}
