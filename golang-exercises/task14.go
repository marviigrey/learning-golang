package main

import "fmt"

func add(x, y int) (z int) {
	z = x + y
	return
}

func main() {
	fmt.Print(add(3, 2))
}
