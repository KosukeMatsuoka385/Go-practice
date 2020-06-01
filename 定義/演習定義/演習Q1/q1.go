package main

import (
	"fmt"
)


func main() {
	f := 1.11
	ff := int(f)
	fmt.Println(ff)
	fmt.Printf("%T, %v", ff, ff)
}