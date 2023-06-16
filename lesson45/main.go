package main

import (
	"fmt"
)

//構造体
//メソッド

type User struct {
	Name string
	Age  int
}

// 構造体の関数を宣言するには、「func (レシーバー名 構造体の型)」とすることで構造体メソッドとして使用できる
func (u User) SayName() {
	fmt.Println(u.Name)
}

func (u User) SetName(name string) {
	u.Name = name
}

func (u *User) SetName2(name string) {
	u.Name = name
}

func main() {
	//構造体の型で暗黙的に宣言
	user1 := User{Name: "user1"}
	//構造体のメソッドを使用してコンソールに表示
	user1.SayName()

	//名前を変更する構造体メソッドを宣言
	user1.SetName("A")
	//値渡しため、処理を行っても実態の値は不動のため表示しても「user1」のまま
	user1.SayName()

	//参照渡しの構造体メソッドを宣言
	user1.SetName2("A")
	//参照渡しのため、処理後の値は変動する
	//この時、変数はポインタ型じゃなくても参照渡しが可能となる
	//uses1は普通のUser型である
	user1.SayName()

	//ポインタ型で代入した場合は当然のこと、実態の値が変動する
	user2 := &User{Name: "user2"}
	user2.SetName2("B")
	user2.SayName()
}
