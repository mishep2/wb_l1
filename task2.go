package main

import "fmt"

func pow(c chan int, arr []int) {
	for _, v := range arr {
		c <- v * v
	}
	close(c)
}

func main() {
	arr := [5]int{2, 4, 6, 8, 10}
	c := make(chan int)

	go pow(c, arr[:])

	arr_squared := make([]int, 0, len(arr))
	for {
		val, ok := <-c
		if !ok {
			break
		}
		arr_squared = append(arr_squared, val)
	}

	for i := range arr {
		fmt.Printf("%d * %d is %d\n", arr[i], arr[i], arr_squared[i])
	}
}
