package main

import (
	"fmt"
	"os"
	"ksm-lang/repl"
)

func main() {
	if len(os.Args) > 1 {
		fileName := os.Args[1]
		content, err := os.ReadFile(fileName)
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			return
		}
 	    repl.Run(string(content))
	} else {
		repl.Start(os.Stdin, os.Stdout)
	}
}
