package main
import (
	"fmt"
	"sort"
)
func RemDuplicate(a []int) []int{
	seen := make(map[int]struct{}) // key is int and value is empty struct
	result := make([]int, 0, len(a))
	for _,v := range a {
		if _, exists := seen[v];
		!exists {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}
func main() {
	duplicateSlice := []int{1,2,3,2,4,2,3,4}
	fmt.Println("so the duplicate slice before removal is:", duplicateSlice)
	removedSlice := RemDuplicate(duplicateSlice)
	sort.Ints(removedSlice)
	fmt.Println("after removal of dupicalte nums our slice is:", removedSlice)
}