package main

import (
	"fmt"
	"sync"
	"time"
)

func test(a int, b int, wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	sum := a + b
	fmt.Println(sum)
	wg.Done()

}

func second(wg *sync.WaitGroup) {
	for i := 1; i <= 10; i++ {

		a := 5
		fmt.Println(a * i)
	}
	wg.Done()
}

func testing(wg *sync.WaitGroup) {
	time := time.Now()
	fmt.Println(time)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go test(4, 5, &wg)
	go second(&wg)
	go testing(&wg)
	wg.Wait()
	fmt.Println("all go routines have been completed.")
}
