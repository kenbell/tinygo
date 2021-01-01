// +build stm32generic

package machine

// LEDs - assume all boards have 1 built-in LED attached to
// a pin.  Declared as var (not const) since the Pin number
// will be initialized late.
var (
	LED         Pin
	LED1        Pin
	LED_BUILTIN Pin
)

// Basic peripherals
var (
	UART0 *UART
)
