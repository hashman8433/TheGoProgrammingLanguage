package main

import "fmt"

func main() {

	w := Wheel{Circle{Point{8, 8}, 5}, 20}

	w = Wheel{
		Circle: Circle{
			Center: Point{X: 8, Y: 8},
			Redius: 5,
		},
		Spokes: 20,
	}

	fmt.Println(w)
}
