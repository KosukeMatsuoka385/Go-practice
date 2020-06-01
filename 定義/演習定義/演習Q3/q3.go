package main

import (
	"fmt"
)

func main() {

	m := map[string]int{"Mike": 20, "Naccy": 24, "Messi": 30}
	fmt.Println(m)

	//バージョン1.12からは map をキー順で表示する。
	//https://tip.golang.org/doc/go1.12#fmt

	fmt.Printf("%T, %v", m, m)
}