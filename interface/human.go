package main

import "fmt"

func main() {
	s := Student{Grade: 1, Major: "English", Human: Human{Name: "lily", Age: 3}}
	s.Eat()
	drinkers := []IDrink{s}
	for _, v := range drinkers {
		v.Drink()
	}

	eaters := []IEat{s}
	for _, v := range eaters {
		v.Eat()
	}
}

type Being struct {
}

type Human struct {
	Being
	Name string
	Age  int
}

func (h Human) Eat() {
	fmt.Println("human eating...")
	h.Drink()
}

func (h Human) Drink() {
	fmt.Println("human drinking...")
}

type Student struct {
	Human
	Grade int
	Major string
}

func (s Student) Eat() {
	fmt.Println("student eating...")
}

type IDrink interface {
	Drink()
}

type IEat interface {
	Eat()
}
