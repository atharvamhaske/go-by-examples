package main

import "fmt"

func leader(nums []int) []int {
	var ans []int

	for i := range nums {
		leader := true

		for j := i + 1; j < len(nums); j++ {
			if nums[j] >= nums[i] {
				leader = false
				break
			}
		}

		if leader {
			ans = append(ans, nums[i])
		}
	}
	return ans
}

func main() {
	nums := []int{1, 2, 5, 3, 1, 2}

	ans := leader(nums)

	fmt.Print("Leaders in the array are: ")
	for _, v := range ans {
		fmt.Print(v, " ")
	}
	fmt.Println()

}
