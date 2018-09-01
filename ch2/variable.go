package main

import "fmt"

func main() {
	p := new(int)
	fmt.Println(*p)
	*p = 2
	fmt.Println(*p)

	count := []int{1, 2, 6}
	x := 1
	count[x] = count[x] * 6
	fmt.Println("count[x]: ", count[x])

	v := 1
	v++
	v--
	fmt.Println("fib: ", fib(5))
}

// greatest common divisor
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func fib(n int) int {
	x, y := 0, 0
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
