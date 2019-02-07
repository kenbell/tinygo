package main

import "math"

func main() {
	for _, n := range []float64{0.3, 1.5, 2.6, -1.1, -3.1, -3.8} {
		println("n:", n)
		println("  asin:     ", math.Asin(n))
		println("  asinh:    ", math.Asinh(n))
		println("  acos:     ", math.Acos(n))
		println("  acosh:    ", math.Acosh(n))
		println("  atan:     ", math.Atan(n))
		println("  atanh:    ", math.Atanh(n))
		println("  atan2:    ", math.Atan2(n, 0.2))
		println("  cbrt:     ", math.Cbrt(n))
		println("  ceil:     ", math.Ceil(n))
		println("  cos:      ", math.Cos(n))
		println("  cosh:     ", math.Cosh(n))
		println("  erf:      ", math.Erf(n))
		println("  erfc:     ", math.Erfc(n))
		println("  exp:      ", math.Exp(n))
		println("  expm1:    ", math.Expm1(n))
		println("  exp2:     ", math.Exp2(n))
		println("  floor:    ", math.Floor(n))
		f, e := math.Frexp(n)
		println("  frexp:    ", f, e)
		println("  hypot:    ", math.Hypot(n, n*2))
		println("  ldexp:    ", math.Ldexp(n, 2))
		println("  log:      ", math.Log(n))
		println("  log1p:    ", math.Log1p(n))
		println("  log10:    ", math.Log10(n))
		println("  log2:     ", math.Log2(n))
		println("  max:      ", math.Max(n, n+1))
		println("  min:      ", math.Min(n, n+1))
		println("  mod:      ", math.Mod(n, n+1))
		i, f := math.Modf(n)
		println("  modf:     ", i, f)
		println("  pow:      ", math.Pow(n, n))
		println("  remainder:", math.Remainder(n, n+0.2))
		println("  sin:      ", math.Sin(n))
		println("  sinh:     ", math.Sinh(n))
		println("  tan:      ", math.Tan(n))
		println("  tanh:     ", math.Tanh(n))
		println("  trunc:    ", math.Trunc(n))
	}
}