package main 
import "fmt"

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}
func cube(m int) int {
	if m == 1 {
		return 1
	}
	return m * m * m
}

func main() {
	m := 3
	total := cube(m)
	fmt.Println("The cube of ", m, "is:", total)
	n := 5
	result := factorial(n)
	fmt.Println("Factorial of", n, "is:", result)
}