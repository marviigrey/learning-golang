package main

import (
	//"errors"
	"fmt"
)

func process(i int) error { //an error function that returns the error message if the statement is false.
	if i%2 == 0 {
		return fmt.Errorf("only odd numbers are allowed, got: %d", i)
	}
	return nil
}
func checkError(e error) {
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println("operation successful")
}
func main() {
	err := process(3)
	checkError(err)
	err = process(2)
	checkError(err)
}
