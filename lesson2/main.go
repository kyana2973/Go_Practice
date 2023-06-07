package main

import (
	"fmt"
)

// i5 := 450
var i5 int = 450

func outer() {
	var s4 string = "outer"
	fmt.Println(s4)
}

func main() {
	//明示的な定義
	//var 変数名 型 = 値
	var i int = 100
	fmt.Println(i)

	var s string = "Hello Go!"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int    = 200
		s2 string = "Golang"
	)
	fmt.Println(i2, s2)

	var (
		i3 int
		s3 string
	)
	fmt.Println(i3, s3)

	i3 = 1000
	s3 = "GoGo"
	fmt.Println(i3, s3)

	i = 5000
	fmt.Println(i)

	//暗黙的な定義
	//変数名 := 値
	i4 := 400
	fmt.Println(i4)

	fmt.Println(i5)

	outer()

	s5 := "Not Use"
	fmt.Println(s5)
}
