package main

import (
	"fmt"
	//	"github.com/marviigrey/learning-golang/projects/cryptit/decrypt"
	"github.com/marviigrey/learning-golang/projects/cryptit/encrypt"
	"github.com/pborman/uuid"
)

func main() {
	uuid := uuid.NewRandom()
	fmt.Println(uuid)
	encryptedStr := encrypt.Nimbus("Kodekloud")
	fmt.Println(encryptedStr)
	// fmt.Println(decrypt.Nimbus(encryptedStr))
}
