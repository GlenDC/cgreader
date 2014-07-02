package cgreader

const (
	PT_RAGNAROK        = "ragnarok"
	PT_RAGNAROK_GIANTS = "ragnarok_giants"

	PT_KIRK = "kirk"

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

func ErrorUnknownLevel(level string) {
	Printf("Error: The \"%s\" level is unkown and can not be excecuted. Please try again...\n", level)
}

func ErrorLevelMissing(level string) {
	Printf("Error: The \"%s\" level is not yet supported. Please try again later or implement it yourself @ GitHub...\n", level)
}

func RunInteractiveProgram(programType, input string, trace bool, initialize UserInitializeFunction, update UserUpdateFunction) {
	switch programType {
	default:
		ErrorUnknownLevel(programType)

	case PT_RAGNAROK:
		RunRagnarokProgram(input, trace, initialize, update)

	case PT_RAGNAROK_GIANTS:
		RunRagnarokGiantsProgram(input, trace, initialize, update)

	case PT_KIRK:
		RunKirkProgram(input, trace, initialize, update)

	case PT_SKYNET:
		ErrorLevelMissing("Skynet")

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
	switch programType {
	default:
		ErrorUnknownLevel(programType)

	case PT_RAGNAROK:
		RunRagnarokPrograms(input, trace, initialize, update)

	case PT_RAGNAROK_GIANTS:
		RunRagnarokGiantsPrograms(input, trace, initialize, update)

	case PT_KIRK:
		RunKirkPrograms(input, trace, initialize, update)

	case PT_SKYNET:
		ErrorLevelMissing("Skynet")

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
