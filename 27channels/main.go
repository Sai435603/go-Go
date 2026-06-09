package main

import (
	"fmt"
	"sync"
)

var (
	counter = 0
	mu      sync.Mutex
)

func IncrementCnt(mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()

	mu.Lock()
	counter++
	mu.Unlock()
}

func main() {
	channel := make(chan int, 2) // this is buffered channel
	// channel2 := make(chan int) // this is unbuffered channel
	var wg sync.WaitGroup
	wg.Add(2)
	// Read only routine and it is not allowed to write channel close here
	go func(channel <-chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		val, IsChannelOpen := <-channel
		fmt.Println(IsChannelOpen)
		fmt.Println(val)
	}(channel, &wg)
	// Write only routine
	go func(channel chan<- int, wg *sync.WaitGroup) {
		defer wg.Done()
		defer close(channel)
		channel <- 69
	}(channel, &wg)

	wg.Wait() //it will wait until done with the internal count

	//experimenting with the counter variable [shared variable]
	wg.Add(2)
	go IncrementCnt(&mu, &wg)
	go IncrementCnt(&mu, &wg)
	wg.Wait()
	fmt.Println(counter)
}
