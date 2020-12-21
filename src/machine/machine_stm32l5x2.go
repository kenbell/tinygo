// +build stm32l5x2

package machine

// Peripheral abstraction layer for the stm32f407

import (
	"device/stm32"
	"runtime/volatile"
)

// System clocks
const (
	CLOCK_LSI     = "LSI"
	CLOCK_LSE     = "LSE"
	CLOCK_MSI     = "MSI"
	CLOCK_HSI     = "HSI"
	CLOCK_HSE     = "HSE"
	CLOCK_HSE_RTC = "HSE_RTC"
	CLOCK_HSI48   = "HSI48"
	CLOCK_SYSCLK  = "SYSCLK"
	CLOCK_HCLK    = "HCLK"
)

// PLL provided clocks
const (
	CLOCK_PLLCLK = "PLLCLK"
	CLOCK_PLLP   = "PLLP"
	CLOCK_PLLQ   = "PLLQ"

	CLOCK_PLLSAI1P = "PLLSAI1P"
	CLOCK_PLLSAI1Q = "PLLSAI1Q"
	CLOCK_PLLSAI1R = "PLLSAI1R"

	CLOCK_PLLSAI2P = "PLLSAI2P"
)

// Peripheral clocks
const (
	CLOCK_PCLK1     = "PCLK1"
	CLOCK_PCLK1_TIM = "PLCLK1_TIM"
	CLOCK_PCLK2     = "PCLK2"
	CLOCK_PCLK2_TIM = "PCLK2_TIM"

	CLOCK_CLK48   = "CLK48"
	CLOCK_RTC     = "RTC"
	CLOCK_RNG     = "RNG"
	CLOCK_SDMMC1  = "SDMMC1"
	CLOCK_USART1  = "USART1"
	CLOCK_USART2  = "USART2"
	CLOCK_USART3  = "USART3"
	CLOCK_UART4   = "UART4"
	CLOCK_UART5   = "UART5"
	CLOCK_LPUART1 = "LPUART1"
	CLOCK_ADC     = "ADC"
)

const (
	PLL_OUTPUT_P = 0
	PLL_OUTPUT_Q = 1
	PLL_OUTPUT_R = 2
)

const (
	LSE_FREQ   = 32768    // 32.768 KHz
	LSI_FREQ   = 32000    // 32 KHz
	HSI_FREQ   = 16000000 // 16 MHz
	HSI48_FREQ = 48000000 // 48 MHz
)

var Clocks = []string{
	CLOCK_LSI, CLOCK_LSE, CLOCK_MSI, CLOCK_HSI, CLOCK_HSE, CLOCK_HSE_RTC, CLOCK_HSI48, CLOCK_RTC, CLOCK_SYSCLK, CLOCK_HCLK,
	CLOCK_PLLCLK, CLOCK_PLLP, CLOCK_PLLQ, CLOCK_PLLSAI1P, CLOCK_PLLSAI1Q, CLOCK_PLLSAI1R, CLOCK_PLLSAI2P,
	CLOCK_PCLK1, CLOCK_PCLK2, CLOCK_CLK48, CLOCK_RTC, CLOCK_RNG, CLOCK_SDMMC1, CLOCK_USART1, CLOCK_USART2, CLOCK_USART3,
	CLOCK_UART4, CLOCK_UART5, CLOCK_LPUART1, CLOCK_ADC,
}

//---------- Clock related code
type peripheralClockSource struct {
	Reg     *volatile.Register32
	Mask    uint32
	Shift   uint32
	Sources []string
}

