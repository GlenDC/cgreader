package main

const (
	SYNOPSIS            = "bfreader [command] [program] [input] [output]\n\tcommand: a subcommand that defines the type of program to run\n\tprogram: the path to the brainfuck program file\n\tinput: the path to the input test file\n\toutput: the path to the output test file (optional)"
	CMD_MANUAL          = "manual"
	CMD_KIRK            = "kirk"
	CMD_RAGNAROK        = "ragnarok"
	CMD_RAGNAROK_GIANTS = "ragnarokGiants"
	SEPERATOR           = "###"
	ERROR_PREFIX        = "Error: "
)

const (
	FLAG_VERBOSE  = 0x76 // v
	FLAG_EMBEDDED = 0x65 // e
	JSON_START    = 0x7B // {
	JSON_STOP     = 0x7D // }
)

const (
	INFO_TYPE   = "type"
	INFO_INPUT  = "input"
	INFO_OUTPUT = "output"
)

const (
	PI    = 0x3E // >
	PD    = 0x3C // <
	VI    = 0x2B // +
	VD    = 0x2D // -
	IN    = 0x2C // ,
	NOUT  = 0x23 // #
	COUT  = 0x2E // .
	START = 0x5B // [
	STOP  = 0x5D // ]
	TRACE = 0x3F // ?
	LF    = 0x0A // \n
	CR    = 0x0D // \r
	DASH  = 0x2D // -
)

const PROGRAM_SIZE = 30000

const EOF = -1
