package main

import "fmt"

type List struct {
	head *Node
}

type Node struct {
	value int
	next *Node
}

func (l *List) Insert(v int) *List {
	if l == nil {
		fmt.Println("List is not created yett...")
		return nil
	}
	temp := new(Node)
	temp.value = v
	temp.next = l.head
	l.head = temp
	return l
}

func (l *List) Sum() int {
	if l == nil {
		fmt.Println("List is not created yett...")
		return 0
	}
	
	temp := l.head
	sum := 0
	for temp != nil {
		sum += temp.value
		temp = temp.next
	}
	return sum
}

func (list *List) Print() {
	if list == nil {
		fmt.Println("list is not created.")
		return
	}
	temp := list.head
	for temp != nil {
		fmt.Print(temp.value, " ")
		temp = temp.next
	}
}

func main() {
	lst := new(List)
	lst.Insert(1)
	lst.Insert(2)
	lst.Insert(3)
	lst.Insert(1)
	lst.Insert(2)
	lst.Insert(3)
	fmt.Println(lst.Sum())
	lst.Print()
}