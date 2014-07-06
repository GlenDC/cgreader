package cgreader

import (
	"fmt"
	"github.com/glendc/cgreader/codingame"
)

// levels

const (
	PT_RAGNAROK        = "ragnarok"
	PT_RAGNAROK_GIANTS = "ragnarok_giants"

	PT_KIRK        = "kirk"
	PS_KIRK_BRIDGE = "skynet_bridge"

	PT_SKYNET         = "skynet"
	PT_SKYNET_FINAL_1 = "skynet_final_1"
	PT_SKYNET_FINAL_2 = "skynet_final_2"

	PT_MARS_LANDER_1 = "mars_lander_1"
	PT_MARS_LANDER_2 = "mars_lander_2"
	PT_MARS_LANDER_3 = "mars_lander_3"

	PT_INDIANA_1 = "indiana_1"
	PT_INDIANA_2 = "indiana_2"
	PT_INDIANA_3 = "indiana_3"

	PT_KIRK_LABYRINTH = "labyrinth"

	PT_SHADOW_KNIGHT_1 = "shadow_knight_1"
	PT_SHADOW_KNIGHT_2 = "shadow_knight_2"
)

// typedefinitions

type (
	ProgramMain            codingame.ProgramMain
	UserInitializeFunction codingame.UserInitializeFunction
	UserUpdateFunction     codingame.UserUpdateFunction
	ProgramResetCallback   codingame.ProgramResetCallback
	PrintfCallback         codingame.PrintfCallback
	SandboxProgramFunction codingame.SandboxProgramFunction
)

// configuration

func SetBuffer(size int) {
	codingame.SetBuffer(size)
}

func SetFrameRate(fps int) {
	if fps == 0 {
		codingame.SetDelay(0)
	} else {
		codingame.SetDelay(1000 / fps)
	}
}

func SetDelay(ms int) {
	codingame.SetDelay(ms)
}

func SetTimeout(seconds float64) {
	codingame.SetTimeout(seconds)
}

func SetPrintfCallback(callback PrintfCallback) {
	codingame.SetPrintfCallback(codingame.PrintfCallback(callback))
}

func SetResetProgramCallback(callback ProgramResetCallback) {
	codingame.SetResetProgramCallback(codingame.ProgramResetCallback(callback))
}

// trace

func Trace(msg string) {
	codingame.Print(msg)
}

func Traceln(msg string) {
	codingame.Println(msg)
}

func Tracef(format string, a ...interface{}) {
	codingame.Printf("%s", fmt.Sprintf(format, a...))
}

// help

func GetFileList(format string, n int) (files []string) {
	files = make([]string, n)
	for i := range files {
		files[i] = fmt.Sprintf(format, i+1)
	}
	return
}

// sandbox

func RunProgram(main SandboxProgramFunction) {
	codingame.RunSandboxProgram(codingame.SandboxProgramFunction(main))
}

// static

func RunStaticProgram(input, output string, trace bool, main ProgramMain) {
	codingame.RunAndValidateManualProgram(input, output, trace, codingame.ProgramMain(main))
}

func RunStaticPrograms(input, output []string, trace bool, main ProgramMain) {
	codingame.RunAndValidateManualPrograms(input, output, trace, codingame.ProgramMain(main))
}

// interactive

type levelMissingFunction func(string)

var ErrorUnknownLevel levelMissingFunction = func(level string) {
	codingame.Printf("Error: The \"%s\" level is unkown and can not be excecuted. Please try again...\n", level)
}

var ErrorLevelMissing levelMissingFunction = func(level string) {
	codingame.Printf("Error: The \"%s\" level is not yet supported. Please try again later or implement it yourself @ GitHub...\n", level)
}

func RunInteractiveProgram(programType, input string, trace bool, initialize UserInitializeFunction, update UserUpdateFunction) {
	userInit := codingame.UserInitializeFunction(initialize)
	userUpdate := codingame.UserUpdateFunction(update)

	switch programType {
	default:
		ErrorUnknownLevel(programType)

	case PT_RAGNAROK:
		codingame.RunRagnarokProgram(input, trace, userInit, userUpdate)

	case PT_RAGNAROK_GIANTS:
		codingame.RunRagnarokGiantsProgram(input, trace, userInit, userUpdate)

	case PT_KIRK:
		codingame.RunKirkProgram(input, trace, userInit, userUpdate)

	case PS_KIRK_BRIDGE:
		ErrorLevelMissing("Kirk Bridge")

	case PT_SKYNET:
		codingame.RunSkynetProgram(input, trace, userInit, userUpdate)

	case PT_SKYNET_FINAL_1:
		ErrorLevelMissing("Skynet Final #1")

	case PT_SKYNET_FINAL_2:
		ErrorLevelMissing("Skynet Final #2")

	case PT_MARS_LANDER_1:
		ErrorLevelMissing("Mars Lander #1")

	case PT_MARS_LANDER_2:
		ErrorLevelMissing("Mars Lander #2")

	case PT_MARS_LANDER_3:
		ErrorLevelMissing("Mars Lander #3")

	case PT_INDIANA_1:
		ErrorLevelMissing("Indiana #1")

	case PT_INDIANA_2:
		ErrorLevelMissing("Indiana #2")

	case PT_INDIANA_3:
		ErrorLevelMissing("Indiana #3")

	case PT_KIRK_LABYRINTH:
		ErrorLevelMissing("Kirk Labyrinth")

	case PT_SHADOW_KNIGHT_1:
		ErrorLevelMissing("Shadow of the Knight #1")

	case PT_SHADOW_KNIGHT_2:
		ErrorLevelMissing("Shadow of the Knight #2")
	}
}

func RunInteractivePrograms(programType string, input []string, trace bool, initialize UserInitializeFunction, update UserUpdateFunction) {
	userInit := codingame.UserInitializeFunction(initialize)
	userUpdate := codingame.UserUpdateFunction(update)

	switch programType {
	default:
		ErrorUnknownLevel(programType)

	case PT_RAGNAROK:
		codingame.RunRagnarokPrograms(input, trace, userInit, userUpdate)

	case PT_RAGNAROK_GIANTS:
		codingame.RunRagnarokGiantsPrograms(input, trace, userInit, userUpdate)

	case PT_KIRK:
		codingame.RunKirkPrograms(input, trace, userInit, userUpdate)

	case PS_KIRK_BRIDGE:
		ErrorLevelMissing("Kirk Bridge")

	case PT_SKYNET:
		codingame.RunSkynetPrograms(input, trace, userInit, userUpdate)

	case PT_SKYNET_FINAL_1:
		ErrorLevelMissing("Skynet Final #1")

	case PT_SKYNET_FINAL_2:
		ErrorLevelMissing("Skynet Final #2")

	case PT_MARS_LANDER_1:
		ErrorLevelMissing("Mars Lander #1")

	case PT_MARS_LANDER_2:
		ErrorLevelMissing("Mars Lander #2")

	case PT_MARS_LANDER_3:
		ErrorLevelMissing("Mars Lander #3")

	case PT_INDIANA_1:
		ErrorLevelMissing("Indiana #1")

	case PT_INDIANA_2:
		ErrorLevelMissing("Indiana #2")

	case PT_INDIANA_3:
		ErrorLevelMissing("Indiana #3")

	case PT_KIRK_LABYRINTH:
		ErrorLevelMissing("Kirk Labyrinth")

	case PT_SHADOW_KNIGHT_1:
		ErrorLevelMissing("Shadow of the Knight #1")

	case PT_SHADOW_KNIGHT_2:
		ErrorLevelMissing("Shadow of the Knight #2")
	}
}
