package main

import (
	"fmt"
	"strconv"
)

// IF文
// エラーハンドリング

func main() {
	var s string = "hello"

	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("i = %T\n", i)
	fmt.Println(i)
}
