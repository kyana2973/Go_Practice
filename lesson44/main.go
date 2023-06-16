package main

import (
	"fmt"
)

//構造体

type User struct {
	Name string
	Age  int
	// 以下のように２つまとめてフィールド名を定義することもできる
	//X, Y int
}

func UserUpdate(user User) {
	user.Name = "A"
	user.Age = 1000
}

func UserUpdate2(user *User) {
	user.Name = "A"
	user.Age = 1000
}

func main() {
	//以下は明示的な宣言
	var user User
	//Name,Ageの初期値は空文字と0になる
	fmt.Println(user)
	//以下のようにプロパティを指定して値を代入することができる
	user.Name = "user1"
	user.Age = 10
	fmt.Println(user)

	//暗黙的にも宣言することができる
	user2 := User{}
	fmt.Println(user2)
	user2.Name = "user2"
	user2.Age = 20
	fmt.Println(user2)

	//以下は暗黙的に宣言して、同時に初期を設定する方法
	user3 := User{Name: "user3", Age: 30}
	fmt.Println(user3)

	//以下のようにフィールド名なしでも代入できるが、structの順番通りに値を並べる必要がある
	user4 := User{"user4", 40}
	fmt.Println(user4)

	//順番を間違えて値を代入すると当然エラーとなる
	/*
		user5 := User{50, "user5"}
		fmt.Println(user5)
	*/

	user6 := User{Name: "user6"}
	fmt.Println(user6)

	user7 := new(User)
	fmt.Println(user7)

	user8 := &User{}
	fmt.Println(user8)

	//以下は値渡しのため、処理を実行しても実態の値は不動となる
	//※関数内では、変動した値を使用することができる
	UserUpdate(user)
	fmt.Println(user)
	//以下は参照渡しでアドレスを引数にしているため、処理後に実態の値が変動する
	UserUpdate2(user8)
	fmt.Println(user8)
}
