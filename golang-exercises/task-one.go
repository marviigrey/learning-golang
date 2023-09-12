package main

import "fmt"

func main() {
	var name string = "marvellous"
	age := 22
	fmt.Printf("my name is %s and i am %d\n", name, age)

	var objective = "Defend Humanity"
	fmt.Println(objective)

	const squares = 2
	var circles = 0

	fmt.Println("Squares:", squares)
	fmt.Println("Circles:", circles)

	circles = 7

	fmt.Println("Squares:", squares)
	fmt.Println("Circles:", circles)
}
