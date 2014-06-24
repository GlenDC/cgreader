package cgreader

import (
	"fmt"
	"math"
	"strings"
)

type RagnarokGiants struct {
	thor, dimensions       Vector
	energy, turn, maxTurns int
	giants                 []Vector
	UserInitialize         UserInitializeFunction
	UserUpdate             UserUpdateFunction
	trace                  bool
}

var WAIT string = "WAIT"
var STRIKE string = "STRIKE"

func (ragnarok *RagnarokGiants) IsPositionAvailable(x, y int) bool {
	for i := range ragnarok.giants {
		if x == ragnarok.giants[i].x && y == ragnarok.giants[i].y {
			return false
		}
	}
	return true
}

func (ragnarok *RagnarokGiants) RemoveGiant(x, y int) {
	i := 0
	for ; i < len(ragnarok.giants); i++ {
		if x == ragnarok.giants[i].x && y == ragnarok.giants[i].y {
			ragnarok.giants = append(ragnarok.giants[:i], ragnarok.giants[i+1:]...)
			return
		}
	}
}

func GetADL(x, y int) (int, int) {
	switch {
	default:
		return x, y + x
	case x == 0:
		return y * -1, y
	case x == y:
		return 0, y
	}
	return 0, 0
}

func GetADR(x, y int) (int, int) {
	switch {
	default:
		return x, y - x
	case x == 0:
		return y, y
	case x != y:
		return 0, y
	}
	return 0, 0
}

func (ragnarok *RagnarokGiants) MoveGiant(giant, target *Vector) {
	channel_a := GetDirection(target.x, giant.x)
	channel_b := GetDirection(target.y, giant.y)

	dx, dy := <-channel_a, <-channel_b
	x, y := giant.x+dx, giant.y+dy

	for i := 0; i < 2 && !ragnarok.IsPositionAvailable(x, y); i++ {
		if i == 0 {
			dx, dy = GetADL(dx, dy)
		} else {
			dx, dy = GetADR(dx, dy)
		}
		x, y = giant.x+dx, giant.y+dy
	}

	if ragnarok.IsPositionAvailable(x, y) {
		giant.x, giant.y = x, y
	}
}

func (ragnarok *RagnarokGiants) MoveGiants() {
	for i := range ragnarok.giants {
		ragnarok.MoveGiant(&ragnarok.giants[i], &ragnarok.thor)
	}
}

func (ragnarok *RagnarokGiants) ParseInitialData(ch <-chan string) {
	fmt.Sscanf(
		<-ch,
		"%d %d %d \n",
		&ragnarok.dimensions.x,
		&ragnarok.dimensions.y,
		&ragnarok.maxTurns)

	var giants int

	fmt.Sscanf(
		<-ch,
		"%d %d %d %d \n",
		&ragnarok.energy,
		&ragnarok.thor.x,
		&ragnarok.thor.y,
		&giants)

	output := make(chan string)
	go func() {
		output <- fmt.Sprintf("%d %d", ragnarok.thor.x, ragnarok.thor.y)
	}()
	ragnarok.UserInitialize(output)

	ragnarok.giants = make([]Vector, giants)

	for i := range ragnarok.giants {
		fmt.Sscanf(
			<-ch,
			"%d %d \n",
			&ragnarok.giants[i].x,
			&ragnarok.giants[i].y)
		ragnarok.giants[i].icon = "G"
	}

	ragnarok.thor.icon = "H"
}

func (ragnarok *RagnarokGiants) GetInput() (ch chan string) {
	ch = make(chan string)
	go func() {
		ch <- fmt.Sprintf("%d %d", ragnarok.energy, len(ragnarok.giants))
		for _, giant := range ragnarok.giants {
			ch <- fmt.Sprintf("%d %d", giant.x, giant.y)
		}
	}()
	return
}

func Sqrt(x int) int {
	return int(math.Sqrt(float64(x)))
}

func Pow(x int) int {
	return int(math.Pow(float64(x), 2.0))
}

func (ragnarok *RagnarokGiants) Update(input <-chan string, output chan string) {
	ragnarok.UserUpdate(input, output)
}

func (ragnarok *RagnarokGiants) SetOutput(output []string) string {
	ragnarok.MoveGiants()

	var hotspots []Vector
	if output[0] == STRIKE {
		for i := 0; i < 9; i++ {
			x, y := 0, 1
			for u := 0; u < 2; u++ {
				rx, ry := ragnarok.thor.x+(x*i), ragnarok.thor.y+(y*i)
				lx, ly := ragnarok.thor.x-(x*i), ragnarok.thor.y-(y*i)

				ragnarok.RemoveGiant(lx, ly)
				ragnarok.RemoveGiant(rx, ry)

				hotspots = append(hotspots, Vector{lx, ly, "X"})
				hotspots = append(hotspots, Vector{rx, ry, "X"})

				x, y = GetADR(GetADR(x, y))
			}
		}
		ragnarok.energy -= 1
	} else if output[0] != WAIT {
		if strings.Contains(output[0], "N") {
			ragnarok.thor.y -= 1
		} else if strings.Contains(output[0], "S") {
			ragnarok.thor.y += 1
		}

		if strings.Contains(output[0], "E") {
			ragnarok.thor.x += 1
		} else if strings.Contains(output[0], "W") {
			ragnarok.thor.x -= 1
		}
	}

	ragnarok.turn++

	if ragnarok.trace {
		hotspots = append(hotspots, ragnarok.thor)
		hotspots = append(hotspots, ragnarok.giants...)

		map_info := make([]MapObject, len(hotspots))
		for i, v := range hotspots {
			map_info[i] = MapObject(v)
		}

		DrawMap(
			ragnarok.dimensions.x,
			ragnarok.dimensions.y,
			".",
			map_info...)

		return fmt.Sprintf(
			"Turn = %d\nAmount of Giants = %d\nThor = (%d,%d)\nEnergy = %d",
			ragnarok.turn,
			len(ragnarok.giants),
			ragnarok.thor.x,
			ragnarok.thor.y,
			ragnarok.energy)
	}

	return ""
}

func (ragnarok *RagnarokGiants) LoseConditionCheck() bool {
	if ragnarok.energy <= 0 || ragnarok.turn >= ragnarok.maxTurns {
		return true
	}

	x, y := ragnarok.thor.x, ragnarok.thor.y
	dx, dy := ragnarok.dimensions.x, ragnarok.dimensions.y

	for _, giant := range ragnarok.giants {
		if giant.x == x && giant.y == y {
			return true
		}
	}

	if x < 0 || x >= dx || y < 0 || y >= dy {
		return true
	}

	return false
}

func (ragnarok *RagnarokGiants) WinConditionCheck() bool {
	return len(ragnarok.giants) == 0
}

func RunRagnarokGiantsProgram(input string, trace bool, initialize UserInitializeFunction, update UserUpdateFunction) bool {
	ragnarok := RagnarokGiants{}
	ragnarok.UserInitialize = initialize
	ragnarok.UserUpdate = update
	ragnarok.trace = trace

	return RunTargetProgram(input, trace, &ragnarok)
}

func RunRagnarokGiantsPrograms(input []string, trace bool, initialize UserInitializeFunction, update UserUpdateFunction) {
	var counter int
	for i := range input {
		if RunRagnarokGiantsProgram(input[i], trace, initialize, update) {
			counter++
		}
		Printf("\n")
	}
	ReportTotalResult(counter, len(input))
}
