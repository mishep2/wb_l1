package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree", "duck", "duck"}

	set := set(arr)
	fmt.Println(set)
}

func set(arr []string) []string {
	set := make([]string, 0)
	hash := make(map[string]struct{})

	for _, v := range arr {
		hash[v] = struct{}{}
	}

	for key := range hash {
		set = append(set, key)
	}

	return set
}
