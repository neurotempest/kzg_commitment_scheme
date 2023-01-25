package kzg_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/neurotempest/kzg_commitment_scheme/kzg"
)

func TestPFPoint_NewFromInt64(t *testing.T) {

	testCases := []struct{
		Name string
		X int64
		Y int64
		Expected *kzg.PFPoint
	}{
		{
			Name: "No wrap",
			X: 2,
			Y: 4,
			Expected: &kzg.PFPoint{
				X: *kzg.NewPFIntFromInt64(2),
				Y: *kzg.NewPFIntFromInt64(4),
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			actual := kzg.NewPFPointFromInt64(test.X, test.Y)
			require.Equal(t, test.Expected, actual)
		})
	}
}
