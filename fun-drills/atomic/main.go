// package main

// import (
// 	"sync/atomic"
// 	"sync"
// 	"fmt"
// )

// func main() {
// 	var x int32
// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		atomic.AddInt32(&x, 2)
// 		}()
// 		wg.Wait()
// 		fmt.Printf("This is new value of counter: %v \n", x)
// }

//now using channel

package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var x int32
	done := make(chan struct{})

	go func() {
		atomic.AddInt32(&x, 2)
		close(done)
	}()

	<-done
	fmt.Println("Value updated to:", x)
}
