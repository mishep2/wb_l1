package main

import (
	"fmt"
	"sync"
	"time"
)

type Mapa struct {
	sync.RWMutex
	data map[string]string
}

func NewMapa() *Mapa {
	return &Mapa{
		data: make(map[string]string),
	}
}

func (m *Mapa) Set(key, val string) {
	m.Lock()
	defer m.Unlock()
	m.data[key] = val
}

func (m *Mapa) Get(key string) string {
	m.RLock()
	defer m.RUnlock()
	return m.data[key]
}

func main() {
	myMapa := NewMapa()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 7; i++ {
			myMapa.Set(fmt.Sprintf("%d", i), fmt.Sprintf("%d", i+100))
			time.Sleep(time.Microsecond)
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 7; i++ {
			res := myMapa.Get(fmt.Sprintf("%d", i))
			if res == "" {
				fmt.Println("first empty " + fmt.Sprintf("%d", i))
			} else {
				fmt.Println("1 : " + res)
			}
			time.Sleep(2 * time.Second)
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 7; i++ {
			res := myMapa.Get(fmt.Sprintf("%d", i))
			if res == "" {
				fmt.Println("second empty " + fmt.Sprintf("%d", i))
			} else {
				fmt.Println("2 : " + res)
			}
			time.Sleep(3 * time.Second)
		}
	}()
	wg.Wait()
}
