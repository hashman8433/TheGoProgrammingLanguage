package main

import "fmt"

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}


func main() {

	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", data)
	//data = nonempty(data) // method 3
	data = nonempty2(data) // method 2
	//fmt.Printf("%q\n", nonempty(data)) // method 1
	fmt.Printf("%q\n", data)
}
