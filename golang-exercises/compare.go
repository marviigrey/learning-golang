package main

import "fmt"

func printResult(radius float64, calcFunction func(r float64) float64) { /*this is a function
	takes another function as an arguement. the first arguement takes float64 and the second input is a function
	that returns a float64 output. the first function must be called first as it is needed for
	the second function to be called.
	*/
	result := calcFunction(radius)  //this variable stores the input of our calcFunction when we call the function.
	fmt.Println("Result: ", result) //when we call the calcFunction, we need to pass in a float64 number type which
	//is stored under the variable "result".
	fmt.Println("thank you!")
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

func getfunction(query int) func(r float64) float64 {
	query_to_func := map[int]func(r float64) float64{ /*the key type is int and the value of the
		map is a function which is mapped to the functions we created and they return is float64*/
		1: calcArea,
		2: calcPerimeter,
		3: calcDiameter,
	}
	return query_to_func[query] // it returns the integer  from the query_to_func map
}

func main() {
	var query int
	var radius float64

	fmt.Print("Enter the radius of circle: ")

	fmt.Scanf("%f", &radius)

	fmt.Printf("Enter \n 1 - area \n 2 - perimeter \n 3- diameter \n 4 - areaOfTriangle: ")

	fmt.Scanf("%d", &query)

	printResult(radius, getfunction(query))
}
