package main

import (
	"fmt"
	"os"
)

func foo() {
	defer fmt.Println("World foo")
	fmt.Println("Hello foo")
}

func main() {
	/*
	defer fmt.Println("World")
	foo()
	fmt.Println("Hello")
	*/
	/*
	fmt.Println("Run")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("Success")
	*/


	file, _ := os.Open("../switch/switch.go")
	defer file.Close()
	data := make([]byte, 100)
	file.Read(data)
	fmt.Println(string(data))
	
}