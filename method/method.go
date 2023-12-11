package method

import "fmt"

type Circle struct {
	radious float64
	area    float64
}

func (c Circle) CalcArea() {
	c.radious = 5
	c.area = 3.14 * c.radious * c.radious
	fmt.Printf("% + v", c)
}

type Student struct {
	name   string
	grades []int
}

func (s *Student) DisplayName() {
	fmt.Print(s.name, "  ")
}

func (s *Student) CalcPerst() float64 {
	sum := 0
	for _, v := range s.grades {
		sum = sum + v
	}
	return float64(sum*100) / float64(len(s.grades)*100)
}

func Run() {
	s := Student{name: "joe", grades: []int{90, 75, 80}}
	s.DisplayName()
	fmt.Printf("%.2f%%", s.CalcPerst())
	fmt.Println()
}
