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

	arr_squared_sum := 0
	for {
		val, ok := <-c
		if !ok {
			break
		}
		arr_squared_sum += val
	}

	fmt.Printf("squared sum of %v  is %d\n", arr, arr_squared_sum)

}
