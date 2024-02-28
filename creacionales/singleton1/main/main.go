package main

import (
	"desingPatterns/creacionales/singleton1/singleton"
	client_one "desingPatterns/creacionales/singleton1/singleton/client-one"
	client_two "desingPatterns/creacionales/singleton1/singleton/client-two"
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(200)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			client_one.IncrementAge()
		}()
		go func() {
			defer wg.Done()
			client_two.IncrementAge()
		}()
	}
	wg.Wait()
	p := singleton.GetInstance()
	age := p.GetAge()
	fmt.Println(age)
}
