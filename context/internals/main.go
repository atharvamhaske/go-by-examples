package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	ctx1 := context.WithValue(ctx, "one", 1)

	ctx2 := context.WithValue(ctx, "two", 2)

	ctxval := ctx.Value("two")

	val2 := ctx1.Value("one")

	val3 := ctx2.Value("two")

	val4 := ctx2.Value(2)

	fmt.Println(ctxval)
	fmt.Println(val2)
	fmt.Println(val4)
	fmt.Println(val3)
}

//wow this is so good to learn