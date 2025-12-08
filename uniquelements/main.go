// package main
// import (
// 	"fmt"
// 	"sort"
// )
// func RemDuplicate(a []int) []int{
// 	seen := make(map[int]struct{}) // key is int and value is empty struct
// 	result := make([]int, 0, len(a))
// 	for _,v := range a {
// 		if _, exists := seen[v];
// 		!exists {
// 			seen[v] = struct{}{}
// 			result = append(result, v)
// 		}
// 	}
// 	return result
// }
// func main() {
// 	duplicateSlice := []int{1,2,3,2,4,2,3,4}
// 	fmt.Println("so the duplicate slice before removal is:", duplicateSlice)
// 	removedSlice := RemDuplicate(duplicateSlice)
// 	sort.Ints(removedSlice)
// 	fmt.Println("after removal of dupicalte nums our slice is:", removedSlice)
// }

package main

import (
	"fmt"
	"sort"
)

func sorted(x []int) []int {
	if len(x) == 0 {
		return x
	}

	final := []int{x[0]} // 2, 1, 2, 3

	for i := 1; i < len(x); i++ {
		if x[i] != x[i-1] {
			final = append(final, x[i])
		}
	}
	return final
}
func main() {
	nums := []int{1, 2, 2, 3, 4, 1}
	fmt.Println("nums slice before removal:", nums)
	sort.Ints(nums)
	sortnums := sorted(nums)
	fmt.Println("nums slice after removal:", sortnums)
}
