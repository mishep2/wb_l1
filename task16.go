package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 9, 2, 8, 3, 7, 4, 6, 5}
	fmt.Println(arr)
	fmt.Println(quicksort(arr))
}

func quicksort(arr []int) []int {
	l, h := 0, len(arr)-1
	return quicksortUnit(arr, l, h)
}

func quicksortUnit(arr []int, l, h int) []int {
	if l < h {
		var p int
		arr, p = partition(arr, l, h)
		arr = quicksortUnit(arr, l, p-1)
		arr = quicksortUnit(arr, p+1, h)
	}
	return arr
}

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}
