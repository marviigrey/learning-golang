package main

import "fmt"

func main() {
	var sum = 0

	for i := 1; i < 100; i++ {

		if i == 50 {
			continue
		}
		sum += i
		fmt.Println(sum)
	}
}
