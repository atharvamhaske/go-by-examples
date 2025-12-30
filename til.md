today i learned: len vs capacity

s := make([]int, 0, 5)
fmt.Println(len(s)) // 0
fmt.Println(cap(s)) // 5

// imp
s := []int{}
s = append(s, 1, 2)
//here if capacity is exceeded then new array is allocated and old data is copied so O(n)