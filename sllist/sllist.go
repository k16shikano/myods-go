package main

import (
	"fmt"
)

type Node struct {
	X int
	Next *Node
}

var head *Node
var tail *Node
var n int = 0

func fmtsllist (u Node) {
	fmt.Print("{")
	for i := n; i > 0; i-- {
		if u.Next != nil {
			fmt.Printf("%d, ", u.X)
			u = *u.Next
		} else {
			fmt.Printf("%d}\n", u.X)
		}
	}
}

func main() {
	push(10)
	push(1)
	remove()
	add(3)
	push(4)
	fmtsllist(*head)
}

func push(x int) int {
	w := head
	u := Node {X:x, Next:w}
	head = &u
	if (n == 0) {
		tail = &u
	}
//	fmt.Println(*head)
//	fmt.Println(*tail)
	n++
	return x
}

func pop() int {
	if (n == 0) {
		return 0
	}
	x := head.X
	head = head.Next
	n = n - 1
	if (n == 0) {
		tail = &Node {}
	}
	return x
}
		
func remove() int {
	if (n == 0) {
		return 0
	}
	x := head.X
	head = head.Next
	n = n - 1
	if (n == 0) {
		tail = &Node {}
	}
	return x
}
	
func add(x int) bool {
	u := Node {X:x, Next:nil}
	if (n == 0) {
		head = &u
	} else {
		tail.Next = &u
	}
	tail = &u
	n++
	return true
}
