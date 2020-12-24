package main

import (
	"machine"
	"runtime"
	"time"
)

const (
	portA machine.Pin = iota * 16
	portB
	portC
	portD
	portE
	portF
	portG
	portH
	portI
	portJ
)

const PB0 = portB + 0
const PB7 = portB + 7
const PB14 = portB + 14

const (
	LED_GREEN = PB0
	LED_BLUE  = PB7
	LED_RED   = PB14
)

func main() {
	LED_GREEN.Configure(machine.PinConfig{Mode: machine.PinOutput})
	LED_BLUE.Configure(machine.PinConfig{Mode: machine.PinOutput})
	LED_RED.Configure(machine.PinConfig{Mode: machine.PinOutput})

	board := board{}
	board.Initialize()
	runtime.Board = &board

	for {
		time.Sleep(500 * time.Millisecond)
		LED_GREEN.High()
		time.Sleep(500 * time.Millisecond)
		LED_GREEN.Low()
	}
}
