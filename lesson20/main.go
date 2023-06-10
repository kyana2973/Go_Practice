package main

//クロージャー

import (
	"fmt"
)

func main() {
	a := 0
	if a == 2 {
		fmt.Println("two")
	} else if a == 1 {
		fmt.Println("one")
	} else {
		fmt.Println("I don't know")
	}

	if b := 100; b == 200{
		fmt.Println("one hundred")
	}

	x := 0
	if x := 10; true{
		fmt.Println(x)
	}
	fmt.Println(x)
}
