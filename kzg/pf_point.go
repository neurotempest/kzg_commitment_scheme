package kzg

import (

)

type PFPoint struct {
	X PFInt
	Y PFInt
}

func NewPFPointFromInt64(x, y int64) *PFPoint {

	ret := new(PFPoint)

	ret.X.SetInt64(x)
	ret.Y.SetInt64(y)

	return ret
}
