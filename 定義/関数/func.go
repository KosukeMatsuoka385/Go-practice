package main

import (
	"fmt"
)

//func add(引数) (返り値の型)
func add(x float64, y float64, z float64) (float64, float64) {
	return x * y * z, x + y + z
}

//変数を決めて返り値にすることができるresult
func cal(price, item int) (result int) {
	result = price * item
	return result//resultはなくてもいい
}

func convert(price int) float64 {
	return float64(price)
}

func main() {
	r1, r2 := add(10.87, 20, 30)
	fmt.Println(r1, r2)

	r3 := cal(300, 5)
	fmt.Println(r3)

	f := func(x int) {
		fmt.Println("inner func", x)
	}
	f(1)

	func(x int) {
		fmt.Println("inner func", x)
	}(100)
}