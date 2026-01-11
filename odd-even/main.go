package main

import (
	"fmt"
	"sync"
)

func main() {
	even, odd := make(chan bool), make(chan bool) //two unbuffered channels
	wg := sync.WaitGroup{}

	n := 10

	wg.Add(2)

	go printOdd(n, &wg, odd, even)
	go printEven(n, &wg, odd, even)

	odd <- true //tells odd goroutine you go first

	wg.Wait()
}

func printOdd(n int, wg *sync.WaitGroup, odd chan bool, even chan bool) {
	defer wg.Done()

	for i := 1; i <= n; i += 2 {
		<-odd //odd goroutine will wait until someone gives it a permission
		fmt.Println("Odd Number:", i)

		even <- true //odd is saying im done its your turn now
	}
	fmt.Println("odd loop finished")
}

func printEven(n int, wg *sync.WaitGroup, odd chan bool, even chan bool) {
	defer wg.Done()

	for i := 2; i <= n; i += 2 {
		<-even
		fmt.Println("Even Number:", i)

		if i == n {
			return
		}
		odd <- true
	}

	fmt.Println("even loop finished")
}
