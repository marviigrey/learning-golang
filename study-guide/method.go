package main

import "fmt"

type Circle struct {
	radius float64
	area   float64
}

func (c *Circle) calcArea() { /*"c *Circle" is the reciever and a pointer to the circle struct
	calcArea is the method */
	c.area = c.radius * c.radius * 3.14
}

func main() {
	c := Circle{radius: 6}
	c.calcArea() // to reference the method we use the reciever(c) and the method name "calcArea()"
	fmt.Printf("%+v \n", c)
}
