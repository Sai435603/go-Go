package main

import (
	"fmt"
	"sync"
)

func main() {
	channel := make(chan int, 2)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	// Read only routine and it is not allowed to write channel close here
	go func(channel <-chan int, wg *sync.WaitGroup) {
		val, IsChannelOpen := <-channel
		fmt.Println(IsChannelOpen)
		fmt.Println(val)
		wg.Done()
	}(channel, wg)
	// Write only routine
	go func(channel chan<- int, wg *sync.WaitGroup) {
		channel <- 1
		close(channel)
		wg.Done()
	}(channel, wg)

	wg.Wait()
}
