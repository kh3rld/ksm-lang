package eval

type Condition struct {
	zero bool
	one  bool
}

func (c *Condition) BooL(s *Condition) *Condition {
	if c.zero {
		return &Condition{zero: c.zero || s.zero}
	}
	return &Condition{one: c.one || s.one}
}
