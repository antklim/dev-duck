package app

import "github.com/antklim/dev-duck/app/iface"

var _ iface.Service = (*Add)(nil)

type Add struct {
	x int
}

func NewAdd(x int) *Add {
	return &Add{x: x}
}

func (o *Add) Do(operand int) int {
	return operand + o.x
}

var _ iface.Service = (*Mul)(nil)

type Mul struct {
	x int
}

func NewMul(x int) *Mul {
	return &Mul{x: x}
}

func (o *Mul) Do(operand int) int {
	return operand * o.x
}
