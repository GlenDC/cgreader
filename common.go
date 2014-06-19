package cgreader

import (
	"fmt"
)

type UserInitializeFunction func(<-chan string)
type UserUpdateFunction func(<-chan string, chan string)

type Vector struct {
	x, y int
	icon string
}

func (v Vector) GetMapCoordinates() string {
	return fmt.Sprintf("%d;%d", v.x, v.y)
}

func (v Vector) GetMapIcon() string {
	return v.icon
}
