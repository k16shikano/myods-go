package arraydeque

import(
//	"fmt"
)

type T string

type ArrayDeque struct {
	Array []T
}

var a ArrayDeque
var j int
var n int

func main() {
//	// figure 2.3
	a = ArrayDeque{make([]T, 12)}
	a.Add(0, "a")
//	a.add(1, "b")
//	a.add(2, "c")
//	a.add(3, "d")
//	a.add(4, "e")
//	a.add(5, "f")
//	a.add(6, "g")
//	a.add(7, "h")
//	fmt.Println(a)
//	a.remove(2)
//	fmt.Println(a)
//	a.add(4, "x")
//	fmt.Println(a)
//	a.add(3, "y")
//	fmt.Println(a)
//	a.add(3, "z")
//	fmt.Println(a)
}

func (ad ArrayDeque) Get(i int) T {
//	fmt.Print(ad.Array)
	return ad.Array[(j + i) % len(ad.Array)]
}

func (ad ArrayDeque) Set(i int, x T) T{
	y := ad.Array[(j + i) % len(ad.Array)]
	ad.Array[(j + i) % len(ad.Array)] = x
	return y
}

func resize() []T {
	b := make([]T, 2*n)
	b = a.Array
	return b
}

func (ad ArrayDeque) Add(i int, x T) {
	if (n + 1 > len(ad.Array)) {
		ad.Array = resize()
	}
	if (i < n/2) {
		if (j == 0) {
			j = len(ad.Array) - 1
		} else {
			j = j - 1
		}
		for k := 0; k <= i-1; k++ {
			ad.Array[(j+k) % len(ad.Array)] = ad.Array[(j+k+1) % len(ad.Array)]
		}
	} else {
		for k := n; k > i; k-- {
			ad.Array[(j+k) % len(ad.Array)] = ad.Array[(j+k-1) % len(ad.Array)]
		}
	}
	ad.Array[(j+i) % len(ad.Array)] = x
	n++
}

func (ad ArrayDeque) Remove(i int) T {
	x := ad.Array[(j+i) % len(ad.Array)]
	if (i < n/2) {
		for k := i; k > 0; k-- {
			ad.Array[(j+k) % len(ad.Array)] = ad.Array[(j+k-1) % len(ad.Array)]
		}
		j = (j + 1) % len(ad.Array)
	} else {
		for k := i; k < n-1; k++ {
			ad.Array[(j+k) % len(ad.Array)] = ad.Array[(j+k+1) % len(ad.Array)]
		}
	}
	n--
	if (3*n < len(ad.Array)) {
		resize()
	}
	return x
}
