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

package main
import (
	"context"
	"fmt"
	"time"
)

func doWork(ctx context.Context) {
	select {
		case <-time.After(2 * time.Second):
		fmt.Println("work done")
		case <-ctx.Done():
		fmt.Println("cancelled:", ctx.Err())
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1 * time.Second)
	defer cancel() //cancel the context to release resources when done
	
	//start go routine to perform some processing
	go doWork(ctx)
	
	//wait for a few seconds to allow proccessing to complete
	time.Sleep(3 * time.Second)
}