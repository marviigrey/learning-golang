package main

import "fmt"

func calc(n int) int {
	if n == 1 {
		return 1
	}
	return n * calc(n+1)
}
func recursion(n int) int {
	if n == 0 {
		return 1
	}
	return n * recursion(n+1)

}

func main() {
	d := 6
	calc(d)
	fmt.Print(calc(d))
	p := 8
	recursion(p)
	fmt.Println(recursion(p))

}
