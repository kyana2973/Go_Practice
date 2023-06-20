package main

import (
	"fmt"
)

//構造体
//コンストラクタ

type User struct {
	Name string
	Age  int
}

// コンストラクタを作成
func NewUser(name string, age int) *User {
	return &User{Name: name, Age: age}
}

func main() {
	//NewUserによってポインタ型で代入している
	user1 := NewUser("user1", 10)

	//普通に表示するとアドレスから値を参照して表示する
	fmt.Println(user1)

	//以下は、実態にアクセスして表示している
	fmt.Println(*user1)
}
