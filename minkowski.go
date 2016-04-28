package qtr

import (
	"fmt"
	"math"
	"strings"
)

var symbM = [4]string{"", "s", "t", "u"}

// A Minkowski represents a Minkowski quaternion (also known as a hyperbolic
// quaternion) as an ordered list of four float64 values.
type Minkowski [4]float64

// String returns the string representation of a Minkowski value. If z
// corresponds to the Minkowski quaternion a + bs + ct + du, then the string
// is "(a+bs+ct+du)", similar to complex128 values.
func (z *Minkowski) String() string {
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
func (z *Minkowski) Equals(y *Minkowski) bool {
	for i, v := range y {
		if notEquals(v, z[i]) {
			return false
		}
	}
	return true
}

// Copy copies x onto z, and returns z.
func (z *Minkowski) Copy(x *Minkowski) *Minkowski {
	for i, v := range x {
		z[i] = v
	}
	return z
}

// NewMinkowski returns a pointer to an Minkowski value made from four given
// float64 values.
func NewMinkowski(a, b, c, d float64) *Minkowski {
	z := new(Minkowski)
	z[0] = a
	z[1] = b
	z[2] = c
	z[3] = d
	return z
}

// IsMinkowskiInf returns true if any of the components of z are infinite.
func (z *Minkowski) IsMinkowskiInf() bool {
	for _, v := range z {
		if math.IsInf(v, 0) {
			return true
		}
	}
	return false
}

// MinkowskiInf returns a pointer to a Minkowski quaternionic infinity value.
func MinkowskiInf(a, b, c, d int) *Minkowski {
	return NewMinkowski(math.Inf(a), math.Inf(b), math.Inf(c), math.Inf(d))
}

// IsMinkowskiNaN returns true if any component of z is NaN and neither is an
// infinity.
func (z *Minkowski) IsMinkowskiNaN() bool {
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

// MinkowskiNaN returns a pointer to a Minkowski quaternionic NaN value.
func MinkowskiNaN() *Minkowski {
	nan := math.NaN()
	return NewMinkowski(nan, nan, nan, nan)
}

// Scal sets z equal to y scaled by a, and returns z.
func (z *Minkowski) Scal(y *Minkowski, a float64) *Minkowski {
	for i, v := range y {
		z[i] = a * v
	}
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Minkowski) Neg(y *Minkowski) *Minkowski {
	return z.Scal(y, -1)
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Minkowski) Conj(y *Minkowski) *Minkowski {
	z[0] = y[0]
	for i, v := range y[1:] {
		z[i+1] = -v
	}
	return z
}

// Add sets z equal to the sum of x and y, and returns z.
func (z *Minkowski) Add(x, y *Minkowski) *Minkowski {
	for i, v := range x {
		z[i] = v + y[i]
	}
	return z
}

// Sub sets z equal to the difference of x and y, and returns z.
func (z *Minkowski) Sub(x, y *Minkowski) *Minkowski {
	for i, v := range x {
		z[i] = v - y[i]
	}
	return z
}

// Mul sets z equal to the product of x and y, and returns z.
//
// The multiplication rule for the basis elements s := Minkowski{0, 1, 0, 0},
// t := Minkowski{0, 0, 1, 0}, and u := Minkowski{0, 0, 0, 1} is:
// 		Mul(s, s) = Mul(t, t) = Mul(u, u) = Minkowski{1, 0, 0, 0}
// 		Mul(s, t) = -Mul(t, s) = +u
// 		Mul(t, u) = -Mul(u, t) = +s
// 		Mul(u, s) = -Mul(s, u) = +t
func (z *Minkowski) Mul(x, y *Minkowski) *Minkowski {
	p := new(Minkowski).Copy(x)
	q := new(Minkowski).Copy(y)
	z[0] = (p[0] * q[0]) + (p[1] * q[1]) + (p[2] * q[2]) + (p[3] * q[3])
	z[1] = (p[0] * q[1]) + (p[1] * q[0]) + (p[2] * q[3]) - (p[3] * q[2])
	z[2] = (p[0] * q[2]) - (p[1] * q[3]) + (p[2] * q[0]) + (p[3] * q[1])
	z[3] = (p[0] * q[3]) + (p[1] * q[2]) - (p[2] * q[1]) + (p[3] * q[0])
	return z
}

// Commutator sets z equal to the commutator of x and y, and returns z.
func (z *Minkowski) Commutator(x, y *Minkowski) *Minkowski {
	return z.Sub(new(Minkowski).Mul(x, y), new(Minkowski).Mul(y, x))
}

// Associator sets z equal to the associator of w, x, and y, and returns z.
func (z *Minkowski) Associator(w, x, y *Minkowski) *Minkowski {
	return z.Sub(
		new(Minkowski).Mul(new(Minkowski).Mul(w, x), y),
		new(Minkowski).Mul(w, new(Minkowski).Mul(x, y)),
	)
}

// AlternatorL sets z equal to the left alternator of x and y, and returns z.
func (z *Minkowski) AlternatorL(x, y *Minkowski) *Minkowski {
	return z.Associator(x, x, y)
}

// AlternatorR sets z equal to the right alternator of x and y, and returns z.
func (z *Minkowski) AlternatorR(x, y *Minkowski) *Minkowski {
	return z.Associator(x, y, y)
}

// Quad returns the quadrance of z, which can be either positive, negative or
// zero.
func (z *Minkowski) Quad() float64 {
	return (new(Minkowski).Mul(z, new(Minkowski).Conj(z)))[0]
}

// IsZeroDiv returns true if z is a zero divisor (i.e. it has zero quadrance).
func (z *Minkowski) IsZeroDiv() bool {
	return !notEquals(z.Quad(), 0)
}

// Inv sets z equal to the inverse of x, and returns z. If x is a zero divisor,
// then Inv panics.
func (z *Minkowski) Inv(x *Minkowski) *Minkowski {
	if x.IsZeroDiv() {
		panic("inverse of zero divisor")
	}
	return z.Scal(new(Minkowski).Conj(x), 1/x.Quad())
}

// Quo sets z equal to the quotient of x and y, and returns z. If y is a zero
// divisor, then Quo panics.
func (z *Minkowski) Quo(x, y *Minkowski) *Minkowski {
	if y.IsZeroDiv() {
		panic("denominator is zero divisor")
	}
	return z.Scal(new(Minkowski).Mul(x, new(Minkowski).Conj(y)), 1/y.Quad())
}

// IsIndempotent returns true if z is an indempotent (i.e. if z = z*z).
func (z *Minkowski) IsIndempotent() bool {
	return z.Equals(new(Minkowski).Mul(z, z))
}

// RectMinkowski returns a Minkowski value made from given curvilinear
// coordinates and quadrance sign.
func RectMinkowski(r, ξ, θ1, θ2 float64, sign int) *Minkowski {
	z := new(Minkowski)
	if sign > 0 {
		z[0] = r * math.Cosh(ξ)
		z[1] = r * math.Sinh(ξ) * math.Cos(θ1)
		z[2] = r * math.Sinh(ξ) * math.Sin(θ1) * math.Cos(θ2)
		z[3] = r * math.Sinh(ξ) * math.Sin(θ1) * math.Sin(θ2)
		return z
	}
	if sign < 0 {
		z[0] = r * math.Sinh(ξ)
		z[1] = r * math.Cosh(ξ) * math.Cos(θ1)
		z[2] = r * math.Cosh(ξ) * math.Sin(θ1) * math.Cos(θ2)
		z[3] = r * math.Cosh(ξ) * math.Sin(θ1) * math.Sin(θ2)
		return z
	}
	z[0] = r
	z[1] = r * math.Cos(θ1)
	z[2] = r * math.Sin(θ1) * math.Cos(θ2)
	z[3] = r * math.Sin(θ1) * math.Sin(θ2)
	return z
}

// Curv returns the curvilinear coordinates of a Minkowski value, along with
// the sign of the quadrance.
func (z *Minkowski) Curv() (r, ξ, θ1, θ2 float64, sign int) {
	quad := z.Quad()
	h := math.Hypot(z[2], z[3])
	θ1 = math.Atan2(z[1], h)
	θ2 = math.Atan2(z[3], z[2])
	if quad > 0 {
		r = math.Sqrt(quad)
		ξ = math.Atanh(math.Hypot(z[1], h) / z[0])
		sign = +1
		return
	}
	if quad < 0 {
		r = math.Sqrt(-quad)
		ξ = math.Atanh(z[0] / math.Hypot(z[1], h))
		sign = -1
		return
	}
	r = z[0]
	ξ = math.NaN()
	sign = 0
	return
}
