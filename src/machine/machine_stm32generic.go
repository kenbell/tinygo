// +build stm32generic

package machine

type UART struct {
	Buffer *RingBuffer
}

// WriteByte writes a byte of data to the UART.
func (uart UART) WriteByte(c byte) error {
	// Discard
	return nil
}
