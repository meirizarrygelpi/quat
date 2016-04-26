package qtr

import (
	"fmt"
	"math"
	"strings"
)

var symbH = [4]string{"", "i", "j", "k"}

// An H represents a Hamilton quaternion (i.e. a traditional quaternion) as
// an ordered array of four float64 values.
type H [4]float64

// String returns the string representation of an H value. If z corresponds to
// the Hamilton quaternion a + bi + cj + dk, then the string is "(a+bi+cj+dk)",
// similar to complex128 values.
func (z *H) String() string {
	a := make([]string, 9)
	a[0] = "("
	a[1] = fmt.Sprintf("%g", z[0])
	i := 1
	for j := 2; j < 8; j = j + 2 {
		switch {
		case math.Signbit(z[i]):
			a[j] = fmt.Sprintf("%g", z[i])
		case math.IsInf(z[i], +1):
			a[j] = "+Inf"
		default:
			a[j] = fmt.Sprintf("+%g", z[i])
		}
		a[j+1] = symbH[i]
		i++
	}
	a[8] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *H) Equals(y *H) bool {
	for i, v := range y {
		if notEquals(v, z[i]) {
			return false
		}
	}
	return true
}

// Copy copies y onto z, and returns z.
func (z *H) Copy(y *H) *H {
	for i, v := range y {
		z[i] = v
	}
	return z
}

// NewH returns a pointer to an H value made from four given float64
// values.
func NewH(a, b, c, d float64) *H {
	z := new(H)
	z[0] = a
	z[1] = b
	z[2] = c
	z[3] = d
	return z
}

// IsHInf returns true if any of the components of z are infinite.
func (z *H) IsHInf() bool {
	for _, v := range z {
		if math.IsInf(v, 0) {
			return true
		}
	}
	return false
}

// HInf returns a pointer to a Hamilton quaternionic infinity value.
func HInf(a, b, c, d int) *H {
	return NewH(math.Inf(a), math.Inf(b), math.Inf(c), math.Inf(d))
}

// IsHNaN returns true if any component of z is NaN and neither is an
// infinity.
func (z *H) IsHNaN() bool {
	for _, v := range z {
		if math.IsInf(v, 0) {
			return false
		}
	}
	for _, v := range z {
		if math.IsNaN(v) {
			return true
		}
	}
	return false
}

// HNaN returns a pointer to a Hamilton quaternionic NaN value.
func HNaN() *H {
	nan := math.NaN()
	return NewH(nan, nan, nan, nan)
}

// Scal sets z equal to y scaled by a, and returns z.
func (z *H) Scal(y *H, a float64) *H {
	for i, v := range y {
		z[i] = a * v
	}
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *H) Neg(y *H) *H {
	return z.Scal(y, -1)
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *H) Conj(y *H) *H {
	z[0] = y[0]
	for i, v := range y[1:] {
		z[i+1] = -v
	}
	return z
}

// Add sets z equal to the sum of x and y, and returns z.
func (z *H) Add(x, y *H) *H {
	for i, v := range x {
		z[i] = v + y[i]
	}
	return z
}

// Sub sets z equal to the difference of x and y, and returns z.
func (z *H) Sub(x, y *H) *H {
	for i, v := range x {
		z[i] = v - y[i]
	}
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
//
// The multiplication rule for the basis elements i := H{0, 1, 0, 0},
// j := H{0, 0, 1, 0}, and k := H{0, 0, 0, 1} is:
// 		Mul(i, i) = Mul(j, j) = Mul(k, k) = H{-1, 0, 0, 0}
// 		Mul(i, j) = -Mul(j, i) = k
// 		Mul(j, k) = -Mul(k, j) = i
// 		Mul(k, i) = -Mul(i, k) = j
func (z *H) Mul(x, y *H) *H {
	p := new(H).Copy(x)
	q := new(H).Copy(y)
	z[0] = (p[0] * q[0]) - (p[1] * q[1]) - (p[2] * q[2]) - (p[3] * q[3])
	z[1] = (p[0] * q[1]) + (p[1] * q[0]) + (p[2] * q[3]) - (p[3] * q[2])
	z[2] = (p[0] * q[2]) - (p[1] * q[3]) + (p[2] * q[0]) + (p[3] * q[1])
	z[3] = (p[0] * q[3]) + (p[1] * q[2]) - (p[2] * q[1]) + (p[3] * q[0])
	return z
}

// Commutator sets z equal to the commutator of x and y, and returns z.
func (z *H) Commutator(x, y *H) *H {
	return z.Sub(new(H).Mul(x, y), new(H).Mul(y, x))
}

// Quad returns the non-negative quadrance of z.
func (z *H) Quad() float64 {
	return (new(H).Mul(z, new(H).Conj(z)))[0]
}

// Inv sets z equal to the inverse of y, and returns z. If y is zero, then Inv
// panics.
func (z *H) Inv(y *H) *H {
	if y.Equals(NewH(0, 0, 0, 0)) {
		panic("inverse of zero")
	}
	return z.Scal(new(H).Conj(y), 1/y.Quad())
}

// Quo sets z equal to the quotient of x and y, and returns z. If y is zero,
// then Quo panics.
func (z *H) Quo(x, y *H) *H {
	if y.Equals(NewH(0, 0, 0, 0)) {
		panic("denominator is zero")
	}
	return z.Scal(new(H).Mul(x, new(H).Conj(y)), 1/y.Quad())
}
