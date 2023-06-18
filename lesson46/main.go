package main

import (
	"fmt"
)

//構造体
//埋め込み

// Tの構造体にUserの構造体を埋め込む
type T struct {
	//フィールド名と型名が同じなのはGoではよく見られるパターン
	User User
	//以下のように型名を省略してフィールド名だけで宣言もできる
	//User
}

type User struct {
	Name string
	Age  int
}

func (u *User) SetName() {
	u.Name = "A"
}

func main() {
	//TのUser構造体のName、Ageフィールドに初期値を代入
	t := T{User: User{Name: "user1", Age: 10}}

	//tを表示
	fmt.Println(t)
	//tの中にあるUserの構造体を表示
	fmt.Println(t.User)
	//tの中のUser構造体のNameを表示
	fmt.Println(t.User.Name)
	//T構造体でUserの型名を省略した場合のみ以下のように、tから直にNameを参照できる
	//fmt.Println(t.Name)

	//tの中のUser構造体のメソッド「SetName」を実行
	t.User.SetName()
	//参照渡しなので、値が変動する
	fmt.Println(t.User.Name)
}
