package main 

import "fmt"

func main() {
	
	z := 2

	for {
		z++
		if z == 40 {
			continue
		} else if z == 60 {
                break
        }
	fmt.Println(z)
	}
}
