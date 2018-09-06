package main

import "fmt"

func reverse(slice *[]int) *[]int {
	intArray := *slice
	for i, j := 0, len(intArray )-1; i < j; i, j = i+1, j-1 {
		intArray[i], intArray[j] = intArray[j], intArray[i]
	}
	return &intArray
}

func main() {
	numbers := []int{0, 1, 2, 3, 4}
	reverse(&numbers)
	fmt.Printf("%d", numbers)
}
