package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	set1 := make([]int, 0, 100)
	set2 := make([]int, 0, 100)

	for i := 0; i < 100; i++ {
		r := rand.Intn(200)
		if !containsGeneric(set1, r) {
			set1 = append(set1, r)
		}

		r = rand.Intn(200)
		if !containsGeneric(set2, r) {
			set2 = append(set2, r)
		}
	}
	fmt.Println(set1)
	fmt.Println(set2)

	start := time.Now()
	fmt.Println(SimpleGeneric(set1, set2))
	fmt.Println(time.Since(start))
	time.Sleep(time.Second)
	start2 := time.Now()
	fmt.Println(HashGeneric(set1, set2))
	fmt.Println(time.Since(start2))

}

// Simple Intersection: Compare each element in A to each in B (O(n^2))
func SimpleGeneric[T comparable](a []T, b []T) []T {
	set := make([]T, 0)

	for _, v := range a {
		if containsGeneric(b, v) {
			set = append(set, v)
		}
	}

	return set
}

func containsGeneric[T comparable](b []T, e T) bool {
	for _, v := range b {
		if v == e {
			return true
		}
	}
	return false
}

// Hash has complexity: O(n * x) where x is a factor of hash function efficiency (between 1 and 2)
func HashGeneric[T comparable](a []T, b []T) []T {
	set := make([]T, 0)
	hash := make(map[T]struct{})

	for _, v := range a {
		hash[v] = struct{}{}
	}

	for _, v := range b {
		if _, ok := hash[v]; ok {
			set = append(set, v)
		}
	}

	return set
}
