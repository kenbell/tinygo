// +build bluepill

package machine

import (
	"device/stm32"
	"runtime/interrupt"
)

const (
	LED = PC13
)

// UART pins
const (
	UART_TX_PIN     = PA9
	UART_RX_PIN     = PA10
	UART_ALT_TX_PIN = PB6
	UART_ALT_RX_PIN = PB7
)

var (
	// USART1 is the first hardware serial port on the STM32.
	// Both UART0 and UART1 refer to USART1.
	UART0 = UART{
		Buffer: NewRingBuffer(),
		Bus:    stm32.USART1,
	}
	UART1 = UART{
		Buffer: NewRingBuffer(),
		Bus:    stm32.USART2,
	}
)

func init() {
	UART0.Interrupt = interrupt.New(stm32.IRQ_USART1, UART0.handleInterrupt)
	UART1.Interrupt = interrupt.New(stm32.IRQ_USART2, UART1.handleInterrupt)
}

// SPI pins
const (
	SPI0_SCK_PIN = PA5
	SPI0_SDO_PIN = PA7
	SPI0_SDI_PIN = PA6
)

// I2C pins
const (
	SDA_PIN = PB7
	SCL_PIN = PB6
)
