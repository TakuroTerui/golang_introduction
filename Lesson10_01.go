package main

import (
	"fmt"
)

func main() {
	x := 10
	y := 12
	if age := x + y ; age >= 20 {
		fmt.Println("adlut")
	} else if age == 0 {
		fmt.Println("baby")
	} else {
		fmt.Println("child")
	}

	if num := 12 ; num >= 5 && num < 10 {
		fmt.Println(true)
	} else {
			fmt.Println(false)
	}
}