package main

import (
	"fmt"
	"time"
)

func pop(c chan string) {
	msg := <- c
	fmt.Println(msg)
}

func main() {

	c := make(chan string, 4)

	c <- "hello 1"
	c <- "hello 2"
	c <- "hello 3"
	c <- "hello 4"
	c <- "hello 5"

	pop(c)
	pop(c)
	pop(c)
	pop(c)
	pop(c)


	time.Sleep(1 * time.Second)
}