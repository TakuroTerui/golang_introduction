package main

import (
	"fmt"
)

type Student struct {
	name string
	math, english float64
}

type User struct {
	gender string
	age int
}

func main() {
	// 構造体の初期化①
	var s Student
	s.name = "sato"
	s.math = 80
	s.english = 70

	// 構造体の初期化②
	// s := Student{"sato", 80, 70}

	// 構造体の初期化③
	// s := Student{name:"sato", math:80, english:70}

	user := User{gender:"male", age:20}

	fmt.Println(s)
	fmt.Println(user)
}