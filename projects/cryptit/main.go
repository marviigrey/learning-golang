package main

import (
	"fmt"

	"github.com/marviigrey/learning-golang/projects/cryptit/decrypt"
	"github.com/marviigrey/learning-golang/projects/cryptit/encrypt"
)

func main() {
	encryptedStr := encrypt.Nimbus("Kodekloud")
	fmt.Println(encryptedStr)
	// fmt.Println(decrypt.Nimbus(encryptedStr))
}
