package main

import "fmt"
import "reflect"
 
func main() {

   var grades int = 78
   var message string = "good score"
   var isCheck bool = true
   var amount float32 = 5466.3
   var car string = "RollsRoyce"
   fmt.Printf("variable grades = %v is of type %T \n", grades, grades)
   fmt.Printf("variable message = %v is of type %T \n", message, message)
   fmt.Printf("variable isCheck = %v is of type %T \n", isCheck, isCheck)
   fmt.Printf("variable amount = %v is of type %T \n", amount, amount)
   fmt.Printf("Type: %v \n", reflect.TypeOf("marviigrey"))
   fmt.Printf("Type: %v \n", reflect.TypeOf(12))
   fmt.Printf("Type: %v \n", reflect.TypeOf(12.0))
   fmt.Printf("Type: %v \n", reflect.TypeOf(true))
   fmt.Printf("variable car = %v \n", reflect.TypeOf(car))
}
