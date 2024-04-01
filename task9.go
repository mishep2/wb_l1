package main

import (
	"fmt"
	"sync"
)

func main() {

	arr := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ch := make(chan int)
	ch2 := make(chan int)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range ch {
			ch2 <- val * 2
		}
		close(ch2)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range ch2 {
			fmt.Println(val)
		}
	}()

	for _, val := range arr {
		ch <- val
	}
	close(ch)
	wg.Wait()
}
