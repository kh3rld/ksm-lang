package main

import (
	"fmt"
	"ksm-lang/lexer"
)

func main() {
	input := `let nums = [1, 2, 3];`
	l := lexer.New(input)

	for tok := l.NextToken(); tok.Type != lexer.EOF; tok = l.NextToken() {
		fmt.Printf("%+v\n", tok)
	}
}
