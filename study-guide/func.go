package main
import "fmt"
func test(a int, b int) (int, int){
	sum := a + b
	diff := a - b
	return sum, diff
}
func newItem(shopping string, item  ...string){
	fmt.Println("hey i  want ", shopping, "like - ")
	for _, shop := range item{
		fmt.Println("%s, ", shop)
	}

}
func printDetails(student string, subject ...string) {
	fmt.Println("hey ", student, ", here are your subject - ")
	for _, sub := range subject {
		fmt.Printf(" %s ", sub)
	}
}
func variadic(numbers ...int)int {
	values := 0
	for _, value := range numbers {
		values += value
	}
	return values

}

func test2() (int, int) {
	return 50, 59

}
func main() {
	addNumbers, subtractNumbers := test(2, 5)
	fmt.Println(addNumbers, " ", subtractNumbers)
	v, _ := test2()
	fmt.Print(v)
	fmt.Println(variadic(34))
	fmt.Println(variadic(34, 45))

	printDetails("marv", "physics", "chemistry")
	newItem("egg", "beans",)


}