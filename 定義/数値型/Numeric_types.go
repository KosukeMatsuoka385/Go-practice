package main

import "fmt"

//fmt
//https://golang.org/pkg/fmt/

//数値型
//https://golang.org/ref/spec#Numeric_types

func main() {
	/*
		var (
			u8  uint8     = 255
			i8  int8      = 127
			f32 float32   = 0.2
			c64 complex64 = -5 + 12i
		)

		fmt.Println(u8, i8, f32, c64)
		//Printfじゃないと調べれないよ
		fmt.Printf("%T %v\n", u8, u8) //%vで値を表示
		fmt.Printf("type=%T value=%v\n", c64, c64)
	*/
	//gofmt -w Numeric_types.goでフォーマットが自動で出来る
	/*
		x := 1 + 1
		fmt.Println(x)
		fmt.Println(1+1, 2+2)
		fmt.Println("1 + 1 =", 1+1)
		fmt.Println("1 - 1 =", 1-1)
		fmt.Println("100 * 3 =", 100*3)
		fmt.Println("100 / 3 =", 100/3)
		fmt.Println("100.0 / 3 =", 100.0/3)
		fmt.Println("100 / 3.0 =", 100/3.0)
		fmt.Println("100 % 3 =", 100%3)
		fmt.Println("100 % 2 =", 100%2)
	*/
	/*
	x := 0
	fmt.Println(x)
	// x = x + 1
	x++
	fmt.Println(x)
	// x = x - 1
	x--
	fmt.Println(x)
	*/

	fmt.Println(1 << 0)// 0001 0001
	fmt.Println(1 << 1)// 0001 0010
	fmt.Println(1 << 2)// 0001 0100
	fmt.Println(1 << 3)// 0001 1000

}
