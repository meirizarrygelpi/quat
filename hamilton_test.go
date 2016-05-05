package quat

import (
	"fmt"
	"testing"
)

func ExampleHamiltonInf() {
	fmt.Println(HamiltonInf(-1, 0, 0, 0))
	fmt.Println(HamiltonInf(0, -1, 0, 0))
	fmt.Println(HamiltonInf(0, 0, -1, 0))
	fmt.Println(HamiltonInf(0, 0, 0, -1))
	// Output:
	// (-Inf+Infi+Infj+Infk)
	// (+Inf-Infi+Infj+Infk)
	// (+Inf+Infi-Infj+Infk)
	// (+Inf+Infi+Infj-Infk)
}

func ExampleHamiltonNaN() {
	fmt.Println(HamiltonNaN())
	// Output:
	// (NaN+NaNi+NaNj+NaNk)
}

func ExampleNewHamilton() {
	fmt.Println(NewHamilton(1, 0, 0, 0))
	fmt.Println(NewHamilton(0, 1, 0, 0))
	fmt.Println(NewHamilton(0, 0, 1, 0))
	fmt.Println(NewHamilton(0, 0, 0, 1))
	fmt.Println(NewHamilton(1, 2, 3, 4))
	// Output:
	// (1+0i+0j+0k)
	// (0+1i+0j+0k)
	// (0+0i+1j+0k)
	// (0+0i+0j+1k)
	// (1+2i+3j+4k)
}

func ExampleRectHamilton() {}

func TestHamiltonAdd(t *testing.T) {}

func TestHamiltonCommutator(t *testing.T) {}

func TestHamiltonConj(t *testing.T) {}

func TestHamiltonCopy(t *testing.T) {}

func TestHamiltonCurv(t *testing.T) {}

func TestHamiltonEquals(t *testing.T) {}

func TestHamiltonInv(t *testing.T) {}

func TestIsHamiltonInf(t *testing.T) {}

func TestIsHamiltonNaN(t *testing.T) {}

func TestHamiltonMul(t *testing.T) {}

func TestHamiltonNeg(t *testing.T) {}

func TestHamiltonQuad(t *testing.T) {}

func TestHamiltonQuo(t *testing.T) {}

func TestHamiltonScal(t *testing.T) {}

func TestHamiltonString(t *testing.T) {}

func TestHamiltonSub(t *testing.T) {}
