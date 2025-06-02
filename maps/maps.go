package main

import "fmt"

func main() {
	var chicken map[string]int
	chicken = map[string]int{}
	chicken["januari"] = 58
	chicken["februari"] = 48
	fmt.Println("februari", chicken["februari"])
	fmt.Println("januari", chicken["januari"])
}
