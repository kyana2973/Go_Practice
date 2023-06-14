package main

import (
	"fmt"
)

//map

func main() {
	m := map[string]int{
		"Apple":  100,
		"Banana": 200,
	}

	//欲しくない情報はアンダーバーにすることで取得しないことも可能
	//以下はキー、バリューどっちも取得するパターン
	for k, v := range m {
		fmt.Println(k, v)
	}

	//以下はバリューのみ取得するパターン
	for _, v := range m {
		fmt.Println(v)
	}
}
