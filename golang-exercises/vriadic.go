package main

import "fmt"

func myFunc(student string, subject ...string){
    fmt.Println("hello my name is %s", student, "my subjects are ")
    for _, sum := range subject{
        fmt.Printf("%s", sum)
    }
}

func main() {
    
    
    myFunc("marvellous", "physics", "maths")
}