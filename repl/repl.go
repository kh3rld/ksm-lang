package repl

import (
	"bufio"
	"fmt"
	"io"
	"ksm-lang/interpreter"
	"ksm-lang/lexer"
	"ksm-lang/object"
	"ksm-lang/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := interpreter.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func Run(input string) {
    l := lexer.New(input)
    p := parser.New(l)

    program := p.ParseProgram()
    if len(p.Errors()) != 0 {
        printParserErrors(io.Discard, p.Errors()) // Use io.Discard to suppress error output
        return
    }

    env := object.NewEnvironment()
    evaluated := interpreter.Eval(program, env)
    if evaluated != nil {
        fmt.Println(evaluated.Inspect())
    }
}
func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
