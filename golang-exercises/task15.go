package main

import "fmt"

func discountedPrice(product string, price float64) float64 {
	var discount float64
	switch product {
	case "apples":
		return price - (0.1 * price)

	case "bananas":
		return price - (0.2 * price)

	case "oranges":
		return price - (0 * price)
	case "orange":
		return price - (0 * price)

	default:
		fmt.Println(price)

	}
	return discount
}

func main() {
	fmt.Println(discountedPrice("apples", 100))
	fmt.Println(discountedPrice("orange", 100))
	fmt.Println(discountedPrice("bananas", 100))
	fmt.Println(discountedPrice("bananas", 100))
	fmt.Println(discountedPrice("oranges", 100))
}
