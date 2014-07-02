package cgreader

import (
	"fmt"
	"math/rand"
	"time"
)

var buffer int = 2048
var delay time.Duration
var timeout time.Duration = time.Second

func SetBuffer(size int) {
	buffer = size
}

func SetFrameRate(fps int) {
	if fps == 0 {
		SetDelay(0)
	} else {
		SetDelay(1000 / fps)
	}
}

func SetDelay(ms int) {
	t := fmt.Sprintf("%dms", ms)
	d, err := time.ParseDuration(t)
	if err == nil {
		delay = d
	}
}

func SetTimeout(seconds float64) {
	dur, err := time.ParseDuration(fmt.Sprintf("%fs", seconds))
	if err == nil {
		timeout = dur
	}
}

func InitializeCGReader() {
	rand.Seed(time.Now().UTC().UnixNano())
}

type ProgramResetCallback func()

var ResetProgram ProgramResetCallback

func SetResetProgramCallback(callback ProgramResetCallback) {
	ResetProgram = callback
}
