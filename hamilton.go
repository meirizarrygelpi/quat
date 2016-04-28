package qtr

import (
	"fmt"
	"math"
	"strings"
)

var symbH = [4]string{"", "i", "j", "k"}

// A Hamilton represents a Hamilton quaternion (i.e. a traditional quaternion)
// as an ordered array of four float64 values.
type Hamilton [4]float64

// String returns the string representation of a Hamilton value. If z
// corresponds to the Hamilton quaternion a + bi + cj + dk, then the string is
// "(a+bi+cj+dk)", similar to complex128 values.
func (z *Hamilton) String() string {
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
func (z *Hamilton) Equals(y *Hamilton) bool {
	for i, v := range y {
		if notEquals(v, z[i]) {
			return false
		}
	}
	return true
}

// Copy copies y onto z, and returns z.
func (z *Hamilton) Copy(y *Hamilton) *Hamilton {
	for i, v := range y {
		z[i] = v
	}
	return z
}

// NewHamilton returns a pointer to a Hamilton value made from four given
// float64 values.
func NewHamilton(a, b, c, d float64) *Hamilton {
	z := new(Hamilton)
	z[0] = a
	z[1] = b
	z[2] = c
	z[3] = d
	return z
}

// IsHamiltonInf returns true if any of the components of z are infinite.
func (z *Hamilton) IsHamiltonInf() bool {
	for _, v := range z {
		if math.IsInf(v, 0) {
			return true
		}
	}
	return false
}

// HamiltonInf returns a pointer to a Hamilton quaternionic infinity value.
func HamiltonInf(a, b, c, d int) *Hamilton {
	return NewHamilton(math.Inf(a), math.Inf(b), math.Inf(c), math.Inf(d))
}

// IsHamiltonNaN returns true if any component of z is NaN and neither is an
// infinity.
func (z *Hamilton) IsHamiltonNaN() bool {
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

// HamiltonNaN returns a pointer to a Hamilton quaternionic NaN value.
func HamiltonNaN() *Hamilton {
	nan := math.NaN()
	return NewHamilton(nan, nan, nan, nan)
}

// Scal sets z equal to y scaled by a, and returns z.
func (z *Hamilton) Scal(y *Hamilton, a float64) *Hamilton {
	for i, v := range y {
		z[i] = a * v
	}
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Hamilton) Neg(y *Hamilton) *Hamilton {
	return z.Scal(y, -1)
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Hamilton) Conj(y *Hamilton) *Hamilton {
	z[0] = y[0]
	for i, v := range y[1:] {
		z[i+1] = -v
	}
	return z
}

// Add sets z equal to the sum of x and y, and returns z.
func (z *Hamilton) Add(x, y *Hamilton) *Hamilton {
	for i, v := range x {
		z[i] = v + y[i]
	}
	return z
}

// Sub sets z equal to the difference of x and y, and returns z.
func (z *Hamilton) Sub(x, y *Hamilton) *Hamilton {
	for i, v := range x {
		z[i] = v - y[i]
	}
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
	z[0] = (p[0] * q[0]) - (p[1] * q[1]) - (p[2] * q[2]) - (p[3] * q[3])
	z[1] = (p[0] * q[1]) + (p[1] * q[0]) + (p[2] * q[3]) - (p[3] * q[2])
	z[2] = (p[0] * q[2]) - (p[1] * q[3]) + (p[2] * q[0]) + (p[3] * q[1])
	z[3] = (p[0] * q[3]) + (p[1] * q[2]) - (p[2] * q[1]) + (p[3] * q[0])
	return z
}

// Commutator sets z equal to the commutator of x and y, and returns z.
func (z *Hamilton) Commutator(x, y *Hamilton) *Hamilton {
	return z.Sub(new(Hamilton).Mul(x, y), new(Hamilton).Mul(y, x))
}

// Quad returns the non-negative quadrance of z.
func (z *Hamilton) Quad() float64 {
	return (new(Hamilton).Mul(z, new(Hamilton).Conj(z)))[0]
}

// Inv sets z equal to the inverse of y, and returns z. If y is zero, then Inv
// panics.
func (z *Hamilton) Inv(y *Hamilton) *Hamilton {
	if y.Equals(&Hamilton{0, 0, 0, 0}) {
		panic("inverse of zero")
	}
	return z.Scal(new(Hamilton).Conj(y), 1/y.Quad())
}

// Quo sets z equal to the quotient of x and y, and returns z. If y is zero,
// then Quo panics.
func (z *Hamilton) Quo(x, y *Hamilton) *Hamilton {
	if y.Equals(&Hamilton{0, 0, 0, 0}) {
		panic("denominator is zero")
	}
	return z.Scal(new(Hamilton).Mul(x, new(Hamilton).Conj(y)), 1/y.Quad())
}

// RectHamilton returns a Hamilton value made from given curvilinear
// coordinates.
func RectHamilton(r, θ1, θ2, θ3 float64) *Hamilton {
	z := new(Hamilton)
	z[0] = r * math.Cos(θ1)
	z[1] = r * math.Sin(θ1) * math.Cos(θ2)
	z[2] = r * math.Sin(θ1) * math.Sin(θ2) * math.Cos(θ3)
	z[3] = r * math.Sin(θ1) * math.Sin(θ2) * math.Sin(θ3)
	return z
}

// Curv returns the curvilinear coordinates of a Hamilton value.
func (z *Hamilton) Curv() (r, θ1, θ2, θ3 float64) {
	h := math.Hypot(z[2], z[3])
	r = math.Sqrt(z.Quad())
	θ1 = math.Atan2(math.Hypot(z[1], h), z[0])
	θ2 = math.Atan2(h, z[1])
	θ3 = math.Atan2(z[3], z[2])
	return
}
