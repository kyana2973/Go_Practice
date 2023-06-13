package main

import (
	"fmt"
)

//slice for

func main() {
	sl := []string{"A", "B", "C"}
	fmt.Println(sl)

	/*
		//インデックス番号含めた表示
		for i, v := range sl {
			fmt.Println(i, v)
		}

		//インデックス番号含めない表示
		for _, v := range sl {
			fmt.Println(v)
		}
	*/

	for i := 0; i < len(sl); i++ {
		fmt.Println(sl[i])
	}
}
