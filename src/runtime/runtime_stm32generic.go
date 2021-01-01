// +build stm32generic

package runtime

var (
	SleepTicksFn         func(d int64)
	TicksFn              func() int64
	TicksToNanosecondsFn func(ticks int64) int64
	NanosecondsToTicksFn func(ns int64) int64
	PutCharFn            func(c byte)
)

const asyncScheduler = false

func putchar(c byte) {
	if PutCharFn == nil {
		return
	}

	PutCharFn(c)
}

func ticksToNanoseconds(ticks timeUnit) int64 {
	if TicksToNanosecondsFn == nil {
		// Default to tick = 1ms
		return int64(ticks) * 1000000
	}

	return TicksToNanosecondsFn(int64(ticks))
}

func nanosecondsToTicks(ns int64) timeUnit {
	if NanosecondsToTicksFn == nil {
		// Default to tick = 1ms
		return timeUnit(ns / 1000000)
	}

	return timeUnit(NanosecondsToTicksFn(ns))
}

// sleepTicks should sleep for specific number of microseconds.
func sleepTicks(d timeUnit) {
	if SleepTicksFn == nil {
		return
	}

	SleepTicksFn(int64(d))
}

// number of ticks (microseconds) since start.
func ticks() timeUnit {
	if TicksFn == nil {
		return 0
	}

	return timeUnit(TicksFn())
}
