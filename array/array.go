package main // array adalah sebuah buku / penyimpanan arsip

import "fmt" // array terdiri dari beberapa index , index dimulai dari 0

func main() {
	var fruits = [4]string{"apple", "grape", "banana", "melon"}
	fmt.Println("Jumlah Element \t\t", len(fruits))
	fmt.Println("isi semua element \t", fruits)
	fmt.Println(fruits[3])
}
