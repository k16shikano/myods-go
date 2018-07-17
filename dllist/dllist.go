package main

import (
	"fmt"
)

type Node struct {
	X int
	Next *Node
	Prev *Node
}

var dummy Node
var n int = 0

func fmtdllist () {
	fmt.Print("{")
	for i := 0; i < n-1; i++ {
		fmt.Printf("%d, ", get(i))
	}
	fmt.Printf("%d}\n", get(n-1))
}

func main() {
	dllist()
	add(0, 3)
	fmt.Println(get(0))
	add(1, 2)
	set(0, 10)
	remove(1)
	fmtdllist()
}

func dllist() {
	dummy.Next = &dummy
	dummy.Prev = &dummy
	n = 0
}

func getNode(i int) *Node {
	var p *Node
	if (i < n/2) {
		p = dummy.Next
		for j := 0; j < i; j++ {
			p = p.Next
		}
	} else {
		p = &dummy
		for j := n; j > i; j-- {
			p = p.Prev
		}
	}
	return p
}

func get(i int) int {
	return getNode(i).X
}

func set(i int, x int) int {
	u := getNode(i)
	y := u.X
	u.X = x
	return y
}

func addBefore(w *Node, x int) *Node {
	u := &Node{X:x, Prev:w.Prev, Next:w}
	u.Next.Prev = u
	u.Prev.Next = u
	n++
	return u
}

func add(i int, x int) {
	addBefore(getNode(i), x)
}

func removeNode (w *Node) {
	w.Prev.Next = w.Next
	w.Next.Prev = w.Prev
	n--
}
	
func remove(i int) int {
	w := getNode(i)
	x := w.X
	removeNode(w)
	return x
}
