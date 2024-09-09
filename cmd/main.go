package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kh3rld/ksm-lang/lexer"
	t "github.com/kh3rld/ksm-lang/token"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ksm <file.ksm>")
		os.Exit(1)
	}

	// Read source file
	sourceFile := os.Args[1]
	input, err := os.ReadFile(sourceFile)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	l := lexer.New(string(input))

	for token := l.NextToken(); token.Type != t.EOF; token = l.NextToken() {
		fmt.Printf("%+v\n", token)
	}
}
