//passing by values in a function & passing by reference using pointers.

package main

import "fmt"

func mo(m map[string]int) { //function that will be used to modify the original value of our map.
	m["A"] = 65 //adds the "A" into the map
}

func modi(d []int) {
	d[0] = 120
}

func mod(s *string) {
	*s = "i love you"
}

func modify(a int) {
	a += 100
}

func main() {
	ascii_codes := make(map[string]int) // we created a map of values

	ascii_codes["F"] = 70

	ascii_codes["K"] = 75

	fmt.Println(ascii_codes) //prints the original values of the map

	mo(ascii_codes) //modified a value in the map

	fmt.Println(ascii_codes)

	slice := []int{10, 20, 30} // original values  of the slice

	fmt.Println(slice) //prints out the same/original value

	modi(slice) // slice value has been modified by passing the value into a function without referencing a pointer.

	fmt.Println(slice) //prints out the modified value without using pointers.

	a := 10 // original value and address of the variable

	fmt.Println(a) /* when we run the program,the value of is doesnt change, because
	the address where the variable is stored is still the same and doesnt change after
	modification.
	*/
	modify(a) //stored in another memory address of the variable "a" but it's doesnt change the actual value of a.
	fmt.Println(a)
	s := "i hate you" //the actual value of s
	fmt.Println(s)    //prints "i hate you"
	mod(&s)           //a new value is added to the s address and therefore s is modified.
	fmt.Println(s)    //prints out a new value of the variable s because we passed in a
	//new value by referencing a pointer.
}
