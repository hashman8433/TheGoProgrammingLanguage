package main

import (
	"fmt"
	"sort"
)

func main() {

	ages := map[string]int {
		"yemi": 29,
		"alice": 31,
		"charlie": 34,
	}
	fmt.Println(ages["alice"])
	delete(ages, "alice")
	ages["bob"] += 1
	ages["bob"]++
	fmt.Println(ages["bob"])
	//_ = &ages["bob"]

	var names []string
	for name := range ages {
		names = append(names, name)
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
	sort.Strings(names)
	fmt.Println("after sort string +++++++++++")
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}

	example3()
}

func example1() {
	ages := make(map[string]int)
	ages["alice"] = 31
	ages["charlie"] = 34
}

func example2() {
	var ages map[string]int
	fmt.Println(ages == nil)
	fmt.Println(len(ages) == 0)
	ages["carol"] = 21 // panic: assignment to entry in nil map
}

func example3() {
	ages := map[string]int {
		"yemi": 29,
		"alice": 31,
		"charlie": 34,
	}
	age, ok := ages["bobo"]
	if !ok {
		fmt.Println("yemi is not a key in this map; age == 0:" ,age == 0)

	}
}