var peripheralClockSources = map[string]peripheralClockSource{
	CLOCK_CLK48: {
		Reg:     &stm32.RCC.CCIPR1,
		Mask:    stm32.RCC_CCIPR1_CLK48MSEL_Msk,
		Shift:   stm32.RCC_CCIPR1_CLK48MSEL_Pos,
		Sources: []string{CLOCK_HSI48, CLOCK_PLLQ, CLOCK_PLLSAI1Q, CLOCK_MSI},
	},
	CLOCK_RTC: {
		Reg:     &stm32.RCC.BDCR,
		Mask:    stm32.RCC_BDCR_RTCSEL_Msk,
		Shift:   stm32.RCC_BDCR_RTCSEL_Pos,
		Sources: []string{"", CLOCK_LSE, CLOCK_LSI, CLOCK_HSE_RTC},
	},
	CLOCK_RNG: {
		Sources: []string{CLOCK_CLK48},
	},
	CLOCK_SDMMC1: {
		Reg:     &stm32.RCC.CCIPR2,
		Mask:    stm32.RCC_CCIPR2_SDMMCSEL_Msk,
		Shift:   stm32.RCC_CCIPR2_SDMMCSEL_Pos,
		Sources: []string{CLOCK_CLK48, CLOCK_PLLP},
	},
	CLOCK_USART1: {
		Reg:     &stm32.RCC.CCIPR1,
		Mask:    stm32.RCC_CCIPR1_USART1SEL_Msk,
		Shift:   stm32.RCC_CCIPR1_USART1SEL_Pos,
		Sources: []string{CLOCK_PCLK2, CLOCK_SYSCLK, CLOCK_HSI, CLOCK_LSE},
	},
	CLOCK_USART2: {
		Reg:     &stm32.RCC.CCIPR1,
		Mask:    stm32.RCC_CCIPR1_USART2SEL_Msk,
		Shift:   stm32.RCC_CCIPR1_USART2SEL_Pos,
		Sources: []string{CLOCK_PCLK1, CLOCK_SYSCLK, CLOCK_HSI, CLOCK_LSE},
	},
	CLOCK_USART3: {
		Reg:     &stm32.RCC.CCIPR1,
		Mask:    stm32.RCC_CCIPR1_USART3SEL_Msk,
		Shift:   stm32.RCC_CCIPR1_USART3SEL_Pos,
		Sources: []string{CLOCK_PCLK1, CLOCK_SYSCLK, CLOCK_HSI, CLOCK_LSE},
	},
	CLOCK_UART4: {
		Reg:     &stm32.RCC.CCIPR1,
		Mask:    stm32.RCC_CCIPR1_UART4SEL_Msk,
		Shift:   stm32.RCC_CCIPR1_UART4SEL_Pos,
		Sources: []string{CLOCK_PCLK1, CLOCK_SYSCLK, CLOCK_HSI, CLOCK_LSE},
	},
	CLOCK_UART5: {
		Reg:     &stm32.RCC.CCIPR1,
		Mask:    stm32.RCC_CCIPR1_UART5SEL_Msk,
		Shift:   stm32.RCC_CCIPR1_UART5SEL_Pos,
		Sources: []string{CLOCK_PCLK1, CLOCK_SYSCLK, CLOCK_HSI, CLOCK_LSE},
	},
	CLOCK_LPUART1: {
		Reg:     &stm32.RCC.CCIPR1,
		Mask:    stm32.RCC_CCIPR1_LPUART1SEL_Msk,
		Shift:   stm32.RCC_CCIPR1_LPUART1SEL_Pos,
		Sources: []string{CLOCK_PCLK1, CLOCK_SYSCLK, CLOCK_HSI, CLOCK_LSE},
	},
	CLOCK_ADC: {
		Reg:     &stm32.RCC.CCIPR1,
		Mask:    stm32.RCC_CCIPR1_ADCSEL_Msk,
		Shift:   stm32.RCC_CCIPR1_ADCSEL_Pos,
		Sources: []string{"", CLOCK_PLLSAI1R, "", CLOCK_SYSCLK},
	},
}

func GetClockFreq(clock string) uint32 {
	// Primitive clocks
	switch clock {
	case CLOCK_LSI:
		return getLSIFreq()
	case CLOCK_LSE:
		return getLSEFreq()
	case CLOCK_MSI:
		return getMSIFreq()
	case CLOCK_HSI:
		return getHSIFreq()
	case CLOCK_HSE:
		return getHSEFreq()
	case CLOCK_HSE_RTC:
		return getHSEFreq() / 32
	case CLOCK_HSI48:
		return getHSI48Freq()
	case CLOCK_SYSCLK:
		return getSYSCLKFreq()
	case CLOCK_HCLK:
		return getHCLKFreq()
	case CLOCK_PLLCLK:
		return getPLLFreq(PLL_OUTPUT_R)
	case CLOCK_PLLP:
		return getPLLFreq(PLL_OUTPUT_P)
	case CLOCK_PLLQ:
		return getPLLFreq(PLL_OUTPUT_Q)
	case CLOCK_PLLSAI1P:
		return getPLLSAI1Freq(PLL_OUTPUT_P)
	case CLOCK_PLLSAI1Q:
		return getPLLSAI1Freq(PLL_OUTPUT_Q)
	case CLOCK_PLLSAI1R:
		return getPLLSAI1Freq(PLL_OUTPUT_R)
	case CLOCK_PLLSAI2P:
		return getPLLSAI2Freq()
	case CLOCK_PCLK1:
		return getPCLK1Freq()
	case CLOCK_PCLK2:
		return getPCLK2Freq()
	}

	// Peripheral Clocks
	src, ok := peripheralClockSources[clock]
	if ok {
		var i int

		if src.Reg != nil {
			i = int((src.Reg.Get() & src.Mask) >> src.Shift)
		}

		if i < len(src.Sources) {
			return GetClockFreq(src.Sources[i])
		}
	}

	return 0
}

