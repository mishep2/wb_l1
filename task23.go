package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(arr)
	fmt.Println(delA(arr, 4))
	fmt.Println(delB(arr, 4))
}

func delA(arr []int, i int) []int {
	arr[i] = arr[0]
	return arr[1:]
}

func delB(arr []int, i int) []int {
	return append(arr[:i], arr[i+1:]...)
}
