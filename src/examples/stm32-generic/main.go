package main

import (
	"machine"
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

const PB7 = portB + 7

func main() {
	PB7.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		for i := 0; i < 1000000; i++ {
			PB7.Low()
		}
		for i := 0; i < 1000000; i++ {
			PB7.High()
		}
	}
}
