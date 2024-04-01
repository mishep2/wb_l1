package main

import (
	"fmt"
	"time"
)

func sleep(x int) {
	<-time.After(time.Second * time.Duration(x))
}

func main() {
	start := time.Now()
	sleep(4)
	fmt.Println(time.Since(start))
}
