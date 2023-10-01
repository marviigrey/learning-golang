package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan int, 6)
	go sell(ch, &wg)
	wg.Wait()

}
func sell(ch chan int, s *sync.WaitGroup) {
	ch <- 10
	ch <- 11
	ch <- 12
	go buy(ch, s)
	fmt.Println("sent all data to the channel")
	s.Done()
}
func buy(ch chan int, s *sync.WaitGroup) {
	fmt.Println("waiting for data")
	fmt.Println("Received data: ", <-ch)
	fmt.Println(len(ch), cap(ch))
	s.Done()
}
