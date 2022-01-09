package main

import (
	"fmt"
	"reflect"
)

func main() {
	// [...]で要素数省略可能
	a := [...]string{"sato", "suzuki", "takahashi"}
	a[1] = "tanaka"

	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a[2])
	fmt.Println(reflect.TypeOf(a))

	b := [2][2]string{{"sato", "suzuki"}, {"takahashi", "tanaka"}}
	fmt.Println(b[0][0])
	fmt.Println(b[0][1])
	fmt.Println(b[1][0])
	fmt.Println(b[1][1])
}