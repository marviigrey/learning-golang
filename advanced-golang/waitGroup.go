package main

import (
	"fmt"
	"time" //package for time
	"sync"
)

func calculateSquare(i int, wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second) //time should sleep for one second before doing the next execution.

	fmt.Println(i * i)
	wg.Done()

}
func main() {
    var wg sync.WaitGroup
    wg.Add(10)
	start := time.Now() // shows the current time that the program started
	for i := 1; i <= 10000; i++ {

		go calculateSquare(i, &wg) //using a go routine to perform concurrent processing.

	}
	elapsed := time.Since(start) // time between the program execution,the sleep, and the end result.
	// that calculates the time
	wg.Wait()
	fmt.Println(elapsed)
}
