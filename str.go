package main

import "fmt"

type Circle struct { //we created a struct that takes in the necessary values needed
	//to calculate the area of a circle
	x      int
	y      int
	radius float64
	area   float64
}

type Triangle struct { //we created a struct that takes in the necessary values needed
	//to calculate the area of a triangle
	b            float64
	h            float64
	areaTriangle float64
}

func calcArea(c Circle) { //created a function to pass in the calculation.
	const PI float64 = 3.14 //a const for the value of PI
	var area float64        // area which will return a float64 number when we do the calculation.

	area = (PI * c.radius * c.radius) //here is the calculation of the area

	c.area = area // c.area: we are calling the defined area in our struct "Circle".
	// and the area is equal to the calculation of area in our calcArea function.
}

func calcTri(t Triangle) {
	const half float64 = 0.5
	var areaTriangle float64
	areaTriangle = (half * t.h * t.b)
	t.areaTriangle = areaTriangle
}

func main() {
	c := Circle{x: 5, y: 5, radius: 5, area: 0}
	t := Triangle{h: 4, b: 4}
	fmt.Printf("%+v \n", c)
	fmt.Printf("%+v \n", t)
	calcArea(c)
	calcTri(t)
	fmt.Printf("%+v \n", c)
	fmt.Printf("%+v \n", t)
}
