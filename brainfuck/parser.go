package main

import (
	"fmt"
	"strings"
)

var rawProgramStream []byte
var currentStreamIndex int
var lineCounter, characterCounter, startLoopCounter, stopLoopCounter uint64
var streamIsValid bool

func RecursiveParser(command Command) Command {
	for ; streamIsValid && currentStreamIndex < len(rawProgramStream); currentStreamIndex++ {
		characterCounter++

		switch rawProgramStream[currentStreamIndex] {
		case PI:
			command.add(Command(AddressIncrementCommand{}))

		case PD:
			command.add(Command(AddressDecrementCommand{}))

		case VI:
			command.add(Command(ValueIncrementCommand{}))
			fmt.Printf("Lenght is now: %d\n", len(command.(LinearGroup).commands))

		case VD:
			command.add(Command(ValueDecrementCommand{}))

		case IN:
			command.add(Command(InputCommand{}))

		case NOUT:
			command.add(Command(NumericalOutputCommand{}))

		case COUT:
			command.add(Command(AlfabeticalOutputCommand{}))

		case START:
			startLoopCounter++
			currentStreamIndex++
			loop := LinearGroup{}
			baseCommand := RecursiveParser(Command(loop))
			loop = baseCommand.(LinearGroup)
			command.add(loop)

		case STOP:
			if stopLoopCounter > startLoopCounter {
				fmt.Printf("ERROR! Parsing failed on Line %d (%d): encountered \"]\" while expecting ><+-,.#[\n", lineCounter, characterCounter)
				streamIsValid = false
			}

			stopLoopCounter++
			currentStreamIndex++
			return command

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

					command.add(TraceCommand{ifi, ila})

					currentStreamIndex = io
				}
			}
		}
	}
	return command
}

func InitializeParser(input []byte) {
	rawProgramStream, currentStreamIndex = input, 0
	startLoopCounter, stopLoopCounter = 0, 0
	streamIsValid = true
}

func ParseLinearProgram(input []byte) *LinearGroup {
	InitializeParser(input)
	command := RecursiveParser(Command(LinearGroup{}))
	if program, ok := command.(LinearGroup); ok {
		fmt.Println(len(program.commands))
		return &program
	} else {
		return nil
	}
}

func ParseLoopingProgram(input []byte) *LoopingGroup {
	InitializeParser(input)
	command := RecursiveParser(Command(LoopingGroup{}))
	if program, ok := command.(LoopingGroup); ok {
		return &program
	} else {
		return nil
	}
}

func ParseManualProgram(stream []byte) (*LinearGroup, bool) {
	lineCounter, characterCounter = 0, 0
	program := ParseLinearProgram(stream)
	return program, streamIsValid
}

func ParseTargetProgram(stream []byte) (initial *LinearGroup, update *LoopingGroup, result bool) {
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
