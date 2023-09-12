package main
import "fmt"
func main() {
var name string
var age int
fmt.Print("Enter your name: ")
fmt.Scanf("%s", &name)
fmt.Println("Welcome ", name)
fmt.Print("Enter your age: ")
fmt.Scanf("%d", &age)
fmt.Println("ok got it!", age)
}
