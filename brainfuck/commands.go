package main

import (
	"fmt"
	"github.com/glendc/cgreader"
)

type Command interface {
	add(Command)
	run()
}

// >
type AddressIncrementCommand struct{}

func (command AddressIncrementCommand) add(Command) {}
func (command AddressIncrementCommand) run() {
	programIndex++
}

// <
type AddressDecrementCommand struct{}

func (command AddressDecrementCommand) add(Command) {}
func (command AddressDecrementCommand) run() {
	programIndex--
}

// +
type ValueIncrementCommand struct{}

func (command ValueIncrementCommand) add(Command) {}
func (command ValueIncrementCommand) run() {
	programBuffer[programIndex]++
}

// -
type ValueDecrementCommand struct{}

func (command ValueDecrementCommand) add(Command) {}
func (command ValueDecrementCommand) run() {
	programBuffer[programIndex]--
}

// ,
type InputCommand struct{}

func (command InputCommand) add(Command) {}
func (command InputCommand) run() {
	if len(programInput) == 0 {
		programInput = []byte(<-inputChannel)
	}

	programBuffer[programIndex] = int64(programInput[0])
	programInput = programInput[1:]
}

// .
type NumericalOutputCommand struct{}

func (command NumericalOutputCommand) add(Command) {}
func (command NumericalOutputCommand) run() {
	outputChannel <- fmt.Sprintf("%d", programBuffer[programIndex])
}

// #
type AlfabeticalOutputCommand struct{}

func (command AlfabeticalOutputCommand) add(Command) {}
func (command AlfabeticalOutputCommand) run() {
	outputChannel <- fmt.Sprintf("%s", string(programBuffer[programIndex]))
}

// ?
type TraceCommand struct{ startIndex, stopIndex int64 }

func (command TraceCommand) add(Command) {}
func (command TraceCommand) run() {
	for index := command.startIndex; index <= command.stopIndex; index++ {
		cgreader.Tracef("%d ", programBuffer[index])
	}
	cgreader.Traceln("")
}

// Loop
type LoopingGroup struct{ commands []Command }

func (command LoopingGroup) add(cmd Command) {
	command.commands = append(command.commands, cmd)
}
func (command LoopingGroup) run() {
	for programBuffer[programIndex] != 0 {
		var cmd Command
		for _, cmd = range command.commands {
			cmd.run()
		}
	}
}

type LinearGroup struct{ commands []Command }

func (command LinearGroup) add(cmd Command) {
	command.commands = append(command.commands, cmd)
}
func (command LinearGroup) run() {
	var cmd Command
	for _, cmd = range command.commands {
		cmd.run()
	}
}
