package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 4; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}

	arr := [...]int{2, 4, 6, 8, 10}
	for i := 0; i <= 4; i++ {
		fmt.Println(arr[i])
	}

	for i := 1; i <= 10; i++ {
		fmt.Println("i:", i)
	}


	j := 1
	for j <= 10 {
		fmt.Println("j:", j)
		j++
	}

	k := 1
	for {
		if (k == 11) {
			break
		}
		fmt.Println("k:", k)
		k++
	}
}