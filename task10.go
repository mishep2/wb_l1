package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	group := group(arr)
	fmt.Println(group)
}

func group(arr []float64) map[int][]float64 {
	group := make(map[int][]float64)

	var (
		wg sync.WaitGroup
		mu sync.Mutex
	)
	wg.Add(len(arr))
	for _, val := range arr {
		go func(val float64) {
			defer wg.Done()
			key := int(val/10) * 10
			mu.Lock()
			group[key] = append(group[key], val)
			mu.Unlock()
		}(val)
	}
	wg.Wait()
	return group
}
