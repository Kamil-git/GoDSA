package main

// z rekurencja
// func quickSort(arr []int, low, high int) {
// 	if low < high {
// 		pivotIndex := partition(arr, low, high)
// 		quickSort(arr, low, pivotIndex-1)
// 		quickSort(arr, pivotIndex+1, high)
// 	}
// }

// func partition(arr []int, low, high int) int {
// 	pivot := arr[high]
// 	i := low - 1

// 	for j := low; j <= high-1; j++ {
// 		if arr[j] < pivot {
// 			i++
// 			arr[i], arr[j] = arr[j], arr[i]
// 		}
// 	}

// 	arr[i+1], arr[high] = arr[high], arr[i+1]
// 	return i + 1
// }

func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	stack := make([]int, len(arr))
	top := -1

	top++
	stack[top] = 0
	top++
	stack[top] = len(arr) - 1

	for top >= 0 {
		high := stack[top]
		top--
		low := stack[top]
		top--

		pivotIndex := partition(arr, low, high)

		if pivotIndex-1 > low {
			top++
			stack[top] = low
			top++
			stack[top] = pivotIndex - 1
		}

		if pivotIndex+1 < high {
			top++
			stack[top] = pivotIndex + 1
			top++
			stack[top] = high
		}
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j <= high-1; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
