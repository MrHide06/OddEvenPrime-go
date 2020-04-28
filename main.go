package main

import (
	"fmt"
)

func ganjil(array [20]int) []int {
	var odd []int
	for _, i := range array {
		if i%2 == 0 {
			fmt.Printf("% d", i)
		}
	}
	fmt.Println("\n")
	return odd
}

func genap(array [20]int) []int {
	var even []int
	for _, i := range array {
		if i%2 != 0 {
			fmt.Printf("% d", i)
		}
	}
	fmt.Println("\n")
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
	ganjil(array)
	genap(array)
	prime(array)
}
