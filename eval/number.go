package eval

import "fmt"

type Number struct {
	Value float64
}

func (n *Number) Add(other *Number) *Number {
	return &Number{Value: n.Value + other.Value}
}

func (n *Number) Subtract(other *Number) *Number {
	return &Number{Value: n.Value - other.Value}
}

func (n *Number) ToInt() int {
	return int(n.Value)
}

func (n *Number) String() string {
	return fmt.Sprintf("%f", n.Value)
}
