package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "amazing words reverse"
	fmt.Println(str)
	fmt.Println(reverseWords(str))
}

func reverseWords(str string) string {
	wordsarr := strings.Split(str, " ")
	for i, j := 0, len(wordsarr)-1; i < j; i, j = i+1, j-1 {
		wordsarr[i], wordsarr[j] = wordsarr[j], wordsarr[i]
	}
	return strings.Join(wordsarr, " ")
}
