// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/kh3rld/ksm-lang/lexer"
// 	t "github.com/kh3rld/ksm-lang/token"
// )

// func main() {
// 	if len(os.Args) < 2 {
// 		fmt.Println("Usage: ksm <file.ksm>")
// 		os.Exit(1)
// 	}

// 	// Read source file
// 	sourceFile := os.Args[1]
// 	input, err := os.ReadFile(sourceFile)
// 	if err != nil {
// 		log.Fatalf("Error reading file: %v", err)
// 	}

// 	l := lexer.New(string(input))

// 	for token := l.NextToken(); token.Type != t.EOF; token = l.NextToken() {
// 		fmt.Printf("%+v\n", token)
// 	}
// }

package main

import (
	"fmt"
	"log"

	"github.com/kh3rld/ksm-lang/lexer"
	"github.com/kh3rld/ksm-lang/parser"
)

func main() {
	inputs := []string{
		"42",
		"-42",
		"3.14",
		"-3.14",
		"not_a_number",
		"0",
	}

	for _, input := range inputs {
		l := lexer.New(input)
		p := parser.New(l)

		numberExpr := p.ParseNumber()

		if numberExpr == nil {
			if len(p.Errors()) > 0 {
				log.Printf("Error parsing input '%s': %v\n", input, p.Errors())
			} else {
				log.Printf("Unexpected error for input '%s': nil returned but no errors recorded.\n", input)
			}
		} else {
			fmt.Printf(" %v\n", numberExpr.Value)
		}
	}
}
