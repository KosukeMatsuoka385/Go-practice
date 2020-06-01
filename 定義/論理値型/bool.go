package main

import (
	"fmt"
)

func main() {
	// var t, f bool = true, false
	t, f := true, false
	fmt.Println(t, f)
	//おさらい%Tは型、%vは値
	fmt.Printf("%T, %v %t\n", t, 13, t)
	fmt.Printf("%T, %v %t\n", f, -65, f)

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && false)

	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || false)

	fmt.Println(!true)
	fmt.Println(!false)

}
