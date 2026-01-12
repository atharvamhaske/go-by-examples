package main

import "fmt"

func main() {
	s := []int{}

	fmt.Println(len(s)) // 0 initially
	fmt.Println(cap(s)) // this too 0 initially

	s = append(s, 1, 2, 3)

	fmt.Println(len(s))
	fmt.Println(cap(s))

	slice := make([]int, 5, 10)
	fmt.Println(cap(slice)) // Output: 10
	fmt.Println(len(slice))

	a := sumofArr([]int{1, 2, 3})
	fmt.Println(a)

}

// The length of a slice is the number of elements it contains.
// The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
// The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).

func sumofArr(arr []int) int {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return sum
}
