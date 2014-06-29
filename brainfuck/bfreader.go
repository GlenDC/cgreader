package main

import (
	"fmt"
	"github.com/glendc/cgreader"
	"io/ioutil"
	"os"
)

var programInput []byte
var programBuffer []int64
var programIndex int

var isVerbose bool

var inputChannel <-chan string
var outputChannel chan string

func InitializeProgram() {
	programIndex = 0
	programBuffer = make([]int64, PROGRAM_SIZE)
}

func main() {
	isVerbose = false

	var arguments []string
	for _, argument := range os.Args {
		if argument[0] != DASH {
			arguments = append(arguments, argument)
		} else {
			if argument == "-v" || argument == "--verbose" {
				isVerbose = true
			}
		}
	}

	switch len(arguments) {
	case 1:
		fmt.Printf("ERROR! Please provide a command\n%s\n", SYNOPSIS)
		return
	case 2:
		fmt.Printf("ERROR! Please provide the path to the brainfuck program\n%s\n", SYNOPSIS)
		return
	case 3:
		fmt.Printf("ERROR! Please provide the path to an input file...\n%s\n", SYNOPSIS)
		return
	default:
		command, program, input := arguments[1], arguments[2], arguments[3]

		if file, err := ioutil.ReadFile(program); err == nil {
			switch command {
			case CMD_MANUAL:
				if len(arguments) < 5 {
					fmt.Printf("ERROR! Please provide the path to an output file...\n%s\n", SYNOPSIS)
				} else {
					output := arguments[4]
					if main, result := ParseManualProgram(file); result {
						InitializeProgram()
						cgreader.RunAndValidateManualProgram(
							input,
							output,
							isVerbose,
							func(input <-chan string, output chan string) {
								inputChannel, outputChannel = input, output
								main.run()
							})
					}
				}
			case CMD_KIRK, CMD_RAGNAROK, CMD_RAGNAROK_GIANTS:
				if initial, update, result := ParseTargetProgram(file); result {
					InitializeProgram()

					initialFunction := func(input <-chan string) {
						inputChannel = input
						initial.run()
					}

					updateFunction := func(input <-chan string, output chan string) {
						inputChannel, outputChannel = input, output
						update.run()
					}

					switch command {
					case CMD_KIRK:
						cgreader.RunKirkProgram(input, isVerbose, initialFunction, updateFunction)
					case CMD_RAGNAROK:
						cgreader.RunRagnarokProgram(input, isVerbose, initialFunction, updateFunction)
					case CMD_RAGNAROK_GIANTS:
						cgreader.RunRagnarokGiantsProgram(input, isVerbose, initialFunction, updateFunction)
					}
				}
			default:
				fmt.Printf(
					"ERROR! \"%s\" is not recognized as a valid command\nLegal commands: %s, %s, %s, %s\n",
					command,
					CMD_MANUAL,
					CMD_KIRK,
					CMD_RAGNAROK,
					CMD_RAGNAROK_GIANTS)
			}
		} else {
			fmt.Printf("ERROR! \"%s\" is not recognized as a valid path\n", program)
		}
	}
}
