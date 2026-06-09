package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	rw      sync.RWMutex
	wg      sync.WaitGroup
)

func Reader(rw *sync.RWMutex, wg *sync.WaitGroup) {
	defer wg.Done()
	rw.RLock()
	fmt.Println(counter)
	rw.RUnlock()
}

func main() {
	counter = 0
	// read routine [allows multiple things to read the data safely without wating]
	wg.Add(1)
	go Reader(&rw, &wg)
	wg.Add(1)
	go Reader(&rw, &wg)

	// it allows only one thing at a time
	wg.Add(1)
	go func(rw *sync.RWMutex, wg *sync.WaitGroup) {
		defer wg.Done()
		rw.Lock()
		counter += 2
		rw.Unlock()
	}(&rw, &wg)
	wg.Add(1)
	go Reader(&rw, &wg)
	wg.Add(1)
	go Reader(&rw, &wg)

	wg.Wait()
}
