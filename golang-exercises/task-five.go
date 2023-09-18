package main

import "fmt"

func main() {
	var name string
	var age int
	fmt.Print("Enter Your name: ")
	fmt.Scanf("%s", &name)

	fmt.Print("and your age: ")
	fmt.Scanf("%d", &age)
	fmt.Println("welcome: ", name)
}
