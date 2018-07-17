package main

var a [] int
var n int

func get(i int) int {
	return a[i]
}

func set(i int, x int) int {
	y := a[i]
	a[i] = x
	return y
}

func add(i int, x int) {
	if (n+1 > len(a)) {
		resize()
	}
	for j := n; j > i; j-- {
		a[j] = a[j-1]
	}
	a[i] = x
	n++
}

func remove(i int) int {
	var x int = a[i]
	for j := i; j < n-1; j++ {
		a[j] = a[j+1]
	}
	n--
	if (len(a) >= 3*n) {
		resize()
	}
	return x
}

func resize() {
	newlen := max(n*2, 1)
	b := make([]int, newlen)
	for i := 0; i < n; i++ {
		b[i] = a[i]
	}
	a = b
}

func max(a, b int) int {
   if a < b {
      return b
   }
   return a
}

