package main

import (
	"fmt"
	"sync"
)

// check if the code has a race condition with the command [go run --race .]
// TODO: explore ReadWrite Locks in go
func main() {
	wg := &sync.WaitGroup{}
	mtx := &sync.Mutex{}
	scores := []int{0}
	wg.Add(3)
	go func(wg *sync.WaitGroup, mtx *sync.Mutex) {
		fmt.Println("One adds 1")
		mtx.Lock()
		scores = append(scores, 1)
		mtx.Unlock()
		wg.Done()

	}(wg, mtx)

	go func(wg *sync.WaitGroup, mtx *sync.Mutex) {
		fmt.Println("Two adds 2")
		mtx.Lock()
		scores = append(scores, 2)
		mtx.Unlock()
		wg.Done()

	}(wg, mtx)

	go func(wg *sync.WaitGroup, mtx *sync.Mutex) {
		fmt.Println("Three adds 3")
		mtx.Lock()
		scores = append(scores, 3)
		mtx.Unlock()
		wg.Done()
	}(wg, mtx)
	wg.Wait()
}
