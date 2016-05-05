package quat

import (
	"fmt"
	"math"
	"math/cmplx"
	"strings"
)

// A Cockle represents a Cockle quaternion (also known as a split-quaternion) as
// an ordered array of two complex128 values.
type Cockle [2]complex128

var (
	symbCockle = [4]string{"", "i", "t", "u"}

	zeroK = &Cockle{0, 0}
	oneK  = &Cockle{1, 0}
	iK    = &Cockle{1i, 0}
	tK    = &Cockle{0, 1}
	uK    = &Cockle{0, 1i}
)

// String returns the string representation of a Cockle value. If z corresponds
// to the Cockle quaternion a + bi + ct + du, then the string is "(a+bi+ct+du)",
// similar to complex128 values.
func (z *Cockle) String() string {
	v := make([]float64, 4)
	v[0], v[1] = real(z[0]), imag(z[0])
	v[2], v[3] = real(z[1]), imag(z[1])
	a := make([]string, 9)
	a[0] = "("
	a[1] = fmt.Sprintf("%g", v[0])
	i := 1
	for j := 2; j < 8; j = j + 2 {
		switch {
		case math.Signbit(v[i]):
			a[j] = fmt.Sprintf("%g", v[i])
		case math.IsInf(v[i], +1):
			a[j] = "+Inf"
		default:
			a[j] = fmt.Sprintf("+%g", v[i])
		}
		a[j+1] = symbCockle[i]
		i++
	}
	a[8] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Cockle) Equals(y *Cockle) bool {
	if notEquals(real(z[0]), real(y[0])) || notEquals(imag(z[0]), imag(y[0])) {
		return false
	}
	if notEquals(real(z[1]), real(y[1])) || notEquals(imag(z[1]), imag(y[1])) {
		return false
	}
	return true
}

// Copy copies y onto z, and returns z.
func (z *Cockle) Copy(y *Cockle) *Cockle {
	z[0] = y[0]
	z[1] = y[1]
	return z
}

// NewCockle returns a pointer to a Cockle value made from four given float64
// values.
func NewCockle(a, b, c, d float64) *Cockle {
	z := new(Cockle)
	z[0] = complex(a, b)
	z[1] = complex(c, d)
	return z
}

// IsInf returns true if any of the components of z are infinite.
func (z *Cockle) IsInf() bool {
	if cmplx.IsInf(z[0]) || cmplx.IsInf(z[1]) {
		return true
	}
	return false
}

// CockleInf returns a pointer to a Cockle quaternionic infinity value.
func CockleInf(a, b, c, d int) *Cockle {
	z := new(Cockle)
	z[0] = complex(math.Inf(a), math.Inf(b))
	z[1] = complex(math.Inf(c), math.Inf(d))
	return z
}

// IsNaN returns true if any component of z is NaN and neither is an
// infinity.
func (z *Cockle) IsNaN() bool {
	if cmplx.IsInf(z[0]) || cmplx.IsInf(z[1]) {
		return false
	}
	if cmplx.IsNaN(z[0]) || cmplx.IsNaN(z[1]) {
		return true
	}
	return false
}

// CockleNaN returns a pointer to a Cockle quaternionic NaN value.
func CockleNaN() *Cockle {
	nan := cmplx.NaN()
	z := new(Cockle)
	z[0] = nan
	z[1] = nan
	return z
}

// Scal sets z equal to y scaled by a (with a being a complex128), and returns
// z.
//
// This is a special case of Mul:
// 		Scal(y, a) = Mul(y, Cockle{a, 0})
func (z *Cockle) Scal(y *Cockle, a complex128) *Cockle {
	z[0] = y[0] * a
	z[1] = y[1] * a
	return z
}

// Dil sets z equal to the dilation of y by a, and returns z.
//
// This is a special case of Mul:
// 		Dil(y, a) = Mul(y, Cockle{complex(a, 0), 0})
func (z *Cockle) Dil(y *Cockle, a float64) *Cockle {
	z[0] = y[0] * complex(a, 0)
	z[1] = y[1] * complex(a, 0)
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Cockle) Neg(y *Cockle) *Cockle {
	return z.Dil(y, -1)
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Cockle) Conj(y *Cockle) *Cockle {
	z[0] = cmplx.Conj(y[0])
	z[1] = -y[1]
	return z
}

// Add sets z equal to the sum of x and y, and returns z.
func (z *Cockle) Add(x, y *Cockle) *Cockle {
	z[0] = x[0] + y[0]
	z[1] = x[1] + y[1]
	return z
}

