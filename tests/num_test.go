package test

import (
	"testing"

	"github.com/kh3rld/ksm-lang/eval"
)

func TestNumber_Add(t *testing.T) {
	num1 := &eval.Number{Value: 10.5}
	num2 := &eval.Number{Value: 5.5}
	expected := 16.0

	result := num1.Add(num2)
	if result.Value != expected {
		t.Errorf("Expected %f, got %f", expected, result.Value)
	}
}

func TestNumber_Subtract(t *testing.T) {
	num1 := &eval.Number{Value: 10.5}
	num2 := &eval.Number{Value: 5.5}
	expected := 5.0

	result := num1.Subtract(num2)
	if result.Value != expected {
		t.Errorf("Expected %f, got %f", expected, result.Value)
	}
}

func TestNumber_ToInt(t *testing.T) {
	num := &eval.Number{Value: 10.5}
	expected := 10

	result := num.ToInt()
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
