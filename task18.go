package main

import (
	"fmt"
	"sync"
)

type counter struct {
	counter int
	sync.Mutex
}

func main() {
	cnt := counter{}
	var wg sync.WaitGroup

	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			cnt.Lock()
			cnt.counter += 1
			cnt.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(cnt.counter, " = ", 1000)
}
