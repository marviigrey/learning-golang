package main

import (
	"fmt"
	
)
func returnFunc(f string) func(){
	sum := 0
	return func(){
		fmt.Println(sum)
	}
}
func use() {
	fmt.Println("hello World!")
}
func main() {
	d := use
	d()
	tesla := func(c int) {
		fmt.Println(c)
	}
	returnFunc("hello")()
	tesla(5)
	grey := func(e int) int {
		return e * -2
	}(4)
	fmt.Println(grey)
}