// +build stm32f7x2

package main

import (
	"device/arm"
	"runtime/interrupt"
	"runtime/volatile"

	"github.com/kenbell/tinygo-stm32/clock"
	"github.com/kenbell/tinygo-stm32/clock/hse"
	"github.com/kenbell/tinygo-stm32/clock/pll"
	"github.com/kenbell/tinygo-stm32/power"
	"github.com/kenbell/tinygo-stm32/timer"
	"github.com/kenbell/tinygo-stm32/uart"
)

const PB0 = portB + 0
const PB7 = portB + 7
const PB14 = portB + 14

const (
	LED_GREEN = PB0
	LED_BLUE  = PB7
	LED_RED   = PB14
)

const TICK_RATE = 1000 // 1 KHz
const TICKS_PER_NS = 1000000000 / TICK_RATE

// Adjust these formulas to match the clock configuration (see initClocks)
//
// Use STM32CubeMX to design & visualize the clocks
const (
	HSE_FREQ    = 8000000 // 8 MHz
	PLL_M       = 4
	PLL_N       = 216
	PLL_P       = 2
	PLL_Q       = 9
	SYSCLK_FREQ = ((HSE_FREQ / PLL_M) * PLL_N) / PLL_P
	HCLK_FREQ   = SYSCLK_FREQ / 1
	PCLK1_FREQ  = HCLK_FREQ / 4
	PCLK2_FREQ  = HCLK_FREQ / 2
)

var tickCount volatile.Register64
var timerWakeup volatile.Register8

var myBoard *board

type board struct {
	TickTimer  *timer.Timer
	SleepTimer *timer.Timer
	UART       *uart.UART
}

func (b *board) Initialize() {
	myBoard = b

	b.initClocks()
	b.initTickTimer()
	b.initSleepTimer()

	b.UART = uart.USART3
	b.UART.Configure(uart.Config{
		BaudRate: 115200,
		Clock:    clock.SourcePCLK1,
	})
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

func (b *board) PutChar(c byte) {
	if b.UART == nil {
		return
	}

	b.UART.WriteByte(c)
}

func (b *board) initClocks() {
	power.EnableOverdrive()

	oscCfg := clock.OscillatorConfig{
		HSE: &hse.Config{State: hse.StateBypass},
		PLL: &pll.Config{
			Source: pll.SourceHSE,
			State:  pll.StateOn,
			M:      PLL_M,
			N:      PLL_N,
			P:      PLL_P,
			Q:      PLL_Q,
		},
	}
	oscCfg.Apply()

	clkCfg := clock.Config{
		Types:          clock.TypeSYSCLK | clock.TypeHCLK | clock.TypePCLK1 | clock.TypePCLK2,
		SYSCLKSource:   clock.SYSCLKSourcePLL,
		AHBCLKDivider:  clock.HPREDividerDiv1,
		APB1CLKDivider: clock.PPREDividerDiv4,
		APB2CLKDivider: clock.PPREDividerDiv2,
	}
	clkCfg.Apply(7)

	// Store the configured peripheral clock frequencies, so peripherals
	// can configure relative to the peripheral clock
	clock.HSE.ClockFrequency = HSE_FREQ
	clock.SYSCLK.ClockFrequency = SYSCLK_FREQ
	clock.HCLK.ClockFrequency = HCLK_FREQ
	clock.PCLK1.ClockFrequency = PCLK1_FREQ
	clock.PCLK2.ClockFrequency = PCLK2_FREQ
}

func (b *board) initSleepTimer() {
	b.SleepTimer = timer.TIM3

	intr := b.SleepTimer.NewInterrupt(handleWakeup)
	intr.SetPriority(0xc3)
	intr.Enable()
}

func (b *board) startSleepTimer(ticks int64) {
	cfg := timer.Config{}
	cfg.SetDelay(b.TicksToNanoseconds(ticks), b.SleepTimer.Clock)
	b.SleepTimer.ConfigureBasic(&cfg)
	b.SleepTimer.StartWithInterrupts()
}

func (b *board) initTickTimer() {
	b.TickTimer = timer.TIM7

	// Repeating timer, with prescale and period calculated
	// from the tick rate
	cfg := timer.Config{}
	cfg.SetFrequency(TICK_RATE, b.TickTimer.Clock)

	b.TickTimer.ConfigureBasic(&cfg)

	intr := b.TickTimer.NewInterrupt(handleTick)
	intr.SetPriority(0xc1)
	intr.Enable()

	b.TickTimer.StartWithInterrupts()
}

func handleWakeup(interrupt.Interrupt) {
	if myBoard.SleepTimer.GetAndClearUpdateFlag() {
		// Repeat is disable, but we also stop the timer when
		// not waiting
		myBoard.SleepTimer.Stop()

		// timer was triggered
		timerWakeup.Set(1)
	}
}

func handleTick(interrupt.Interrupt) {
	if myBoard.TickTimer.GetAndClearUpdateFlag() {
		c := tickCount.Get()
		tickCount.Set(c + 1)
	}
}
