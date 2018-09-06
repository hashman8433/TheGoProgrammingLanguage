package main

import "fmt"

func main() {

	word := "hello"
	word1 := word[0:2]
	word1 += word1
	fmt.Println(word)
	fmt.Println(word1)


	var runes []rune
	for _, r := range "hello world" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)

}
