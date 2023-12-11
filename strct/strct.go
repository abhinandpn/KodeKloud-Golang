package strct

import "fmt"

type Circle struct {
	X       int
	Y       int
	Radious float64
	Area    float64
}

const PI float64 = 3.14

func CalcArea(c *Circle) {

	area := (PI * c.Radious * c.Radious)
	fmt.Println(area)
}

func ComStr() {
	type S1 struct {
		x int
	}

	c := S1{x: 5}
	d := S1{x: 5}

	if c == d {
		fmt.Println("yes")
	}
}
