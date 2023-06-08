package main

import (
	"fmt"
	//"strconv"
)

func main() {
	/*
		var i int = 1
		fmt.Println(i)
		fl64 := float64(i)
		fmt.Println(fl64)
		fmt.Printf("i = %T\n", i)
		fmt.Printf("fl = %T\n", fl64)

		var i2 int = int(fl64)
		fmt.Println(i2)
		fmt.Printf("i = %T\n", i)

		fl := 10.5
		i3 := int(fl)
		fmt.Println(i3)
		fmt.Printf("i3 = %T\n", i3)
	*/

	/*
		var s string = "100"
		fmt.Printf("s = %T\n", s)

		i, _ := strconv.Atoi(s)
		fmt.Println(i)
		fmt.Printf("i = %T\n", i)

		i2 := 200
		s2 := strconv.Itoa(i2)
		fmt.Println(s2)
		fmt.Printf("s2 = %T\n", s2)
	*/

	var h string = "Hello World"
	b := []byte(h)
	fmt.Println(b)
	fmt.Printf("b = %T\n", b)

	h2 := string(b)
	fmt.Println(h2)
	fmt.Printf("h2 = %T\n", h2)

}
