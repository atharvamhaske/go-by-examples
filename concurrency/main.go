package main

import "fmt"

func main() {
    var data int
    go func() {
        data++
    }()

    if data == 0 {
        fmt.Printf("the value is %d", data)
    }
}

// Race condition occur when two or more operations must execute in the correct order,
// but the program has not been written so that this order is guaranteed to be maintained.

// Data race is when one concurrent operation attempts to read a variable
// while at some undetermined time another concurrent operation is attempting to write to the same variable.
// The main func is the main goroutine.