// +build stm32generic

package machine

import (
	"device/stm32"
	"errors"
	"runtime"
)

var ErrNotImplemented = errors.New("Not implemented")

// CustomBoard is an interface that enables applications to provide their own
// board & mcu initialization logic.
type CustomBoard interface {
	SleepTicks(d int64)
	Ticks() int64
	TicksToNanoseconds(ticks int64) int64
	NanosecondsToTicks(ns int64) int64
	UART() GenericUART
	LED() Pin
}

// Keep board private, so the only way to initialize is via function
var board CustomBoard

// InitializeBoard sets the custom board implementation
func InitializeBoard(b CustomBoard) {
	board = b
	runtime.SleepTicksFn = b.SleepTicks
	runtime.TicksFn = b.Ticks
	runtime.TicksToNanosecondsFn = b.TicksToNanoseconds
	runtime.NanosecondsToTicksFn = b.NanosecondsToTicks

	LED = b.LED()

	UART0 = &UART{
		Buffer: NewRingBuffer(),
		impl:   board.UART(),
	}
	UART0.impl.SetReceiveCallback(UART0.Receive)
	runtime.PutCharFn = func(ch byte) {
		if UART0 != nil {
			UART0.Write([]byte{ch})
		}
	}
}

// GetBoard returns the custom board implementation (if any)
func GetBoard() CustomBoard {
	return board
}

// GenericUART provides a completely abstracted UART (just
// read/write byte).  It is intended to be implemented by
// custom board implementations to enable `machine.UART0`
// to function as expected by code implemented against the
// standard tinygo 'machine' model.
type GenericUART interface {
	WriteByte(c byte) error
	SetReceiveCallback(func(b byte))
}

// UART implements the standard tinygo UART model, common
// across all MCU types.
type UART struct {
	Buffer *RingBuffer
	impl   GenericUART
}

// WriteByte writes a byte of data to the UART.
func (uart *UART) WriteByte(c byte) error {
	if uart.impl == nil {
		return nil
	}

	return uart.impl.WriteByte(c)
}

type SPI struct {
	Bus *stm32.SPI_Type
}

func (spi SPI) Transfer(w byte) (byte, error) {
	return 0, ErrNotImplemented
}

type I2C struct {
	Bus *stm32.I2C_Type
}

func (i2c I2C) Tx(addr uint16, w, r []byte) error {
	return ErrNotImplemented
}
