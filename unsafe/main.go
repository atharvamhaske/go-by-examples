package main

import (
	"fmt"
	// "reflect"
	"unsafe"
)

func main() {
	i := 10
	iptr := unsafe.Pointer(&i)
	fmt.Println(*(*int)(iptr)) // we here telling go that this iptr is a pointer to int
}