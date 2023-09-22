package main
import "fmt"
func main() {
    //1. we need slices to be able to cut through the capacity and length of an array. slices are flexible than array.
    //2. How to define an empty slice of strings called someSlice?
    
    someSlice := [] string {}
    //3. 
    someSlice = append(someSlice, "abc") //adds the string abc to the slice someSlice.
    
    //4. How to retrieve the first element of the slice someSlice?
    fmt.Println(someSlice[0])
}