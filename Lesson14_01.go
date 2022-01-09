package main

import (
	"fmt"
)

type Student struct {
	name string
	math, english float64
}

type User struct {
	name string
}

func (s Student) avg(math, english float64) (avgResult float64) {
	avgResult = (math + english) / 2
	return
}

func (user User) cal(weight, height float64) (bmi float64) {
	bmi = (weight * 10000) / (height * height)
	return
}

func main() {
	a001 := Student{name:"sato"}
	fmt.Println(a001.avg(80, 70))

	user := User{name:"takuro"}
	fmt.Println(user.name, user.cal(60, 174))
}