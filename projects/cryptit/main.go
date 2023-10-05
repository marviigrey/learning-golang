package main

import (
	"fmt"
	"github.com/marviigrey/learning-golang/projects/cryptit/encrypt"
	"github.com/marviigrey/learning-golang/projects/decrypt"
)

func main() {
	fmt.Println(encrypt.Nimbus("marviigrey"))
	fmt.Println(decrypt.Nimbus(encrtptedStr))
}
