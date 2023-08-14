package main

import (
	"fmt"
	
)

//line below we created a function named test.
func test(id, string) {
	fmt.Print("Hello", id)
}
func grey() {
	var name string = "marvellous"
	var age int = 22
	fmt.Println(name, age)
	
}
func label(a, b int) (a1 int, b1 int) {
	defer fmt.Println("hello, run before you return the statement")
	a1 = a + b
	b1 = a - b
	fmt.Println("print this after the defer function")
	return
}
func para(new string) {
	fmt.Println(new)
}
func add(x, y int) (int, int){
	return x + y, x-y
}
//line below we called the function test and we did it multiple times to show we can actually reuse functions.
func main() {
	test()
	test()
	grey()
	para("samuel")
	ans1, ans2 := add(14, 7)
	fmt.Println(ans1, ans2)
	ans3, ans4 := label(30, 20)
	fmt.Println(ans3, ans4)
	z := test
	z("grey")


}