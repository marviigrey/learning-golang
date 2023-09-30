package main

import "fmt"

func addFunc(x, y int) int {
	return x + y
}
func add(x, y int) int { //fixed the code by adding int as return data type.
	return x + y
}

func main() {
	d := addFunc(4, 5)
	fmt.Println(d)
	fmt.Println("Result:", add(2, 3))
	///question 2: no return data type was specified.

}
