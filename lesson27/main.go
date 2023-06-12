package main

import (
	"fmt"
)

//パニックとリカバーについて

func main() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()
	panic("runtaime error")
	fmt.Println("START")
}
