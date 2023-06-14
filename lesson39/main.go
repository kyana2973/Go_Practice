package main

import (
	"fmt"
	"time"
)

//channel
//close

func reciever(name string, ch chan int) {
	for {
		i, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + "END")
}

func main() {
	ch1 := make(chan int, 2)

	/*
		ch1 <- 1
		close(ch1)

		//ch1 <- 1

		i, ok := <-ch1
		fmt.Println(i, ok)

		i2, ok := <-ch1
		fmt.Println(i2, ok)
	*/

	//ch1に送られた値を以下３つのゴルーチンで処理する
	//処理する順番はランダム
	go reciever("1.goroutin", ch1)
	go reciever("2.goroutin", ch1)
	go reciever("3.goroutin", ch1)

	i := 0
	//このfor文でch1に値を送信する。それより受信を待機しているゴルーチンのreciever関数が起動する。
	//起動する順番はランダム
	for i < 100 {
		ch1 <- i
		i++
	}
	//最後にi=100となったらch1をクローズする。その際、ch1にクローズした時のch1のキューを送信する。
	close(ch1)
	//最後に送ったキューは中身なしで、クローズしているので「...END」と表示される。
	//以下のtime関数は、最後のgoroutinが完了するのを待つための処理。
	//今回は後続で行う処理はないが、ある場合はここで待ってから次の処理に移行する必要がある。
	time.Sleep(3 * time.Second)
}
