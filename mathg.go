package mathg

import (
	"math"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Float | constraints.Integer
}

// Defines generic wrappers of the standard math package,
// expect for the following functions which probably don't need a generic version:
// Float32bits, Float32frombits, Float64bits, Float64frombits, Inf, IsInf, NaN, Nextafter, Nextafter32, Pow10

func Abs[T Number](x T) T {
	return T(math.Abs(float64(x)))
}

func Acos[T Number](x T) float64 {
	return math.Acos(float64(x))
}

func Acosh[T Number](x T) float64 {
	return math.Acosh(float64(x))
}

func Asin[T Number](x T) float64 {
	return math.Asin(float64(x))
}

func Asinh[T Number](x T) float64 {
	return math.Asinh(float64(x))
}

func Atan[T Number](x T) float64 {
	return math.Atan(float64(x))
}

func Atan2[T Number](y, x T) float64 {
	return math.Atan2(float64(y), float64(x))
}

func Atanh[T Number](x T) float64 {
	return math.Atanh(float64(x))
}

func Cbrt[T Number](x T) float64 {
	return math.Cbrt(float64(x))
}

func Ceil[T Number](x T) T {
	return T(math.Ceil(float64(x)))
}

func Copysign[T Number](f, sign T) T {
	return T(math.Copysign(float64(f), float64(sign)))
}

func Cos[T Number](x T) float64 {
	return math.Cos(float64(x))
}

func Cosh[T Number](x T) float64 {
	return math.Cosh(float64(x))
}

func Dim[T Number](x, y T) T {
	return T(math.Dim(float64(x), float64(y)))
}

func Erf[T Number](x T) float64 {
	return math.Erf(float64(x))
}

func Erfc[T Number](x T) float64 {
	return math.Erfc(float64(x))
}

func Erfcinv[T Number](x T) float64 {
	return math.Erfcinv(float64(x))
}

func Erfinv[T Number](x T) float64 {
	return math.Erfinv(float64(x))
}

func Exp[T Number](x T) float64 {
	return math.Exp(float64(x))
}

func Exp2[T Number](x T) float64 {
	return math.Exp2(float64(x))
}

func Expm1[T Number](x T) float64 {
	return math.Expm1(float64(x))
}

func FMA[T Number](x, y, z T) T {
	return T(math.FMA(float64(x), float64(y), float64(z)))
}

func Floor[T Number](x T) T {
	return T(math.Floor(float64(x)))
}

func Frexp[T Number](x T) (frac float64, exp int) {
	return math.Frexp(float64(x))
}

func Gamma[T Number](x T) float64 {
	return math.Gamma(float64(x))
}

func Hypot[T Number](p, q T) float64 {
	return math.Hypot(float64(p), float64(q))
}

func Ilogb[T Number](x T) int {
	return math.Ilogb(float64(x))
}

func IsNaN[T Number](x T) (is bool) {
	return math.IsNaN(float64(x))
}

func J0[T Number](x T) float64 {
	return math.J0(float64(x))
}

func J1[T Number](x T) float64 {
	return math.J1(float64(x))
}

func Jn[T Number](n int, x T) float64 {
	return math.Jn(n, float64(x))
}

func Ldexp[T Number](frac T, exp int) float64 {
	return math.Ldexp(float64(frac), exp)
}

func Lgamma[T Number](x T) (lgamma float64, sign int) {
	return math.Lgamma(float64(x))
}

func Log[T Number](x T) float64 {
	return math.Log(float64(x))
}

func Log10[T Number](x T) float64 {
	return math.Log10(float64(x))
}

func Log1p[T Number](x T) float64 {
	return math.Log1p(float64(x))
}

func Log2[T Number](x T) float64 {
	return math.Log2(float64(x))
}

func Logb[T Number](x T) float64 {
	return math.Logb(float64(x))
}

func Max[T Number](x, y T) T {
	return T(math.Max(float64(x), float64(y)))
}

func Min[T Number](x, y T) T {
	return T(math.Min(float64(x), float64(y)))
}

func Mod[T Number](x, y T) float64 {
	return math.Mod(float64(x), float64(y))
}

func Modf[T Number](x T) (int float64, frac float64) {
	return math.Modf(float64(x))
}

func Pow[T Number](x, y T) float64 {
	return math.Pow(float64(x), float64(y))
}

func Remainder[T Number](x, y T) float64 {
	return math.Remainder(float64(x), float64(y))
}

func Round[T Number](x T) T {
	return T(math.Round(float64(x)))
}

func RoundToEven[T Number](x T) T {
	return T(math.RoundToEven(float64(x)))
}

func Signbit[T Number](x T) bool {
	return math.Signbit(float64(x))
}

func Sin[T Number](x T) float64 {
	return math.Sin(float64(x))
}

func Sincos[T Number](x T) (sin, cos float64) {
	return math.Sincos(float64(x))
}

func Sinh[T Number](x T) float64 {
	return math.Sinh(float64(x))
}

func Sqrt[T Number](x T) float64 {
	return math.Sqrt(float64(x))
}

func Tan[T Number](x T) float64 {
	return math.Tan(float64(x))
}

func Tanh[T Number](x T) float64 {
	return math.Tanh(float64(x))
}

func Trunc[T Number](x T) T {
	return T(math.Trunc(float64(x)))
}

func Y0[T Number](x T) float64 {
	return math.Y0(float64(x))
}

func Y1[T Number](x T) float64 {
	return math.Y1(float64(x))
}

func Yn[T Number](n int, x T) float64 {
	return math.Yn(n, float64(x))
}
