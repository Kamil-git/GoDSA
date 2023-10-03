package main

import "fmt"

func main() {

	arr := []int{2, 1, 4, 3, 6, 5, 8, 7, 10, 9, 12, 11, 14, 13}
	fmt.Println("Tablica przed sortowaniem:", arr)

	sortedArr := mergeSort(arr)

	fmt.Println("Tablica po sortowaniu:", sortedArr)
}
