package main

import (
	"fmt"
)

// switch文
// 型スイッチ

func anything(a interface{}){
	//fmt.Println(a)
	switch v := a.(type){
	case string:
		fmt.Println(v + "!?")
	case int:
		fmt.Println(v + 10000)
	case bool:
		fmt.Println(v)
	default:
		fmt.Println("I don't Know")
	}
}

func main() {
	anything("aaa")
	anything(1)
	anything(false)

	var x interface{} = 3
	i, isInt := x.(int)
	fmt.Println(i, isInt)

	f, isFoart64 := x.(float64)
	fmt.Println(f, isFoart64)

	if x == nil{
		fmt.Println("None")
	}else if i ,isInt := x.(int); isInt{
		fmt.Println(i, "x Is Int")
	}else if s, isString := x.(string); isString{
		fmt.Println(s, "x Is String")
	}else{
		fmt.Println("I don't Know")
	}

	switch x.(type){
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("I don't Know")
	}

	switch v := x.(type){
	case int:
		fmt.Println(v, "int")
	case string:
		fmt.Println(v, "string")
	case bool:
		fmt.Println(v, "bool")
	default:
		fmt.Println("I don't Know")
	}
}
