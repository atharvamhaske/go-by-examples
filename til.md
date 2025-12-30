today i learned: len vs capacity

s := make([]int, 0, 5)
fmt.Println(len(s)) // 0
fmt.Println(cap(s)) // 5

// imp
s := []int{}
s = append(s, 1, 2)
//here if capacity is exceeded then new array is allocated and old data is copied so O(n)

left slice and right slice in go

for ex:

arr := []int{1, 2, 3, 4, 5, 6}

1. left slice [:high] //rule : start from 0 and end (exculsive)
ex: s := arr[:3] //here start 0 and end 3 so output s = [1, 2, 3]

2. right slice [low:]
ex: s:= arr[2:] //start from 2 and end len(arr)

//len  = high - low
//cap  = cap(original) - low
