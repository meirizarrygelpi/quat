// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package quat

import (
	"fmt"
	"testing"
)

func ExampleCockleInf() {
	fmt.Println(CockleInf(-1, 0, 0, 0))
	fmt.Println(CockleInf(0, -1, 0, 0))
	fmt.Println(CockleInf(0, 0, -1, 0))
	fmt.Println(CockleInf(0, 0, 0, -1))
	// Output:
	// (-Inf+Infi+Inft+Infu)
	// (+Inf-Infi+Inft+Infu)
	// (+Inf+Infi-Inft+Infu)
	// (+Inf+Infi+Inft-Infu)
}

func ExampleCockleNaN() {
	fmt.Println(CockleNaN())
	// Output:
	// (NaN+NaNi+NaNt+NaNu)
}

func ExampleNewCockle() {
	fmt.Println(NewCockle(1, 0, 0, 0))
	fmt.Println(NewCockle(0, 1, 0, 0))
	fmt.Println(NewCockle(0, 0, 1, 0))
	fmt.Println(NewCockle(0, 0, 0, 1))
	fmt.Println(NewCockle(1, 2, 3, 4))
	// Output:
	// (1+0i+0t+0u)
	// (0+1i+0t+0u)
	// (0+0i+1t+0u)
	// (0+0i+0t+1u)
	// (1+2i+3t+4u)
}

func TestCockleAdd(t *testing.T) {}

func TestCockleCommutator(t *testing.T) {}

func TestCockleConj(t *testing.T) {}

func TestCockleCopy(t *testing.T) {}

func TestCockleEquals(t *testing.T) {}

func TestCockleInv(t *testing.T) {}

func TestIsCockleInf(t *testing.T) {}

func TestCockleIsIndempotent(t *testing.T) {}

func TestIsCockleNaN(t *testing.T) {}

func TestCockleIsNilpotent(t *testing.T) {}

func TestCockleIsZeroDiv(t *testing.T) {}

func TestCockleMul(t *testing.T) {}

func TestCockleNeg(t *testing.T) {}

func TestCockleQuad(t *testing.T) {}

func TestCockleQuo(t *testing.T) {}

func TestCockleScal(t *testing.T) {}

func TestCockleString(t *testing.T) {}

func TestCockleSub(t *testing.T) {}
