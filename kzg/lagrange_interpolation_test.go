package kzg_test

import (
	"testing"
	//"strings"
	//"math/big"

	//"github.com/stretchr/testify/require"

	"github.com/neurotempest/kzg_commitment_scheme/kzg"
)

func Test_Temp(t *testing.T) {

	kzg.SetFieldModulus(t, 1021)

	type PFPoint struct {
		X *kzg.PFInt
		Y *kzg.PFInt
	}

	_ = []PFPoint{
		PFPoint{
			X: pfInt(10),
			Y: pfInt(20),
		},
		PFPoint{
			X: pfInt(35),
			Y: pfInt(234),
		},
		PFPoint{
			X: pfInt(123),
			Y: pfInt(12),
		},
	}

/*
	x := pfInt(30)

	y := new(kzg.PFInt)

	for iY := range points {

		for iX := range points {


		}
	}
*/

}
