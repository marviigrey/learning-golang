package main
import (
    "fmt"
    )
    func main() {
        nums := [5]int{1, 2, 3, 4, 5}
        for i, item := range nums {
            fmt.Println(i, item)
        }
    }