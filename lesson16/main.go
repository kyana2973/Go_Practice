package main

import (
	"fmt"
)

func returnFunc() func(){
	return func() {
		fmt.Println("I am Function")
	}
}

func main(){
	f := returnFunc()
	f()
}
