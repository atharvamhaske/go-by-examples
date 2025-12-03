package main

import (
	"errors"
	"fmt"
)

func foo(x int) (int, error) {
	if x == 42 {
		return  -1, errors.New("cant work with 42")
	}
	return x + 3, nil
}

var ErroutOfTea = errors.New("no more tea available")
var ErrPower = errors.New("cannot boil water")

func makeTea(arg int) error {
	if arg == 2 {
		return ErroutOfTea
	} else if arg == 4 {
		return fmt.Errorf("making tea: %w", ErrPower)
	}
	return nil
}

func main () {
	for _, i := range []int{7, 42} {
		if r, e := foo(i); e != nil {
			fmt.Println("foo failed", e)
		} else {
			fmt.Println("foo worked", r)
		}
	} // here if we have for x, y then x is here index and y is value
	
	for i := range 5 {
		if err := makeTea(i); err != nil {
			if errors.Is(err, ErroutOfTea) {
				fmt.Println("we should buy a new tea")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("now its dark")
			} else {
				fmt.Printf("unknown error: %s\n", err)
			}
			break
		}
		fmt.Println("tea is ready")
	}
}