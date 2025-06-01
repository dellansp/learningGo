package main

import "fmt"

func main() {
	var point = 8
	if point == 10 {
		fmt.Println("lulus")
	} else if point > 5 {
		fmt.Println("Lulus")
	} else if point == 4 {
		fmt.Println("Hampir Lulus")
	} else {
		fmt.Printf("tidak lulus, nilai anda %d\n", point)
	}
}
