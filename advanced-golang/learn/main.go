package main
import (
    "github.com/pborman/uuid"
    "fmt"
)
func main() {
    uuid := uuid.NewRandom()
    fmt.Println(uuid)
}