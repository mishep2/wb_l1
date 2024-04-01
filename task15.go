package main

import (
	"fmt"
	"math/rand"
	"time"
)

// var justString string

// // непонятно для чего использовали глобалиную переменную

// func someFunc() {
// 	v := createHugeString(1 << 10)
// 	justString = v[:100]
// 	// если бы len(v) < 100 то была бы ошибка
// 	// мы используем только 100 из 1024 для чего?
// 	// для больших string лучше использовать strings.Builder{} он работает быстрее

// }
// func main() {
// 	someFunc()
// }

func main() {
	var justString string
	someFunc(&justString)
	fmt.Println(justString)
}

func someFunc(justString *string) {
	*justString = createHugeString(100)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func createHugeString(n int) string {
	rand.Seed(time.Now().Unix())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
