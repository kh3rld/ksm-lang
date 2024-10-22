package strings

import (
	"testing"
)

func TestStrOperations(t *testing.T) {
	tests := []struct {
		name     string
		testFunc func(*testing.T)
	}{
		{"TestNew", func(t *testing.T) {
			s := New("hello")
			if s.Value != "hello" {
				t.Errorf("Expected 'hello', got %s", s.Value)
			}
		}},
		{"TestAdd", func(t *testing.T) {
			s1 := New("hello")
			s2 := New("world")
			result := s1.Concatenate(s2)
			if result.Value != "helloworld" {
				t.Errorf("Expected 'helloworld', got %s", result.Value)
			}
		}},
		{"TestLen", func(t *testing.T) {
			s := New("hello")
			if s.Len() != 5 {
				t.Errorf("Expected length 5, got %d", s.Len())
			}
		}},
		{"TestSub", func(t *testing.T) {
			s := New("hello")
			result := s.Sub(0, 2)
			if result.Value != "he" {
				t.Errorf("Expected 'he', got %s", result.Value)
			}
		}},
		{"TestUp", func(t *testing.T) {
			s := New("hello")
			result := s.Up()
			if result.Value != "HELLO" {
				t.Errorf("Expected 'HELLO', got %s", result.Value)
			}
		}},
		{"TestLow", func(t *testing.T) {
			s := New("HELLO")
			result := s.Low()
			if result.Value != "hello" {
				t.Errorf("Expected 'hello', got %s", result.Value)
			}
		}},
		{"TestHas", func(t *testing.T) {
			s := New("hello world")
			substr := New("world")
			if !s.Has(substr) {
				t.Error("Expected string to contain 'world'")
			}
		}},
		{"TestRev", func(t *testing.T) {
			s := New("hello")
			result := s.Rev()
			if result.Value != "olleh" {
				t.Errorf("Expected 'olleh', got %s", result.Value)
			}
		}},
		{"TestRep", func(t *testing.T) {
			s := New("ha")
			result := s.Rep(3)
			if result.Value != "hahaha" {
				t.Errorf("Expected 'hahaha', got %s", result.Value)
			}
		}},
		{"TestCut", func(t *testing.T) {
			s := New("a-b-c")
			delim := New("-")
			parts := s.Cut(delim)
			if len(parts) != 3 {
				t.Errorf("Expected 3 parts, got %d", len(parts))
			}
			if parts[0].Value != "a" || parts[1].Value != "b" || parts[2].Value != "c" {
				t.Errorf("Cut results incorrect")
			}
		}},
		{"TestTrim", func(t *testing.T) {
			s := New("  hello  ")
			result := s.Trim()
			if result.Value != "hello" {
				t.Errorf("Expected 'hello', got %s", result.Value)
			}
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.testFunc)
	}
}
