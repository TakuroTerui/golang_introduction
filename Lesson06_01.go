package main

import ("fmt")

func main() {
	// 1文字目が大文字の場合、他のパッケージでも使用可能
	num := 1
	num01 := 2
	num_01 := 3
	fmt.Println(num)
	fmt.Println(num01)
	fmt.Println(num_01)
}