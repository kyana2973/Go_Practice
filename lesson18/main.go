package main

//クロージャー

import (
	"fmt"
)

func Later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func main() {
	f := Later()
	fmt.Println(f("Hello"))
	fmt.Println(f("My"))
	fmt.Println(f("Name"))
	fmt.Println(f("Is"))
	fmt.Println(f("Golang"))
}
