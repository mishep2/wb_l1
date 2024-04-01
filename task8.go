package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	num := rand.Intn(257)
	fmt.Println("number : ", num)
	bitString := fmt.Sprintf("%b", num)
	fmt.Println("bit number : ", bitString)
	lenBit := len(bitString)
	pos := rand.Intn(lenBit)
	fmt.Println("pos bit : ", pos)
	new_bit := rand.Intn(2)
	fmt.Println("val new bit : ", new_bit)
	changeNbit(&num, lenBit-pos, new_bit)
	fmt.Println("new number : ", num)
	fmt.Printf("new bit number %b\n", num)
}

func changeNbit(num *int, pos int, new_bit int) {
	if new_bit == 1 {
		*num = *num | (1 << pos)
	} else {
		*num = *num & ^(1 << pos)
	}
}
