package main
import "fmt"
func main() {
var x, y int = 7, 9
var a, b int = 4,6
var c, d int = 3, 5
var e int = 8 
z := x & y
fmt.Println(z)
fmt.Println(a | b)
fmt.Println(c ^ d)
fmt.Println(a << 1)
fmt.Println(e >> 1)
}
