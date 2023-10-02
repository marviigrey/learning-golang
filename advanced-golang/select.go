package main
import (
    "fmt"
    "time"
    
    )
    func main() {
        ch1 := make(chan string)
        ch2 := make(chan string)
        
        go goOne(ch1)
        go goTwo(ch2)
        
        select {
            case val1 := <- ch1:
            fmt.Println(val1)
            
            case val2 := <- ch2:
            fmt.Println(val2)
        }
        time.Sleep(2 * time.Second)
    }
    
    func goOne(ch1 chan string) {
        ch1 <- "channel one"
        
    }
    func goTwo(ch2 chan string) {
       ch2 <- "channel two"
    }