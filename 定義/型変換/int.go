package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int = 1
	xx := float64(x)
	fmt.Printf("%T, %v, %f\n", xx, xx, xx)

	var y float64 = 1.2
	yy := int(y)
	fmt.Printf("%T, %v, %d\n", yy, yy, yy)

	var s string = "14"
	i, _ := strconv.Atoi(s)
	/*
		i, err := strconv.Atoi(s)//Atoiは関数
		if err != nil{
			fmt.Println("ERROR")
		}
	*/
	fmt.Printf("%T, %v\n", i, i)

	h := "Hello World"
	fmt.Println(h[0])//ASC2コード//72
	fmt.Println(string(h[0]))
}
