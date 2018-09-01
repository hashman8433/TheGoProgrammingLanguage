package main

import (
	"fmt"
	"os"
)

func main() {

	s, sep := "", ""
	for index, arg := range os.Args[0:] {
		fmt.Println("index:", index, " value:", arg)
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)

}
