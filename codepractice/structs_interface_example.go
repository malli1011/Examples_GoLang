package codepractice

import "fmt"

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (c circle) area() (area float64) {
	area = 3.14 * c.radius * c.radius
	return
}

type shape interface {
	area() float64
}

func (s square) area() (area float64) {
	area = s.side * s.side
	return
}

func info(msg string, s shape) {
	fmt.Println(msg, s.area())
}

func StructsInterfaceDemo() {

	c := circle{25}
	s := square{25}

	info("Area of Circle", c)
	info("Area of Square", s)
}
