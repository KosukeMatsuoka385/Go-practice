package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Hello, " + "World!")
	fmt.Println("Hello, World!"[0])
	fmt.Println("Hello, World!"[1])
	fmt.Println("Hello, World!"[8])
	fmt.Println(string("Hello, World"[0]))
	fmt.Println(string("Hello, World"[1]))
	fmt.Println(string("Hello, World"[2]))
	fmt.Println(string("Hello, World"[3]))
	fmt.Println(string("Hello, World"[4]))
	fmt.Println(string("Hello, World"[5]))
	fmt.Println(string("Hello, World"[6]))

	var s string = "Hello, Hello"
	//https://golang.org/pkg/strings/#Replace
	s = strings.Replace(s, "ll", "oo", 2) //package stringsを使用してReplaceして変換
	fmt.Println(s)
	fmt.Println(strings.Contains(s, "Hello")) //変換されたのでfalse
	fmt.Println(strings.Contains(s, "Heooo")) //変換されているのtrue

	//文字列書き方
	fmt.Println(`Test
													Test
	Test
	`)

	fmt.Println("\"")
	fmt.Println(`"`)
}
