package main 
import (
    "sort"
    "fmt"
    )

func main() {
    vars := []int{3, 4, 5, 6, 1, 8 , 9}
    strings := []string{"jesus", "has", "risen"}
    sort.Strings(strings)
    sort.Ints(vars)
    fmt.Println(vars)
    fmt.Println(strings)
}