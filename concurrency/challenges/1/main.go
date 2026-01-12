package main

import (
	"fmt"
	"time"
)

func g1(c chan string) {
	c <- "Hello from func g1"
}

func g2(c chan string) {
	msg := <- c

	fmt.Println("I am g2 and ...", msg)


}
func main() {
    ch := make(chan string)
	go g1(ch)
	go g2(ch)

	time.Sleep(1 * time.Second)

}