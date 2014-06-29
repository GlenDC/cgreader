package main

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
	LF    = 0x0A
	CR    = 0x0D
	DASH  = 0x2D
	TIN   = 0x28
	TOUT  = 0x29
	TSE   = 0x3A
)

const PROGRAM_SIZE = 30000