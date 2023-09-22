package main
import "fmt"

func main() {
	
 pizza := [] string {}
     
    pizza = append(pizza, "hamburger", "salad")
	for _, value := range pizza  {
		fmt.Printf("food: %v\n", value)
	}
	shapes := [3]string{"square", "circle", "triangle"}
	for index, value := range shapes{
		
		fmt.Printf("This is %v and its index in the array is %v\n",  value, index)
	}
	var name string
	var age int
	for  {
	fmt.Printf("your name is ___ and your age is: ")
	fmt.Scan(&name, &age)
	fmt.Printf("hello %v, you're %v years old", name, age)
	}
}