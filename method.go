package main 
import "fmt"
type Circle struct {
	radius float64
	area float64
}

type Triangle struct {
	area float64
	base float64
	height float64
}

func (t *Triangle) calcAr() {
	t.area =  t.base * t.height * 0.5
}


func (c *Circle) calcArea() {
c.area = 3.14 * c.radius * c.radius
}

func main() {
	t := Triangle{base: 6, height:8}
	t.calcAr()
	fmt.Printf("%+v \n", t)
	c := Circle{radius: 4}
	c.calcArea()
	fmt.Printf("%+v \n", c)
}