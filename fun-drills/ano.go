package main

import "fmt"

type Notifier struct {
	ch chan string
}

func (n *Notifier) Notify(msg string) {
	n.ch <- msg
}

func main() {
	n := Notifier{
		ch: make(chan string),
	}

	go func(){
		fmt.Println("Message received", <-n.ch)
	}()

	n.Notify("Hello")
}