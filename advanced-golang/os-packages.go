package main 
import (
    "fmt"
    "os"
    )
    
    func main() {
    file, err := os.OpenFile("/home/ec2-user/environment/learning-golang/test", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()
        _, err = file.WriteString("I love God\n")
        if err != nil {
            fmt.Println(err)
        }
        
    }