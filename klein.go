package qtr

import (
	"fmt"
	"math"
	"strings"
)

// A Klein represents a Klein quaternion (also known as a split-quaternion) as
// an ordered array of four float64 values.
type Klein [4]float64

var (
	symbK = [4]string{"", "i", "t", "u"}

	zeroK = &Klein{0, 0, 0, 0}
	oneK  = &Klein{1, 0, 0, 0}
	iK    = &Klein{0, 1, 0, 0}
	tK    = &Klein{0, 0, 1, 0}
	uK    = &Klein{0, 0, 0, 1}
)

// String returns the string representation of a Klein value. If z corresponds
// to the Klein quaternion a + bi + ct + du, then the string is "(a+bi+ct+du)",
// similar to complex128 values.
func (z *Klein) String() string {
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
func (z *Klein) Equals(y *Klein) bool {
	for i, v := range y {
		if notEquals(v, z[i]) {
			return false
		}
	}
	return true
}

// Copy copies y onto z, and returns z.
func (z *Klein) Copy(y *Klein) *Klein {
	for i, v := range y {
		z[i] = v
	}
	return z
}

// NewKlein returns a pointer to a Klein value made from four given float64
// values.
func NewKlein(a, b, c, d float64) *Klein {
	z := new(Klein)
	z[0] = a
	z[1] = b
	z[2] = c
	z[3] = d
	return z
}

// IsKleinInf returns true if any of the components of z are infinite.
func (z *Klein) IsKleinInf() bool {
	for _, v := range z {
		if math.IsInf(v, 0) {
			return true
		}
	}
	return false
}

// KleinInf returns a pointer to a Klein quaternionic infinity value.
func KleinInf(a, b, c, d int) *Klein {
	return NewKlein(math.Inf(a), math.Inf(b), math.Inf(c), math.Inf(d))
}

// IsKleinNaN returns true if any component of z is NaN and neither is an
// infinity.
func (z *Klein) IsKleinNaN() bool {
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

// KleinNaN returns a pointer to a Klein quaternionic NaN value.
func KleinNaN() *Klein {
	nan := math.NaN()
	return NewKlein(nan, nan, nan, nan)
}

// Scal sets z equal to y scaled by a, and returns z.
func (z *Klein) Scal(y *Klein, a float64) *Klein {
	for i, v := range y {
		z[i] = a * v
	}
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Klein) Neg(y *Klein) *Klein {
	return z.Scal(y, -1)
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Klein) Conj(y *Klein) *Klein {
	z[0] = y[0]
	for i, v := range y[1:] {
		z[i+1] = -v
	}
	return z
}

// Add sets z equal to the sum of x and y, and returns z.
func (z *Klein) Add(x, y *Klein) *Klein {
	for i, v := range x {
		z[i] = v + y[i]
	}
	return z
}

// Sub sets z equal to the difference of x and y, and returns z.
func (z *Klein) Sub(x, y *Klein) *Klein {
	for i, v := range x {
		z[i] = v - y[i]
	}
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
//
// The multiplication rule for the basis elements i := Klein{0, 1, 0, 0},
// t := Klein{0, 0, 1, 0}, and u := Klein{0, 0, 0, 1} is:
// 		Mul(i, i) = Klein{-1, 0, 0, 0}
// 		Mul(t, t) = Mul(u, u) = Klein{1, 0, 0, 0}
// 		Mul(i, t) = -Mul(t, i) = +u
// 		Mul(t, u) = -Mul(u, t) = -i
// 		Mul(u, i) = -Mul(i, u) = +t
func (z *Klein) Mul(x, y *Klein) *Klein {
	p := new(Klein).Copy(x)
	q := new(Klein).Copy(y)
	z[0] = (p[0] * q[0]) - (p[1] * q[1]) + (p[2] * q[2]) + (p[3] * q[3])
	z[1] = (p[0] * q[1]) + (p[1] * q[0]) - (p[2] * q[3]) + (p[3] * q[2])
	z[2] = (p[0] * q[2]) - (p[1] * q[3]) + (p[2] * q[0]) + (p[3] * q[1])
	z[3] = (p[0] * q[3]) + (p[1] * q[2]) - (p[2] * q[1]) + (p[3] * q[0])
	return z
}

// Commutator sets z equal to the commutator of x and y, and returns z.
func (z *Klein) Commutator(x, y *Klein) *Klein {
	return z.Sub(new(Klein).Mul(x, y), new(Klein).Mul(y, x))
}

// Quad returns the quadrance of z, which can be either positive, negative or
// zero.
func (z *Klein) Quad() float64 {
	return (new(Klein).Mul(z, new(Klein).Conj(z)))[0]
}

// IsZeroDiv returns true if z is a zero divisor (i.e. it has zero quadrance).
func (z *Klein) IsZeroDiv() bool {
	return !notEquals(z.Quad(), 0)
}

// Inv sets z equal to the inverse of x, and returns z. If x is a zero divisor,
// then Inv panics.
func (z *Klein) Inv(x *Klein) *Klein {
	if x.IsZeroDiv() {
		panic("inverse of zero divisor")
	}
	return z.Scal(new(Klein).Conj(x), 1/x.Quad())
}

// Quo sets z equal to the quotient of x and y, and returns z. If y is a zero
// divisor, then Quo panics.
func (z *Klein) Quo(x, y *Klein) *Klein {
	if y.IsZeroDiv() {
		panic("denominator is zero divisor")
	}
	return z.Scal(new(Klein).Mul(x, new(Klein).Conj(y)), 1/y.Quad())
}

// IsIndempotent returns true if z is an indempotent (i.e. if z = z*z).
func (z *Klein) IsIndempotent() bool {
	return z.Equals(new(Klein).Mul(z, z))
}

// IsNilpotent returns true if z raised to the nth power vanishes.
func (z *Klein) IsNilpotent(n int) bool {
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
