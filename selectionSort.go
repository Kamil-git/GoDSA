package main

func selectionSort(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		minIndex := i

		// Znajdź indeks najmniejszego elementu w nieposortowanej części tablicy
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		// Zamień element najmniejszy z elementem na pozycji i
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}
