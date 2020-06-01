package main

import (
	"fmt"
	"os/user"
	//https://ashitani.jp/golangtips/tips_time.html//日付を時刻
	"time"//時刻を取得するライブラリ
)

func init() {//一番最初に実行される特別な関数
	fmt.Println("Init");
}

func buzz() {//main()で実行しなければ、表示されない
	fmt.Println("BUZZ");
}

func main() {
	buzz();
	fmt.Println("Hello, World!", "TEST TEST", time.Now().Year())
	fmt.Println("Hello, World!")
	fmt.Println(user.Current())
}