package main

import (
	"fmt"
	"machine"
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
	machine.InitializeBoard(&board)

	go func() {
		buf := []byte{0}

		for {
			n, _ := machine.UART0.Read(buf)
			if n != 0 {
				fmt.Printf("Got 0x%x\r\n", buf[0])
			}
			time.Sleep(50*time.Millisecond)
		}
	}()

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
