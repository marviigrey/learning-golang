package main 
import "fmt"
//line below we created a function named test.
func test() {
	fmt.Println("test")
}
func grey() {
	var name string = "marvellous"
	var age int = 22
	fmt.Printf("my name is %v and I am %d years old", name, age)
}
//line below we called the function test and we did it multiple times to show we can actually reuse functions.
func main() {
	test()
	test()
	grey()
}