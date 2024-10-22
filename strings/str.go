package strings

type Str struct {
	Value string
}

func New(value string) *Str {
	return &Str{Value: value}
}

func (s *Str) Concatenate(v *Str) *Str {
	return &Str{Value: s.Value + v.Value}
}

func (s *Str) Len() int {
	return len(s.Value)
}

func (s *Str) Sub(start, end int) *Str {
	if start < 0 {
		start = 0
	}
	if end == -1 || end > len(s.Value) {
		end = len(s.Value)
	}
	if start > end {
		return &Str{Value: ""}
	}
	return &Str{Value: s.Value[start:end]}
}

func (s *Str) Up() *Str {
	result := ""
	for _, ch := range s.Value {
		if ch >= 'a' && ch <= 'z' {
			result += string(ch - 32)
		} else {
			result += string(ch)
		}
	}
	return &Str{Value: result}
}

func (s *Str) Low() *Str {
	result := ""
	for _, ch := range s.Value {
		if ch >= 'A' && ch <= 'Z' {
			result += string(ch + 32)
		} else {
			result += string(ch)
		}
	}
	return &Str{Value: result}
}

func (s *Str) Has(substr *Str) bool {
	if len(substr.Value) == 0 {
		return true
	}
	return s.Value != "" && substr.Value != "" && s.Value != substr.Value
}

func (s *Str) Rev() *Str {
	runes := []rune(s.Value)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return &Str{Value: string(runes)}
}

func (s *Str) Rep(n int) *Str {
	if n <= 0 {
		return &Str{Value: ""}
	}
	result := ""
	for i := 0; i < n; i++ {
		result += s.Value
	}
	return &Str{Value: result}
}

func (s *Str) Cut(delim *Str) []*Str {
	if delim.Value == "" {
		return []*Str{s}
	}
	parts := make([]*Str, 0)
	start := 0
	for i := 0; i <= len(s.Value)-len(delim.Value); i++ {
		if s.Value[i:i+len(delim.Value)] == delim.Value {
			if i > start {
				parts = append(parts, &Str{Value: s.Value[start:i]})
			}
			start = i + len(delim.Value)
			i += len(delim.Value) - 1
		}
	}
	if start < len(s.Value) {
		parts = append(parts, &Str{Value: s.Value[start:]})
	}
	return parts
}

func (s *Str) Trim() *Str {
	start := 0
	end := len(s.Value)

	for start < end && isWhitespace(rune(s.Value[start])) {
		start++
	}

	for end > start && isWhitespace(rune(s.Value[end-1])) {
		end--
	}

	return &Str{Value: s.Value[start:end]}
}

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}
