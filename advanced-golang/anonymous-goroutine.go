package main
import (
    "time"
    "fmt"
    )
    
    func main() {
         
      go func() {
          
          fmt.Println("An anonymous method")
          
      }() 
      time.Sleep(2 * time.Second)
     
    }