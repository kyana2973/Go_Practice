package main

import (
	"fmt"
)

//map

func main() {
	var m = map[string]int{"A":100, "B":200}
	fmt.Println(m)

	m2 := map[string]int{"A":100, "B":200}
	fmt.Println(m2)

	m3 := map[int]string{
		1: "A",
		2: "B",
	}
	fmt.Println(m3)

	m4 := make(map[int]string)
	fmt.Println(m4)
	m4[1] = "Japan"
	m4[2] = "USA"
	fmt.Println(m4)

	fmt.Println(m["A"])
	fmt.Println(m4[1])
	//要素がないインデックスを指定しても空文字で表示される
	fmt.Println(m4[3])

	//空文字を表示されて困る場合はエラーハンドリングを行うこともできる
	//sに値を代入すると同時に、第２引数で代入に成功したかの判定結果を取得する
	s, ok := m4[3]
	//判定により処理を分岐させることができる
	if !ok{
		fmt.Println("error")
	}else{
		fmt.Println(s, ok)
	}

	m4[2] = "US"
	fmt.Println(m4)
	m4[3] = "CHINA"
	fmt.Println(m4)

	//削除関数を使って要素を削除
	delete(m4, 3)
	fmt.Println(m4)

	fmt.Println(len(m4))
}
