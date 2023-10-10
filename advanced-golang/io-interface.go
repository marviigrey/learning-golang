package main 
import "fmt"
type Shape interface {
  Area() float64
  Perimeter() float64
}

type Rectangle struct {
    Length float64
    Width float64
}
func (r Rectangle) Area() float64 {
    return r.Length * r.Width
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Length + r.Width)
}
func main() {
    var s Shape = Rectangle{Length: 8, Width: 5.0}
    fmt.Println(s.Area())
    fmt.Println(s.Perimeter())
}