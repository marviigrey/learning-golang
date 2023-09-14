package main

import "fmt"

func calc(a, b int) int {
	var c int
	c = a - b //return variable c with an integer as c
	return c

}

func main() {
	var packagesToDeliver int = 100
	var packagesDelivered int = 20
	var customers = 4
	fmt.Println("i have ", packagesToDeliver, "packages")

	fmt.Println("i have delivered ", packagesDelivered, "packages")

	fmt.Println("i have to deliver ", packagesToDeliver, "to ", customers, "equally,which is ", packagesToDeliver/customers, "each")

	c := calc(packagesToDeliver, packagesDelivered)

	fmt.Println("number of packages remaining: ", c)

}
