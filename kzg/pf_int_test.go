package kzg_test

import (
	"runtime"
	"reflect"
	"strings"
	"testing"
	"math/big"

	"github.com/stretchr/testify/require"

	"github.com/neurotempest/kzg_commitment_scheme/kzg"
)

func pfInt(v int64) *kzg.PFInt {

	return kzg.NewPFIntFromInt64(v)
}

func TestPFInt_BinaryOperators(t *testing.T) {

	kzg.SetFieldModulus(t, 1021)

	testCases := []struct{
		Name string
		Method func(*kzg.PFInt, *kzg.PFInt, *kzg.PFInt)*kzg.PFInt
		A *kzg.PFInt
		B *kzg.PFInt
		Expected *kzg.PFInt
	}{
		{
			Name: "Add initalised too big",
			Method: (*kzg.PFInt).Add,
			A: pfInt(1230),
			B: pfInt(456),
			Expected: pfInt(665),
		},
		{
			Name: "Add no overflow",
			Method: (*kzg.PFInt).Add,
			A: pfInt(123),
			B: pfInt(456),
			Expected: pfInt(579),
		},
		{
			Name: "Add with overflow",
			Method: (*kzg.PFInt).Add,
			A: pfInt(783),
			B: pfInt(456),
			Expected: pfInt(218),
		},
		{
			Name: "Sub no underflow",
			Method: (*kzg.PFInt).Sub,
			A: pfInt(783),
			B: pfInt(456),
			Expected: pfInt(327),
		},
		{
			Name: "Sub with underflow",
			Method: (*kzg.PFInt).Sub,
			A: pfInt(123),
			B: pfInt(456),
			Expected: pfInt(688),
		},
		{
			Name: "Mul no underflow",
			Method: (*kzg.PFInt).Mul,
			A: pfInt(12),
			B: pfInt(45),
			Expected: pfInt(540),
		},
		{
			Name: "Mul with underflow",
			Method: (*kzg.PFInt).Mul,
			A: pfInt(123),
			B: pfInt(456),
			Expected: pfInt(954),
		},
		{
			Name: "Div no underflow",
			Method: (*kzg.PFInt).Div,
			A: pfInt(428),
			B: pfInt(4),
			Expected: pfInt(107),
		},
		{
			Name: "Div with underflow",
			Method: (*kzg.PFInt).Div,
			A: pfInt(324),
			B: pfInt(567),
			Expected: pfInt(584),
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {

			originalA := kzg.NewPFIntDeepCopy(test.A)
			originalB := kzg.NewPFIntDeepCopy(test.B)

			workspace := new(kzg.PFInt)
			actual := test.Method(workspace, test.A, test.B)
			require.Equal(t, test.Expected, actual)
			require.Equal(t, test.Expected, workspace)

			require.Equal(t, originalA, test.A, "input value changed")
			require.Equal(t, originalB, test.B, "input value changed")
		})
	}
}

func TestPFInt_Exp(t *testing.T) {

	kzg.SetFieldModulus(t, 1019)

	testCases := []struct{
		Name string
		X *kzg.PFInt
		Exponent *big.Int
		Expected *kzg.PFInt
	}{
		{
			Name: "No wrap around",
			X: pfInt(9),
			Exponent: big.NewInt(3),
			Expected: pfInt(729),
		},
		{
			Name: "With wrap around",
			X: pfInt(123),
			Exponent: big.NewInt(456),
			Expected: pfInt(239),
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {

			originalX := kzg.NewPFIntDeepCopy(test.X)
			originalExponent := new(big.Int)
			originalExponent.Set(test.Exponent)

			workspace := new(kzg.PFInt)
			actual := workspace.Exp(test.X, test.Exponent)
			require.Equal(t, test.Expected, actual)
			require.Equal(t, test.Expected, workspace)

			require.Equal(t, originalX, test.X, "input value changed")
			require.Equal(t, originalExponent, test.Exponent, "input value changed")
		})
	}
}

