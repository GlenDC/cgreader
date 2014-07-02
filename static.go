package cgreader

func RunStaticProgram(input, output string, trace bool, main ProgramMain) {
	RunAndValidateManualProgram(input, output, trace, main)
}

func RunStaticPrograms(input, output []string, trace bool, main ProgramMain) {
	RunAndValidateManualPrograms(input, output, trace, main)
}
