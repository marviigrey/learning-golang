package main 
import (
    "fmt"
    //"sync"
    )
    
func main() {
    ch := make(chan int, 5)
    ch <- 100
    ch <- 200
    ch <- 300
    close(ch)
    
    for val := range ch {
        fmt.Println(val)
    }
}