// Sub sets z equal to the difference of x and y, and returns z.
func (z *Cockle) Sub(x, y *Cockle) *Cockle {
	z[0] = x[0] - y[0]
	z[1] = x[1] - y[1]
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
//
// The multiplication rule for the basis elements i := Cockle{0, 1, 0, 0},
// t := Cockle{0, 0, 1, 0}, and u := Cockle{0, 0, 0, 1} is:
// 		Mul(i, i) = Cockle{-1, 0, 0, 0}
// 		Mul(t, t) = Mul(u, u) = Cockle{1, 0, 0, 0}
// 		Mul(i, t) = -Mul(t, i) = +u
// 		Mul(t, u) = -Mul(u, t) = -i
// 		Mul(u, i) = -Mul(i, u) = +t
func (z *Cockle) Mul(x, y *Cockle) *Cockle {
	p := new(Cockle).Copy(x)
	q := new(Cockle).Copy(y)
	z[0] = (p[0] * q[0]) + (cmplx.Conj(q[1]) * p[1])
	z[1] = (p[0] * q[1]) + (p[1] * cmplx.Conj(q[0]))
	return z
}

// Commutator sets z equal to the commutator of x and y, and returns z.
func (z *Cockle) Commutator(x, y *Cockle) *Cockle {
	return z.Sub(new(Cockle).Mul(x, y), new(Cockle).Mul(y, x))
}

// Quad returns the quadrance of z, which can be either positive, negative or
// zero.
func (z *Cockle) Quad() float64 {
	a, b := cmplx.Abs(z[0]), cmplx.Abs(z[1])
	return (a * a) - (b * b)
}

// IsZeroDiv returns true if z is a zero divisor (i.e. it has zero quadrance).
func (z *Cockle) IsZeroDiv() bool {
	return !notEquals(z.Quad(), 0)
}

// Inv sets z equal to the inverse of x, and returns z. If x is a zero divisor,
// then Inv panics.
func (z *Cockle) Inv(x *Cockle) *Cockle {
	if x.IsZeroDiv() {
		panic("inverse of zero divisor")
	}
	return z.Dil(new(Cockle).Conj(x), 1/x.Quad())
}

// Quo sets z equal to the quotient of x and y, and returns z. If y is a zero
// divisor, then Quo panics.
func (z *Cockle) Quo(x, y *Cockle) *Cockle {
	if y.IsZeroDiv() {
		panic("denominator is zero divisor")
	}
	return z.Dil(new(Cockle).Mul(x, new(Cockle).Conj(y)), 1/y.Quad())
}

// IsIndempotent returns true if z is an indempotent (i.e. if z = z*z).
func (z *Cockle) IsIndempotent() bool {
	return z.Equals(new(Cockle).Mul(z, z))
}

// IsNilpotent returns true if z raised to the nth power vanishes.
func (z *Cockle) IsNilpotent(n int) bool {
	if z.Equals(zeroK) {
		return true
	}
	p := oneK
	for i := 0; i < n; i++ {
		p.Mul(p, z)
		if p.Equals(zeroK) {
			return true
		}
	}
	return false
}

// RectCockle returns a Cockle value made from given curvilinear coordinates and
// quadrance sign.
func RectCockle(r, ξ, θ1, θ2 float64, sign int) *Cockle {
	z := new(Cockle)
	if sign > 0 {
		z[0] = complex(
			r*math.Cosh(ξ)*math.Cos(θ1),
			r*math.Cosh(ξ)*math.Sin(θ1),
		)
		z[1] = complex(
			r*math.Sinh(ξ)*math.Cos(θ2),
			r*math.Sinh(ξ)*math.Sin(θ2),
		)
		return z
	}
	if sign < 0 {
		z[0] = complex(
			r*math.Sinh(ξ)*math.Cos(θ1),
			r*math.Sinh(ξ)*math.Sin(θ1),
		)
		z[1] = complex(
			r*math.Cosh(ξ)*math.Cos(θ2),
			r*math.Cosh(ξ)*math.Sin(θ2),
		)
		return z
	}
	z[0] = cmplx.Rect(r, θ1)
	z[1] = cmplx.Rect(r, θ2)
	return z
}

// Curv returns the curvilinear coordinates of a Cockle value, along with the
// sign of the quadrance.
func (z *Cockle) Curv() (r, ξ, θ1, θ2 float64, sign int) {
	quad := z.Quad()
	θ1 = cmplx.Phase(z[0])
	θ2 = cmplx.Phase(z[1])
	if quad > 0 {
		r = math.Sqrt(quad)
		ξ = math.Atanh(cmplx.Abs(z[1]) / cmplx.Abs(z[0]))
		sign = +1
		return
	}
	if quad < 0 {
		r = math.Sqrt(-quad)
		ξ = math.Atanh(cmplx.Abs(z[0]) / cmplx.Abs(z[1]))
		sign = -1
		return
	}
	r = cmplx.Abs(z[0])
	ξ = math.NaN()
	sign = 0
	return
}
