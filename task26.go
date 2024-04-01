package main

import (
	"fmt"
	"strings"
)

func main() {
	strs := []string{"abcd", "abCdefAaf", "aabcd", "aA"}
	for _, val := range strs {
		fmt.Println(val, " - ", isUnic(val))
	}
}

func isUnic(str string) bool {
	str = strings.ToLower(str)
	m := make(map[int32]struct{})

	for _, val := range str {
		if _, ok := m[val]; ok {
			return false
		}
		m[val] = struct{}{}
	}
	return true
}
