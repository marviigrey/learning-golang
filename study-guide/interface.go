package main

import "fmt"

type shape interface {
	area() float64
	perimeter() float64
}
type square struct {
	side float64
}
type rect struct {
	length, breadth float64
}

func (s square) area() float64 {
	return s.side * s.side
}
func (s square) perimeter() float64 {
	return 4 * s.side
}
func (r rect) area() float64 {
	return r.length * r.breadth
}
func (r rect) perimeter() float64 {
	return 2 * (r.breadth + r.length)
}
func printData(s shape) {
	fmt.Println(s)
	fmt.Println(s.area())
	fmt.Println(s.perimeter())
}
func main() {
	c := square{side: 4}
	r := rect{length: 5, breadth: 6}
	printData(c)
	printData(r)
}
