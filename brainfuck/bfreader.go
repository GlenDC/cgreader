package main

import (
	//"github.com/glendc/cgreader"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	SYNOPSIS     = "bfreader [command] [program] [input] [output]\n\tcommand: a subcommand that defines the type of program to run\n\tprogram: the path to the brainfuck program file\n\tinput: the path to the input test file\n\toutput: the path to the output test file (optional)"
	CMD_MANUAL   = "manual"
	CMD_RAGNAROK = "ragnarok"
	SEPERATOR    = "###"
)

const (
	PI    = 0x3E
	PD    = 0x3C
	VI    = 0x2B
	VD    = 0x2D
	IN    = 0x2C
	OUT   = 0x2E
	START = 0x5B
	STOP  = 0x5D
	LF    = 0x0A
	CR    = 0x0D
)

func ParseProgram(input []byte) (string, bool) {
	var output string
	var loopStartCounter, loopStopCounter, l, c uint64
	var cmd byte

	l, c = 1, 1

	for i := range input {
		switch cmd = input[i]; cmd {
		case PI, PD, VI, VD, IN, OUT, START, STOP:
			if cmd == START {
				loopStartCounter++
			} else if cmd == STOP {
				loopStopCounter++
				if loopStopCounter > loopStartCounter {
					fmt.Printf("ERROR! Parsing failed on Line %d (%d): encountered \"]\" while expecting ><+-,.[\n", l, c)
					return "", false
				}
			}

			output += string(cmd)
		case LF, CR:
			l, c = l+1, 1
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
		if initial, result = ParseProgram([]byte(input[:index-1])); result {
			update, result = ParseProgram([]byte(input[index+3:]))
		} else {
			result = false
		}
	} else {
		fmt.Printf("ERROR! Please seperate your intial and update logic with \"%s\"\n", SEPERATOR)
		result = false
	}
	return
}

func main() {
	arguments := os.Args
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
		command, program := arguments[1], arguments[2]
		if file, err := ioutil.ReadFile(program); err == nil {
			switch command {
			case CMD_MANUAL:
				if main, result := ParseProgram(file); result {
					fmt.Println(main)
				}
			case CMD_RAGNAROK:
				if initial, update, result := ParseTargetProgram(string(file)); result {
					fmt.Println(initial)
					fmt.Println(update)
				}
			default:
				fmt.Printf(
					"ERROR! \"%s\" is not recognized as a valid command\nLegal commands: %s, %s\n",
					command,
					CMD_MANUAL,
					CMD_RAGNAROK)
			}
		} else {
			fmt.Printf("ERROR! \"%s\" is not recognized as a valid path\n", program)
		}
	}
}
