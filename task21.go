package main

import "fmt"

type Duck struct {
	inWater bool
}

func (duck *Duck) Swim() {
	fmt.Println("duck swims")
}

func (duck *Duck) Fly() {
	fmt.Println("duck flyes")
}

type Dog struct {
}

func (dog *Dog) Run() {
	fmt.Println("dog runs")
}

// Целевой интерфейс
type Mover interface {
	Move()
}

// адаптер для утки
type DuckAdapter struct {
	*Duck
}

// движения утки
func (adaptee *DuckAdapter) Move() {
	if adaptee.inWater {
		adaptee.Swim()
	} else {
		adaptee.Fly()
	}
}

// конструктор адаптера для утки
func NewDuckAdapter(duck *Duck) Mover {
	return &DuckAdapter{duck}
}

// адаптер для собаки
type DogAdapter struct {
	*Dog
}

// движения собаки
func (adaptee *DogAdapter) Move() {
	adaptee.Run()
}

// конструктор адаптера для собаки
func NewDogAdapter(dog *Dog) Mover {
	return &DogAdapter{dog}
}

func main() {
	animals := []Mover{NewDuckAdapter(&Duck{inWater: true}), NewDogAdapter(&Dog{}), NewDuckAdapter(&Duck{inWater: false})}

	for _, val := range animals {
		val.Move()
	}
}
