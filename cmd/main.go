package main

import (
	"bufio"
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
	sourceFile := os.Args[1]
	file, err := os.Open(sourceFile)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
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
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
}
