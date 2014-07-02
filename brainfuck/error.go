package main

import (
	"fmt"
)

func ErrorMessage(message string) {
	fmt.Printf("%s%s\n", ERROR_PREFIX, message)
}

func ErrorTextParseError(message string, line, character int) {
	fmt.Printf("%s%s at line %d (%d)\n", ERROR_PREFIX, message, line, character)
}

func ErrorMissingProgramType() {
	fmt.Printf("%sPlease provide a program type\n%s\n", ERROR_PREFIX, SYNOPSIS)
}

func ErrorMissingBrainfuckProgram() {
	fmt.Printf("%sPlease provide the path to the brainfuck program\n%s\n", ERROR_PREFIX, SYNOPSIS)
}

func ErrorMissingInputFile() {
	fmt.Printf("%sPlease provide the path to an input file...\n%s\n", ERROR_PREFIX, SYNOPSIS)
}

func ErrorMissingOutputFile() {
	fmt.Printf("%sPlease provide the path to an output file...\n%s\n", ERROR_PREFIX, SYNOPSIS)
}

func ErrorIllegalEmbbedFormat() {
	fmt.Printf("%sIllegal Embedded Info Format\n", ERROR_PREFIX)
}

func ErrorIllegalProgramType(programType string) {
	fmt.Printf(
		"%s\"%s\" is not recognized as a valid program type\nLegal program types: %s, %s, %s, %s\n",
		ERROR_PREFIX,
		programType,
		CMD_MANUAL,
		CMD_KIRK,
		CMD_RAGNAROK,
		CMD_RAGNAROK_GIANTS)
}

func ErrorManualProgramInputAndOutFilesNotEqual() {
	fmt.Printf("%sThe amount of input and output files given is not equal\n", ERROR_PREFIX)
}

func ErrorIllegalTargetProgramFormat() {
	fmt.Printf("%sPlease seperate your intial and update logic with \"%s\"\n", ERROR_PREFIX, SEPERATOR)
}

func ErrorIllegalProgramFilePath(program string) {
	fmt.Printf("%s\"%s\" is not recognized as a valid path\n", ERROR_PREFIX, program)
}

func ErrorIllegalEmbbedFormatValueType(fileType string) {
	fmt.Printf("%sUnsupported embbed %s information format.\n", ERROR_PREFIX, fileType)
}

func ErrorIllegalEmbbedFormatSmartPath() {
	fmt.Printf("%sInsert a '*' in your string where the numeric value should be placed.\n", ERROR_PREFIX)
}

func ErrorMissingInEmbedFormat(subject string) {
	fmt.Printf("%sMissing %s in the embed info format\n", subject)
}
