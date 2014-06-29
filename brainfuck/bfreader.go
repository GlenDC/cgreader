package main

import (
	"flag"
	"fmt"
	"github.com/glendc/cgreader"
	"io/ioutil"
	"os"
	"strings"
)

const (
	SYNOPSIS            = "bfreader [command] [program] [input] [output]\n\tcommand: a subcommand that defines the type of program to run\n\tprogram: the path to the brainfuck program file\n\tinput: the path to the input test file\n\toutput: the path to the output test file (optional)"
	CMD_MANUAL          = "manual"
	CMD_KIRK            = "kirk"
	CMD_RAGNAROK        = "ragnarok"
	CMD_RAGNAROK_GIANTS = "ragnarokGiants"
	SEPERATOR           = "###"
)

const (
	PI    = 0x3E
	PD    = 0x3C
	VI    = 0x2B
	VD    = 0x2D
	IN    = 0x2C
	NOUT  = 0x23
	COUT  = 0x2E
	START = 0x5B
	STOP  = 0x5D
	TO    = 0x3F
	TL    = 0x21
	LF    = 0x0A
	CR    = 0x0D
	DASH  = 0x2D
	TIN   = 0x28
	TOUT  = 0x29
	TSE   = 0x3A
)

const PROGRAM_SIZE = 30000

func ParseProgram(input []byte, trace byte) (string, bool) {
	var output string
	var loopStartCounter, loopStopCounter, l, c uint64
	var cmd byte

	l, c = 1, 1

	for i := 0; i < len(input); i++ {
		switch cmd = input[i]; cmd {
		case PI, PD, VI, VD, IN, NOUT, COUT, START, STOP:
			if cmd == START {
				loopStartCounter++
			} else if cmd == STOP {
				loopStopCounter++
				if loopStopCounter > loopStartCounter {
					fmt.Printf("ERROR! Parsing failed on Line %d (%d): encountered \"]\" while expecting ><+-,.#[\n", l, c)
					return "", false
				}
			}

			output += string(cmd)
		case LF, CR:
			l, c = l+1, 1
		case TIN:
			var io, is int
			io = i+1+strings.Index(string(input[i+1:]), string(TOUT))
			if io != i {
				is = i+1+strings.Index(string(input[i+1:io-1]), string(TSE))
				if is != i && io-i <= 14 {
					var ifi, ila int
					i++

					if is-i == 1 {
						ifi = 0
					} else {
						fmt.Sscanf(string(input[i:is-1]), "%d", &ifi)
					}

					if is = is + 1; io-is == 1 {
						ila = PROGRAM_SIZE - 1
					} else {
						fmt.Sscanf(string(input[is:io-1]), "%d", &ila)
					}

					if trace == TO || trace == TL {
						output += string(trace)
						TraceQueue.Push(&QueuedFunction{func() {
							cgreader.Tracef("Tracing from %d to %d.\n", ifi, ila)
							for i := ifi; i <= ila; i++ {
								cgreader.Tracef("%d ", programBuffer[i])
							}
							cgreader.Traceln("")
						}})
					} else {
						fmt.Printf("ERROR! Parsing failed due to unrecognized trace type \"%d\"\n", trace)
						return "", false
					}

					i = io
				}
			}
		}
		c++
	}

	if result := loopStopCounter == loopStartCounter; result {
		return output, true
	} else {
		fmt.Println("ERROR! Parsing failed due to EOF encounter while expecting \"]\"")
		return "", false
	}
}

func ParseTargetProgram(input string) (initial, update string, result bool) {
	if index := strings.Index(input, SEPERATOR); index != -1 {
		if initial, result = ParseProgram([]byte(input[:index-1]), TO); result {
			update, result = ParseProgram([]byte(input[index+3:]), TL)
		} else {
			result = false
		}
	} else {
		fmt.Printf("ERROR! Please seperate your intial and update logic with \"%s\"\n", SEPERATOR)
		result = false
	}
	return
}

