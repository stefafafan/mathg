package mathg

import (
	"math"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Float | constraints.Integer
}

func Abs[T Number](x T) T {
	return T(math.Abs(float64(x)))
}

func Acos[T Number](x T) T {
	return T(math.Acos(float64(x)))
}

func Acosh[T Number](x T) T {
	return T(math.Acosh(float64(x)))
}

func Asin[T Number](x T) T {
	return T(math.Asin(float64(x)))
}

func Asinh[T Number](x T) T {
	return T(math.Asinh(float64(x)))
}

func Atan[T Number](x T) T {
	return T(math.Atan(float64(x)))
}

func Atan2[T Number](y, x T) T {
	return T(math.Atan2(float64(y), float64(x)))
}

func Atanh[T Number](x T) T {
	return T(math.Atanh(float64(x)))
}

func Cbrt[T Number](x T) T {
	return T(math.Cbrt(float64(x)))
}

func Ceil[T Number](x T) T {
	return T(math.Ceil(float64(x)))
}

func Copysign[T Number](f, sign T) T {
	return T(math.Copysign(float64(f), float64(sign)))
}

func Cos[T Number](x T) T {
	return T(math.Cos(float64(x)))
}

func Cosh[T Number](x T) T {
	return T(math.Cosh(float64(x)))
}

func Dim[T Number](x, y T) T {
	return T(math.Dim(float64(x), float64(y)))
}

func Erf[T Number](x T) T {
	return T(math.Erf(float64(x)))
}

func Erfc[T Number](x T) T {
	return T(math.Erfc(float64(x)))
}

func Erfcinv[T Number](x T) T {
	return T(math.Erfcinv(float64(x)))
}

func Erfinv[T Number](x T) T {
	return T(math.Erfinv(float64(x)))
}

func Exp[T Number](x T) T {
	return T(math.Exp(float64(x)))
}

func Exp2[T Number](x T) T {
	return T(math.Exp2(float64(x)))
}
