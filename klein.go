package qtr

import (
	"fmt"
	"math"
	"strings"
)

var symbK = [4]string{"", "i", "t", "u"}

// A K represents a Klein quaternion (also known as a split-quaternion) as an
// ordered array of four float64 values.
type K [4]float64

// String returns the string representation of a K value. If z corresponds to
// the Klein quaternion a + bi + ct + du, then the string is "(a+bi+ct+du)",
// similar to complex128 values.
func (z *K) String() string {
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
		a[j+1] = symbK[i]
		i++
	}
	a[8] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *K) Equals(y *K) bool {
	for i, v := range y {
		if notEquals(v, z[i]) {
			return false
		}
	}
	return true
}

// Copy copies y onto z, and returns z.
func (z *K) Copy(y *K) *K {
	for i, v := range y {
		z[i] = v
	}
	return z
}

// NewK returns a pointer to a K value made from four given float64 values.
func NewK(a, b, c, d float64) *K {
	z := new(K)
	z[0] = a
	z[1] = b
	z[2] = c
	z[3] = d
	return z
}

// IsKInf returns true if any of the components of z are infinite.
func (z *K) IsKInf() bool {
	for _, v := range z {
		if math.IsInf(v, 0) {
			return true
		}
	}
	return false
}

// KInf returns a pointer to a Klein quaternionic infinity value.
func KInf(a, b, c, d int) *K {
	return NewK(math.Inf(a), math.Inf(b), math.Inf(c), math.Inf(d))
}

// IsKNaN returns true if any component of z is NaN and neither is an infinity.
func (z *K) IsKNaN() bool {
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

// KNaN returns a pointer to a Klein quaternionic NaN value.
func KNaN() *K {
	nan := math.NaN()
	return NewK(nan, nan, nan, nan)
}

// Scal sets z equal to y scaled by a, and returns z.
func (z *K) Scal(y *K, a float64) *K {
	for i, v := range y {
		z[i] = a * v
	}
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *K) Neg(y *K) *K {
	return z.Scal(y, -1)
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *K) Conj(y *K) *K {
	z[0] = y[0]
	for i, v := range y[1:] {
		z[i+1] = -v
	}
	return z
}

// Add sets z equal to the sum of x and y, and returns z.
func (z *K) Add(x, y *K) *K {
	for i, v := range x {
		z[i] = v + y[i]
	}
	return z
}

// Sub sets z equal to the difference of x and y, and returns z.
func (z *K) Sub(x, y *K) *K {
	for i, v := range x {
		z[i] = v - y[i]
	}
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
//
// The multiplication rule for the basis elements i := K{0, 1, 0, 0},
// t := K{0, 0, 1, 0}, and u := K{0, 0, 0, 1} is:
// 		Mul(i, i) = K{-1, 0, 0, 0}
// 		Mul(t, t) = Mul(u, u) = K{1, 0, 0, 0}
// 		Mul(i, t) = -Mul(t, i) = u
// 		Mul(t, u) = -Mul(u, t) = -i
// 		Mul(u, i) = -Mul(i, u) = t
func (z *K) Mul(x, y *K) *K {
	p := new(K).Copy(x)
	q := new(K).Copy(y)
	z[0] = (p[0] * q[0]) - (p[1] * q[1]) + (p[2] * q[2]) + (p[3] * q[3])
	z[1] = (p[0] * q[1]) + (p[1] * q[0]) - (p[2] * q[3]) + (p[3] * q[2])
	z[2] = (p[0] * q[2]) - (p[1] * q[3]) + (p[2] * q[0]) + (p[3] * q[1])
	z[3] = (p[0] * q[3]) + (p[1] * q[2]) - (p[2] * q[1]) + (p[3] * q[0])
	return z
}

// Commutator sets z equal to the commutator of x and y, and returns z.
func (z *K) Commutator(x, y *K) *K {
	return z.Sub(new(K).Mul(x, y), new(K).Mul(y, x))
}

// Quad returns the quadrance of z, which can be either positive,
// negative or zero.
func (z *K) Quad() float64 {
	return (new(K).Mul(z, new(K).Conj(z)))[0]
}

// IsZeroDiv returns true if z is a zero divisor (i.e. it has zero
// quadrance).
func (z *K) IsZeroDiv() bool {
	return !notEquals(z.Quad(), 0)
}

// Inv sets z equal to the inverse of x, and returns z. If x is a zero
// divisor, then Inv panics.
func (z *K) Inv(x *K) *K {
	if x.IsZeroDiv() {
		panic("inverse of zero divisor")
	}
	return z.Scal(new(K).Conj(x), 1/x.Quad())
}

// Quo sets z equal to the quotient of x and y, and returns z. If y is a
// zero divisor, then Quo panics.
func (z *K) Quo(x, y *K) *K {
	if y.IsZeroDiv() {
		panic("denominator is zero divisor")
	}
	return z.Scal(new(K).Mul(x, new(K).Conj(y)), 1/y.Quad())
}

// IsIndempotent returns true if z is an indempotent (i.e. if z = z*z).
func (z *K) IsIndempotent() bool {
	return z.Equals(new(K).Mul(z, z))
}
