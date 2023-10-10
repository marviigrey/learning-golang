package main

import (
	"fmt"
	"strings"
)

func main() {
	r := strings.NewReader("this is a string") //r calls the NewReader function of the io.reader interface and stores the string
	buf := make([]byte, 4)                     // a slice is created to store 4 integers in the byte slice, the characters are read from the varaible r.
	for {
		n, err := r.Read(buf)
		fmt.Println(string(buf[:n]), err)
		if err != nil { //a loop to continue printing until there's no value to print and finally returns an error EOF.
			fmt.Println("breaking out")
			break
		}
	}
}
