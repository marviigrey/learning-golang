package main
import "fmt"

func main() {
    var c string = "emy"
    asciiCode := int(c)
    character := string(asciiCode + 3)
    fmt.Println(character)
    
}