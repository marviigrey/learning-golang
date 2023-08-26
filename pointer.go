package main
import "fmt"
func main() {
    s := "hello"
    var b *string = &s
    var a = &s
    c := &s
    fmt.Println(c)
    fmt.Println(a)
    fmt.Println(b)
}