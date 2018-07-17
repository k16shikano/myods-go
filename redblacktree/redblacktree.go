package main

import (
	"fmt"
)

type T int

type N struct {
	left *N
	right *N
	parent *N
	data T
	color string
}

type RedBlackTree struct {
	root *N
}

var root *N
var n int

func main () {
	n = 0
	root = &N {left:nil, right:nil, parent:nil}
	add(1)
	add(2)
	fmt.Print(*root.right.right)
}

func depth (u *N) int {
	d := 0
	for u != root {
		u = u.parent
		d++
	}
	return d
}

func size (u *N) int {
	if (u == nil) {
		return 0
	} else {
		return 1 + size(u.left) + size(u.right)
	}
}

func height (u *N) int {
	if (u == nil) {
		return -1
	} else {
		if (height(u.left) > height(u.right)) {			
			return 1 + height(u.left)
		} else {
			return 1 + height(u.right)
		}
	}
}

func traverse (u *N) {
	if (u != nil) {
		fmt.Println(u)
		traverse(u.left)
		traverse(u.right)
	}
}

func compare (a T, b T) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	} else {
		return 0
	}
}

func findEQ(x T) T {
	var w *N
	w = root
	for w != nil {
		comp := compare(x, w.data)
		if  comp < 0 {
			w = w.left
		} else if comp > 0 {
			w = w.right
		} else {
			return w.data
		}
	}
	return 0
}

func find(x T) T {
	var w *N
	var z *N
	w = root
	z = nil
	for w != nil {
		comp := compare(x, w.data)
		if (comp < 0) {
			z = w
			w = w.left
		} else if comp > 0 {
			w = w.right
		} else {
			return w.data
		}
	}
	if z == nil {
		return 0
	} else {
		return z.data
	}
}

func add(x T) bool {
	var p *N
	p = findLast(x)
	u := &N {data:x}
	return addChild(p, u)
}

func findLast(x T) *N {
	var w *N
	var prev *N
	w = root
	prev = nil
	for w != nil {
		prev = w
		comp := compare(x, w.data)
		if (comp < 0) {
			w = w.left
		} else if (comp > 0) {
			w = w.right
		} else {
			return w
		}
	}
	return prev
}

func addChild(p *N, u *N) bool {
	if (p == nil) {
		root = u
	} else {
		comp := compare(u.data, p.data)
		if (comp <0) {
			p.left = u
		} else if (comp > 0) {
			p.right = u
		} else {
			return false
		}
		u.parent = p
	}
	n++
	return true
}
	
func splice (u *N) {
	var s *N
	var p *N
	if (u.left != nil) {
		s = u.left
	} else {
		s = u.right
	}
	if (u == root) {
		root = s
		p = nil
	} else {
		p = u.parent
		if (p.left == u) {
			p.left = s
		} else {
			p.right = s
		}
	}
	if (s != nil) {
		s.parent = p
	}
	n--
}

func remove(u *N) {
	if (u.left == nil || u.right == nil) {
		splice(u)
	} else {
		var w *N
		w = u.right
		for w.left != nil {
			w = w.left
		}
		u.data = w.data
		splice(w)
	}
}

