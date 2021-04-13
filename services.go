package devduck

type Service interface {
	Do(int) int
}

type add struct {
	x int
}

func NewAdd(x int) Service {
	return &add{x: x}
}

func (o *add) Do(operand int) int {
	return operand + o.x
}

type mul struct {
	x int
}

func NewMul(x int) Service {
	return &mul{x: x}
}

func (o *mul) Do(operand int) int {
	return operand * o.x
}
