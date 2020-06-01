//ネットワーク、ファイル読込関係

package main

import (
	"fmt"
)

func main() {
	b := []byte{72, 73}
	fmt.Println(b)
	fmt.Println(string(b))//アスキーコードに変換

	c := []byte("HI")
	fmt.Println(c)
	fmt.Println(string(c))
}