// +build stm32generic

package runtime

type CustomBoard interface {
	SleepTicks(d int64)
	Ticks() int64
	TicksToNanoseconds(ticks int64) int64
	NanosecondsToTicks(ns int64) int64
	PutChar(c byte)
}

var Board CustomBoard

const asyncScheduler = false

func putchar(c byte) {
	if Board == nil {
		return
	}

	Board.PutChar(c)
}

func ticksToNanoseconds(ticks timeUnit) int64 {
	if Board == nil {
		// Default to tick = 1ms
		return int64(ticks) * 1000000
	}

	return Board.TicksToNanoseconds(int64(ticks))
}

func nanosecondsToTicks(ns int64) timeUnit {
	if Board == nil {
		// Default to tick = 1ms
		return timeUnit(ns / 1000000)
	}

	return timeUnit(Board.NanosecondsToTicks(ns))
}

// sleepTicks should sleep for specific number of microseconds.
func sleepTicks(d timeUnit) {
	if Board == nil {
		return
	}

	Board.SleepTicks(int64(d))
}

// number of ticks (microseconds) since start.
func ticks() timeUnit {
	if Board == nil {
		return 0
	}

	return timeUnit(Board.Ticks())
}
