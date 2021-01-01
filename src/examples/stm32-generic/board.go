package main

import (
	"device/arm"
	"machine"
	"runtime/interrupt"

	"github.com/kenbell/tinygo-stm32/timer"
	"github.com/kenbell/tinygo-stm32/uart"
)

type board struct {
	tickTimer  *timer.Timer
	sleepTimer *timer.Timer
	uart       *uart.UART
}

func (b *board) SleepTicks(d int64) {
	timerWakeup.Set(0)

	b.startSleepTimer(d)

	// wait till timer wakes up
	for timerWakeup.Get() == 0 {
		arm.Asm("wfi")
	}
}

func (b *board) Ticks() int64 {
	return int64(tickCount.Get())
}

func (b *board) TicksToNanoseconds(ticks int64) int64 {
	return ticks * TICKS_PER_NS
}

func (b *board) NanosecondsToTicks(ns int64) int64 {
	return ns / TICKS_PER_NS
}

func (b *board) UART() machine.GenericUART {
	return b.uart
}

func (b *board) initSleepTimer(t *timer.Timer) {
	b.sleepTimer = t

	intr := b.sleepTimer.NewInterrupt(handleWakeup)
	intr.SetPriority(0xc3)
	intr.Enable()
}

func (b *board) startSleepTimer(ticks int64) {
	cfg := timer.Config{}
	cfg.SetDelay(b.TicksToNanoseconds(ticks), b.sleepTimer.Clock)
	b.sleepTimer.ConfigureBasic(&cfg)
	b.sleepTimer.StartWithInterrupts()
}

func (b *board) initTickTimer(t *timer.Timer) {
	b.tickTimer = t

	// Repeating timer, with prescale and period calculated
	// from the tick rate
	cfg := timer.Config{}
	cfg.SetFrequency(TICK_RATE, b.tickTimer.Clock)

	b.tickTimer.ConfigureBasic(&cfg)

	intr := b.tickTimer.NewInterrupt(handleTick)
	intr.SetPriority(0xc1)
	intr.Enable()

	b.tickTimer.StartWithInterrupts()
}

func handleWakeup(interrupt.Interrupt) {
	if myBoard.sleepTimer.GetAndClearUpdateFlag() {
		// Repeat is disable, but we also stop the timer when
		// not waiting
		myBoard.sleepTimer.Stop()

		// timer was triggered
		timerWakeup.Set(1)
	}
}

func handleTick(interrupt.Interrupt) {
	if myBoard.tickTimer.GetAndClearUpdateFlag() {
		c := tickCount.Get()
		tickCount.Set(c + 1)
	}
}
