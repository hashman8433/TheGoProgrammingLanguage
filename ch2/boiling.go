package main

import "fmt"

const bilingF = 212.0

func main() {
	var f = bilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("biling point = %g°F or %g°C\n", f, c)
}
