package main

import "fmt"

func main() {
	a, b := 14, 4
	fmt.Println("a = ", a, " b = ", b)
	a, b = b, a
	fmt.Println("a = ", a, " b = ", b)
}