func getMSIFreq() uint32 {
	if !stm32.RCC.CR.HasBits(stm32.RCC_CR_MSIRDY) {
		return 0
	}

	var r uint32
	if stm32.RCC.CR.HasBits(stm32.RCC_CR_MSIRGSEL) {
		r = (stm32.RCC.CR.Get() & stm32.RCC_CR_MSIRANGE_Msk) >> stm32.RCC_CR_MSIRANGE_Pos
	} else {
		r = (stm32.RCC.CSR.Get() & (0xF << 8)) >> 0x8
	}

	switch r {
	case 0:
		return 100000 // 100 KHz
	case 1:
		return 200000 // 200 KHz
	case 2:
		return 400000 // 400 KHz
	case 3:
		return 800000 // 800 KHz
	case 4:
		return 1000000 // 1 MHz
	case 5:
		return 2000000 // 2 MHz
	case 6:
		return 4000000 // 4 MHz
	case 7:
		return 8000000 // 8 MHz
	case 8:
		return 16000000 // 16 MHz
	case 9:
		return 24000000 // 24 MHz
	case 10:
		return 32000000 // 32 MHz
	case 11:
		return 48000000 // 48 MHz
	}

	return 0
}

func getPLLFreq(output int) uint32 {
	if !stm32.RCC.CR.HasBits(stm32.RCC_CR_PLLRDY) {
		return 0
	}

	pllSrc := stm32.RCC.PLLCFGR.Get() & stm32.RCC_PLLCFGR_PLLSRC_Msk

	var pllvco uint32
	switch pllSrc {
	case 1: // MSI
		pllvco = getMSIFreq()
	case 2: // HSI
		pllvco = getHSIFreq()
	case 3: // HSE
		pllvco = getHSEFreq()
	}

	pllm := ((stm32.RCC.PLLCFGR.Get() & stm32.RCC_PLLCFGR_PLLM_Msk) >> stm32.RCC_PLLCFGR_PLLM_Pos) + 1
	plln := (stm32.RCC.PLLCFGR.Get() & stm32.RCC_PLLCFGR_PLLN_Msk) >> stm32.RCC_PLLCFGR_PLLN_Pos

	var outdiv uint32

	switch output {
	case PLL_OUTPUT_P:
		outdiv = (stm32.RCC.PLLCFGR.Get() & stm32.RCC_PLLCFGR_PLLPDIV_Msk) >> stm32.RCC_PLLCFGR_PLLPDIV_Pos
		if outdiv == 0 {
			if stm32.RCC.PLLCFGR.Get()&stm32.RCC_PLLCFGR_PLLP != 0 {
				outdiv = 17
			} else {
				outdiv = 7
			}
		}

	case PLL_OUTPUT_Q:
		outdiv = (((stm32.RCC.PLLCFGR.Get() & stm32.RCC_PLLCFGR_PLLQ_Msk) >> stm32.RCC_PLLCFGR_PLLQ_Pos) + 1) << 1

	case PLL_OUTPUT_R:
		outdiv = (((stm32.RCC.PLLCFGR.Get() & stm32.RCC_PLLCFGR_PLLR_Msk) >> stm32.RCC_PLLCFGR_PLLR_Pos) + 1) * 2

	default:
		// Avoid divide-by-zero
		outdiv = 1
	}

	return ((pllvco / pllm) * plln) / outdiv
}

