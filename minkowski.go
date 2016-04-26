package qtr

import (
	"fmt"
	"math"
	"strings"
)

var symbM = [4]string{"", "s", "t", "u"}

// An M represents a Minkowski quaternion (also known as a hyperbolic
// quaternion) as an ordered list of four float64 values.
type M [4]float64

// String returns the string representation of an M value. If
// z corresponds to the Minkowski quaternion a + bs + ct + du, then the string
// is "(a+bs+ct+du)", similar to complex128 values.
func (z *M) String() string {
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
		a[j+1] = symbM[i]
		i++
	}
	a[8] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *M) Equals(y *M) bool {
	for i, v := range y {
		if notEquals(v, z[i]) {
			return false
		}
	}
	return true
}

// Copy copies x onto z, and returns z.
func (z *M) Copy(x *M) *M {
	for i, v := range x {
		z[i] = v
	}
	return z
}

// NewM returns a pointer to an M value made from four given float64 values.
func NewM(a, b, c, d float64) *M {
	z := new(M)
	z[0] = a
	z[1] = b
	z[2] = c
	z[3] = d
	return z
}

// IsMInf returns true if any of the components of z are infinite.
func (z *M) IsMInf() bool {
	for _, v := range z {
		if math.IsInf(v, 0) {
			return true
		}
	}
	return false
}

// MInf returns a pointer to a Minkowski quaternionic infinity value.
func MInf(a, b, c, d int) *M {
	return NewM(math.Inf(a), math.Inf(b), math.Inf(c), math.Inf(d))
}

// IsMNaN returns true if any component of z is NaN and neither is an infinity.
func (z *M) IsMNaN() bool {
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

// MNaN returns a pointer to a Minkowski quaternionic NaN value.
func MNaN() *M {
	nan := math.NaN()
	return NewM(nan, nan, nan, nan)
}

// Scal sets z equal to y scaled by a, and returns z.
func (z *M) Scal(y *M, a float64) *M {
	for i, v := range y {
		z[i] = a * v
	}
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *M) Neg(y *M) *M {
	return z.Scal(y, -1)
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *M) Conj(y *M) *M {
	z[0] = y[0]
	for i, v := range y[1:] {
		z[i+1] = -v
	}
	return z
}

// Add sets z equal to the sum of x and y, and returns z.
func (z *M) Add(x, y *M) *M {
	for i, v := range x {
		z[i] = v + y[i]
	}
	return z
}

// Sub sets z equal to the difference of x and y, and returns z.
func (z *M) Sub(x, y *M) *M {
	for i, v := range x {
		z[i] = v - y[i]
	}
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
//
// The multiplication rule for the basis elements s := M{0, 1, 0, 0},
// t := M{0, 0, 1, 0}, and u := M{0, 0, 0, 1} is:
// 		Mul(s, s) = Mul(t, t) = Mul(u, u) = M{1, 0, 0, 0}
// 		Mul(s, t) = -Mul(t, s) = u
// 		Mul(t, u) = -Mul(u, t) = s
// 		Mul(u, s) = -Mul(s, u) = t
func (z *M) Mul(x, y *M) *M {
	p := new(M).Copy(x)
	q := new(M).Copy(y)
	z[0] = (p[0] * q[0]) + (p[1] * q[1]) + (p[2] * q[2]) + (p[3] * q[3])
	z[1] = (p[0] * q[1]) + (p[1] * q[0]) + (p[2] * q[3]) - (p[3] * q[2])
	z[2] = (p[0] * q[2]) - (p[1] * q[3]) + (p[2] * q[0]) + (p[3] * q[1])
	z[3] = (p[0] * q[3]) + (p[1] * q[2]) - (p[2] * q[1]) + (p[3] * q[0])
	return z
}

// Commutator sets z equal to the commutator of x and y, and returns z.
func (z *M) Commutator(x, y *M) *M {
	return z.Sub(new(M).Mul(x, y), new(M).Mul(y, x))
}

// Quad returns the quadrance of z, which can be either positive, negative or
// zero.
func (z *M) Quad() float64 {
	return (new(M).Mul(z, new(M).Conj(z)))[0]
}

// IsZeroDiv returns true if z is a zero divisor (i.e. it has zero quadrance).
func (z *M) IsZeroDiv() bool {
	return !notEquals(z.Quad(), 0)
}

// Inv sets z equal to the inverse of x, and returns z. If x is a zero divisor,
// then Inv panics.
func (z *M) Inv(x *M) *M {
	if x.IsZeroDiv() {
		panic("inverse of zero divisor")
	}
	return z.Scal(new(M).Conj(x), 1/x.Quad())
}

// Quo sets z equal to the quotient of x and y, and returns z. If y is a zero
// divisor, then Quo panics.
func (z *M) Quo(x, y *M) *M {
	if y.IsZeroDiv() {
		panic("denominator is zero divisor")
	}
	return z.Scal(new(M).Mul(x, new(M).Conj(y)), 1/y.Quad())
}

// IsIndempotent returns true if z is an indempotent (i.e. if z = z*z).
func (z *M) IsIndempotent() bool {
	return z.Equals(new(M).Mul(z, z))
}
