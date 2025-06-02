package main

import (
	"fmt"
	"strings"
)

func main() { // fungsi function adalah untuk mengeksekusi code tanpa harus menulis ulang code yang panajbg , jadi hanya dipanggil functionnya saja
	var names = []string{"Della", "Novita"}
	printMessage("halo", names)
	printMessage("Nama Saya", names)
}
func printMessage(message string, arr []string) { // ini adalah function untuk memanggil names
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}
