package main

import (
 	
	"fmt"
	"strconv"
) 
func main() {
	var m string = "122"
	c, err := strconv.Atoi(m)
	var a int = 89
	var f float64 = float64(a)
	var b int = 78
	var s string = strconv.Itoa(b)
	fmt.Printf("%q \n", s)
	fmt.Printf("%.2f \n", f)
	fmt.Printf("%d \n", c, err)
}

