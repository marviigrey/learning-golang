package main

import "fmt"

type Circle struct {
	y      int
	x      int
	radius float64
	area   float64
}

func calcArea() {
	const PI float64 = 3.14
	var area float64
	area = (PI * c.radius * c.radius)
	c.area = area
}
func main() {
	c := Circle{x: 4, y: 5, radius: 5, area: 0} /*we declared a variable using the
	data type we created from our struct program */
	fmt.Printf("%+v \n", c)
	calcArea(c)
	fmt.Printf("%+v \n", c)
}
