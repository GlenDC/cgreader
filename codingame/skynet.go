package codingame

import (
	"fmt"
)

type Skynet struct {
	bike           Vector
	speed, turns   int
	platforms      []Vector
	UserInitialize UserInitializeFunction
	UserUpdate     UserUpdateFunction
	trace, crashed bool
}

const (
	ICON_BIKE     = "P"
	ICON_ROAD     = "+"
	ICON_PLATFORM = "-"
	ICON_GAP      = "."
)

const (
	MAX_TURNS = 50
	MIN_SPEED = 0
	MAX_SPEED = 50
)

const (
	CMD_SPEED = "SPEED"
	CMD_SLOW  = "SLOW"
	CMD_JUMP  = "JUMP"
	CMD_WAIT  = "WAIT"
)

func (skynet *Skynet) ParseInitialData(ch <-chan string) {
	output := make(chan string)

	var R, G, L int

	fmt.Sscanf(<-ch, "%d", &skynet.speed)
	fmt.Sscanf(<-ch, "%d", &R)
	fmt.Sscanf(<-ch, "%d", &G)
	fmt.Sscanf(<-ch, "%d", &L)

	i, T := 0, R+G+L

	skynet.platforms = make([]Vector, T)

	for ; i < R; i++ {
		skynet.platforms[i] = Vector{i, 0, ICON_ROAD}
	}

	for T = R + G; i < T; i++ {
		skynet.platforms[i] = Vector{i, 0, ICON_GAP}
	}

	for T = T + L; i < T; i++ {
		skynet.platforms[i] = Vector{i, 0, ICON_PLATFORM}
	}

	go func() {
		output <- fmt.Sprintf("%d\n", R)
		output <- fmt.Sprintf("%d\n", G)
		output <- fmt.Sprintf("%d\n", L)
	}()
	skynet.UserInitialize(output)

	skynet.bike = Vector{0, 0, ICON_BIKE}
	skynet.crashed = false
}

func (skynet *Skynet) GetInput() (ch chan string) {
	ch = make(chan string)
	go func() {
		ch <- fmt.Sprintf("%d\n", skynet.speed)
		ch <- fmt.Sprintf("%d\n", skynet.bike.x)
	}()
	return
}

func (skynet *Skynet) Update(input <-chan string, output chan string) {
	skynet.UserUpdate(input, output)
}

func (skynet *Skynet) SetOutput(output []string) string {
	isJumping := false
	var userCommand string

	switch output[0] {
	default:
		userCommand = CMD_WAIT

	case CMD_SPEED:
		userCommand = CMD_SPEED
		if skynet.speed < MAX_SPEED {
			skynet.speed++
		}

	case CMD_SLOW:
		userCommand = CMD_SLOW
		if skynet.speed > MIN_SPEED {
			skynet.speed--
		}

	case CMD_JUMP:
		isJumping, userCommand = true, CMD_JUMP
	}

	skynet.turns++

	for i, p, t := 0, skynet.bike.x, len(skynet.platforms); i < skynet.speed; i++ {
		if p = p + 1; p < t {
			if !isJumping && skynet.platforms[p].icon == ICON_GAP {
				skynet.crashed = true
			} else {
				skynet.bike.x = p
			}
		} else {
			skynet.crashed = true
		}
	}

	if skynet.trace {
		mapInfo := make([]MapObject, len(skynet.platforms)+1)
		mapInfo[len(skynet.platforms)] = MapObject(skynet.bike)
		for i, platform := range skynet.platforms {
			mapInfo[i] = MapObject(platform)
		}

		DrawMap(
			len(mapInfo),
			1,
			"?",
			mapInfo...)

		return fmt.Sprintf(
			"Bike = (%d)\nCommand = %s\nTurn %d/%d",
			skynet.bike.x,
			userCommand,
			skynet.turns,
			MAX_TURNS)
	}

	return ""
}

func (skynet *Skynet) LoseConditionCheck() bool {
	return skynet.crashed || skynet.turns >= 50 ||
		skynet.bike.x >= len(skynet.platforms) ||
		skynet.platforms[skynet.bike.x].icon == ICON_GAP
}

func (skynet *Skynet) WinConditionCheck() bool {
	return skynet.speed == 0 &&
		skynet.platforms[skynet.bike.x].icon == ICON_PLATFORM
}

func RunSkynetProgram(input string, trace bool, initialize UserInitializeFunction, update UserUpdateFunction) bool {
	SetTimeout(0.150)

	skynet := Skynet{}
	skynet.UserInitialize = initialize
	skynet.UserUpdate = update
	skynet.trace = trace

	return RunTargetProgram(input, trace, &skynet)
}

func RunSkynetPrograms(input []string, trace bool, initialize UserInitializeFunction, update UserUpdateFunction) {
	var counter int
	for i := range input {
		if RunSkynetProgram(input[i], trace, initialize, update) {
			counter++
		}
		Println("")
	}
	ReportTotalResult(counter, len(input))
}
