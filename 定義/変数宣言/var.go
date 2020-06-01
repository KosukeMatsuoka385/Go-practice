package main

import "fmt"

var (//関数外で宣言して使用出来る
		i int = 1
		f64 float64 = 1.2
		s string = "test"
		t, f bool = true, false
	)

func foo() {
	//変数を省略記法//省略して書くと関数内でしか宣言出来ない
	xi := 1
	// var xi int = 5
	xf64 := 1.22516481438546389238475//小数点16桁まで計算して表示する、繰り上げ
	var xf32 float32 = 1.5
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf64, xf32, xs, xt, xf)
	fmt.Printf("%T\n", xf32)//"%T"で型を調べれる
	fmt.Printf("%T\n", xi)//改行するには\n
}
func main() {
	fmt.Println(i, f64, s, t, f)
	foo()
}