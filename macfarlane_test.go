package quat

import (
	"fmt"
	"testing"
)

func ExampleMacfarlaneInf() {
	fmt.Println(MacfarlaneInf(-1, 0, 0, 0))
	fmt.Println(MacfarlaneInf(0, -1, 0, 0))
	fmt.Println(MacfarlaneInf(0, 0, -1, 0))
	fmt.Println(MacfarlaneInf(0, 0, 0, -1))
	// Output:
	// (-Inf+Infs+Inft+Infu)
	// (+Inf-Infs+Inft+Infu)
	// (+Inf+Infs-Inft+Infu)
	// (+Inf+Infs+Inft-Infu)
}

func ExampleMacfarlaneNaN() {
	fmt.Println(MacfarlaneNaN())
	// Output:
	// (NaN+NaNs+NaNt+NaNu)
}

func ExampleNewMacfarlane() {
	fmt.Println(NewMacfarlane(1, 0, 0, 0))
	fmt.Println(NewMacfarlane(0, 1, 0, 0))
	fmt.Println(NewMacfarlane(0, 0, 1, 0))
	fmt.Println(NewMacfarlane(0, 0, 0, 1))
	fmt.Println(NewMacfarlane(1, 2, 3, 4))
	// Output:
	// (1+0s+0t+0u)
	// (0+1s+0t+0u)
	// (0+0s+1t+0u)
	// (0+0s+0t+1u)
	// (1+2s+3t+4u)
}

func TestMacfarlaneAdd(t *testing.T) {}

func TestMacfarlaneAlternatorL(t *testing.T) {}

func TestMacfarlaneAlternatorR(t *testing.T) {}

func TestMacfarlaneAssociator(t *testing.T) {}

func TestMacfarlaneCommutator(t *testing.T) {}

func TestMacfarlaneConj(t *testing.T) {}

func TestMacfarlaneCopy(t *testing.T) {}

func TestMacfarlaneEquals(t *testing.T) {}

func TestMacfarlaneInv(t *testing.T) {}

func TestIsMacfarlaneInf(t *testing.T) {}

func TestMacfarlaneIsIndempotent(t *testing.T) {}

func TestIsMacfarlaneNaN(t *testing.T) {}

func TestMacfarlaneIsZeroDiv(t *testing.T) {}

func TestMacfarlaneMul(t *testing.T) {}

func TestMacfarlaneNeg(t *testing.T) {}

func TestMacfarlaneQuad(t *testing.T) {}

func TestMacfarlaneQuo(t *testing.T) {}

func TestMacfarlaneScal(t *testing.T) {}

func TestMacfarlaneString(t *testing.T) {}

func TestMacfarlaneSub(t *testing.T) {}
