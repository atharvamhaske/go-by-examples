// package main

// import "fmt"

// func leader(nums []int) []int {
// 	var ans []int

// 	for i := range nums {
// 		leader := true

// 		for j := i + 1; j < len(nums); j++ {
// 			if nums[j] >= nums[i] {
// 				leader = false
// 				break
// 			}
// 		}

// 		if leader {
// 			ans = append(ans, nums[i])
// 		}
// 	}
// 	return ans
// }

// func main() {
// 	nums := []int{1, 2, 5, 3, 1, 2}

// 	ans := leader(nums)

// 	fmt.Print("Leaders in the array are: ")
// 	for _, v := range ans {
// 		fmt.Print(v, " ")
// 	}
// 	fmt.Println()

// }

// Optimal solution:

	package main

	import "fmt"

	func leader(nums []int) []int {
		n := len(nums)

		if n == 0 {
			return nil
		}

		var ans []int

		currLdr := nums[n-1]
		ans = append(ans, currLdr)

		for i := n - 2; i >= 0; i-- {
			if nums[i] > currLdr {
				currLdr = nums[i]
				ans = append(ans, currLdr)
			}
		}

		//reverse the answer to maintain order
		for l, r := 0, len(ans)-1; l < r; l, r = l+1, r-1 {
			ans[l], ans[r] = ans[r], ans[l]
		}

		return ans
	}

	func main() {
		nums := []int{1, 2, 5, 3, 1, 2}

		ans := leader(nums)

		fmt.Print("Leaders in array are: ")
		for _, val := range ans {
			fmt.Print(val, " ")
		}
		fmt.Println()
	}
