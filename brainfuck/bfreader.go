package main

import (
	//"github.com/glendc/cgreader"
	"flag"
	"fmt"
)

const (
	SYNOPSIS = "bfreader [command] [input] [output]\n\tcommand: a subcommand that defines the type of program to run\n\tinput: the path to the program input file\n\toutput: the path to the program output file (optional)\n"
	COMMANDS = "manual"
)

func main() {

	arguments := flag.Args()
	switch len(arguments) {
	case 0:
		fmt.Println("ERROR! Please provide a command...\n", SYNOPSIS)
		return
	case 1:
		fmt.Println("ERROR! Please provide the path to an input file...\n", SYNOPSIS)
		return
	default:
		command := arguments[0]
		switch command {
		case "manual":
			// define logic here
			// 1. Parse
			// 2. Report errors if there are, if so also quit
			// 3. If parsed ok => Run program
		default:
			fmt.Println("ERROR! \"", command, "\" is not recognized as a valid command...\nLegal commands: ", COMMANDS)
		}
	}
}
