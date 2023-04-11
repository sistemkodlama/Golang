package main

import (
	"fmt"
)

func topla(sayi1 int, sayi2 int) int {
	return sayi1 + sayi2
}

// fonksiyondan çoklu değer dönülebilinir
func swap(x int, y int) (int, int) {
	return y, x
}

func main() {
	fmt.Println("2 + 4 = ", topla(2, 4))

	a, b := swap(1, 99)
	fmt.Println("a : ", a, " ... b : ", b)
}
