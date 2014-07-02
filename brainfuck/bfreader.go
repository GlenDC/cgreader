package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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

	if len(arguments) >= 2 && programHasEmbeddedInfo {
		var program string
		var rawInfo []byte

		program = arguments[1]

		if file, err := ioutil.ReadFile(program); err == nil {
			var char byte
			var programIndex int
			infoIsOK := false

			for ac, oc, c, l, i := 0, 0, 0, 0, 0; i < len(file); i, c = i+1, c+1 {
				if char = file[i]; char == JSON_START {
					ac++
				} else if char == JSON_STOP {
					if oc = oc + 1; oc > ac {
						fmt.Printf("ERROR! Illegal JSON format, encountered '}' at line %d (%d)\n", l, c)
						return
					}
				} else if char == LF || char == CR {
					l, c = l+1, 1
				}

				rawInfo = append(rawInfo, char)

				if ac > 0 && ac == oc {
					infoIsOK, programIndex = true, i
					break
				}
			}

			if infoIsOK {
				file = file[programIndex:]

				decoder := json.NewDecoder(strings.NewReader(string(rawInfo)))
				var jsonInfo map[string]interface{}

				if err := decoder.Decode(&jsonInfo); err != nil {
					fmt.Printf("ERROR! %s\n", err)
					return
				}

				var programType string
				var inputFiles, outputFiles []string

				for key, value := range jsonInfo {
					switch key {
					case INFO_TYPE:
						programType = value.(string)

					case INFO_INPUT:
						inputFiles = append(inputFiles, value.(string))

					case INFO_OUTPUT:
						outputFiles = append(outputFiles, value.(string))

					}
				}

				if programType == "" {
					fmt.Printf("ERROR! Please provide a program type\n%s\n", SYNOPSIS)
				} else if len(inputFiles) == 0 {
					fmt.Printf("ERROR! Please provide the path to an input file...\n%s\n", SYNOPSIS)
				}

				switch programType {
				case CMD_MANUAL:
					if len(outputFiles) == 0 {
						fmt.Printf("ERROR! Please provide the path to an output file...\n%s\n", SYNOPSIS)
						return
					} else {
						CreateAndRunManulProgram(file, inputFiles[0], outputFiles[0])
					}
				case CMD_KIRK, CMD_RAGNAROK, CMD_RAGNAROK_GIANTS:
					CreateAndRunTargetProgram(file, programType, inputFiles[0])
				default:
					fmt.Printf(
						"ERROR! \"%s\" is not recognized as a valid program type\nLegal program types: %s, %s, %s, %s\n",
						programType,
						CMD_MANUAL,
						CMD_KIRK,
						CMD_RAGNAROK,
						CMD_RAGNAROK_GIANTS)
				}
			} else {
				fmt.Println("ERROR! Illegal Embedded Info Format.")
			}
		}
	} else {
		switch len(arguments) {
		case 1:
			fmt.Printf("ERROR! Please provide a program type\n%s\n", SYNOPSIS)
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
						"ERROR! \"%s\" is not recognized as a valid program type\nLegal program types: %s, %s, %s, %s\n",
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
