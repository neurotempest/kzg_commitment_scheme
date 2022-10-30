package kzg_test

import (
	"runtime"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/neurotempest/kzg_commitment_scheme/kzg"
)

func pfInt(v int64) kzg.PFInt {

	return kzg.NewPFIntFromInt64(v)
}

func TestPFInt_BinaryOperators(t *testing.T) {

	kzg.SetFieldModulus(t, 1021)

	testCases := []struct{
		Name string
		Method func(kzg.PFInt, kzg.PFInt)kzg.PFInt
		A kzg.PFInt
		B kzg.PFInt
		Expected kzg.PFInt
	}{
		{
			Name: "Add initalised too big",
			Method: (kzg.PFInt).Add,
			A: pfInt(1230),
			B: pfInt(456),
			Expected: pfInt(665),
		},
		{
			Name: "Add no overflow",
			Method: (kzg.PFInt).Add,
			A: pfInt(123),
			B: pfInt(456),
			Expected: pfInt(579),
		},
		{
			Name: "Add with overflow",
			Method: (kzg.PFInt).Add,
			A: pfInt(783),
			B: pfInt(456),
			Expected: pfInt(218),
		},
		{
			Name: "Sub no underflow",
			Method: (kzg.PFInt).Sub,
			A: pfInt(783),
			B: pfInt(456),
			Expected: pfInt(327),
		},
		{
			Name: "Sub with underflow",
			Method: (kzg.PFInt).Sub,
			A: pfInt(123),
			B: pfInt(456),
			Expected: pfInt(688),
		},
		{
			Name: "Mul no underflow",
			Method: (kzg.PFInt).Mul,
			A: pfInt(12),
			B: pfInt(45),
			Expected: pfInt(540),
		},
		{
			Name: "Mul with underflow",
			Method: (kzg.PFInt).Mul,
			A: pfInt(123),
			B: pfInt(456),
			Expected: pfInt(954),
		},
		{
			Name: "Div no underflow",
			Method: (kzg.PFInt).Div,
			A: pfInt(428),
			B: pfInt(4),
			Expected: pfInt(107),
		},
		{
			Name: "Div with underflow",
			Method: (kzg.PFInt).Div,
			A: pfInt(324),
			B: pfInt(567),
			Expected: pfInt(584),
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {

			originalA := kzg.NewPFIntCopy(test.A)
			originalB := kzg.NewPFIntCopy(test.B)

			actual := test.Method(test.A, test.B)
			require.Equal(t, test.Expected, actual)

			require.Equal(t, originalA, test.A, "input value changed")
			require.Equal(t, originalB, test.B, "input value changed")
		})
	}
}

func TestSome(t *testing.T) {

	kzg.SetFieldModulus(t, 7)

	testCases := []struct{
		Name string
		Method func(kzg.PFInt, kzg.PFInt)kzg.PFInt
		A kzg.PFInt
		B kzg.PFInt
		Expected kzg.PFInt
	}{
		{
			Name: "Div no underflow",
			Method: (kzg.PFInt).Div,
			//A: pfInt(12),
			//B: pfInt(45),
			A: kzg.NewPFIntFromInt64(3),
			B: kzg.NewPFIntFromInt64(2),
			Expected: kzg.NewPFIntFromInt64(5),
		},
		//{
		//	Name: "Div with underflow",
		//	Method: (kzg.PFInt).Div,
		//	A: pfInt(123),
		//	B: pfInt(456),
		//	Expected: pfInt(954),
		//},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {

			originalA := kzg.NewPFIntCopy(test.A)
			originalB := kzg.NewPFIntCopy(test.B)

			actual := test.Method(test.A, test.B)
			require.Equal(t, test.Expected, actual)

			require.Equal(t, originalA, test.A, "input value changed")
			require.Equal(t, originalB, test.B, "input value changed")
		})
	}
}


func funcName(i interface{}) string {
	splitName := strings.Split(runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name(), ".")
	return splitName[len(splitName)-1]
}
