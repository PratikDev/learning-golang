package main

import (
	"fmt"
	"sync"
)

func main() {
	myChannel := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(2)

	// receive ONLY
	go func(channel <-chan int, wg *sync.WaitGroup) {
		value, isChannelOpen := <-channel

		fmt.Println(isChannelOpen)
		fmt.Println(value)

		wg.Done()
	}(myChannel, wg)

	// send ONLY
	go func(channel chan<- int, wg *sync.WaitGroup) {
		channel <- 5
		close(channel)

		wg.Done()
	}(myChannel, wg)

	wg.Wait()
}
