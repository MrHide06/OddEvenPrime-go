package main

import (
	"fmt"
)

func main() {
	array := [30]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	fmt.Printf("Bilangan awal : %+v, Jumlah = %d\n", array, len(array))
	fmt.Println("-----------------------------------------------------------------------------------")
	fmt.Printf("Bilangan ganjil : %+v, jumlah = %d\n", ganjil(array), len(ganjil(array)))
	fmt.Println("-----------------------------------------------------------------------------------")
	fmt.Printf("Bilangan genap : %+v, jumlah = %d\n", genap(array), len(genap(array)))
	fmt.Println("-----------------------------------------------------------------------------------")
	fmt.Printf("Bilangan Prima : %+v, jumlah = %d\n", prime(array), len(prime(array)))
}

func genap(array [30]int) []int {
	var odd []int
	for _, i := range array {
		if i != 0 {
			if i%2 == 0 {
				odd = append(odd, i)
			}
		}
	}

	return odd
}

func ganjil(array [30]int) []int {
	var even []int
	for _, i := range array {
		if i%2 != 0 {
			even = append(even, i)
		}
	}
	return even
}

func prime(array [30]int) []int {
	var prima []int
	for x := 1; x < len(array); x++ {
		i := 0
		for y := 1; y < len(array); y++ {
			if x%y == 0 {
				i++
			}
		}
		if (i == 2) && (x != 1) {
			prima = append(prima, x)
		}
	}
	return prima
}
