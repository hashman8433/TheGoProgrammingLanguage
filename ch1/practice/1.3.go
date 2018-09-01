package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[0:], " "))
	fmt.Println(os.Args[0:])

	s := []string{"a", "b", "c"}
	fmt.Println(strings.Join(s[0:], " "))
	fmt.Println(s)
}
