// Package mathg defines generic wrapper functions supplied by the standard math package.
//
// The following functions are not defined since either they are intended to be used for specific types.
// Float32bits, Float32frombits, Float64bits, Float64frombits, Inf, IsInf, NaN, Nextafter, Nextafter32, Pow10
package mathg

import (
	"math"

	"golang.org/x/exp/constraints"
)

// A Number type represents a Float or Integer.
type Number interface {
	constraints.Float | constraints.Integer
}

// Abs is a wrapper of math.Abs.
func Abs[T Number](x T) T {
	return T(math.Abs(float64(x)))
}

// Acos is a wrapper of math.Acos.
func Acos[T Number](x T) float64 {
	return math.Acos(float64(x))
}

// Acosh is a wrapper of math.Acosh.
func Acosh[T Number](x T) float64 {
	return math.Acosh(float64(x))
}

// Asin is a wrapper of math.Asin.
func Asin[T Number](x T) float64 {
	return math.Asin(float64(x))
}

// Asinh is a wrapper of math.Asinh.
func Asinh[T Number](x T) float64 {
	return math.Asinh(float64(x))
}

// Atan is a wrapper of math.Atan.
func Atan[T Number](x T) float64 {
	return math.Atan(float64(x))
}

// Atan2 is a wrapper of math.Atan2.
func Atan2[T Number](y, x T) float64 {
	return math.Atan2(float64(y), float64(x))
}

// Atanh is a wrapper of math.Atanh.
func Atanh[T Number](x T) float64 {
	return math.Atanh(float64(x))
}

// Cbrt is a wrapper of math.Cbrt.
func Cbrt[T Number](x T) float64 {
	return math.Cbrt(float64(x))
}

// Ceil is a wrapper of math.Ceil.
func Ceil[T Number](x T) T {
	return T(math.Ceil(float64(x)))
}

// Copysign is a wrapper of math.Copysign.
func Copysign[T Number](f, sign T) T {
	return T(math.Copysign(float64(f), float64(sign)))
}

// Cos is a wrapper of math.Cos.
func Cos[T Number](x T) float64 {
	return math.Cos(float64(x))
}

// Cosh is a wrapper of math.Cosh.
func Cosh[T Number](x T) float64 {
	return math.Cosh(float64(x))
}

// Dim is a wrapper of math.Dim.
func Dim[T Number](x, y T) T {
	return T(math.Dim(float64(x), float64(y)))
}

// Erf is a wrapper of math.Erf.
func Erf[T Number](x T) float64 {
	return math.Erf(float64(x))
}

// Erfc is a wrapper of math.Erfc.
func Erfc[T Number](x T) float64 {
	return math.Erfc(float64(x))
}

// Erfcinv is a wrapper of math.Erfcinv.
func Erfcinv[T Number](x T) float64 {
	return math.Erfcinv(float64(x))
}

// Erfinv is a wrapper of math.Erfinv.
func Erfinv[T Number](x T) float64 {
	return math.Erfinv(float64(x))
}

// Exp is a wrapper of math.Exp.
func Exp[T Number](x T) float64 {
	return math.Exp(float64(x))
}

// Exp2 is a wrapper of math.Exp2.
func Exp2[T Number](x T) float64 {
	return math.Exp2(float64(x))
}

// Expm1 is a wrapper of math.Expm1.
func Expm1[T Number](x T) float64 {
	return math.Expm1(float64(x))
}

// FMA is a wrapper of math.FMA.
func FMA[T Number](x, y, z T) T {
	return T(math.FMA(float64(x), float64(y), float64(z)))
}

// Floor is a wrapper of math.Floor.
func Floor[T Number](x T) T {
	return T(math.Floor(float64(x)))
}

// Frexp is a wrapper of math.Frexp.
func Frexp[T Number](x T) (frac float64, exp int) {
	return math.Frexp(float64(x))
}

// Gamma is a wrapper of math.Gamma.
func Gamma[T Number](x T) float64 {
	return math.Gamma(float64(x))
}

// Hypot is a wrapper of math.Hypot.
func Hypot[T Number](p, q T) float64 {
	return math.Hypot(float64(p), float64(q))
}

// Ilogb is a wrapper of math.Ilogb.
func Ilogb[T Number](x T) int {
	return math.Ilogb(float64(x))
}

