// +build stm32generic

package runtime

var tickCount timeUnit

func putchar(c byte) {
}

func ticksToNanoseconds(ticks timeUnit) int64 {
	return int64(ticks) * 1000
}

func nanosecondsToTicks(ns int64) timeUnit {
	return timeUnit(ns / 1000)
}

const asyncScheduler = false

// sleepTicks should sleep for specific number of microseconds.
func sleepTicks(d timeUnit) {
}

// number of ticks (microseconds) since start.
func ticks() timeUnit {
	return 0
}