func TestPFInt_Set(t *testing.T) {

	kzg.SetFieldModulus(t, 1019)

	testCases := []struct{
		Name string
		A *kzg.PFInt
		Expected *kzg.PFInt
	}{
		{
			Name: "No wrap around",
			A: pfInt(235),
			Expected: pfInt(235),
		},
		{
			Name: "With wrap around",
			A: pfInt(245775),
			Expected: pfInt(196),
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {

			originalA := kzg.NewPFIntDeepCopy(test.A)

			workspace := new(kzg.PFInt)
			actual := workspace.Set(test.A)
			require.Equal(t, test.Expected, actual)
			require.Equal(t, test.Expected, workspace)

			require.Equal(t, originalA, test.A, "input value changed")
		})
	}
}

func TestPFInt_SetBigInt(t *testing.T) {

	kzg.SetFieldModulus(t, 1019)

	testCases := []struct{
		Name string
		A *big.Int
		Expected *kzg.PFInt
	}{
		{
			Name: "No wrap around",
			A: big.NewInt(235),
			Expected: pfInt(235),
		},
		{
			Name: "With wrap around",
			A: big.NewInt(245775),
			Expected: pfInt(196),
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {

			originalA := new(big.Int)
			originalA.Set(test.A)

			workspace := new(kzg.PFInt)
			actual := workspace.SetBigInt(test.A)
			require.Equal(t, test.Expected, actual)
			require.Equal(t, test.Expected, workspace)

			require.Equal(t, originalA, test.A, "input value changed")
		})
	}
}

func TestPFInt_SetInt64(t *testing.T) {

	kzg.SetFieldModulus(t, 1019)

	testCases := []struct{
		Name string
		A int64
		Expected *kzg.PFInt
	}{
		{
			Name: "No wrap around",
			A: 235,
			Expected: pfInt(235),
		},
		{
			Name: "With wrap around",
			A: 245775,
			Expected: pfInt(196),
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {

			workspace := new(kzg.PFInt)
			actual := workspace.SetInt64(test.A)
			require.Equal(t, test.Expected, actual)
			require.Equal(t, test.Expected, workspace)
		})
	}
}

func TestPFInt_BinaryOperatorOnNewPFInt(t *testing.T) {

	kzg.SetFieldModulus(t, 1021)

	a := kzg.NewPFIntFromInt64(1234)
	b := kzg.NewPFIntFromInt64(5678)
	expected := kzg.NewPFIntFromInt64(786)

	originalA := kzg.NewPFIntDeepCopy(a)
	originalB := kzg.NewPFIntDeepCopy(b)

	actual := new(kzg.PFInt).Add(a, b)

	require.Equal(t, expected, actual)

	require.Equal(t, originalA, a, "input value changed")
	require.Equal(t, originalB, b, "input value changed")
}

func TestPFInt_NewPFIntFromInt64(t *testing.T) {

	kzg.SetFieldModulus(t, 1031)

	testCases := []struct{
		Name string
		I int64
		Expected *kzg.PFInt
	}{
		{
			Name: "No wrap around",
			I: 123,
			Expected: pfInt(123),
		},
		{
			Name: "With wrap around",
			I: 123456788,
			Expected: pfInt(724),
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			actual := kzg.NewPFIntFromInt64(test.I)
			require.Equal(t, test.Expected, actual)
		})
	}
}

func TestPFInt_NewPFIntDeepCopy(t *testing.T) {

	kzg.SetFieldModulus(t, 1031)

	testCases := []struct{
		Name string
		I *kzg.PFInt
	}{
		{
			Name: "When original updated copy does not change",
			I: pfInt(123),
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			iCopy := kzg.NewPFIntDeepCopy(test.I)

			test.I.Add(test.I, kzg.NewPFIntFromInt64(2))

			require.NotEqual(t, iCopy, test.I)
		})
	}
}

func funcName(i interface{}) string {
	splitName := strings.Split(runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name(), ".")
	return splitName[len(splitName)-1]
}
