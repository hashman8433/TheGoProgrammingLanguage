package main

import "fmt"

func main() {
	var u uint8 = 255
	fmt.Println(u, u+1, u*u)

	var i int8 = 127
	fmt.Println(i, i+1, i*i)

	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2
	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)


	fmt.Printf("%08b\n", x&y)
	fmt.Printf("%08b\n", x|y)
	fmt.Printf("%08b\n", x^y)
	fmt.Printf("%08b\n", x&^y)

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 {
			fmt.Println(i)
		}
	}

	fmt.Printf("%08b\n", x<<1)
	fmt.Printf("%08b\n", x>>1)

	medals := []string{"gold", "silver", "bronze"}
	for i := len(medals ) - 1; i > 1; i-- {
		fmt.Println(medals [i])
	}

	var apples int32 = 1
	var oranges int16 = 2
	//var compote int = apples + oranges // compile error
	var compote int = int(apples) + int(oranges)
	fmt.Println("compote:", compote )

	f := 3.141
	interger := int(f)
	fmt.Println(f, interger )

	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o)
	x1 := int64(0xdeadbeff)
	fmt.Printf("%d %[1]x %#[1]X\n", x1)

	ascii := 'a'
	unicode := '国'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)
	fmt.Printf("%d %[1]c %[1]q\n", unicode )
	fmt.Printf("%d %[1]q\n", newline )

}
