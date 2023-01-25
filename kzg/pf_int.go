package kzg

import (
	"math/big"
)

// PFInt represents an integer in a prime field with modulus `FieldModulus()`
type PFInt struct {
	v big.Int
}

func NewPFIntFromInt64(i int64) *PFInt {
	ret := new(PFInt)
	ret.v.SetInt64(i)
	ret.v.Mod(&ret.v, FieldModulus())
	return ret
}

func NewPFIntDeepCopy(i *PFInt) *PFInt {
	ret := NewPFIntFromInt64(0)
	ret.v.Set(&i.v)
	return ret
}

func (i *PFInt) Add(a, b *PFInt) *PFInt {
	i.v.Add(&a.v, &b.v)
	i.mod()
	return i
}

func (i *PFInt) Div(a, b *PFInt) *PFInt {

	// TODO make Div more memory efficient
	x := new(big.Int)
	y := new(big.Int)
	new(big.Int).GCD(x, y, &b.v, FieldModulus())

	multiplier := new(big.Int).Mod(x, FieldModulus())

	i.v.Mul(&a.v, multiplier)
	i.v.Mod(&i.v, FieldModulus())
	return i
}

func (i *PFInt) Exp(x *PFInt, exponent *big.Int) *PFInt {
	i.v.Exp(&x.v, exponent, FieldModulus())
	return i
}

func (i *PFInt) Mul(a, b *PFInt) *PFInt {
	i.v.Mul(&a.v, &b.v)
	i.mod()
	return i
}

func (i *PFInt) Set(a *PFInt) *PFInt {
	i.v.Set(&a.v)
	return i
}

func (i *PFInt) SetBigInt(a *big.Int) *PFInt {
	i.v.Set(a)
	i.mod()
	return i
}

func (i *PFInt) SetInt64(a int64) *PFInt {
	i.v.SetInt64(a)
	i.mod()
	return i
}

func (i *PFInt) Sub(a, b *PFInt) *PFInt {
	i.v.Sub(&a.v, &b.v)
	i.mod()
	return i
}

func (i *PFInt) mod() {
	i.v.Mod(&i.v, FieldModulus())
}
