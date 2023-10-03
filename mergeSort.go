package main

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// Dziel tablicę na dwie części
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	// Scalaj dwie posortowane części
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0)

	// Indeksy do śledzenia pozycji w lewej i prawej tablicy
	l, r := 0, 0

	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}

	// Dodaj pozostałe elementy z lewej i prawej tablicy
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)

	return result
}
