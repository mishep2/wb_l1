package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	stopByChannel()
	stopByContext()
	stopByBool()
	stopByWg()
}

func stopByChannel() {
	fmt.Println("Start stopByChannel")
	ch := make(chan int)
	done := make(chan struct{})
	go func() {
		for {
			select {
			case <-done:
				close(ch)
				return

			default:
				ch <- 1
			}
			time.Sleep(time.Second)
		}
	}()

	go func() {
		time.Sleep(4 * time.Second)
		done <- struct{}{}
	}()

	for val := range ch {
		fmt.Println("receive ", val)
	}
	fmt.Println("Fin stopByChannel")
}

func stopByContext() {
	fmt.Println("Start stopByContext")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(4))
	defer cancel()

	ch := make(chan int)
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(ch)
				return
			default:
				ch <- 2
			}
			time.Sleep(time.Second)
		}
	}()

	for val := range ch {
		fmt.Println("receive ", val)
	}
	fmt.Println("Fin stopByContext")
}

func stopByBool() {
	fmt.Println("Start stopByBool")

	ch := make(chan int)
	done := false
	go func() {
		for {
			if done {
				close(ch)
				return
			} else {
				ch <- 3
			}
			time.Sleep(time.Second)
		}
	}()

	go func() {
		time.Sleep(4 * time.Second)
		done = true
	}()

	for val := range ch {
		fmt.Println("receive ", val)
	}

	fmt.Println("Fin stopByBool")
}

func stopByWg() {
	fmt.Println("Start stopByWg")

	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		for {
			val, ok := <-ch
			if !ok {
				wg.Done()
				return
			}
			fmt.Println("receive ", val)
			time.Sleep(time.Second)
		}
	}()

	for i := 0; i < 4; i++ {

		ch <- 4

	}
	close(ch)
	wg.Wait()
	fmt.Println("Fin stopByWg")
}
