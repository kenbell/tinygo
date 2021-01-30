// +build stm32,stm32f103xx

package machine

// Peripheral abstraction layer for the stm32.

import (
	"device/stm32"
	"unsafe"
)

const (
	PinInput       PinMode = 0 // Input mode
	PinOutput10MHz PinMode = 1 // Output mode, max speed 10MHz
	PinOutput2MHz  PinMode = 2 // Output mode, max speed 2MHz
	PinOutput50MHz PinMode = 3 // Output mode, max speed 50MHz
	PinOutput      PinMode = PinOutput2MHz

	PinInputModeAnalog     PinMode = 0  // Input analog mode
	PinInputModeFloating   PinMode = 4  // Input floating mode
	PinInputModePullUpDown PinMode = 8  // Input pull up/down mode
	PinInputModeReserved   PinMode = 12 // Input mode (reserved)

	PinOutputModeGPPushPull   PinMode = 0  // Output mode general purpose push/pull
	PinOutputModeGPOpenDrain  PinMode = 4  // Output mode general purpose open drain
	PinOutputModeAltPushPull  PinMode = 8  // Output mode alt. purpose push/pull
	PinOutputModeAltOpenDrain PinMode = 12 // Output mode alt. purpose open drain
)

const (
	PA0  = portA + 0
	PA1  = portA + 1
	PA2  = portA + 2
	PA3  = portA + 3
	PA4  = portA + 4
	PA5  = portA + 5
	PA6  = portA + 6
	PA7  = portA + 7
	PA8  = portA + 8
	PA9  = portA + 9
	PA10 = portA + 10
	PA11 = portA + 11
	PA12 = portA + 12
	PA13 = portA + 13
	PA14 = portA + 14
	PA15 = portA + 15

	PB0  = portB + 0
	PB1  = portB + 1
	PB2  = portB + 2
	PB3  = portB + 3
	PB4  = portB + 4
	PB5  = portB + 5
	PB6  = portB + 6
	PB7  = portB + 7
	PB8  = portB + 8
	PB9  = portB + 9
	PB10 = portB + 10
	PB11 = portB + 11
	PB12 = portB + 12
	PB13 = portB + 13
	PB14 = portB + 14
	PB15 = portB + 15

	PC0  = portC + 0
	PC1  = portC + 1
	PC2  = portC + 2
	PC3  = portC + 3
	PC4  = portC + 4
	PC5  = portC + 5
	PC6  = portC + 6
	PC7  = portC + 7
	PC8  = portC + 8
	PC9  = portC + 9
	PC10 = portC + 10
	PC11 = portC + 11
	PC12 = portC + 12
	PC13 = portC + 13
	PC14 = portC + 14
	PC15 = portC + 15

	PD0  = portD + 0
	PD1  = portD + 1
	PD2  = portD + 2
	PD3  = portD + 3
	PD4  = portD + 4
	PD5  = portD + 5
	PD6  = portD + 6
	PD7  = portD + 7
	PD8  = portD + 8
	PD9  = portD + 9
	PD10 = portD + 10
	PD11 = portD + 11
	PD12 = portD + 12
	PD13 = portD + 13
	PD14 = portD + 14
	PD15 = portD + 15

	PE0  = portE + 0
	PE1  = portE + 1
	PE2  = portE + 2
	PE3  = portE + 3
	PE4  = portE + 4
	PE5  = portE + 5
	PE6  = portE + 6
	PE7  = portE + 7
	PE8  = portE + 8
	PE9  = portE + 9
	PE10 = portE + 10
	PE11 = portE + 11
	PE12 = portE + 12
	PE13 = portE + 13
	PE14 = portE + 14
	PE15 = portE + 15

	PF0  = portF + 0
	PF1  = portF + 1
	PF2  = portF + 2
	PF3  = portF + 3
	PF4  = portF + 4
	PF5  = portF + 5
	PF6  = portF + 6
	PF7  = portF + 7
	PF8  = portF + 8
	PF9  = portF + 9
	PF10 = portF + 10
	PF11 = portF + 11
	PF12 = portF + 12
	PF13 = portF + 13
	PF14 = portF + 14
	PF15 = portF + 15

	PG0  = portG + 0
	PG1  = portG + 1
	PG2  = portG + 2
	PG3  = portG + 3
	PG4  = portG + 4
	PG5  = portG + 5
	PG6  = portG + 6
	PG7  = portG + 7
	PG8  = portG + 8
	PG9  = portG + 9
	PG10 = portG + 10
	PG11 = portG + 11
	PG12 = portG + 12
	PG13 = portG + 13
	PG14 = portG + 14
	PG15 = portG + 15
)

// Configure this pin with the given I/O settings.
// stm32f1xx uses different technique for setting the GPIO pins than the stm32f407
func (p Pin) Configure(config PinConfig) {
	// Configure the GPIO pin.
	p.enableClock()
	port := p.getPort()
	pin := uint8(p) % 16
	pos := (pin % 8) * 4
	if pin < 8 {
		port.CRL.ReplaceBits(uint32(config.Mode), 0xf, pos)
	} else {
		port.CRH.ReplaceBits(uint32(config.Mode), 0xf, pos)
	}
}

func (p Pin) ConfigureAltFunc(config PinConfig, altFunc uint8) {
	p.Configure(config)
}

func (p Pin) getPort() *stm32.GPIO_Type {
	switch p / 16 {
	case 0:
		return stm32.GPIOA
	case 1:
		return stm32.GPIOB
	case 2:
		return stm32.GPIOC
	case 3:
		return stm32.GPIOD
	case 4:
		return stm32.GPIOE
	case 5:
		return stm32.GPIOF
	case 6:
		return stm32.GPIOG
	default:
		panic("machine: unknown port")
	}
}

// enableClock enables the clock for this desired GPIO port.
func (p Pin) enableClock() {
	switch p / 16 {
	case 0:
		stm32.RCC.APB2ENR.SetBits(stm32.RCC_APB2ENR_IOPAEN)
	case 1:
		stm32.RCC.APB2ENR.SetBits(stm32.RCC_APB2ENR_IOPBEN)
	case 2:
		stm32.RCC.APB2ENR.SetBits(stm32.RCC_APB2ENR_IOPCEN)
	case 3:
		stm32.RCC.APB2ENR.SetBits(stm32.RCC_APB2ENR_IOPDEN)
	case 4:
		stm32.RCC.APB2ENR.SetBits(stm32.RCC_APB2ENR_IOPEEN)
	case 5:
		stm32.RCC.APB2ENR.SetBits(stm32.RCC_APB2ENR_IOPFEN)
	case 6:
		stm32.RCC.APB2ENR.SetBits(stm32.RCC_APB2ENR_IOPGEN)
	default:
		panic("machine: unknown port")
	}
}

// Enable peripheral clock. Expand to include all the desired peripherals
func enableAltFuncClock(bus unsafe.Pointer) {
	if bus == unsafe.Pointer(stm32.USART1) {
		stm32.RCC.APB2ENR.SetBits(stm32.RCC_APB2ENR_USART1EN)
	} else if bus == unsafe.Pointer(stm32.USART2) {
		stm32.RCC.APB1ENR.SetBits(stm32.RCC_APB1ENR_USART2EN)
	} else if bus == unsafe.Pointer(stm32.I2C1) {
		stm32.RCC.APB1ENR.SetBits(stm32.RCC_APB1ENR_I2C1EN)
	} else if bus == unsafe.Pointer(stm32.SPI1) {
		stm32.RCC.APB2ENR.SetBits(stm32.RCC_APB2ENR_SPI1EN)
	}
}