// IsNaN is a wrapper of math.IsNaN.
func IsNaN[T Number](x T) (is bool) {
	return math.IsNaN(float64(x))
}

// J0 is a wrapper of math.J0.
func J0[T Number](x T) float64 {
	return math.J0(float64(x))
}

// J1 is a wrapper of math.J1.
func J1[T Number](x T) float64 {
	return math.J1(float64(x))
}

// Jn is a wrapper of math.Jn.
func Jn[T Number](n int, x T) float64 {
	return math.Jn(n, float64(x))
}

// Ldexp is a wrapper of math.Ldexp.
func Ldexp[T Number](frac T, exp int) float64 {
	return math.Ldexp(float64(frac), exp)
}

// Lgamma is a wrapper of math.Lgamma.
func Lgamma[T Number](x T) (lgamma float64, sign int) {
	return math.Lgamma(float64(x))
}

// Log is a wrapper of math.Log.
func Log[T Number](x T) float64 {
	return math.Log(float64(x))
}

// Log10 is a wrapper of math.Log10.
func Log10[T Number](x T) float64 {
	return math.Log10(float64(x))
}

// Log1p is a wrapper of math.Log1p.
func Log1p[T Number](x T) float64 {
	return math.Log1p(float64(x))
}

// Log2 is a wrapper of math.Log2.
func Log2[T Number](x T) float64 {
	return math.Log2(float64(x))
}

// Logb is a wrapper of math.Logb.
func Logb[T Number](x T) float64 {
	return math.Logb(float64(x))
}

// Max is a wrapper of math.Max.
func Max[T Number](x, y T) T {
	return T(math.Max(float64(x), float64(y)))
}

// Min is a wrapper of math.Min.
func Min[T Number](x, y T) T {
	return T(math.Min(float64(x), float64(y)))
}

// Mod is a wrapper of math.Mod.
func Mod[T Number](x, y T) float64 {
	return math.Mod(float64(x), float64(y))
}

// Modf is a wrapper of math.Modf.
func Modf[T Number](x T) (int float64, frac float64) {
	return math.Modf(float64(x))
}

// Pow is a wrapper of math.Pow.
func Pow[T Number](x, y T) float64 {
	return math.Pow(float64(x), float64(y))
}

// Remainder is a wrapper of math.Remainder.
func Remainder[T Number](x, y T) float64 {
	return math.Remainder(float64(x), float64(y))
}

// Round is a wrapper of math.Round.
func Round[T Number](x T) T {
	return T(math.Round(float64(x)))
}

// RoundToEven is a wrapper of math.RoundToEven.
func RoundToEven[T Number](x T) T {
	return T(math.RoundToEven(float64(x)))
}

// Signbit is a wrapper of math.Signbit.
func Signbit[T Number](x T) bool {
	return math.Signbit(float64(x))
}

// Sin is a wrapper of math.Sin.
func Sin[T Number](x T) float64 {
	return math.Sin(float64(x))
}

// Sincos is a wrapper of math.Sincos.
func Sincos[T Number](x T) (sin, cos float64) {
	return math.Sincos(float64(x))
}

// Sinh is a wrapper of math.Sinh.
func Sinh[T Number](x T) float64 {
	return math.Sinh(float64(x))
}

// Sqrt is a wrapper of math.Sqrt.
func Sqrt[T Number](x T) float64 {
	return math.Sqrt(float64(x))
}

// Tan is a wrapper of math.Tan.
func Tan[T Number](x T) float64 {
	return math.Tan(float64(x))
}

// Tanh is a wrapper of math.Tanh.
func Tanh[T Number](x T) float64 {
	return math.Tanh(float64(x))
}

// Trunc is a wrapper of math.Trunc.
func Trunc[T Number](x T) T {
	return T(math.Trunc(float64(x)))
}

// Y0 is a wrapper of math.Y0.
func Y0[T Number](x T) float64 {
	return math.Y0(float64(x))
}

// Y1 is a wrapper of math.Y1.
func Y1[T Number](x T) float64 {
	return math.Y1(float64(x))
}

// Yn is a wrapper of math.Yn.
func Yn[T Number](n int, x T) float64 {
	return math.Yn(n, float64(x))
}
