package main

import (
	"fmt"
)

// For文

func main() {
	/*
		i := 0
		for {
			i++
			if i==3{
				break
			}
			fmt.Println("roop")
		}
	*/

	/*
		point := 0
		for point < 10 {
			fmt.Println(point)
			point++
		}
	*/

	/*
		for i := 0; i < 10; i++ {
			if i == 3 {
				continue
			}
			if i == 6 {
				break
			}
			fmt.Println(i)
		}
	*/

	// arr := [3]int{1,2,3}
	// for i := 0; i < len(arr);i++{
	// 	fmt.Println(arr[i])
	// }

	/*
	arr := [3]int{1, 2, 3}
	for i, v := range arr {
		fmt.Println(i, v)
	}

	sl := []string{"Python", "PHP", "GO"}
	for _, v := range sl {
		fmt.Println(v)
	}
	*/

	m := map[string]int{"apple":100,"banana":200}
	for k,_ := range m{
		fmt.Println(k)
	}
}
