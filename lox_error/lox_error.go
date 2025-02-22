package lox_error

import "fmt"

type ErrorHandler struct {
	HadError bool
}

type Error struct {
	Line    int
	Where   string
	Message string
}

func (e *ErrorHandler) Report(err Error) {
	fmt.Printf("[line %d] Error%s: %s", err.Line, err.Where, err.Message)
	e.HadError = true
}
