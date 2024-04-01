package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h *Human) Talk() {
	fmt.Printf("Hi, my name is %s, I'm %d years old\n", h.Name, h.Age)
}

type Action struct {
	Human
	Movement string
}

func main() {
	h := Human{
		Name: "Mihail",
		Age:  22,
	}
	h.Talk()

	a := Action{
		Human:    Human{"Mihail", 22},
		Movement: "goes",
	}
	// Action имеет все методы Human
	a.Talk()

}
