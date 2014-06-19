package cgreader

/*import (
    "fmt"
    "strings"
)*/

type Kirk struct {
	UserInitialize UserInitializeFunction
	UserUpdate     UserUpdateFunction
	trace          bool
}

func (kirk *Kirk) ParseInitialData(ch <-chan string) {
	//kirk.UserInitialize(output)
}

func (kirk *Kirk) GetInput() (ch chan string) {
	ch = make(chan string)
	go func() {
		ch <- "input"
	}()
	return
}

func (kirk *Kirk) Update(input <-chan string, output chan string) {
	kirk.UserUpdate(input, output)
}

func (kirk *Kirk) SetOutput(output []string) string {
	return ""
}

func (kirk *Kirk) LoseConditionCheck() bool {
	return false
}

func (kirk *Kirk) WinConditionCheck() bool {
	return true
}

func RunKirkProgram(input string, trace bool, initialize UserInitializeFunction, update UserUpdateFunction) {
	kirk := Kirk{}
	kirk.UserInitialize = initialize
	kirk.UserUpdate = update
	kirk.trace = trace

	RunTargetProgram(input, trace, &kirk)
}

func RunKirkPrograms(input []string, trace bool, initialize UserInitializeFunction, update UserUpdateFunction) {
	for i := range input {
		RunKirkProgram(input[i], trace, initialize, update)
		Printf("\n")
	}
}
