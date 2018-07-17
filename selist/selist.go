package main

import (
	"fmt"
	. "./arraydeque"
)

//type T int

var n int
var j int
var b int // size of bounded deque
var dummy *SEList
var se *SEList

func main() {
	b = 3
	dummy = NewSEList(b)
	se = NewSEList(b)
//	se.Set(1, "b")
//	se.Set(2, "c")
//	se.Set(3, "d")
	se.Add(4, "e")
	se.AddToLast("f")
	se.Set(0, "a")
//	fmt.Print(n)
	se.fmtdllist()
}

func (se SEList) fmtdllist () {
	fmt.Print("{")
	for i := 0; i < n*b-1; i++ {
		fmt.Printf("%s, ", se.Get(i))
	}
	fmt.Printf("%s}\n", se.Get(n*b-1))
}

type BDeque struct {
	AD *ArrayDeque
}

type SEList struct {
	d BDeque
	prev *SEList
	next *SEList
}

type Location struct {
	u *SEList
	j int
}

func NewLocation(u *SEList, j int) Location {
	return Location{u:u, j:j}
}

func NewBDeque(b int) BDeque {
	n = 0
	j = 0
	var a ArrayDeque
	var z BDeque
	a = ArrayDeque{make([]T, b+1)}
	z = BDeque{AD: &a}
	return z
}

func (bd BDeque) Add(i int, x T) {
	bd.AD.Add(i,x)
}

func (bd BDeque) Size() int {
	return n //len(bd.AD.Array)
}

func NewSEList(b int) *SEList {
	bd := NewBDeque(b)
	return &SEList{d:bd}
}

func getLocation(i int, ell Location) Location {
//	fmt.Printf("i:%d, n:%d\n",i,n)
	if (i < n / 2) {
		var u *SEList
		u = dummy.next
		for i >= u.d.Size() {
			i -= u.d.Size()
			u = u.next
		}
		ell.u = u
		ell.j = i
	} else {
		var u *SEList
		u = dummy
		idx := n
		for i < idx {
			u = u.prev
			idx -= u.d.Size()
		}
		ell.u = u
		ell.j = i - idx
	}
	return ell
}

func addBefore(w *SEList) *SEList {
	u := NewSEList(b)
	u.prev = w.prev
	u.next = w
	u.next.prev = u
	if w != dummy {
		u.prev.next = u
	}
	return u
}

func removeNode(w *SEList) {
	w.prev.next = w.next
	w.next.prev = w.prev
}

func (node SEList) Get(i int) T {
	var l Location
	l = getLocation(i, l)
	return l.u.d.AD.Get(l.j)
}

func (node SEList) Set(i int, x T) T {
	var l Location
	l = getLocation(i, l)
	var y T
	y = l.u.d.AD.Get(l.j)
	l.u.d.AD.Set(l.j, x)
	return y
}

func (node SEList) AddToLast(x T) {
	var last *SEList
	last = dummy.prev
	if (last == nil || last.d.Size() == b+1) {
		last = addBefore(dummy)
	}
	last.d.AD.Add(node.d.Size(), x)
	n++
}

func (node SEList) Add(i int, x T) {
	if (i == n) {
		node.AddToLast(x)
	} else {
		var l Location
		l = getLocation(i, l)
		var u *SEList
		u = l.u
		r := 0
		for r < b && u != nil && u.d.Size() == b+1 {
			u = u.next
			r++
		}
		if (r == b) {
			spread(l.u)
			u = l.u
		}
		if (u == nil) {
			u = addBefore(u)
		}
		for (u != l.u) {
			u.d.AD.Add(0, u.prev.d.AD.Remove(u.prev.d.Size()-1))
			u = u.prev
		}
		u.d.AD.Add(l.j, x)
		n++
	}
	
}

func (node SEList) Remove(i int) T {
	var l Location
	getLocation(i, l)
	var y T
	y = l.u.d.AD.Get(l.j)
	var u *SEList
	u = l.u
	r := 0
	for r < b && u != dummy && u.d.Size() == b-1 {
		u = u.next
		r++
	}
	if (r == b) {
		gather(l.u)
	}
	u = l.u
	u.d.AD.Remove(l.j)
	for u.d.Size() < b-1 && u.next != dummy {
		u.d.AD.Add(u.d.Size(), u.next.d.AD.Remove(0))
		u = u.next
	}
	if (u.d.Size() == 0) {
		removeNode(u)
	}
	n--
	return y
}
	
func spread(u *SEList) {
	var w *SEList
	w = u
	for j := 0; j < b; j++ {
		w = w.next
	}
	w = addBefore(w)
	for w != u {
		for w.d.Size() < b {
			w.d.AD.Add(0, w.prev.d.AD.Remove(w.prev.d.Size()-1))
		}
		w = w.prev
	}	
}

func gather(u *SEList) {
	var w *SEList
	w = u
	for j :=0; j < b-1; j++ {
		for w.d.Size() < b {
			w.d.AD.Add(w.d.Size(), w.next.d.AD.Remove(0))
		}
		w = w.next
	}
	removeNode(w)
}
	
