package main

import (
	"fmt"
)

func main() {
	var avg = calculate(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	var msg = fmt.Sprintf("Rata - rata : %.2f", avg)
	fmt.Println(msg)
}
func calculate(numbers ...int) float64 { // ini namanya function variatif , dimna numbers akan menjadi array berisi integer
	var total int = 0
	for _, number := range numbers { // ini adalah looping penjumlahan dari 0 +2 = +3 = +4 dst
		total += number
	}
	var avg = float64(total) / float64(len(numbers))
	return avg
}
