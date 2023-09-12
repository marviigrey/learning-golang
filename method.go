<<<<<<< HEAD
package main
import "fmt" 

type Circle struct {
    radius float64
    area float64
}

func (c *Circle) calcArea() { /*"c *Circle" is the reciever and a pointer to the circle struct
calcArea is the method */
    c.area = c.radius * c.radius * 3.14
}

func main() {
 c := Circle{radius: 6}
c.calcArea() // to reference the method we use the reciever(c) and the method name "calcArea()"
fmt.Printf("%+v \n", c)
=======
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
>>>>>>> 1e4dbfb81fbebc65751cf958c2016c63595803df
}