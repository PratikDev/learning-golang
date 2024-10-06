package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	scores := []int{}

	wg.Add(3)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Routine One")

		m.Lock()
		scores = append(scores, 1)
		m.Unlock()

		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Routine Two")

		m.Lock()
		scores = append(scores, 2)
		m.Unlock()

		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Routine Three")

		m.Lock()
		scores = append(scores, 3)
		m.Unlock()

		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(scores)
}
