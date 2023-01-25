package kzg

import (
	"math/big"
)

type PFPolynomial struct {

	coefficients []PFInt
}

func NewPFPolynomialFromInt64(coefficients []int64) *PFPolynomial {

	p := new(PFPolynomial)
	p.coefficients = make([]PFInt, len(coefficients))

	for iC, c := range coefficients {
		p.coefficients[iC].SetInt64(c)
	}

	return p
}

func (p *PFPolynomial) Add(a, b *PFPolynomial) *PFPolynomial {

	aCoeffsLen := len(a.coefficients)
	bCoeffsLen := len(b.coefficients)

	coeffsLen := aCoeffsLen
	if bCoeffsLen > aCoeffsLen {
		coeffsLen = bCoeffsLen
	}

	p.coefficients = make([]PFInt, coeffsLen)

	for iC := range a.coefficients {
		if iC < bCoeffsLen {
			p.coefficients[iC].Add(
				&a.coefficients[iC],
				&b.coefficients[iC],
			)
		} else {
			p.coefficients[iC].Set(
				&a.coefficients[iC],
			)
		}
	}

	for iC := len(a.coefficients); iC < len(b.coefficients); iC++ {
		p.coefficients[iC].Set(
			&b.coefficients[iC],
		)
	}

	return p
}

func (p *PFPolynomial) Eval(x *PFInt) *PFInt {

	y := new(PFInt)
	wsMul := new(PFInt)
	wsExp := new(PFInt)
	exponent := new(big.Int)
	for iCoeff := range p.coefficients {
		if iCoeff == 0 {
			y.Set(&p.coefficients[iCoeff])
			continue
		}

		// y += coeff * x^iCoeff
		y.Add(
			y,
			wsMul.Mul(
				&p.coefficients[iCoeff],
				wsExp.Exp(
					x,
					exponent.SetInt64(int64(iCoeff)),
				),
			),
		)
	}

	return y
}

func (p *PFPolynomial) Set(x *PFPolynomial) *PFPolynomial {

	p.coefficients = make([]PFInt, len(x.coefficients))
	for iC, c := range x.coefficients {
		p.coefficients[iC].Set(&c)
	}

	return p
}

func (p *PFPolynomial) Sub(a, b *PFPolynomial) *PFPolynomial {

	aCoeffsLen := len(a.coefficients)
	bCoeffsLen := len(b.coefficients)

	coeffsLen := aCoeffsLen
	if bCoeffsLen > aCoeffsLen {
		coeffsLen = bCoeffsLen
	}

	p.coefficients = make([]PFInt, coeffsLen)

	for iC := range a.coefficients {
		if iC < bCoeffsLen {
			p.coefficients[iC].Sub(
				&a.coefficients[iC],
				&b.coefficients[iC],
			)
		} else {
			p.coefficients[iC].Set(
				&a.coefficients[iC],
			)
		}
	}

	for iC := len(a.coefficients); iC < len(b.coefficients); iC++ {
		p.coefficients[iC].Sub(
			&p.coefficients[iC],
			&b.coefficients[iC],
		)
	}

	return p
}
