package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	go leak(&wg)
	wg.Wait()
}

func leak(s *sync.WaitGroup) {
	ch := make(chan int)
	go func() {
		val := <-ch
		fmt.Println("Received ", val)
		s.Done()

	}()
	fmt.Println("exiting leak method")
	s.Done()
}
