package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	if _, err := io.Copy(os.Stdout, r); err != nil { //err := (type, r) == Copy(dst Writer, src Reader)
		log.Fatal(err) //Fatal is equivalent to Print() followed by a call to os.Exit(1).
	}
}
