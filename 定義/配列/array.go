package main

import(
	"fmt"
)


func main(){
	var a[2]int
	a[0] = 100
	a[1] = 200
	fmt.Println(a)

	var b[2]string
	b[0] = "100"
	b[1] = "200"
	fmt.Println(b)

	/*
	var c [2]int = [2]int{100,200}
	c = append(c, 300)//配列はサイズを変更出来ないエラーになる
	fmt.Println(c)
	*/

	var c []int = []int{100, 200}
	c = append(c, 300, 400)//append関数
	fmt.Println(c)
}