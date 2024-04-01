package main

import (
	"fmt"
	"math"
)

type point struct {
	x, y float64
}

func newPoint(x, y float64) *point {
	return &point{x, y}
}

func dist(a, b *point) float64 {
	x, y := (a.x - b.x), (a.y - b.y)
	return math.Sqrt(x*x + y*y)
}

func main() {
	a := newPoint(1, 2)
	b := newPoint(3, 4)
	fmt.Println(dist(a, b))
}
