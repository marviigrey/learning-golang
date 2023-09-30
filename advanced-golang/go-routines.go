package main

import (
	"fmt"
	"time" //package for time
)

func calculateSquare(i int) {
	time.Sleep(1 * time.Second) //time should sleep for one second before doing the next execution.

	fmt.Println(i * i)

}
func main() {
	start := time.Now() // shows the current time that the program started
	for i := 1; i <= 10000; i++ {

		go calculateSquare(i) //using a go routine to perform concurrent processing.

	}
	elapsed := time.Since(start) // time between the program execution,the sleep, and the end result.
	// that calculates the time
	time.Sleep(2 * time.Second)
	fmt.Println(elapsed)
}
