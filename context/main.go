// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"time"
// )

// func hello(w http.ResponseWriter, req *http.Request) {
// 	ctx := req.Context()
// 	fmt.Println("server: hello handler started")
// 	defer fmt.Println("server: hello handler ended")

// 	select {
// 	case <-time.After(10 * time.Second):
// 		fmt.Fprintf(w, "hello\n")
// 	case <-ctx.Done():

// 		err := ctx.Err()
// 		fmt.Println("server:", err)
// 		internalError := http.StatusInternalServerError
// 		http.Error(w, err.Error(), internalError)
// 	}
// }

// func main() {
// 	http.HandleFunc("/hello", hello)
// 	http.ListenAndServe(":8090", nil)
// }
//
// example more below

// package main

// import (
// 	"context"
// 	"fmt"
// 	"time"
// )

// func doWork(ctx context.Context) {
// 	select {
// 	case <-time.After(2 * time.Second):
// 		fmt.Println("work done")
// 	case <-ctx.Done():
// 		fmt.Println("cancelled:", ctx.Err())
// 	}
// }

// func main() {
// 	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
// 	defer cancel() //cancel the context to release resources when done

// 	//start go routine to perform some processing
// 	go doWork(ctx)

// 	//wait for a few seconds to allow proccessing to complete
// 	time.Sleep(3 * time.Second)
// }

package main

import (
	"time"
	"context"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	
	go func(){
		select {
			case <-time.After(3 * time.Second):
			fmt.Println("work done")
			case <-ctx.Done(): // this waits untill someone cancels channel it returns c.done named channel internally once cancelled this case runs
			//Because a receive on an unclosed, empty channel ALWAYS blocks.
			fmt.Println("work cancelled")
		}
	}()
	
	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(3 * time.Second)
}

// ctx.Done() returns context’s internal done channel
//<-ctx.Done() blocks until channel is CLOSED
//cancel() closes the channel
//Closing a channel wakes all waiting goroutines
//Select chooses ctx.Done() case
//The println runs immediately after wake-up
// 
// 
// will write blog on this all reminder to it