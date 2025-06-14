package main

import "fmt"

func main() {
	var lastName string = "Novita"
	var firstName string = "Della"

	var bilBulat uint8 = 20

	var numberA int = 4
	var numberB *int = &numberA
	*numberB = 8

	fmt.Print("halo", firstName, lastName, bilBulat, *numberB, "!\n")
}
