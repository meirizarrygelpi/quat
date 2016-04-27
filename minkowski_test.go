package qtr

import (
	"fmt"
	"testing"
)

func ExampleMinkowskiInf() {
	fmt.Println(MinkowskiInf(-1, 0, 0, 0))
	fmt.Println(MinkowskiInf(0, -1, 0, 0))
	fmt.Println(MinkowskiInf(0, 0, -1, 0))
	fmt.Println(MinkowskiInf(0, 0, 0, -1))
	// Output:
	// (-Inf+Infs+Inft+Infu)
	// (+Inf-Infs+Inft+Infu)
	// (+Inf+Infs-Inft+Infu)
	// (+Inf+Infs+Inft-Infu)
}

func ExampleMinkowskiNaN() {
	fmt.Println(MinkowskiNaN())
	// Output:
	// (NaN+NaNs+NaNt+NaNu)
}

func ExampleNewMinkowski() {
	fmt.Println(NewMinkowski(1, 0, 0, 0))
	fmt.Println(NewMinkowski(0, 1, 0, 0))
	fmt.Println(NewMinkowski(0, 0, 1, 0))
	fmt.Println(NewMinkowski(0, 0, 0, 1))
	fmt.Println(NewMinkowski(1, 2, 3, 4))
	// Output:
	// (1+0s+0t+0u)
	// (0+1s+0t+0u)
	// (0+0s+1t+0u)
	// (0+0s+0t+1u)
	// (1+2s+3t+4u)
}

func TestMinkowskiAdd(t *testing.T) {}

func TestMinkowskiAlternatorL(t *testing.T) {}

func TestMinkowskiAlternatorR(t *testing.T) {}

func TestMinkowskiAssociator(t *testing.T) {}

func TestMinkowskiCommutator(t *testing.T) {}

func TestMinkowskiConj(t *testing.T) {}

func TestMinkowskiCopy(t *testing.T) {}

func TestMinkowskiEquals(t *testing.T) {}

func TestMinkowskiInv(t *testing.T) {}

func TestIsMinkowskiInf(t *testing.T) {}

func TestMinkowskiIsIndempotent(t *testing.T) {}

func TestIsMinkowskiNaN(t *testing.T) {}

func TestMinkowskiIsZeroDiv(t *testing.T) {}

func TestMinkowskiMul(t *testing.T) {}

func TestMinkowskiNeg(t *testing.T) {}

func TestMinkowskiQuad(t *testing.T) {}

func TestMinkowskiQuo(t *testing.T) {}

func TestMinkowskiScal(t *testing.T) {}

func TestMinkowskiString(t *testing.T) {}

func TestMinkowskiSub(t *testing.T) {}
