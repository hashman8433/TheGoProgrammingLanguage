package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Center Point
	Redius int
}

type Wheel struct {
	Circle Circle
	Spokes int
}

func main() {

	var w Wheel
	w.Circle.Center.X = 8
	w.Circle.Center.Y = 8
	w.Circle.Redius = 5
	w.Spokes = 20
	fmt.Println("main", w)

	wheel1()
	wheel2()
	wheel3()
}

func wheel1() {

	var w Wheel
	w.Circle.Center.X = 8
	w.Circle.Center.Y = 8
	w.Circle.Redius = 5
	w.Spokes = 20
	fmt.Println("wheel1:", w)
}

func wheel2() {

	var w Wheel
	w.Circle.Center.X = 8
	w.Circle.Center.Y = 8
	w.Circle.Redius = 5
	w.Spokes = 20
	fmt.Println("wheel2:", w)
}

func wheel3() {

	w := Wheel{Circle{Point{8, 8}, 5}, 20}

	w = Wheel{
		Circle: Circle{
			Center: Point{X: 8, Y: 8},
			Redius: 5,
		},
		Spokes: 20,
	}

	w.X = 42
	fmt.Println("wheel3:", w)
}
