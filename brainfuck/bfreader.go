package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type brainfuck_t int32

var programInput []byte
var programBuffer []brainfuck_t
var programIndex int

var isVerbose bool

var inputChannel <-chan string
var outputChannel chan string

var outputIsAvailable, inputIsAvailable bool

func InitializeProgram() {
	programIndex = 0
	programBuffer = make([]brainfuck_t, PROGRAM_SIZE)
}

func main() {
	isVerbose = false

	programHasEmbeddedInfo := false

	var arguments []string
	for _, argument := range os.Args {
		if argument[0] != DASH {
			arguments = append(arguments, argument)
		} else {
			for i := 1; i < len(argument); i++ {
				switch argument[i] {
				case FLAG_VERBOSE:
					isVerbose = true
				case FLAG_EMBEDDED:
					programHasEmbeddedInfo = true
				}
			}
		}
	}

	if programHasEmbeddedInfo {

	} else {
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
			programType, program, input := arguments[1], arguments[2], arguments[3]

			if file, err := ioutil.ReadFile(program); err == nil {
				switch programType {
				case CMD_MANUAL:
					if len(arguments) < 5 {
						fmt.Printf("ERROR! Please provide the path to an output file...\n%s\n", SYNOPSIS)
					} else {
						output := arguments[4]
						CreateAndRunManulProgram(file, input, output)
					}
				case CMD_KIRK, CMD_RAGNAROK, CMD_RAGNAROK_GIANTS:
					CreateAndRunTargetProgram(file, programType, input)
				default:
					fmt.Printf(
						"ERROR! \"%s\" is not recognized as a valid command\nLegal commands: %s, %s, %s, %s\n",
						programType,
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
}
