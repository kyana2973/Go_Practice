package main

import (
	"fmt"
)

func main() {
	var i int = 100
	fmt.Printf("%T\n", i)
	fmt.Printf("%T\n", int32(i))
	var i2 int32 = 150
	fmt.Println(i + int(i2))
}
