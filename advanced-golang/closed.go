package main

import "fmt"

func main() {
	ch := make(chan int, 10)
	ch <- 23
	ch <- 44
	ch <- 45
	val, ok := <-ch
	fmt.Println(val, ok)
	close(ch)
	val, ok = <-ch
	fmt.Println(val, ok)
	close(ch)
	val, ok = <-ch
	fmt.Println(val, ok)

}
