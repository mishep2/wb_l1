package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func worker(mainChan chan<- string, sigChan chan os.Signal) {
	for {
		time.Sleep(1 * time.Second)
		str := RandStringRunes(4)
		select {
		case mainChan <- str:
		case <-sigChan:
			close(mainChan)
			fmt.Println("\nFin Ctrl+C")
			return
		}
	}
}

func main() {
	rand.Seed(42)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	var numWorkers int
	fmt.Println("num workers: ")
	fmt.Scanf("%d\n", &numWorkers)

	mainChan := make(chan string)

	for i := 0; i < numWorkers; i++ {

		go worker(mainChan, sigChan)

	}

	for resVal := range mainChan {
		fmt.Println(resVal)
	}
}
