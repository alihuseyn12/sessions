package custerr

import "errors"

//Homework 1: Custom Error Types

type DivisionError struct {
	message string
}

func (e DivisionError) Error() error {
	return errors.New(e.message)
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		divisionError := DivisionError{message: "Error: Division by zero is not allowed."}
		return 0, divisionError.Error()
	} else {
		return a / b, nil
	}
}