var programStream, programInput string
var programBuffer []int64
var programIndex, streamIndex int

var programCommands map[rune]prco
var TraceQueue Queue

var inputChannel <-chan string
var outputChannel chan string

func InitialzeProgram(stream string) {
	programStream, programIndex, streamIndex = stream, 0, 0
	programBuffer = make([]int64, PROGRAM_SIZE)
	programInput = ""
}

func GetProgramInput() (result int64) {
	if len(programInput) == 0 {
		programInput = <-inputChannel
	}

	result = int64(programInput[0])
	programInput = programInput[1:]
	return
}

func RunLoop(stream string) {
	var cmd rune
	for _, cmd = range stream {
		streamIndex++
		programCommands[cmd]()
	}
}

func main() {
	var arguments []string
	for _, argument := range os.Args {
		if argument[0] != DASH {
			arguments = append(arguments, argument)
		}
	}

	flag.Parse()
	verbose := *flag.Bool("v", false, "verbose")

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
		programCommands = make(map[rune]prco)
		programCommands[PI] = func() { programIndex++ }
		programCommands[PD] = func() { programIndex-- }
		programCommands[VI] = func() { programBuffer[programIndex]++ }
		programCommands[VD] = func() { programBuffer[programIndex]-- }
		programCommands[IN] = func() { programBuffer[programIndex] = GetProgramInput() }
		programCommands[NOUT] = func() {
			outputChannel <- fmt.Sprintf("%d", programBuffer[programIndex])
		}
		programCommands[COUT] = func() {
			outputChannel <- fmt.Sprintf("%s", string(programBuffer[programIndex]))
		}
		programCommands[START] = func() {
			i := strings.Index(programStream[streamIndex:], string(STOP))
			RunLoop(programStream[streamIndex : i-1])
			streamIndex = i + 1
		}
		programCommands[STOP] = func() {
			fmt.Printf("ERROR! Parsing failed: encountered \"]\" while expecting ><+-,.#[\n")
			os.Exit(0)
		}
		programCommands[TO] = func() {
			TraceQueue.Pop().excecute()
		}
		programCommands[TL] = func() {
			cmd := TraceQueue.Pop()
			TraceQueue.Push(cmd)
			cmd.excecute()
		}

		command, program, input := arguments[1], arguments[2], arguments[3]

		if file, err := ioutil.ReadFile(program); err == nil {
			TraceQueue = *NewQueue(1)

			switch command {
			case CMD_MANUAL:
				if len(arguments) < 5 {
					fmt.Printf("ERROR! Please provide the path to an output file...\n%s\n", SYNOPSIS)
				} else {
					output := arguments[4]
					if main, result := ParseProgram(file, TO); result {
						cgreader.RunAndValidateManualProgram(
							input,
							output,
							verbose,
							func(input <-chan string, output chan string) {
								inputChannel, outputChannel = input, output
								InitialzeProgram(main)
								RunLoop(programStream)
							})
					}
				}
			case CMD_KIRK, CMD_RAGNAROK, CMD_RAGNAROK_GIANTS:
				if initial, update, result := ParseTargetProgram(string(file)); result {
					initialFunction := func(input <-chan string) {
						inputChannel = input
						InitialzeProgram(initial)
						RunLoop(programStream)
						InitialzeProgram(update)
					}

					updateFunction := func(input <-chan string, output chan string) {
						inputChannel, outputChannel = input, output
						RunLoop(programStream)
					}

					switch command {
					case CMD_KIRK:
						cgreader.RunKirkProgram(input, verbose, initialFunction, updateFunction)
					case CMD_RAGNAROK:
						cgreader.RunRagnarokProgram(input, verbose, initialFunction, updateFunction)
					case CMD_RAGNAROK_GIANTS:
						cgreader.RunRagnarokGiantsProgram(input, verbose, initialFunction, updateFunction)
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
