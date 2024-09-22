package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kh3rld/ksm-lang/lexer"
	"github.com/kh3rld/ksm-lang/parser"
	t "github.com/kh3rld/ksm-lang/token"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ksm <file.ksm>")
		os.Exit(1)
	}

	// Read source file
	sourceFile := os.Args[1]
	inputs, err := os.ReadFile(sourceFile)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	l := lexer.New(string(inputs))

	for token := l.NextToken(); token.Type != t.EOF; token = l.NextToken() {
		fmt.Printf("%+v\n", token)
	}
	for _, input := range []string{string(inputs)} {
		l := lexer.New(input)
		p := parser.New(l)

		numberExpr := p.ParseNumber()

		if numberExpr == nil {
			if len(p.Errors()) > 0 {
				log.Printf("Error parsing input %s: %v", input, p.Errors())
			} else {
				log.Printf("Unexpected error for input '%s': nil returned but no errors recorded.\n", input)
			}
		} else {
			fmt.Printf(" %v\n", numberExpr.Value)
		}
	}
}
