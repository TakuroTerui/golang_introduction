package main

import (
	"fmt"
)

func main() {
	sayHello("takuro!")
	result01, result02 := cal(6, 3)
	fmt.Println(result01, result02)
	add := calAdd(10, 5)
	fmt.Println(add)
}

func sayHello(greeting string) {
	fmt.Println(greeting)
}
func cal(x int, y int) (a int, b int) {
	a = x / y
	b = x * y
	return
}

func calAdd(x int, y int) (addNum int) {
	addNum = x + y
	return
}