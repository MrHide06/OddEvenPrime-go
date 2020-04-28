package main

import (
	"fmt"
)

func genap(array [20]int) []int {
	var odd []int
	for _, i := range array {
		if i%2 == 0 {
			odd = append(odd, i)
		}
	}

	return odd
}

func ganjil(array [20]int) []int {
	var even []int
	for _, i := range array {
		if i%2 != 0 {
			even = append(even, i)
		}
	}
	return even
}

func prime(array [20]int) []int {
	var prima []int
	for _, angka := range array {
		if (angka%2 == 0) || (angka != 1) {
			fmt.Printf("% d", angka)
		}
	}
	fmt.Println("\n")
	return prima
}
func main() {
	array := [20]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	fmt.Printf("Bilangan awal : %+v, Jumlah = %d\n", array, len(array))
	fmt.Println("-----------------------------------------------------------------------------------")
	fmt.Printf("Bilangan ganjil : %+v, jumlah = %d\n", ganjil(array), len(ganjil(array)))
	fmt.Println("-----------------------------------------------------------------------------------")
	fmt.Printf("Bilangan genap : %+v, jumlah = %d\n", genap(array), len(genap(array)))
	fmt.Println("-----------------------------------------------------------------------------------")
	prime(array)
}
