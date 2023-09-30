package main
import (
    "fmt"
    "time"
    )
func main() {
    go start() //main go-routine qhich can strt other go routines.
    time.Sleep(1 * time.Second)
}
func start() {
    go end() //other go routine initialized after the main go routine has been executed.
    fmt.Println("In start") //labelled as the main go routine
}

func end() {
    fmt.Println("In progress")
}