package main
import "fmt"
func addNumbers(a int, b int) int {
	sum := a + b

	return sum

}

func main() {
	sumOfNumbers := addNumbers(2, 3)
	fmt.Println(sumOfNumbers)
}