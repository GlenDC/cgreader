package main

import (
	"github.com/bitly/go-simplejson"
	"github.com/glendc/cgreader"
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

	cgreader.SetResetProgramCallback(func() {
		InitializeProgram()
	})

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
		if len(arguments) >= 2 {
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
							ErrorTextParseError("Illegal JSON format, encountered '}'", l, c)
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

					reader := strings.NewReader(string(rawInfo))
					jsonInfo, err := simplejson.NewFromReader(reader)

					if err != nil {
						ErrorIllegalEmbbedFormat()
					}

					var programType string
					var inputFiles, outputFiles []string

					if jsonType := jsonInfo.Get(INFO_TYPE); jsonType != nil {
						if programType, err = jsonType.String(); err != nil {
							ErrorMessage(err.Error())
						}
					} else {
						ErrorMissingProgramType()
						return
					}

					if jsonInput := jsonInfo.Get(INFO_INPUT); jsonInput != nil {
						inputFiles = GetEmbedFiles(jsonInput, inputFiles)
					} else {
						ErrorMissingInputFile()
						return
					}

					if jsonOutput := jsonInfo.Get(INFO_OUTPUT); jsonOutput != nil {
						outputFiles = GetEmbedFiles(jsonOutput, outputFiles)
					} else if programType == CMD_MANUAL {
						ErrorMissingOutputFile()
						return
					}

					switch programType {
					case CMD_MANUAL:
						if len(outputFiles) == 0 {
							ErrorMissingOutputFile()
							return
						} else {
							if len(inputFiles) > 1 {
								if len(outputFiles) == len(inputFiles) {
									CreateAndRunManulPrograms(file, inputFiles, outputFiles)
								} else {
									ErrorManualProgramInputAndOutFilesNotEqual()
								}
							} else {
								CreateAndRunManulProgram(file, inputFiles[0], outputFiles[0])
							}
						}
					case CMD_KIRK, CMD_RAGNAROK, CMD_RAGNAROK_GIANTS:
						CreateAndRunTargetProgram(file, programType, inputFiles[0])
					default:
						ErrorIllegalProgramType(programType)
					}
				} else {
					ErrorIllegalEmbbedFormat()
				}
			} else {
				ErrorIllegalProgramFilePath(program)
			}
		} else {
			ErrorMissingBrainfuckProgram()
		}
	} else {
		switch len(arguments) {
		case 1:
			ErrorMissingProgramType()
			return
		case 2:
			ErrorMissingBrainfuckProgram()
			return
		case 3:
			ErrorMissingInputFile()
			return
		default:
			programType, program, input := arguments[1], arguments[2], arguments[3]

			if file, err := ioutil.ReadFile(program); err == nil {
				switch programType {
				case CMD_MANUAL:
					if len(arguments) < 5 {
						ErrorMissingOutputFile()
					} else {
						output := arguments[4]
						CreateAndRunManulProgram(file, input, output)
					}
				case CMD_KIRK, CMD_RAGNAROK, CMD_RAGNAROK_GIANTS:
					CreateAndRunTargetProgram(file, programType, input)
				default:
					ErrorIllegalProgramType(programType)
				}
			} else {
				ErrorIllegalProgramFilePath(program)
			}
		}
	}
}
