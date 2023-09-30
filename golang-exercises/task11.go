package main
import "fmt"

func main() {
    colors := [3]string{"red", "green", "blue"}
    for _, value :=range colors {
        fmt.Printf("color: %v \n ", value)
    }
    var i int = 1
    for i < 101{
        i += 3
       
        fmt.Println(i)
    }
}