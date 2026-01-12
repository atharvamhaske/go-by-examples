package main

import (
	"errors"
	"fmt"
)

type xError struct {
	x   int
	msg string
}

func (e *xError) Error() string {
	return fmt.Sprintf("%d - %s", e.x, e.msg)
}

func foo(y int) (int, error) {
	if y == 42 {
		return -1, &xError{y, "cannot work with it"}
	}
	return y + 3, &xError{y + 3, "ok lala"}
}

func main() {
	_, err := foo(40)
	var ae *xError
	if errors.As(err, &ae) {
		fmt.Println(ae.x)
		fmt.Println(ae.msg)
	} else {

		// fmt.Println("error dont match our custom error type")
	}
}
