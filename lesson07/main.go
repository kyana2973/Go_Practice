package main

import (
	"fmt"
)

func main() {
	ByteA := []byte{72, 73}
	fmt.Println(ByteA)

	fmt.Println(string(ByteA))

	c := []byte("HI")
	fmt.Println(c)
	fmt.Println(string(c))

}
