package main

import "fmt"

type student struct { //
	name  string // name adalah property
	grade int
}

func main() {
	var s1 student
	s1.name = " Della Novita"
	s1.grade = 10
	fmt.Println("name :", s1.name)
	fmt.Println("grade :", s1.grade)
}
