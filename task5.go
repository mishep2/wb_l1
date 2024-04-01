package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var n int
	fmt.Println("num seconds: ")
	fmt.Scanf("%d\n", &n)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(n))
	defer cancel()

	ch := make(chan string)
	// read ch
	go func() {
		for val := range ch {
			fmt.Println("read : " + val)
		}
		fmt.Println("channel close")
	}()

	for {
		select {
		case <-ctx.Done():
			close(ch)
			fmt.Println("time's up")
			return
		default:
			ch <- RandStringRunes(4)
		}
		time.Sleep(time.Second)
	}

}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
