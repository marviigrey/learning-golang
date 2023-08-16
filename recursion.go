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
	//anonymous function.
	d := func(l int, j int)int  {
		return l * j
	}(9, 34)
	m := 3
	total := cube(m)
	fmt.Println("The cube of ", m, "is:", total)
	n := 5
	result := factorial(n)
	fmt.Println("Factorial of", n, "is:", result)
	fmt.Println(d)
}