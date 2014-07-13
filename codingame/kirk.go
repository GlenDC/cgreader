package codingame

import (
	"fmt"
	"math"
	"strings"
)

const (
	KIRK_M0 = 0x000f
	KIRK_M1 = 0x00f0
	KIRK_M2 = 0x0f00
	KIRK_MX = 0xf000

	KIRK_S0 = 0
	KIRK_S1 = 4
	KIRK_S2 = 8
	KIRK_SX = 12

	KIRK_N = 8
)

type Kirk struct {
	UserInitialize UserInitializeFunction
	UserUpdate     UserUpdateFunction
	trace          bool
	maxHeight      uint8
	mountains      []uint32
	player         Vector
	direction      int8
	canFire        bool
}

func (kirk *Kirk) ParseInitialData(ch <-chan string) {
	fmt.Sscanf(<-ch, "%d", &kirk.maxHeight)
	kirk.player, kirk.maxHeight = Vector{0, 0, "S"}, kirk.maxHeight+1

	for i := 0; i < KIRK_N; i++ {
		heights := strings.Split(<-ch, " ")
		for u, h := range heights {
			var height uint32
			fmt.Sscanf(h, "%d", &height)
			kirk.mountains[i] += height << (uint32(u) * KIRK_S1)
		}
		if kirk.mountains[i]&KIRK_M0 != 0 {
			kirk.mountains[i] += uint32(len(heights)) << KIRK_SX
		}
	}

	kirk.UserInitialize(make(chan string))
}

func (kirk *Kirk) GetHeight(m uint32) uint32 {
	var height uint32
	if id := m >> KIRK_SX; id > 0 {
		for ; id > 0; id-- {
			s := KIRK_S1 * (id - 1)
			height += (m & (KIRK_M0 << s)) >> s
		}
	}
	return uint32(kirk.maxHeight) - height - 1
}

func (kirk *Kirk) GetInput() (ch chan string) {
	ch = make(chan string)
	go func() {
		ch <- fmt.Sprintf("%d %d\n", kirk.player.x, kirk.player.y)
		for _, mountain := range kirk.mountains {
			ch <- fmt.Sprintf("%d\n", uint32(kirk.maxHeight)-kirk.GetHeight(mountain)-1)
		}
	}()
	return
}

func (kirk *Kirk) Update(input <-chan string, output chan string) {
	kirk.UserUpdate(input, output)
}

func (kirk *Kirk) GetMountains(icons []MapObject) []MapObject {
	for i, mountain := range kirk.mountains {
		for height := int(kirk.GetHeight(mountain)); height < int(kirk.maxHeight); height++ {
			icons = append(icons, MapObject(Vector{i, height, "X"}))
		}
	}
	return icons
}

func (kirk *Kirk) SetOutput(output []string) string {
	playerFired, damage := kirk.canFire && output[0] == "FIRE", uint32(0)
	if playerFired {
		kirk.canFire = false
		m := kirk.mountains[kirk.player.x]
		x := m >> KIRK_SX
		if x > 0 {
			id := x - 1
			s := KIRK_S1 * id
			damage = uint32((m >> s) & KIRK_M0)
			x--
			m &= math.MaxUint32 - KIRK_MX - (KIRK_M0 << (id * KIRK_S1))
			kirk.mountains[kirk.player.x] = m | (x << KIRK_SX)
		}
	}

	kirk.player.x += int(kirk.direction)
	if kirk.player.x < 0 || kirk.player.x > KIRK_N-1 {
		kirk.canFire = true
		kirk.player.y, kirk.direction = kirk.player.y+1, -kirk.direction
		kirk.player.x += int(kirk.direction)
	}

	if kirk.trace {
		icons := make([]MapObject, 1)
		icons[0] = MapObject(kirk.player)
		icons = kirk.GetMountains(icons)
		DrawMap(KIRK_N, int(kirk.maxHeight)-1, ".", icons...)

		shipInfo := fmt.Sprintf("Ship = (%d,%d)\n", kirk.player.x, kirk.player.y)
		if playerFired {
			return shipInfo + fmt.Sprintf("Ship fired and did %d damage.", damage)
		} else {
			return shipInfo + "Ship hold fire."
		}
	}

	return ""
}

func (kirk *Kirk) LoseConditionCheck() bool {
	return kirk.player.y >= int(kirk.maxHeight) ||
		kirk.player.y >= int(kirk.GetHeight(kirk.mountains[kirk.player.x]))
}

func (kirk *Kirk) WinConditionCheck() bool {
	for _, mountain := range kirk.mountains {
		if mountain>>KIRK_SX > 0 {
			return false
		}
	}
	return true
}

func RunKirkProgram(input string, trace bool, initialize UserInitializeFunction, update UserUpdateFunction) bool {
	kirk := Kirk{}

	SetTimeout(0.1)

	kirk.UserInitialize, kirk.UserUpdate, kirk.trace = initialize, update, trace
	kirk.mountains, kirk.direction, kirk.canFire = make([]uint32, KIRK_N), 1, true

	return RunTargetProgram(input, trace, &kirk)
}

func RunKirkPrograms(input []string, trace bool, initialize UserInitializeFunction, update UserUpdateFunction) {
	var counter int
	for i := range input {
		if RunKirkProgram(input[i], trace, initialize, update) {
			counter++
		}
		Println("")
	}
	ReportTotalResult(counter, len(input))
}