func getPLLSAI1Freq(output int) uint32 {
	if !stm32.RCC.CR.HasBits(stm32.RCC_CR_PLLSAI1RDY) {
		return 0
	}

	pllSrc := stm32.RCC.PLLSAI1CFGR.Get() & stm32.RCC_PLLSAI1CFGR_PLLSAI1SRC_Msk

	var pllvco uint32
	switch pllSrc {
	case 1: // MSI
		pllvco = getMSIFreq()
	case 2: // HSI
		pllvco = getHSIFreq()
	case 3: // HSE
		pllvco = getHSEFreq()
	}

	pllm := ((stm32.RCC.PLLSAI1CFGR.Get() & stm32.RCC_PLLSAI1CFGR_PLLSAI1M_Msk) >> stm32.RCC_PLLSAI1CFGR_PLLSAI1M_Pos) + 1
	plln := (stm32.RCC.PLLSAI1CFGR.Get() & stm32.RCC_PLLSAI1CFGR_PLLSAI1N_Msk) >> stm32.RCC_PLLSAI1CFGR_PLLSAI1N_Pos

	var outdiv uint32 = 1

	switch output {
	case PLL_OUTPUT_P:
		if stm32.RCC.PLLSAI1CFGR.HasBits(stm32.RCC_PLLSAI1CFGR_PLLSAI1PEN) {
			outdiv = (stm32.RCC.PLLSAI1CFGR.Get() & stm32.RCC_PLLSAI1CFGR_PLLSAI1PDIV_Msk) >> stm32.RCC_PLLSAI1CFGR_PLLSAI1PDIV_Pos
			if outdiv == 0 {
				if stm32.RCC.PLLSAI1CFGR.Get()&stm32.RCC_PLLSAI1CFGR_PLLSAI1P != 0 {
					outdiv = 17
				} else {
					outdiv = 7
				}
			}
		}

	case PLL_OUTPUT_Q:
		if stm32.RCC.PLLSAI1CFGR.HasBits(stm32.RCC_PLLSAI1CFGR_PLLSAI1QEN) {
			outdiv = (((stm32.RCC.PLLSAI1CFGR.Get() & stm32.RCC_PLLSAI1CFGR_PLLSAI1Q_Msk) >> stm32.RCC_PLLSAI1CFGR_PLLSAI1Q_Pos) + 1) << 1
		}

	case PLL_OUTPUT_R:
		if stm32.RCC.PLLSAI1CFGR.HasBits(stm32.RCC_PLLSAI1CFGR_PLLSAI1REN) {
			outdiv = (((stm32.RCC.PLLSAI1CFGR.Get() & stm32.RCC_PLLSAI1CFGR_PLLSAI1R_Msk) >> stm32.RCC_PLLSAI1CFGR_PLLSAI1R_Pos) + 1) * 2
		}
	}

	return ((pllvco / pllm) * plln) / outdiv
}

func getPLLSAI2Freq() uint32 {
	if !stm32.RCC.CR.HasBits(stm32.RCC_CR_PLLSAI2RDY) {
		return 0
	}

	pllSrc := stm32.RCC.PLLSAI2CFGR.Get() & stm32.RCC_PLLSAI2CFGR_PLLSAI2SRC_Msk

	var pllvco uint32
	switch pllSrc {
	case 1: // MSI
		pllvco = getMSIFreq()
	case 2: // HSI
		pllvco = getHSIFreq()
	case 3: // HSE
		pllvco = getHSEFreq()
	}

	pllm := ((stm32.RCC.PLLSAI2CFGR.Get() & stm32.RCC_PLLSAI2CFGR_PLLSAI2M_Msk) >> stm32.RCC_PLLSAI2CFGR_PLLSAI2M_Pos) + 1
	plln := (stm32.RCC.PLLSAI2CFGR.Get() & stm32.RCC_PLLSAI2CFGR_PLLSAI2N_Msk) >> stm32.RCC_PLLSAI2CFGR_PLLSAI2N_Pos

	var outdiv uint32 = 1

	if stm32.RCC.PLLSAI2CFGR.HasBits(stm32.RCC_PLLSAI2CFGR_PLLSAI2PEN) {
		outdiv = (stm32.RCC.PLLSAI2CFGR.Get() & stm32.RCC_PLLSAI2CFGR_PLLSAI2PDIV_Msk) >> stm32.RCC_PLLSAI2CFGR_PLLSAI2PDIV_Pos
		if outdiv == 0 {
			if stm32.RCC.PLLSAI2CFGR.Get()&stm32.RCC_PLLSAI2CFGR_PLLSAI2P != 0 {
				outdiv = 17
			} else {
				outdiv = 7
			}
		}
	}

	return ((pllvco / pllm) * plln) / outdiv
}

