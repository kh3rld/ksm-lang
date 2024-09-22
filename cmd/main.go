package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kh3rld/ksm-lang/eval"
	"github.com/kh3rld/ksm-lang/lexer"
	"github.com/kh3rld/ksm-lang/parser"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <filename.ksm>", os.Args[0])
	}

	filename := os.Args[1]

	// Read the contents of the .ksm file
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Create a lexer for the contents
	l := lexer.New(string(data))

	// Create a parser
	p := parser.New(l)

	// Parse the input into a program
	program := p.ParseProgram()
	if len(p.Errors()) > 0 {
		for _, err := range p.Errors() {
			fmt.Println("Parser error:", err)
		}
		return
	}

	// Create an interpreter and evaluate the program
	interp := &eval.Evaluator{}
	for _, stmt := range program.Statements {
		result := interp.Evaluate(stmt)
		if result != nil {
			fmt.Println("Result:", result)
		} else {
			fmt.Println("Evaluated statement:", stmt.String(), "returned nil")
		}
	}
}
