package main

import "fmt"

type node struct {
	data int
	next *node
}

func main() {
	a := []int{1, 2, 3, 4, 3, 4, 5}
	var head *node
	for _, v := range a {
		n := node{data: v}
		head = add(head, n)
	}
	for head != nil {
		fmt.Print(head.data, "->")
		head = head.next
	}

}

func add(head *node, n2 node) *node {
	if head == nil {
		head = &n2
		return head
	}
	temp := head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = &n2
	return head
}
