package main

import "fmt"

func calcTri(b, h float64) float64 {
	return 0.5 * b * h
}

func calcArea(r float64) float64 {
	return 3.14 * r * r
}
func calcPerimeter(r float64) float64 {
	return 2 * 3.14 * r
}
func calcDiameter(r float64) float64 {
	return 2 * r
}
func main() {
	var query int
	var radius float64

	var base float64

	var height float64

	fmt.Print("Enter the radius of circle: ")

	fmt.Scanf("%f", &radius)

	fmt.Printf("Enter \n 1 - area \n 2 - perimeter \n 3- diameter \n 4 - areaOfTriangle: ")

	fmt.Scanf("%d", &query)

	fmt.Print("Enter the base of the triangle: ")

	fmt.Scanf("%f", &base)

	fmt.Print("Enter the height of the triangle: ")

	fmt.Scanf("%f", &height)

	if query == 1 {
		fmt.Println(calcArea(radius))
	} else if query == 2 {
		fmt.Println(calcPerimeter(radius))
	} else if query == 3 {
		fmt.Println(calcDiameter(radius))
	} else if query == 4 {
		fmt.Println(calcTri(base, height))
	} else {
		fmt.Println("invalid query")
	}
}
