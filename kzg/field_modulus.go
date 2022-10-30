package kzg

import (
	"errors"
	"math/big"
	"testing"
	"strconv"
)

var defaultFieldModulus = "21888242871839275222246405745257275088696311157297823662689037894645226208583"

func FieldModulus() *big.Int {

	fm := new(big.Int)
	err := fm.UnmarshalText([]byte(
		defaultFieldModulus,
	))
	if err != nil {
		panic(err)
	}

	return fm
}

func checkModulusIsPrime(v *big.Int) {

	// check modulus is prime if (2^v)%v == 2
	var check big.Int
	check.Exp(big.NewInt(2), v, v)
	if check.Cmp(big.NewInt(2)) != 0 {
		panic(errors.New("modulus not prime"))
	}
}

func SetFieldModulus(t *testing.T, i int64) {

	checkModulusIsPrime(
		big.NewInt(i),
	)

	old := defaultFieldModulus
	defaultFieldModulus = strconv.FormatInt(i, 10)
	t.Cleanup(func() {
		defaultFieldModulus = old
	})
}

