package main

import (
	//"github.com/glendc/cgreader"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	SYNOPSIS = "bfreader [command] [program] [input] [output]\n\tcommand: a subcommand that defines the type of program to run\n\tprogram: the path to the brainfuck program file\n\tinput: the path to the input test file\n\toutput: the path to the output test file (optional)"
	COMMANDS = "manual"
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
		command := arguments[1]
		switch command {
		case "manual":
			program := arguments[2]
			if file, err := ioutil.ReadFile(program); err == nil {
				if pp, ok := ParseProgram(file); ok {
					fmt.Println(pp)
				}
			} else {
				fmt.Printf("ERROR! \"%s\" is not recognized as a valid path\n", program)
			}
		default:
			fmt.Printf("ERROR! \"%s\" is not recognized as a valid command\nLegal commands: %s\n", command, COMMANDS)
		}
	}
}
