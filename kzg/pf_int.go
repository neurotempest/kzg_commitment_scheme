package kzg

import (
	"math/big"
)

// PFInt represents an integer in a prime field with modulus `FieldModulus()`
type PFInt struct {
	v *big.Int
}

func NewPFIntFromInt64(i int64) PFInt {

	val := big.NewInt(i)
	return PFInt{
		v: val.Mod(val, FieldModulus()),
	}
}

func NewPFIntCopy(i PFInt) PFInt {

	ret := NewPFIntFromInt64(0)
	ret.v.Set(i.v)
	return ret
}


func (i PFInt) Add(j PFInt) PFInt {

	ret := NewPFIntFromInt64(0)
	ret.v.Add(i.v, j.v)
	ret.v.Mod(ret.v, FieldModulus())
	return ret
}

func (i PFInt) Sub(j PFInt) PFInt {

	ret := NewPFIntFromInt64(0)
	ret.v.Sub(i.v, j.v)
	ret.v.Mod(ret.v, FieldModulus())
	return ret
}

func (i PFInt) Mul(j PFInt) PFInt {

	ret := NewPFIntFromInt64(0)
	ret.v.Mul(i.v, j.v)
	ret.v.Mod(ret.v, FieldModulus())
	return ret
}

func (i PFInt) Div(j PFInt) PFInt {

	x := new(big.Int)
	y := new(big.Int)
	new(big.Int).GCD(x, y, j.v, FieldModulus())

	multiplier := new(big.Int).Mod(x, FieldModulus())

	ret := NewPFIntFromInt64(0)
	ret.v.Mul(i.v, multiplier)
	ret.v.Mod(ret.v, FieldModulus())
	return ret
}
