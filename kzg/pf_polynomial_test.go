package kzg_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/neurotempest/kzg_commitment_scheme/kzg"
)


func TestPFPolynomial_Eval(t *testing.T) {

	testCases := []struct{
		Name string
		FieldModulus int64
		P *kzg.PFPolynomial
		X *kzg.PFInt
		Expected *kzg.PFInt
	}{
		{
			Name: "mod 7 quadratic small x",
			FieldModulus: 7,
			P: kzg.NewPFPolynomialFromInt64([]int64{5, 4, 3}),
			X: pfInt(5),
			Expected: pfInt(2),
		},
		{
			Name: "mod 7 quadratic large x",
			FieldModulus: 7,
			P: kzg.NewPFPolynomialFromInt64([]int64{5, 4, 3}),
			X: pfInt(1243653*7 + 5),
			Expected: pfInt(2),
		},
		{
			Name: "mod 7 bigger values",
			FieldModulus: 7,
			P: kzg.NewPFPolynomialFromInt64([]int64{9, 4, 3, 43, 8, 54}),
			X: pfInt(624177),
			Expected: pfInt(2),
		},
		{
			Name: "mod 23 even bigger values",
			FieldModulus: 23,
			P: kzg.NewPFPolynomialFromInt64([]int64{923234, 442323, 3124, 4364657, 8421, 544121}),
			X: pfInt(342453560981256),
			Expected: pfInt(18),
		},
		{
			Name: "mod 1021 even bigger values",
			FieldModulus: 1021,
			P: kzg.NewPFPolynomialFromInt64([]int64{923234, 442323, 3124, 4364657, 8421, 544121}),
			X: pfInt(342453560981256),
			Expected: pfInt(51),
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			if test.FieldModulus != 0 {
				kzg.SetFieldModulus(t, test.FieldModulus)
			}
			actual := test.P.Eval(test.X)
			require.Equal(t, test.Expected, actual)
		})
	}
}

func TestPFPolynomial_BinaryOperators(t *testing.T) {

	kzg.SetFieldModulus(t, 1021)

	testCases := []struct{
		Name string
		Method func(*kzg.PFPolynomial, *kzg.PFPolynomial, *kzg.PFPolynomial)*kzg.PFPolynomial
		A *kzg.PFPolynomial
		B *kzg.PFPolynomial
		Expected *kzg.PFPolynomial
	}{
		{
			Name: "Add when first param has more polynomial terms",
			Method: (*kzg.PFPolynomial).Add,
			A: kzg.NewPFPolynomialFromInt64([]int64{56, 901, 29, 53, 28}),
			B: kzg.NewPFPolynomialFromInt64([]int64{5, 21, 12, 32}),
			Expected: kzg.NewPFPolynomialFromInt64([]int64{61, 922, 41, 85, 28}),
		},
		{
			Name: "Add when second param has more polynomial terms",
			Method: (*kzg.PFPolynomial).Add,
			A: kzg.NewPFPolynomialFromInt64([]int64{546, 9021, 29053, 9328}),
			B: kzg.NewPFPolynomialFromInt64([]int64{5, 21, 290, 12, 32}),
			Expected: kzg.NewPFPolynomialFromInt64([]int64{551, 874, 755, 151, 32}),
		},
		{
			Name: "Sub when first param has more polynomial terms",
			Method: (*kzg.PFPolynomial).Sub,
			A: kzg.NewPFPolynomialFromInt64([]int64{56, 901, 29, 53, 28}),
			B: kzg.NewPFPolynomialFromInt64([]int64{5, 1003, 12, 67}),
			Expected: kzg.NewPFPolynomialFromInt64([]int64{51, 919, 17, 1007, 28}),
		},
		{
			Name: "Sub when second param has more polynomial terms",
			Method: (*kzg.PFPolynomial).Sub,
			A: kzg.NewPFPolynomialFromInt64([]int64{546, 9021, 29053, 932}),
			B: kzg.NewPFPolynomialFromInt64([]int64{5, 832, 590, 1204, 32}),
			Expected: kzg.NewPFPolynomialFromInt64([]int64{541, 21, 896, 749, 989}),
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {

			originalA := new(kzg.PFPolynomial)
			originalA.Set(test.A)
			originalB := new(kzg.PFPolynomial)
			originalB.Set(test.B)

			newWorkspace := new(kzg.PFPolynomial)
			actual := test.Method(newWorkspace, test.A, test.B)
			require.Equal(t, test.Expected, actual)
			require.Equal(t, test.Expected, newWorkspace)

			require.Equal(t, originalA, test.A, "input value changed")
			require.Equal(t, originalB, test.B, "input value changed")

			existingWorkspace := kzg.NewPFPolynomialFromInt64([]int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
			actual = test.Method(existingWorkspace, test.A, test.B)
			require.Equal(t, test.Expected, actual)
			require.Equal(t, test.Expected, existingWorkspace)

			require.Equal(t, originalA, test.A, "input value changed")
			require.Equal(t, originalB, test.B, "input value changed")
		})
	}
}

func TestPFPoly_Set(t *testing.T) {

	kzg.SetFieldModulus(t, 1019)

	testCases := []struct{
		Name string
		A *kzg.PFPolynomial
		Expected *kzg.PFPolynomial
	}{
		{
			Name: "With wrap around",
			A: kzg.NewPFPolynomialFromInt64([]int64{423, 67453, 5433}),
			Expected: kzg.NewPFPolynomialFromInt64([]int64{423, 199, 338}),
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {

			workspace := new(kzg.PFPolynomial)
			actual := workspace.Set(test.A)
			require.Equal(t, test.Expected, actual)
			require.Equal(t, test.Expected, workspace)
		})
	}
}
