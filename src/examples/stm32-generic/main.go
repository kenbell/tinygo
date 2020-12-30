package main

import (
	"fmt"
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

func main() {
	LED_GREEN.Configure(machine.PinConfig{Mode: machine.PinOutput})
	LED_BLUE.Configure(machine.PinConfig{Mode: machine.PinOutput})
	LED_RED.Configure(machine.PinConfig{Mode: machine.PinOutput})

	board := board{}
	board.Initialize()
	runtime.Board = &board

	i := 0
	for {
		time.Sleep(500 * time.Millisecond)
		LED_GREEN.High()
		time.Sleep(500 * time.Millisecond)
		LED_GREEN.Low()
		fmt.Printf("Boo! %v\r\n", i)
		i++
	}
}
