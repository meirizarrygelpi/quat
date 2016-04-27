package qtr

import (
	"fmt"
	"testing"
)

func ExampleKleinInf() {
	fmt.Println(KleinInf(-1, 0, 0, 0))
	fmt.Println(KleinInf(0, -1, 0, 0))
	fmt.Println(KleinInf(0, 0, -1, 0))
	fmt.Println(KleinInf(0, 0, 0, -1))
	// Output:
	// (-Inf+Infi+Inft+Infu)
	// (+Inf-Infi+Inft+Infu)
	// (+Inf+Infi-Inft+Infu)
	// (+Inf+Infi+Inft-Infu)
}

func ExampleKleinNaN() {
	fmt.Println(KleinNaN())
	// Output:
	// (NaN+NaNi+NaNt+NaNu)
}

func ExampleNewKlein() {
	fmt.Println(NewKlein(1, 0, 0, 0))
	fmt.Println(NewKlein(0, 1, 0, 0))
	fmt.Println(NewKlein(0, 0, 1, 0))
	fmt.Println(NewKlein(0, 0, 0, 1))
	fmt.Println(NewKlein(1, 2, 3, 4))
	// Output:
	// (1+0i+0t+0u)
	// (0+1i+0t+0u)
	// (0+0i+1t+0u)
	// (0+0i+0t+1u)
	// (1+2i+3t+4u)
}

func TestKleinAdd(t *testing.T) {}

func TestKleinCommutator(t *testing.T) {}

func TestKleinConj(t *testing.T) {}

func TestKleinCopy(t *testing.T) {}

func TestKleinEquals(t *testing.T) {}

func TestKleinInv(t *testing.T) {}

func TestIsKleinInf(t *testing.T) {}

func TestKleinIsIndempotent(t *testing.T) {}

func TestIsKleinNaN(t *testing.T) {}

func TestKleinIsNilpotent(t *testing.T) {}

func TestKleinIsZeroDiv(t *testing.T) {}

func TestKleinMul(t *testing.T) {}

func TestKleinNeg(t *testing.T) {}

func TestKleinQuad(t *testing.T) {}

func TestKleinQuo(t *testing.T) {}

func TestKleinScal(t *testing.T) {}

func TestKleinString(t *testing.T) {}

func TestKleinSub(t *testing.T) {}
