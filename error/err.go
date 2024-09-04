package error

import "fmt"

type KsmError struct {
	Line    int
	Column  int
	Message string
}

func (e *KsmError) Error() string {
	return fmt.Sprintf("Line %d, Column %d: %s", e.Line, e.Column, e.Message)
}

func NewKsmError(line, column int, message string) *KsmError {
	return &KsmError{
		Line:    line,
		Column:  column,
		Message: message,
	}
}

func handleError(err *KsmError) {
	fmt.Println(err)
}
