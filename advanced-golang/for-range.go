package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan int)
	go sell(ch, &wg)
	go buy(ch, &wg)
	wg.Wait()
}
func sell(ch chan int, wg *sync.WaitGroup) {
	ch <- 1

	ch <- 2

	fmt.Println("sent all data")

	close(ch) //make sure to use close() function when using for-range to avoid a go routine deadlock error.

	wg.Done()
}
func buy(ch chan int, wg *sync.WaitGroup) {

	fmt.Println("waiting for data")

	for val := range ch {

		fmt.Println(val)

	}

	wg.Done()
}
