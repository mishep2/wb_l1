package main

import "fmt"

func main() {
	intA := 14
	stringA := "four"
	boolA := true
	channelA := make(chan int)

	typeOf(intA)
	typeOf(stringA)
	typeOf(boolA)
	typeOf(channelA)
}

func typeOf(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Printf("%v is %T\n", v, v)
	case string:
		fmt.Printf("%v is %T\n", v, v)
	case bool:
		fmt.Printf("%v is %T\n", v, v)
	case chan int:
		fmt.Printf("%v is %T\n", v, v)
	}

}
