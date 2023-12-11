package interfaces

import "fmt"

type shape interface {
	area() float64
	perimeter() float64
}
type squre struct {
	side float64
}

type rect struct {
	length, breadth float64
}

func (s squre) area() float64 {
	return s.side * s.side
}

func (s squre) perimeter() float64 {
	return 4 * s.side
}

func (r rect) area() float64 {
	return r.length * r.breadth
}

func (r rect) perimeter() float64 {
	return 2*r.length + 2*r.breadth
}
func pritData(s shape) {
	fmt.Println(s)
	fmt.Println("areia : ", s.area())
	fmt.Println("perimeter : ", s.perimeter())
}

func Run() {
	r := rect{length: 3, breadth: 6}
	c := squre{side: 5}
	pritData(c)
	pritData(r)
}
