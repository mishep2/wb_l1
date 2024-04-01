package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{-4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(sort.SearchInts(arr, 7))
	fmt.Println(myBinarySearch(arr, 7))
}

func myBinarySearch(arr []int, target int) int {
	min, max := 0, len(arr)-1

	for min < max {
		mid := (min + max) / 2
		switch {
		case arr[mid] > target:
			max = mid - 1
		case arr[mid] < target:
			min = mid + 1
		case arr[mid] == target:
			return mid
		}
	}
	return -1
}