func getHCLKFreq() uint32 {
	sysClk := getSYSCLKFreq()

	i := (stm32.RCC.CFGR.Get() & stm32.RCC_CFGR_HPRE_Msk) >> stm32.RCC_CFGR_HPRE_Pos

	table := []uint32{0, 0, 0, 0, 0, 0, 0, 0, 1, 2, 3, 4, 6, 7, 8, 9}
	return sysClk >> table[i]
}

func getPCLK1Freq() uint32 {
	i := (stm32.RCC.CFGR.Get() & stm32.RCC_CFGR_PPRE1_Msk) >> stm32.RCC_CFGR_PPRE1_Pos

	var shift uint32
	if i > 3 {
		shift = i - 3
	}

	return getHCLKFreq() >> shift
}

func getPCLK2Freq() uint32 {
	i := (stm32.RCC.CFGR.Get() & stm32.RCC_CFGR_PPRE2_Msk) >> stm32.RCC_CFGR_PPRE2_Pos

	var shift uint32
	if i > 3 {
		shift = i - 3
	}

	return getHCLKFreq() >> shift
}

func getSYSCLKFreq() uint32 {
	src := stm32.RCC.CFGR.Get() & stm32.RCC_CFGR_SWS_Msk

	switch src {
	case 0: // MSI
		return getMSIFreq()
	case 1 << stm32.RCC_CFGR_SWS_Pos: // HSI
		return HSI_FREQ
	case 2 << stm32.RCC_CFGR_SWS_Pos: // HSE
		return HSE_FREQ
	case 3 << stm32.RCC_CFGR_SWS_Pos: // PLLCLK
		return getPLLFreq(PLL_OUTPUT_R)
	}

	return 0
}

func getHSI48Freq() uint32 {
	if !stm32.RCC.CRRCR.HasBits(stm32.RCC_CRRCR_HSI48RDY) {
		return 0
	}

	return HSI48_FREQ
}

func getHSIFreq() uint32 {
	if !stm32.RCC.CR.HasBits(stm32.RCC_CR_HSIRDY) {
		return 0
	}

	return HSI_FREQ
}

func getHSEFreq() uint32 {
	if !stm32.RCC.CR.HasBits(stm32.RCC_CR_HSERDY) {
		return 0
	}

	return HSE_FREQ
}

func getLSIFreq() uint32 {
	if !stm32.RCC.CSR.HasBits(stm32.RCC_CSR_LSIRDY) {
		return 0
	}

	return LSI_FREQ
}

func getLSEFreq() uint32 {
	if !stm32.RCC.BDCR.HasBits(stm32.RCC_BDCR_LSERDY) {
		return 0
	}

	return LSE_FREQ
}

func CPUFrequency() uint32 {
	return 110000000
}

//---------- UART related code

// Configure the UART.
func (uart *UART) configurePins(config UARTConfig) {
	if config.RX.getPort() == stm32.GPIOG || config.TX.getPort() == stm32.GPIOG {
		// Enable VDDIO2 power supply, which is an independant power supply for the PGx pins
		stm32.PWR.CR2.SetBits(stm32.PWR_CR2_IOSV)
	}

	// enable the alternate functions on the TX and RX pins
	config.TX.ConfigureAltFunc(PinConfig{Mode: PinModeUARTTX}, uart.AltFuncSelector)
	config.RX.ConfigureAltFunc(PinConfig{Mode: PinModeUARTRX}, uart.AltFuncSelector)
}

// UART baudrate calc based on the bus and clockspeed
// NOTE: keep this in sync with the runtime/runtime_stm32l5x2.go clock init code
func (uart *UART) getBaudRateDivisor(baudRate uint32) uint32 {
	return 256 * (CPUFrequency() / baudRate)
}

// Register names vary by ST processor, these are for STM L5
func (uart *UART) setRegisters() {
	uart.rxReg = &uart.Bus.RDR
	uart.txReg = &uart.Bus.TDR
	uart.statusReg = &uart.Bus.ISR
	uart.txEmptyFlag = stm32.USART_ISR_TXE
}
