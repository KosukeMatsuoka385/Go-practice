package main

import (
	"fmt"
)

func by2(num int) string {
	if num%2 == 0 {
		return "OK"
	} else {
		return "Yoisyo"
	}
}

func main() {

	result := by2(31)
	if result == "OK" {
		fmt.Println("Great")
	}

	fmt.Println(result)

	if result2 := by2(10); result2 == "OK" {
		fmt.Println("great 2")
	}

	// fmt.Println(result2)
	/*
		num := 9
		if num % 2 == 0 {
				fmt.Println("OK")
			} else if num % 3 == 0 {
				fmt.Printf("mamayane")
			} else {
				fmt.Println("DAME")
		}
	*/

	x, y := 11, 12
	if x == 10 && y == 10 {
		fmt.Println("&&")
	}

	if x == 10 || y == 10 {
		fmt.Println("||")
	}
}
