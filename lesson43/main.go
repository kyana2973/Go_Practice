package main

import (
	"fmt"
)

//Pointer

func Double(i int) {
	i = i * 2
}

func DoubleV2(i *int) {
	*i = *i * 2
}

func DoubleV3(s []int){
	for i, v := range s{
		s[i] = v * 2
	}
}

func main() {
	var n int = 100
	fmt.Println(n)
	//アドレスを表示させるには「&」を手前につける
	fmt.Println(&n)

	//nはポインタ型ではないため関数に渡す際に別物として処理される
	Double(n)
	fmt.Println(n)

	//ポインタ型の宣言
	var p *int = &n
	//nのアドレスを参照しているので同じアドレスとなる
	fmt.Println(p)
	//アドレスの値を参照する時は「*」を手前につけると表示できる
	fmt.Println(*p)

	/*
		*p = 300
		fmt.Println(n)
		n = 200
		fmt.Println(*p)
	*/

	//nのアドレスを渡す、関数によって２倍になる
	DoubleV2(&n)
	fmt.Println(n)

	//nのアドレスが格納されているpを渡しても関数によって更に倍になる
	DoubleV2(p)
	fmt.Println(*p)

	var sl []int = []int{1,2,3}
	DoubleV3(sl)
	fmt.Println(sl)
}
