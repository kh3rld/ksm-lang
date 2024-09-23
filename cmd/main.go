package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/kh3rld/ksm-lang/eval"
	"github.com/kh3rld/ksm-lang/lexer"
	"github.com/kh3rld/ksm-lang/parser"
	t "github.com/kh3rld/ksm-lang/token"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ksm <file.ksm>")
		os.Exit(1)
	}
	sourceFile := os.Args[1]
	file, err := os.Open(sourceFile)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	evaluator := &eval.Evaluator{}
	for scanner.Scan() {
		input := scanner.Text()
		le := lexer.New(string(input))

		for token := le.NextToken(); token.Type != t.EOF; token = le.NextToken() {
			fmt.Printf("%+v\n", token)
		}
		l := lexer.New(input)
		p := parser.New(l)
		numberExpr := p.ParseNumber()

		if numberExpr == nil {
			if len(p.Errors()) > 0 {
				log.Printf("Error parsing input %v", p.Errors())
			} else {
				log.Printf("Unexpected error nil returned but no errors recorded.\n")
			}
		} else {
			fmt.Printf("%v\n", numberExpr.Value)
		}
		expression := p.ParseExpression()

		// Check if the parsing was successful
		if expression == nil {
			if len(p.Errors()) > 0 {
				log.Printf("Error parsing input '%s': %v", input, p.Errors())
			} else {
				log.Printf("Unexpected error for input '%s': nil returned but no errors recorded.\n", input)
			}
			continue
		}
		// Evaluate the parsed expression using the evaluator
		result := evaluator.Eval(expression)
		if result == nil {
			log.Printf("Error evaluating input '%s'", input)
			continue
		}

		// Print the evaluated result
		fmt.Printf("Evaluated result: %v\n", result.Value)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
}
