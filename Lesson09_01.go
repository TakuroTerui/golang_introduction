package main

import (
	"fmt"
)

func main() {
	x := 10
	y := 2

	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(x % y)
	fmt.Println(x >= 5 && x <= 10)
	fmt.Println(y >= 5 && y <= 10)
	fmt.Println(x >= 5 || x <= 10)
	fmt.Println(y >= 5 || y <= 10)
}