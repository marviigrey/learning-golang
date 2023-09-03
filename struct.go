package main

import "fmt"

type Student struct { //we declared a struct with two different values and data types.
    name string
    
    rollNo int
}
type Circle struct { //we declared a struct with different values and a single data type
    x int
    y int
    radius float64
    area float64
}
func main() {
     st := Student{ // created a variable st with our struct "Student"
     
         name: "grey",
         rollNo: 34,
     }
    
     var c Circle //created a variable st with our struct "Circle"
     c.x = 20
     c.y =45
     star := Student{"joe", 45} 
     fmt.Printf("%+v \n", st, star)
     fmt.Printf("%+v \n", c.x)
}