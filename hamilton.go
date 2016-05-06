// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package quat

import (
	"fmt"
	"math"
	"math/cmplx"
	"strings"
)

var (
	symbHamilton = [4]string{"", "i", "j", "k"}

	zeroH = &Hamilton{0, 0}
	oneH  = &Hamilton{1, 0}
	iH    = &Hamilton{1i, 0}
	jH    = &Hamilton{0, 1}
	kH    = &Hamilton{0, 1i}
)

// A Hamilton represents a Hamilton quaternion (i.e. a traditional quaternion)
// as an ordered array of two complex128 values.
type Hamilton [2]complex128

// String returns the string representation of a Hamilton value. If z
// corresponds to the Hamilton quaternion a + bi + cj + dk, then the string is
// "(a+bi+cj+dk)", similar to complex128 values.
func (z *Hamilton) String() string {
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
		a[j+1] = symbHamilton[i]
		i++
	}
	a[8] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Hamilton) Equals(y *Hamilton) bool {
	if notEquals(real(z[0]), real(y[0])) || notEquals(imag(z[0]), imag(y[0])) {
		return false
	}
	if notEquals(real(z[1]), real(y[1])) || notEquals(imag(z[1]), imag(y[1])) {
		return false
	}
	return true
}

// Copy copies y onto z, and returns z.
func (z *Hamilton) Copy(y *Hamilton) *Hamilton {
	z[0] = y[0]
	z[1] = y[1]
	return z
}

// NewHamilton returns a pointer to a Hamilton value made from four given
// float64 values.
func NewHamilton(a, b, c, d float64) *Hamilton {
	z := new(Hamilton)
	z[0] = complex(a, b)
	z[1] = complex(c, d)
	return z
}

// IsInf returns true if any of the components of z are infinite.
func (z *Hamilton) IsInf() bool {
	if cmplx.IsInf(z[0]) || cmplx.IsInf(z[1]) {
		return true
	}
	return false
}

// HamiltonInf returns a pointer to a Hamilton quaternionic infinity value.
func HamiltonInf(a, b, c, d int) *Hamilton {
	z := new(Hamilton)
	z[0] = complex(math.Inf(a), math.Inf(b))
	z[1] = complex(math.Inf(c), math.Inf(d))
	return z
}

// IsNaN returns true if any component of z is NaN and neither is an
// infinity.
func (z *Hamilton) IsNaN() bool {
	if cmplx.IsInf(z[0]) || cmplx.IsInf(z[1]) {
		return false
	}
	if cmplx.IsNaN(z[0]) || cmplx.IsNaN(z[1]) {
		return true
	}
	return false
}

// HamiltonNaN returns a pointer to a Hamilton quaternionic NaN value.
func HamiltonNaN() *Hamilton {
	nan := cmplx.NaN()
	z := new(Hamilton)
	z[0] = nan
	z[1] = nan
	return z
}

// Scal sets z equal to y scaled by a (with a being a complex128), and returns
// z.
//
// This is a special case of Mul:
// 		Scal(y, a) = Mul(y, Hamilton{a, 0})
func (z *Hamilton) Scal(y *Hamilton, a complex128) *Hamilton {
	z[0] = y[0] * a
	z[1] = y[1] * a
	return z
}

// Dil sets z equal to the dilation of y by a, and returns z.
//
// This is a special case of Mul:
// 		Dil(y, a) = Mul(y, Hamilton{complex(a, 0), 0})
func (z *Hamilton) Dil(y *Hamilton, a float64) *Hamilton {
	z[0] = y[0] * complex(a, 0)
	z[1] = y[1] * complex(a, 0)
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Hamilton) Neg(y *Hamilton) *Hamilton {
	return z.Dil(y, -1)
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Hamilton) Conj(y *Hamilton) *Hamilton {
	z[0] = cmplx.Conj(y[0])
	z[1] = -y[1]
	return z
}

// Add sets z equal to the sum of x and y, and returns z.
func (z *Hamilton) Add(x, y *Hamilton) *Hamilton {
	z[0] = x[0] + y[0]
	z[1] = x[1] + y[1]
	return z
}

// Sub sets z equal to the difference of x and y, and returns z.
func (z *Hamilton) Sub(x, y *Hamilton) *Hamilton {
	z[0] = x[0] - y[0]
	z[1] = x[1] - y[1]
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
//
// The multiplication rule for the basis elements i := Hamilton{0, 1, 0, 0},
// j := Hamilton{0, 0, 1, 0}, and k := Hamilton{0, 0, 0, 1} is:
// 		Mul(i, i) = Mul(j, j) = Mul(k, k) = Hamilton{-1, 0, 0, 0}
// 		Mul(i, j) = -Mul(j, i) = +k
// 		Mul(j, k) = -Mul(k, j) = +i
// 		Mul(k, i) = -Mul(i, k) = +j
func (z *Hamilton) Mul(x, y *Hamilton) *Hamilton {
	p := new(Hamilton).Copy(x)
	q := new(Hamilton).Copy(y)
	z[0] = (p[0] * q[0]) - (cmplx.Conj(q[1]) * p[1])
	z[1] = (p[0] * q[1]) + (p[1] * cmplx.Conj(q[0]))
	return z
}

// Commutator sets z equal to the commutator of x and y, and returns z.
func (z *Hamilton) Commutator(x, y *Hamilton) *Hamilton {
	return z.Sub(new(Hamilton).Mul(x, y), new(Hamilton).Mul(y, x))
}

// Quad returns the non-negative quadrance of z.
func (z *Hamilton) Quad() float64 {
	a, b := cmplx.Abs(z[0]), cmplx.Abs(z[1])
	return (a * a) + (b * b)
}

// Inv sets z equal to the inverse of y, and returns z. If y is zero, then Inv
// panics.
func (z *Hamilton) Inv(y *Hamilton) *Hamilton {
	if y.Equals(zeroH) {
		panic("inverse of zero")
	}
	return z.Dil(new(Hamilton).Conj(y), 1/y.Quad())
}

// Quo sets z equal to the quotient of x and y, and returns z. If y is zero,
// then Quo panics.
func (z *Hamilton) Quo(x, y *Hamilton) *Hamilton {
	if y.Equals(zeroH) {
		panic("denominator is zero")
	}
	return z.Dil(new(Hamilton).Mul(x, new(Hamilton).Conj(y)), 1/y.Quad())
}

// RectHamilton returns a Hamilton value made from given curvilinear
// coordinates.
func RectHamilton(r, θ1, θ2, θ3 float64) *Hamilton {
	if notEquals(r, 0) {
		z := new(Hamilton)
		z[0] = complex(
			r*math.Cos(θ1),
			r*math.Sin(θ1)*math.Cos(θ2),
		)
		z[1] = complex(
			r*math.Sin(θ1)*math.Sin(θ2)*math.Cos(θ3),
			r*math.Sin(θ1)*math.Sin(θ2)*math.Sin(θ3),
		)
		return z
	}
	return zeroH
}

// Curv returns the curvilinear coordinates of a Hamilton value.
func (z *Hamilton) Curv() (r, θ1, θ2, θ3 float64) {
	if z.Equals(zeroH) {
		return 0, math.NaN(), math.NaN(), math.NaN()
	}
	h := cmplx.Abs(z[1])
	r = math.Sqrt(z.Quad())
	θ1 = math.Atan(math.Hypot(imag(z[0]), h) / real(z[0]))
	θ2 = math.Atan(h / imag(z[0]))
	θ3 = math.Atan2(imag(z[1]), real(z[1]))
	return
}
