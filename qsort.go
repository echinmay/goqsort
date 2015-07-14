package main

import (
	"fmt"
)

func partition(arr []int) int {

	if len(arr) <= 1 {
		return 1
	}

	p := arr[0]
	i := 1
	for j := 1; j < len(arr); j++ {
		if arr[j] < p {
			arr[j], arr[i] = arr[i], arr[j]
			i++
		}
	}
	arr[0], arr[i-1] = arr[i-1], arr[0]
	return i
}

func qsort(arr []int) {
	pos := partition(arr)
	if pos > 1 {
		// Sort left
		qsort(arr[:pos-1])
	}
	if pos < len(arr) {
		// Sort right
		qsort(arr[pos:])
	}
}

func main() {

	arr := make([]int, 0)
	arr = append(arr, 2, 3, 4, 9, 8, 1, 5, 6, 7)
	fmt.Println(arr)
	qsort(arr)
	fmt.Println(arr)

}
