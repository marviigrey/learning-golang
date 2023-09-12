package main
import "fmt" 

type Circle struct {
    radius float64
    area float64
}

func calcArea(c *Circle) {
    c.area = c.radius * c.radius * 3.14
}

func main() {
    c := Circle{radius: 6}
    c.calcArea()
    fmt.Printf("%+v \n", c)
}