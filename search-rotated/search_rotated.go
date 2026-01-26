package main

import "fmt"

func search(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n-1

	for l <= r {
		mid := l + (r-l)/2

		if nums[mid] == target {
			return mid
		}

		//if left part is sorted then we do binary search on left part
		if nums[l] <= nums[mid] {
			if nums[l] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}

		//when right part is sorted then we do binary search on right part of array
		if nums[mid] <= nums[r] {
			if nums[mid] < target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}

	}
	//when we don't find a target return -1
	return -1
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0

	idx := search(nums, target)

	if idx == -1 {
		fmt.Printf("Target %d not found at index %d\n", target, idx)
	} else {
		fmt.Printf("Target %d found at %d\n", target, idx)
	}
}
