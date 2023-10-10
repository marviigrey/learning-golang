package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "God is good,learning go is fun, learning is fun"
	fun := "learning"
	result := strings.Contains(text, fun) //this checks if the string fun can be found in the string text.
	fmt.Println(result)
	countResult := strings.Count(text, fun) //this counts the number of times the string "fun appears in the string "text".
	fmt.Println(countResult)

	test := "learning standard library in python."
	try := "library in python"
	replaceResult := strings.ReplaceAll(test, try, "library in Go") //searches for try in the variable test and replaces it with "library in Go"
	fmt.Println(replaceResult)

}
