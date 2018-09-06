package main

import "fmt"

func reverse(s []int) {
	for i, j := 0, len(s) - 1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {

	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a)

	s := []int{0, 1, 2, 3, 4, 5}
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)
	fmt.Println(s)

	s1 := []string{"word", "hello", "354"}
	fmt.Println(equal(s1[2:], s1[2:]))

	var s2 []int // len(s) == 0, s == nil
	fmt.Println(len(s2) == 0, s2 == nil)
	s2 = nil // len(s) == 0, s == nil
	fmt.Println(len(s2) == 0, s2 == nil)
	s2 = []int(nil) // len(s) == 0, s == nil
	fmt.Println(len(s2) == 0, s2 == nil)
	s2 = []int{} // len(s) == 0, s != nill
	fmt.Println(len(s2) == 0, s2 != nil)
}

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}