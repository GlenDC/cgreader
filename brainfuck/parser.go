package main

import (
	"fmt"
	"github.com/glendc/cgreader"
	"strings"
)

var rawProgramStream []byte
var currentStreamIndex int
var lineCounter, characterCounter, startLoopCounter, stopLoopCounter uint64
var streamIsValid bool

func RecursiveParser(command *Command) {
	for ; streamIsValid && currentStreamIndex < len(rawProgramStream); currentStreamIndex++ {
		characterCounter++

		switch rawProgramStream[currentStreamIndex] {
		case PI:
			command.add(&Command{func([]*Command) {
				programIndex++
			}, nil})

		case PD:
			command.add(&Command{func([]*Command) {
				programIndex--
			}, nil})

		case VI:
			command.add(&Command{func([]*Command) {
				programBuffer[programIndex]++
			}, nil})

		case VD:
			command.add(&Command{func([]*Command) {
				programBuffer[programIndex]--
			}, nil})

		case IN:
			command.add(&Command{func([]*Command) {
				if len(programInput) == 0 {
					programInput = []byte(<-inputChannel)
				}

				programBuffer[programIndex] = int64(programInput[0])
				programInput = programInput[1:]
			}, nil})

		case NOUT:
			command.add(&Command{func([]*Command) {
				outputChannel <- fmt.Sprintf("%d", programBuffer[programIndex])
			}, nil})

		case COUT:
			command.add(&Command{func([]*Command) {
				outputChannel <- fmt.Sprintf("%s", string(programBuffer[programIndex]))
			}, nil})

		case START:
			startLoopCounter++
			currentStreamIndex++

			loop := CreateLoopGroup()
			RecursiveParser(loop)
			command.add(loop)

		case STOP:
			if stopLoopCounter > startLoopCounter {
				fmt.Printf("ERROR! Parsing failed on Line %d (%d): encountered \"]\" while expecting ><+-,.#[\n", lineCounter, characterCounter)
				streamIsValid = false
			}

			stopLoopCounter++
			currentStreamIndex++
			return

		case LF, CR:
			lineCounter, characterCounter = lineCounter+1, 0

		case TIN:
			var io, is int
			io = currentStreamIndex + 1 + strings.Index(string(rawProgramStream[currentStreamIndex+1:]), string(TOUT))
			if io != currentStreamIndex {
				is = currentStreamIndex + 1 + strings.Index(string(rawProgramStream[currentStreamIndex+1:io-1]), string(TSE))
				if is != currentStreamIndex && io-currentStreamIndex <= 14 {
					var ifi, ila int64
					currentStreamIndex++

					if is-currentStreamIndex == 1 {
						ifi = 0
					} else {
						fmt.Sscanf(string(rawProgramStream[currentStreamIndex:is-1]), "%d", &ifi)
					}

					if is = is + 1; io-is == 1 {
						ila = PROGRAM_SIZE - 1
					} else {
						fmt.Sscanf(string(rawProgramStream[is:io-1]), "%d", &ila)
					}

					command.add(&Command{func([]*Command) {
						for index := ifi; index <= ila; index++ {
							cgreader.Tracef("%d ", programBuffer[index])
						}
						cgreader.Traceln("")
					}, nil})

					currentStreamIndex = io
				}
			}
		}
	}
}

func InitializeParser(input []byte) {
	rawProgramStream, currentStreamIndex = input, 0
	startLoopCounter, stopLoopCounter = 0, 0
	streamIsValid = true
}

func ParseLinearProgram(input []byte) *Command {
	InitializeParser(input)
	command := CreateLinearGroup()
	RecursiveParser(command)
	return command
}

func ParseLoopingProgram(input []byte) *Command {
	InitializeParser(input)
	command := CreateLoopGroup()
	RecursiveParser(command)
	return command
}

func ParseManualProgram(stream []byte) (*Command, bool) {
	lineCounter, characterCounter = 0, 0
	program := ParseLinearProgram(stream)
	return program, streamIsValid
}

func ParseTargetProgram(stream []byte) (initial, update *Command, result bool) {
	lineCounter, characterCounter = 0, 0
	if index := strings.Index(string(stream), SEPERATOR); index != -1 {
		if initial = ParseLinearProgram(stream[:index-1]); streamIsValid {
			update = ParseLoopingProgram(stream[index+3:])
			result = streamIsValid
		} else {
			result = false
		}
	} else {
		fmt.Printf("ERROR! Please seperate your intial and update logic with \"%s\"\n", SEPERATOR)
		result = false
	}
	return
}
