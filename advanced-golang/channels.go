package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string) //channel "ch" that can only transfer data type string.
	go sell(ch)             //run the go routine using the ch channel to send and receive data.
	go buy(ch)
	time.Sleep(1 * time.Second)
}

// send data to the channel.
func sell(ch chan string) {
	ch <- "furniture"
	fmt.Println("send data to the channel")
}

// receive data from the channel.
func buy(ch chan string) {
	fmt.Println("waiting for the data")
	val := <-ch
	fmt.Println("Received data: ", val)
	fmt.Println(len(ch))
}
