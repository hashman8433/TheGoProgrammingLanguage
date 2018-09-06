package main

import (
	"fmt"
)

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func appendInts(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	fmt.Println(zlen)
	// ...expand z to at least zlen...
	copy(z[len(x):], y)
	return z
}

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap%d\t%v\n", i, cap(y), y)
		x = y
	}
	x[5] = 6
	fmt.Printf("%d cap%d\t%v\n", 9, cap(y), y)

	x1 := []int{1, 1, 1, 1}
	x2 := appendInt(x1[1:], 1)
	x1[2] = 9
	fmt.Printf("%d cap%d\t%v\n", 9, cap(x2), x2)

	x1 = append(x1, 5, 6)
	fmt.Printf("%d cap%d\t%v\n", 9, cap(x1), x1)

}