package eval

type Str struct {
	Value string
}

func (s *Str) Concatinate(v *Str) *Str {
	return &Str{Value: s.Value + v.Value}
}
